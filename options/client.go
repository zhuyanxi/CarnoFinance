package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	eastmoneyListURL  = "https://push2.eastmoney.com/api/qt/clist/get"
	eastmoneyQuoteURL = "https://push2.eastmoney.com/api/qt/stock/get"
	eastmoneyKlineURL = "https://push2his.eastmoney.com/api/qt/stock/kline/get"
	csi1000SecurityID = "1.000852"
)

var chinaTimeZone = time.FixedZone("Asia/Shanghai", 8*60*60)

type EastmoneyClient struct {
	httpClient *http.Client
}

type MarketSnapshot struct {
	RetrievedAt   time.Time          `json:"retrieved_at"`
	SpotQuoteTime *time.Time         `json:"spot_quote_time,omitempty"`
	SpotName      string             `json:"spot_name"`
	SpotPrice     float64            `json:"spot_price"`
	Contracts     []ContractSnapshot `json:"contracts"`
}

type ContractSnapshot struct {
	Contract    string  `json:"contract"`
	SecurityID  string  `json:"security_id"`
	ExpiryDate  string  `json:"expiry_date"`
	FuturePrice float64 `json:"future_price"`
	BasisMetrics
}

type listedContract struct {
	Code       string
	Market     int
	Name       string
	Price      float64
	ExpiryDate time.Time
}

type futuresListResponse struct {
	RC   int `json:"rc"`
	Data *struct {
		Total int               `json:"total"`
		Diff  []futuresListItem `json:"diff"`
	} `json:"data"`
}

type futuresListItem struct {
	Code   string          `json:"f12"`
	Market int             `json:"f13"`
	Name   string          `json:"f14"`
	Price  json.RawMessage `json:"f2"`
}

type spotQuoteResponse struct {
	RC   int `json:"rc"`
	Data *struct {
		Name      string          `json:"f58"`
		Price     json.RawMessage `json:"f43"`
		QuoteUnix json.RawMessage `json:"f86"`
	} `json:"data"`
}

func NewEastmoneyClient() *EastmoneyClient {
	return &EastmoneyClient{
		httpClient: &http.Client{Timeout: 15 * time.Second},
	}
}

func (c *EastmoneyClient) Snapshot(ctx context.Context) (MarketSnapshot, error) {
	requestContext, cancel := context.WithCancel(ctx)
	defer cancel()

	var contracts []listedContract
	var spotPrice float64
	var spotName string
	var spotQuoteTime *time.Time
	var contractsErr error
	var spotErr error
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		contracts, contractsErr = c.fetchIMContracts(requestContext)
		if contractsErr != nil {
			cancel()
		}
	}()
	go func() {
		defer waitGroup.Done()
		spotPrice, spotName, spotQuoteTime, spotErr = c.fetchSpotQuote(requestContext)
		if spotErr != nil {
			cancel()
		}
	}()
	waitGroup.Wait()
	if contractsErr != nil {
		return MarketSnapshot{}, fmt.Errorf("fetch IM contracts: %w", contractsErr)
	}
	if spotErr != nil {
		return MarketSnapshot{}, fmt.Errorf("fetch CSI 1000 quote: %w", spotErr)
	}

	retrievedAt := time.Now().In(chinaTimeZone)
	snapshot := MarketSnapshot{
		RetrievedAt:   retrievedAt,
		SpotQuoteTime: spotQuoteTime,
		SpotName:      spotName,
		SpotPrice:     spotPrice,
		Contracts:     make([]ContractSnapshot, 0, len(contracts)),
	}
	for _, contract := range contracts {
		metrics, err := calculateBasis(spotPrice, contract.Price, retrievedAt, contract.ExpiryDate)
		if err != nil {
			return MarketSnapshot{}, fmt.Errorf("calculate basis for %s: %w", contract.Name, err)
		}
		snapshot.Contracts = append(snapshot.Contracts, ContractSnapshot{
			Contract:     contract.Name,
			SecurityID:   fmt.Sprintf("%d.%s", contract.Market, contract.Code),
			ExpiryDate:   contract.ExpiryDate.Format(time.DateOnly),
			FuturePrice:  contract.Price,
			BasisMetrics: metrics,
		})
	}
	return snapshot, nil
}

