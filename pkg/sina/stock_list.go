package sina

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	countURL         = "https://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeStockCount?node=hs_a"
	pageURLTemplate  = "https://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData?page=%d&num=%d&sort=symbol&asc=1&node=hs_a&symbol=&_s_r_a=page"
	defaultPageSize  = 100
	defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36"
)

type Client struct {
	hc *http.Client
}

type Stock struct {
	Code       string
	Name       string
	Exchange   string
	MarketType string
}

type stockListItem struct {
	Symbol string `json:"symbol"`
	Code   string `json:"code"`
	Name   string `json:"name"`
}

func New() *Client {
	return &Client{hc: &http.Client{Timeout: 20 * time.Second}}
}

func (c *Client) GetAShareList(ctx context.Context) ([]Stock, error) {
	count, err := c.getStockCount(ctx)
	if err != nil {
		return nil, err
	}

	stocks := make([]Stock, 0, count)
	totalPages := (count + defaultPageSize - 1) / defaultPageSize
	for page := 1; page <= totalPages; page++ {
		items, err := c.getStockPage(ctx, page)
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			exchange := exchangeFromSymbol(item.Symbol)
			if exchange == "" {
				continue
			}
			stocks = append(stocks, Stock{
				Code:       item.Code,
				Name:       item.Name,
				Exchange:   exchange,
				MarketType: exchange,
			})
		}
	}

	sort.Slice(stocks, func(left, right int) bool {
		if stocks[left].Exchange == stocks[right].Exchange {
			return stocks[left].Code < stocks[right].Code
		}
		return stocks[left].Exchange < stocks[right].Exchange
	})
	return stocks, nil
}

func (c *Client) getStockCount(ctx context.Context) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, countURL, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", defaultUserAgent)

	resp, err := c.hc.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	count, err := strconv.Atoi(strings.Trim(string(body), "\"\n\r\t "))
	if err != nil {
		return 0, fmt.Errorf("parse sina stock count: %w", err)
	}
	return count, nil
}

func (c *Client) getStockPage(ctx context.Context, page int) ([]stockListItem, error) {
	requestURL := fmt.Sprintf(pageURLTemplate, page, defaultPageSize)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", defaultUserAgent)

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var items []stockListItem
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, fmt.Errorf("decode sina stock page %d: %w", page, err)
	}
	return items, nil
}

func exchangeFromSymbol(symbol string) string {
	if len(symbol) < 2 {
		return ""
	}
	prefix := strings.ToUpper(symbol[:2])
	switch prefix {
	case "SH", "SZ", "BJ":
		return prefix
	default:
		return ""
	}
}