func (c *EastmoneyClient) fetchIMContracts(ctx context.Context) ([]listedContract, error) {
	contracts := make([]listedContract, 0, 4)
	today := dateOnly(time.Now().In(chinaTimeZone))
	const pageSize = 100

	for page := 1; ; page++ {
		query := url.Values{
			"pn":     {strconv.Itoa(page)},
			"pz":     {strconv.Itoa(pageSize)},
			"po":     {"1"},
			"np":     {"1"},
			"fltt":   {"2"},
			"invt":   {"2"},
			"fid":    {"f3"},
			"fs":     {"m:8"},
			"fields": {"f12,f13,f14,f2"},
		}
		var payload futuresListResponse
		if err := c.getJSON(ctx, eastmoneyListURL, query, &payload); err != nil {
			return nil, err
		}
		if payload.RC != 0 || payload.Data == nil {
			return nil, fmt.Errorf("Eastmoney returned an empty CFFEX contract list (rc=%d)", payload.RC)
		}
		for _, item := range payload.Data.Diff {
			if item.Market != 8 || !isIMDeliveryContract(item.Name) || item.Code == "" {
				continue
			}
			expiryDate, err := contractExpiry(item.Name)
			if err != nil || expiryDate.Before(today) {
				continue
			}
			price, ok := parseEastmoneyNumber(item.Price)
			if !ok || price <= 0 {
				continue
			}
			contracts = append(contracts, listedContract{
				Code:       item.Code,
				Market:     item.Market,
				Name:       item.Name,
				Price:      price,
				ExpiryDate: expiryDate,
			})
		}
		if page*pageSize >= payload.Data.Total || len(payload.Data.Diff) == 0 {
			break
		}
	}

	if len(contracts) == 0 {
		return nil, fmt.Errorf("no active IM delivery contracts with valid quotes found")
	}
	sort.Slice(contracts, func(left, right int) bool {
		if contracts[left].ExpiryDate.Equal(contracts[right].ExpiryDate) {
			return contracts[left].Name < contracts[right].Name
		}
		return contracts[left].ExpiryDate.Before(contracts[right].ExpiryDate)
	})
	return contracts, nil
}

func (c *EastmoneyClient) fetchSpotQuote(ctx context.Context) (float64, string, *time.Time, error) {
	query := url.Values{
		"secid":  {csi1000SecurityID},
		"fields": {"f58,f43,f86"},
	}
	var payload spotQuoteResponse
	if err := c.getJSON(ctx, eastmoneyQuoteURL, query, &payload); err != nil {
		return 0, "", nil, err
	}
	if payload.RC != 0 || payload.Data == nil {
		return 0, "", nil, fmt.Errorf("Eastmoney returned an empty CSI 1000 quote (rc=%d)", payload.RC)
	}
	price, ok := parseEastmoneyNumber(payload.Data.Price)
	if !ok || price <= 0 {
		return 0, "", nil, fmt.Errorf("Eastmoney returned an invalid CSI 1000 price")
	}

	var quoteTime *time.Time
	if quoteUnix, ok := parseEastmoneyNumber(payload.Data.QuoteUnix); ok && quoteUnix > 0 {
		if quoteUnix > 1e12 {
			quoteUnix /= 1000
		}
		value := time.Unix(int64(quoteUnix), 0).In(chinaTimeZone)
		quoteTime = &value
	}
	return price / 100, payload.Data.Name, quoteTime, nil
}

func (c *EastmoneyClient) getJSON(ctx context.Context, endpoint string, query url.Values, target any) error {
	requestURL := endpoint
	if len(query) > 0 {
		requestURL += "?" + query.Encode()
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Referer", "https://quote.eastmoney.com")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(response.Body, 1024))
		return fmt.Errorf("Eastmoney returned %s: %s", response.Status, strings.TrimSpace(string(body)))
	}
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		return fmt.Errorf("decode Eastmoney response: %w", err)
	}
	return nil
}

func isIMDeliveryContract(name string) bool {
	if len(name) != 6 || !strings.HasPrefix(name, "IM") {
		return false
	}
	for _, digit := range name[2:] {
		if digit < '0' || digit > '9' {
			return false
		}
	}
	return true
}

func parseEastmoneyNumber(raw json.RawMessage) (float64, bool) {
	if len(raw) == 0 || string(raw) == "null" {
		return 0, false
	}
	var number float64
	if err := json.Unmarshal(raw, &number); err == nil {
		return number, true
	}
	var text string
	if err := json.Unmarshal(raw, &text); err != nil {
		return 0, false
	}
	number, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
	if err != nil {
		return 0, false
	}
	return number, true
}