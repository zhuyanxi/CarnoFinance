package domain

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

const (
	OptionTrackStatusInProgress = "IN_PROGRESS"
	OptionTrackStatusClosed     = "CLOSED"

	OptionTrackInstrumentTypeOption = "OPTION"
	OptionTrackInstrumentTypeFuture = "FUTURE"

	OptionTrackTradeTypeBuyOpen   = "BUY_OPEN"
	OptionTrackTradeTypeSellOpen  = "SELL_OPEN"
	OptionTrackTradeTypeBuyClose  = "BUY_CLOSE"
	OptionTrackTradeTypeSellClose = "SELL_CLOSE"

	OptionTrackEntrySourceInitial = "INITIAL"
	OptionTrackEntrySourceManual  = "MANUAL"
)

var (
	ErrOptionTrackNotFound        = errors.New("option track plan not found")
	ErrOptionTrackTradeNotFound   = errors.New("option track trade not found")
	ErrOptionTrackClosedPlan      = errors.New("option track plan is closed")
	ErrOptionTrackOpenPositions   = errors.New("option track plan still has open positions")
	ErrOptionTrackInvalidPayload  = errors.New("invalid option track payload")
	ErrOptionTrackInvalidHistory  = errors.New("option track trade history is invalid")
	ErrOptionTrackClosedEditBreak = errors.New("editing closed plan would reopen positions")
)

type OptionTrackService struct {
	ctx context.Context
	db  *bun.DB
}

type OptionTrackPlan struct {
	bun.BaseModel `bun:"table:option_track_plans"`

	ID           int64      `bun:",pk,autoincrement" json:"id"`
	Name         string     `bun:",notnull" json:"name"`
	Status       string     `bun:",notnull" json:"status"`
	CreatedAt    time.Time  `bun:",notnull" json:"created_at"`
	ClosedAt     *time.Time `bun:",nullzero" json:"closed_at,omitempty"`
	LatestTradeAt *time.Time `bun:",nullzero" json:"latest_trade_at,omitempty"`
}

type OptionTrackTradeRecord struct {
	bun.BaseModel `bun:"table:option_track_trade_records"`

	ID                 int64     `bun:",pk,autoincrement" json:"id"`
	PlanID             int64     `bun:",notnull" json:"plan_id"`
	InstrumentType     string    `bun:"-" json:"instrument_type"`
	ContractName       string    `bun:",notnull" json:"contract_name"`
	ContractCode       string    `json:"contract_code,omitempty"`
	ExpiryDate         string    `bun:",notnull" json:"expiry_date"`
	StrikePrice        float64   `bun:",notnull" json:"strike_price"`
	OptionSide         string    `bun:",notnull" json:"option_side"`
	ContractMultiplier int       `bun:",notnull" json:"contract_multiplier"`
	TradeType          string    `bun:",notnull" json:"trade_type"`
	TradeTime          time.Time `bun:",notnull" json:"trade_time"`
	Price              float64   `bun:",notnull" json:"price"`
	Quantity           int       `bun:",notnull" json:"quantity"`
	Amount             float64   `bun:",notnull" json:"amount"`
	EntrySource        string    `bun:",notnull" json:"entry_source"`
	Remark             string    `json:"remark,omitempty"`
	CreatedAt          time.Time `bun:",notnull" json:"created_at"`
	UpdatedAt          time.Time `bun:",notnull" json:"updated_at"`
}

type OptionTrackPlanCreateRequest struct {
	Name      string                         `json:"name"`
	Positions []OptionTrackTradeUpsertInput `json:"positions"`
}

type OptionTrackTradeUpsertInput struct {
	InstrumentType     string  `json:"instrument_type"`
	ContractName       string  `json:"contract_name"`
	ContractCode       string  `json:"contract_code"`
	ExpiryDate         string  `json:"expiry_date"`
	StrikePrice        float64 `json:"strike_price"`
	OptionSide         string  `json:"option_side"`
	ContractMultiplier int     `json:"contract_multiplier"`
	TradeType          string  `json:"trade_type"`
	TradeTime          string  `json:"trade_time"`
	Price              float64 `json:"price"`
	Quantity           int     `json:"quantity"`
	Remark             string  `json:"remark"`
}

type OptionTrackPlanSummary struct {
	ID            int64      `json:"id"`
	Name          string     `json:"name"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	ClosedAt      *time.Time `json:"closed_at,omitempty"`
	LatestTradeAt *time.Time `json:"latest_trade_at,omitempty"`
	ContractCount int        `json:"contract_count"`
}

type OptionTrackPosition struct {
	ContractKey        string  `json:"contract_key"`
	InstrumentType     string  `json:"instrument_type"`
	ContractName       string  `json:"contract_name"`
	ContractCode       string  `json:"contract_code,omitempty"`
	ExpiryDate         string  `json:"expiry_date"`
	StrikePrice        float64 `json:"strike_price"`
	OptionSide         string  `json:"option_side"`
	ContractMultiplier int     `json:"contract_multiplier"`
	LongQuantity       int     `json:"long_quantity"`
	ShortQuantity      int     `json:"short_quantity"`
	NetQuantity        int     `json:"net_quantity"`
	CanSellClose       int     `json:"can_sell_close"`
	CanBuyClose        int     `json:"can_buy_close"`
	LastTradePrice     float64 `json:"last_trade_price"`
	AverageLongPrice   float64 `json:"average_long_price"`
	AverageShortPrice  float64 `json:"average_short_price"`
}

type OptionTrackPlanDetail struct {
	Plan      OptionTrackPlanSummary     `json:"plan"`
	Trades    []OptionTrackTradeRecord   `json:"trades"`
	Positions []OptionTrackPosition      `json:"positions"`
}

type OptionTrackPlanListResponse struct {
	Items     []OptionTrackPlanSummary `json:"items"`
	Page      int                      `json:"page"`
	PageSize  int                      `json:"page_size"`
	Total     int                      `json:"total"`
	TotalPage int                      `json:"total_page"`
}

type OptionTrackPnLContract struct {
	ContractKey        string  `json:"contract_key"`
	InstrumentType     string  `json:"instrument_type"`
	ContractName       string  `json:"contract_name"`
	ContractCode       string  `json:"contract_code,omitempty"`
	ExpiryDate         string  `json:"expiry_date"`
	StrikePrice        float64 `json:"strike_price"`
	OptionSide         string  `json:"option_side"`
	ContractMultiplier int     `json:"contract_multiplier"`
	CurrentPrice       float64 `json:"current_price"`
	RealizedPnL        float64 `json:"realized_pnl"`
	UnrealizedPnL      float64 `json:"unrealized_pnl"`
	TotalPnL           float64 `json:"total_pnl"`
	LongQuantity       int     `json:"long_quantity"`
	ShortQuantity      int     `json:"short_quantity"`
	AverageLongPrice   float64 `json:"average_long_price"`
	AverageShortPrice  float64 `json:"average_short_price"`
}

type OptionTrackPnLResponse struct {
	PlanID         int64                    `json:"plan_id"`
	Status         string                   `json:"status"`
	RealizedPnL    float64                  `json:"realized_pnl"`
	UnrealizedPnL  float64                  `json:"unrealized_pnl"`
	TotalPnL       float64                  `json:"total_pnl"`
	Contracts      []OptionTrackPnLContract `json:"contracts"`
}

type OptionTrackPnLPreviewRequest struct {
	Prices []OptionTrackContractPriceInput `json:"prices"`
}

type OptionTrackContractPriceInput struct {
	ContractKey  string  `json:"contract_key"`
	CurrentPrice float64 `json:"current_price"`
}

type optionTrackLedgerState struct {
	planID        int64
	positions     []OptionTrackPosition
	latestTradeAt *time.Time
	hasOpen       bool
	contractCount int
	contracts     []OptionTrackPnLContract
	priceMap      map[string]float64
	realizedPnL   float64
	unrealizedPnL float64
	totalPnL      float64
}

type optionTrackAggregate struct {
	InstrumentType     string
	ContractName       string
	ContractCode       string
	ExpiryDate         string
	StrikePrice        float64
	OptionSide         string
	ContractMultiplier int
	LongQuantity       int
	ShortQuantity      int
	LongCost           float64
	ShortProceeds      float64
	RealizedPnL        float64
	LastTradePrice     float64
}

func NewOptionTrackService(ctx context.Context, db *bun.DB) *OptionTrackService {
	return &OptionTrackService{ctx: ctx, db: db}
}

func (s *OptionTrackService) Init() error {
	if _, err := s.db.NewCreateTable().Model((*OptionTrackPlan)(nil)).IfNotExists().Exec(s.ctx); err != nil {
		return err
	}
	if _, err := s.db.NewCreateTable().Model((*OptionTrackTradeRecord)(nil)).IfNotExists().Exec(s.ctx); err != nil {
		return err
	}
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_option_track_trade_records_plan_time ON option_track_trade_records(plan_id, trade_time)",
		"CREATE INDEX IF NOT EXISTS idx_option_track_trade_records_contract ON option_track_trade_records(plan_id, contract_name, expiry_date, strike_price, option_side)",
	}
	for _, stmt := range indexes {
		if _, err := s.db.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (s *OptionTrackService) CreatePlan(req OptionTrackPlanCreateRequest) (*OptionTrackPlanDetail, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" || len(req.Positions) == 0 {
		return nil, ErrOptionTrackInvalidPayload
	}

	now := time.Now()
	plan := &OptionTrackPlan{
		Name:      name,
		Status:    OptionTrackStatusInProgress,
		CreatedAt: now,
	}

	trades := make([]OptionTrackTradeRecord, 0, len(req.Positions))
	for _, input := range req.Positions {
		trade, err := s.tradeFromInput(0, input, OptionTrackEntrySourceInitial, now)
		if err != nil {
			return nil, err
		}
		if trade.TradeType != OptionTrackTradeTypeBuyOpen && trade.TradeType != OptionTrackTradeTypeSellOpen {
			return nil, ErrOptionTrackInvalidPayload
		}
		trades = append(trades, trade)
	}
	ledger, err := buildOptionTrackLedger(0, trades, nil)
	if err != nil {
		return nil, err
	}

	err = s.db.RunInTx(s.ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(plan).Exec(ctx); err != nil {
			return err
		}
		for i := range trades {
			trades[i].PlanID = plan.ID
		}
		if _, err := tx.NewInsert().Model(&trades).Exec(ctx); err != nil {
			return err
		}
		if err := s.updatePlanMetadataTx(ctx, tx, plan.ID, ledger.latestTradeAt, nil, OptionTrackStatusInProgress); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.GetPlanDetail(plan.ID)
	}

func (s *OptionTrackService) ListPlans(page, pageSize int) (*OptionTrackPlanListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	total, err := s.db.NewSelect().Model((*OptionTrackPlan)(nil)).Count(s.ctx)
	if err != nil {
		return nil, err
	}

	plans := make([]OptionTrackPlan, 0)
	err = s.db.NewSelect().Model(&plans).
		Order("created_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Scan(s.ctx)
	if err != nil {
		return nil, err
	}

	summaries := make([]OptionTrackPlanSummary, 0, len(plans))
	for _, plan := range plans {
		contractCount, err := s.countPlanContracts(plan.ID)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, toOptionTrackPlanSummary(plan, contractCount))
	}

	totalPage := 0
	if total > 0 {
		totalPage = int(math.Ceil(float64(total) / float64(pageSize)))
	}

	return &OptionTrackPlanListResponse{
		Items:     summaries,
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: totalPage,
	}, nil
}

func (s *OptionTrackService) GetPlanDetail(planID int64) (*OptionTrackPlanDetail, error) {
	plan, trades, err := s.loadPlanAndTrades(planID)
	if err != nil {
		return nil, err
	}
	ledger, err := buildOptionTrackLedger(plan.ID, trades, nil)
	if err != nil {
		return nil, err
	}
	return &OptionTrackPlanDetail{
		Plan:      toOptionTrackPlanSummary(*plan, ledger.contractCount),
		Trades:    trades,
		Positions: ledger.positions,
	}, nil
}

func (s *OptionTrackService) CreateTrade(planID int64, req OptionTrackTradeUpsertInput) (*OptionTrackPlanDetail, error) {
	plan, trades, err := s.loadPlanAndTrades(planID)
	if err != nil {
		return nil, err
	}
	if plan.Status == OptionTrackStatusClosed {
		return nil, ErrOptionTrackClosedPlan
	}

	trade, err := s.tradeFromInput(planID, req, OptionTrackEntrySourceManual, time.Now())
	if err != nil {
		return nil, err
	}
	candidate := append(append([]OptionTrackTradeRecord{}, trades...), trade)
	ledger, err := buildOptionTrackLedger(planID, candidate, nil)
	if err != nil {
		return nil, err
	}

	err = s.db.RunInTx(s.ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(&trade).Exec(ctx); err != nil {
			return err
		}
		return s.updatePlanMetadataTx(ctx, tx, planID, ledger.latestTradeAt, nil, plan.Status)
	})
	if err != nil {
		return nil, err
	}
	return s.GetPlanDetail(planID)
}

func (s *OptionTrackService) UpdateTrade(planID, tradeID int64, req OptionTrackTradeUpsertInput) (*OptionTrackPlanDetail, error) {
	plan, trades, err := s.loadPlanAndTrades(planID)
	if err != nil {
		return nil, err
	}

	updated, err := s.tradeFromInput(planID, req, OptionTrackEntrySourceManual, time.Now())
	if err != nil {
		return nil, err
	}
	updated.ID = tradeID

	found := false
	candidate := make([]OptionTrackTradeRecord, 0, len(trades))
	for _, trade := range trades {
		if trade.ID == tradeID {
			updated.CreatedAt = trade.CreatedAt
			candidate = append(candidate, updated)
			found = true
			continue
		}
		candidate = append(candidate, trade)
	}
	if !found {
		return nil, ErrOptionTrackTradeNotFound
	}

	ledger, err := buildOptionTrackLedger(planID, candidate, nil)
	if err != nil {
		return nil, err
	}
	if plan.Status == OptionTrackStatusClosed && ledger.hasOpen {
		return nil, ErrOptionTrackClosedEditBreak
	}

	err = s.db.RunInTx(s.ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewUpdate().Model(&updated).Where("id = ? AND plan_id = ?", tradeID, planID).Exec(ctx); err != nil {
			return err
		}
		closedAt := plan.ClosedAt
		status := plan.Status
		return s.updatePlanMetadataTx(ctx, tx, planID, ledger.latestTradeAt, closedAt, status)
	})
	if err != nil {
		return nil, err
	}
	return s.GetPlanDetail(planID)
}

func (s *OptionTrackService) DeleteTrade(planID, tradeID int64) (*OptionTrackPlanDetail, error) {
	plan, trades, err := s.loadPlanAndTrades(planID)
	if err != nil {
		return nil, err
	}

	found := false
	candidate := make([]OptionTrackTradeRecord, 0, len(trades))
	for _, trade := range trades {
		if trade.ID == tradeID {
			found = true
			continue
		}
		candidate = append(candidate, trade)
	}
	if !found {
		return nil, ErrOptionTrackTradeNotFound
	}

	ledger, err := buildOptionTrackLedger(planID, candidate, nil)
	if err != nil {
		return nil, err
	}
	if plan.Status == OptionTrackStatusClosed && ledger.hasOpen {
		return nil, ErrOptionTrackClosedEditBreak
	}

	err = s.db.RunInTx(s.ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewDelete().Model((*OptionTrackTradeRecord)(nil)).Where("id = ? AND plan_id = ?", tradeID, planID).Exec(ctx); err != nil {
			return err
		}
		closedAt := plan.ClosedAt
		status := plan.Status
		return s.updatePlanMetadataTx(ctx, tx, planID, ledger.latestTradeAt, closedAt, status)
	})
	if err != nil {
		return nil, err
	}
	return s.GetPlanDetail(planID)
}

func (s *OptionTrackService) ClosePlan(planID int64) (*OptionTrackPlanSummary, error) {
	plan, trades, err := s.loadPlanAndTrades(planID)
	if err != nil {
		return nil, err
	}
	if plan.Status == OptionTrackStatusClosed {
		count, err := s.countPlanContracts(planID)
		if err != nil {
			return nil, err
		}
		summary := toOptionTrackPlanSummary(*plan, count)
		return &summary, nil
	}
	ledger, err := buildOptionTrackLedger(planID, trades, nil)
	if err != nil {
		return nil, err
	}
	if ledger.hasOpen {
		return nil, ErrOptionTrackOpenPositions
	}

	closedAt := time.Now()
	err = s.db.RunInTx(s.ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return s.updatePlanMetadataTx(ctx, tx, planID, ledger.latestTradeAt, &closedAt, OptionTrackStatusClosed)
	})
	if err != nil {
		return nil, err
	}
	plan.Status = OptionTrackStatusClosed
	plan.ClosedAt = &closedAt
	plan.LatestTradeAt = ledger.latestTradeAt
	summary := toOptionTrackPlanSummary(*plan, ledger.contractCount)
	return &summary, nil
}

func (s *OptionTrackService) GetPlanPnL(planID int64) (*OptionTrackPnLResponse, error) {
	plan, trades, err := s.loadPlanAndTrades(planID)
	if err != nil {
		return nil, err
	}
	ledger, err := buildOptionTrackLedger(planID, trades, nil)
	if err != nil {
		return nil, err
	}
	return &OptionTrackPnLResponse{
		PlanID:        plan.ID,
		Status:        plan.Status,
		RealizedPnL:   roundMoney(ledger.realizedPnL),
		UnrealizedPnL: roundMoney(ledger.unrealizedPnL),
		TotalPnL:      roundMoney(ledger.totalPnL),
		Contracts:     ledger.contracts,
	}, nil
}

func (s *OptionTrackService) PreviewPlanPnL(planID int64, req OptionTrackPnLPreviewRequest) (*OptionTrackPnLResponse, error) {
	plan, trades, err := s.loadPlanAndTrades(planID)
	if err != nil {
		return nil, err
	}
	overrides := make(map[string]float64, len(req.Prices))
	for _, item := range req.Prices {
		if strings.TrimSpace(item.ContractKey) == "" || item.CurrentPrice <= 0 {
			return nil, ErrOptionTrackInvalidPayload
		}
		overrides[item.ContractKey] = item.CurrentPrice
	}
	ledger, err := buildOptionTrackLedger(planID, trades, overrides)
	if err != nil {
		return nil, err
	}
	return &OptionTrackPnLResponse{
		PlanID:        plan.ID,
		Status:        plan.Status,
		RealizedPnL:   roundMoney(ledger.realizedPnL),
		UnrealizedPnL: roundMoney(ledger.unrealizedPnL),
		TotalPnL:      roundMoney(ledger.totalPnL),
		Contracts:     ledger.contracts,
	}, nil
}

func (s *OptionTrackService) countPlanContracts(planID int64) (int, error) {
	trades := make([]OptionTrackTradeRecord, 0)
	if err := s.db.NewSelect().Model(&trades).Where("plan_id = ?", planID).Scan(s.ctx); err != nil {
		return 0, err
	}
	keys := make(map[string]struct{})
	for _, trade := range trades {
		trade = hydrateOptionTrackTrade(trade)
		keys[optionTrackContractKey(trade.InstrumentType, trade.ContractName, trade.ContractCode, trade.ExpiryDate, trade.StrikePrice, trade.OptionSide, trade.ContractMultiplier)] = struct{}{}
	}
	return len(keys), nil
}

func (s *OptionTrackService) loadPlanAndTrades(planID int64) (*OptionTrackPlan, []OptionTrackTradeRecord, error) {
	plan := new(OptionTrackPlan)
	if err := s.db.NewSelect().Model(plan).Where("id = ?", planID).Scan(s.ctx); err != nil {
		return nil, nil, ErrOptionTrackNotFound
	}
	trades := make([]OptionTrackTradeRecord, 0)
	if err := s.db.NewSelect().Model(&trades).Where("plan_id = ?", planID).Order("trade_time ASC").Order("id ASC").Scan(s.ctx); err != nil {
		return nil, nil, err
	}
	for i := range trades {
		trades[i] = hydrateOptionTrackTrade(trades[i])
	}
	return plan, trades, nil
}

func (s *OptionTrackService) updatePlanMetadataTx(ctx context.Context, tx bun.Tx, planID int64, latestTradeAt *time.Time, closedAt *time.Time, status string) error {
	plan := &OptionTrackPlan{
		ID:            planID,
		Status:        status,
		ClosedAt:      closedAt,
		LatestTradeAt: latestTradeAt,
	}
	_, err := tx.NewUpdate().Model(plan).
		Column("status", "closed_at", "latest_trade_at").
		WherePK().
		Exec(ctx)
	return err
}

func (s *OptionTrackService) tradeFromInput(planID int64, input OptionTrackTradeUpsertInput, entrySource string, now time.Time) (OptionTrackTradeRecord, error) {
	tradeTime, err := parseOptionTrackTime(input.TradeTime)
	if err != nil {
		return OptionTrackTradeRecord{}, ErrOptionTrackInvalidPayload
	}
	instrumentType := normalizeInstrumentType(input.InstrumentType, input.OptionSide, input.StrikePrice)
	trade := OptionTrackTradeRecord{
		PlanID:             planID,
		InstrumentType:     instrumentType,
		ContractName:       strings.TrimSpace(input.ContractName),
		ContractCode:       strings.TrimSpace(input.ContractCode),
		ExpiryDate:         strings.TrimSpace(input.ExpiryDate),
		StrikePrice:        input.StrikePrice,
		OptionSide:         normalizeOptionSide(input.OptionSide),
		ContractMultiplier: input.ContractMultiplier,
		TradeType:          normalizeTradeType(input.TradeType),
		TradeTime:          tradeTime,
		Price:              input.Price,
		Quantity:           input.Quantity,
		Amount:             roundMoney(input.Price * float64(input.Quantity*input.ContractMultiplier)),
		EntrySource:        entrySource,
		Remark:             strings.TrimSpace(input.Remark),
		CreatedAt:          now,
		UpdatedAt:          now,
	}
	if instrumentType == OptionTrackInstrumentTypeFuture {
		trade.StrikePrice = 0
		trade.OptionSide = ""
	}
	if err := validateOptionTrackTrade(trade); err != nil {
		return OptionTrackTradeRecord{}, err
	}
	return trade, nil
}

func validateOptionTrackTrade(trade OptionTrackTradeRecord) error {
	if strings.TrimSpace(trade.ContractName) == "" {
		return ErrOptionTrackInvalidPayload
	}
	if trade.Price <= 0 || trade.Quantity <= 0 || trade.ContractMultiplier <= 0 {
		return ErrOptionTrackInvalidPayload
	}
	switch trade.TradeType {
	case OptionTrackTradeTypeBuyOpen, OptionTrackTradeTypeSellOpen, OptionTrackTradeTypeBuyClose, OptionTrackTradeTypeSellClose:
	default:
		return ErrOptionTrackInvalidPayload
	}
	if trade.TradeTime.IsZero() {
		return ErrOptionTrackInvalidPayload
	}
	switch optionTrackInstrumentType(trade) {
	case OptionTrackInstrumentTypeOption:
		if strings.TrimSpace(trade.ExpiryDate) == "" || trade.StrikePrice <= 0 {
			return ErrOptionTrackInvalidPayload
		}
		if trade.OptionSide != "CALL" && trade.OptionSide != "PUT" {
			return ErrOptionTrackInvalidPayload
		}
	case OptionTrackInstrumentTypeFuture:
		if trade.OptionSide != "" || trade.StrikePrice != 0 {
			return ErrOptionTrackInvalidPayload
		}
	default:
		return ErrOptionTrackInvalidPayload
	}
	return nil
}

func buildOptionTrackLedger(planID int64, trades []OptionTrackTradeRecord, currentPriceOverrides map[string]float64) (*optionTrackLedgerState, error) {
	state := &optionTrackLedgerState{
		planID:   planID,
		priceMap: make(map[string]float64),
	}
	if len(trades) == 0 {
		return state, nil
	}

	sortedTrades := append([]OptionTrackTradeRecord{}, trades...)
	sort.Slice(sortedTrades, func(i, j int) bool {
		if sortedTrades[i].TradeTime.Equal(sortedTrades[j].TradeTime) {
			return sortedTrades[i].ID < sortedTrades[j].ID
		}
		return sortedTrades[i].TradeTime.Before(sortedTrades[j].TradeTime)
	})

	aggs := make(map[string]*optionTrackAggregate)
	keys := make([]string, 0)

	for _, trade := range sortedTrades {
		trade = hydrateOptionTrackTrade(trade)
		if err := validateOptionTrackTrade(trade); err != nil {
			return nil, err
		}
		key := optionTrackContractKey(trade.InstrumentType, trade.ContractName, trade.ContractCode, trade.ExpiryDate, trade.StrikePrice, trade.OptionSide, trade.ContractMultiplier)
		agg, ok := aggs[key]
		if !ok {
			agg = &optionTrackAggregate{
				InstrumentType:     trade.InstrumentType,
				ContractName:       trade.ContractName,
				ContractCode:       trade.ContractCode,
				ExpiryDate:         trade.ExpiryDate,
				StrikePrice:        trade.StrikePrice,
				OptionSide:         trade.OptionSide,
				ContractMultiplier: trade.ContractMultiplier,
			}
			aggs[key] = agg
			keys = append(keys, key)
		}

		agg.LastTradePrice = trade.Price
		state.priceMap[key] = trade.Price
		switch trade.TradeType {
		case OptionTrackTradeTypeBuyOpen:
			agg.LongQuantity += trade.Quantity
			agg.LongCost += trade.Price * float64(trade.Quantity*trade.ContractMultiplier)
		case OptionTrackTradeTypeSellClose:
			if agg.LongQuantity < trade.Quantity {
				return nil, fmt.Errorf("%w: not enough long position for %s", ErrOptionTrackInvalidHistory, trade.ContractName)
			}
			avg := 0.0
			if agg.LongQuantity > 0 {
				avg = agg.LongCost / float64(agg.LongQuantity*agg.ContractMultiplier)
			}
			agg.RealizedPnL += (trade.Price - avg) * float64(trade.Quantity*trade.ContractMultiplier)
			agg.LongCost -= avg * float64(trade.Quantity*trade.ContractMultiplier)
			agg.LongQuantity -= trade.Quantity
		case OptionTrackTradeTypeSellOpen:
			agg.ShortQuantity += trade.Quantity
			agg.ShortProceeds += trade.Price * float64(trade.Quantity*trade.ContractMultiplier)
		case OptionTrackTradeTypeBuyClose:
			if agg.ShortQuantity < trade.Quantity {
				return nil, fmt.Errorf("%w: not enough short position for %s", ErrOptionTrackInvalidHistory, trade.ContractName)
			}
			avg := 0.0
			if agg.ShortQuantity > 0 {
				avg = agg.ShortProceeds / float64(agg.ShortQuantity*agg.ContractMultiplier)
			}
			agg.RealizedPnL += (avg - trade.Price) * float64(trade.Quantity*trade.ContractMultiplier)
			agg.ShortProceeds -= avg * float64(trade.Quantity*trade.ContractMultiplier)
			agg.ShortQuantity -= trade.Quantity
		}

		latest := trade.TradeTime
		if state.latestTradeAt == nil || latest.After(*state.latestTradeAt) {
			state.latestTradeAt = &latest
		}
	}

	sort.Strings(keys)
	state.contractCount = len(keys)
	state.positions = make([]OptionTrackPosition, 0, len(keys))
	state.contracts = make([]OptionTrackPnLContract, 0, len(keys))

	for _, key := range keys {
		agg := aggs[key]
		avgLong := 0.0
		if agg.LongQuantity > 0 {
			avgLong = agg.LongCost / float64(agg.LongQuantity*agg.ContractMultiplier)
		}
		avgShort := 0.0
		if agg.ShortQuantity > 0 {
			avgShort = agg.ShortProceeds / float64(agg.ShortQuantity*agg.ContractMultiplier)
		}
		currentPrice := agg.LastTradePrice
		if price, ok := currentPriceOverrides[key]; ok {
			currentPrice = price
		}
		unrealized := 0.0
		if agg.LongQuantity > 0 {
			unrealized += (currentPrice - avgLong) * float64(agg.LongQuantity*agg.ContractMultiplier)
		}
		if agg.ShortQuantity > 0 {
			unrealized += (avgShort - currentPrice) * float64(agg.ShortQuantity*agg.ContractMultiplier)
		}
		if agg.LongQuantity > 0 || agg.ShortQuantity > 0 {
			state.hasOpen = true
		}
		position := OptionTrackPosition{
			ContractKey:        key,
			InstrumentType:     agg.InstrumentType,
			ContractName:       agg.ContractName,
			ContractCode:       agg.ContractCode,
			ExpiryDate:         agg.ExpiryDate,
			StrikePrice:        agg.StrikePrice,
			OptionSide:         agg.OptionSide,
			ContractMultiplier: agg.ContractMultiplier,
			LongQuantity:       agg.LongQuantity,
			ShortQuantity:      agg.ShortQuantity,
			NetQuantity:        agg.LongQuantity - agg.ShortQuantity,
			CanSellClose:       agg.LongQuantity,
			CanBuyClose:        agg.ShortQuantity,
			LastTradePrice:     roundMoney(agg.LastTradePrice),
			AverageLongPrice:   roundMoney(avgLong),
			AverageShortPrice:  roundMoney(avgShort),
		}
		state.positions = append(state.positions, position)
		contractPnL := OptionTrackPnLContract{
			ContractKey:        key,
			InstrumentType:     agg.InstrumentType,
			ContractName:       agg.ContractName,
			ContractCode:       agg.ContractCode,
			ExpiryDate:         agg.ExpiryDate,
			StrikePrice:        agg.StrikePrice,
			OptionSide:         agg.OptionSide,
			ContractMultiplier: agg.ContractMultiplier,
			CurrentPrice:       roundMoney(currentPrice),
			RealizedPnL:        roundMoney(agg.RealizedPnL),
			UnrealizedPnL:      roundMoney(unrealized),
			TotalPnL:           roundMoney(agg.RealizedPnL + unrealized),
			LongQuantity:       agg.LongQuantity,
			ShortQuantity:      agg.ShortQuantity,
			AverageLongPrice:   roundMoney(avgLong),
			AverageShortPrice:  roundMoney(avgShort),
		}
		state.contracts = append(state.contracts, contractPnL)
		state.realizedPnL += agg.RealizedPnL
		state.unrealizedPnL += unrealized
	}
	state.totalPnL = state.realizedPnL + state.unrealizedPnL
	state.realizedPnL = roundMoney(state.realizedPnL)
	state.unrealizedPnL = roundMoney(state.unrealizedPnL)
	state.totalPnL = roundMoney(state.totalPnL)
	return state, nil
}

func toOptionTrackPlanSummary(plan OptionTrackPlan, contractCount int) OptionTrackPlanSummary {
	return OptionTrackPlanSummary{
		ID:            plan.ID,
		Name:          plan.Name,
		Status:        plan.Status,
		CreatedAt:     plan.CreatedAt,
		ClosedAt:      plan.ClosedAt,
		LatestTradeAt: plan.LatestTradeAt,
		ContractCount: contractCount,
	}
}

func parseOptionTrackTime(raw string) (time.Time, error) {
	value := strings.TrimSpace(raw)
	if value == "" {
		return time.Time{}, ErrOptionTrackInvalidPayload
	}
	layouts := []string{time.RFC3339, "2006-01-02T15:04", "2006-01-02 15:04", "2006-01-02"}
	for _, layout := range layouts {
		if parsed, err := time.ParseInLocation(layout, value, time.Local); err == nil {
			return parsed, nil
		}
	}
	return time.Time{}, ErrOptionTrackInvalidPayload
}

func optionTrackContractKey(instrumentType, name, code, expiry string, strike float64, side string, multiplier int) string {
	instrumentType = normalizeInstrumentType(instrumentType, side, strike)
	if instrumentType == OptionTrackInstrumentTypeFuture {
		return fmt.Sprintf("%s|%s|%s|%s|%d", instrumentType, strings.TrimSpace(name), strings.TrimSpace(code), strings.TrimSpace(expiry), multiplier)
	}
	return fmt.Sprintf("%s|%s|%s|%s|%.4f|%s|%d", instrumentType, strings.TrimSpace(name), strings.TrimSpace(code), strings.TrimSpace(expiry), strike, strings.TrimSpace(side), multiplier)
}

func normalizeInstrumentType(value, optionSide string, strike float64) string {
	upper := strings.ToUpper(strings.TrimSpace(value))
	upper = strings.ReplaceAll(upper, "-", "_")
	upper = strings.ReplaceAll(upper, " ", "_")
	switch upper {
	case "期权":
		return OptionTrackInstrumentTypeOption
	case "期货":
		return OptionTrackInstrumentTypeFuture
	case OptionTrackInstrumentTypeOption, OptionTrackInstrumentTypeFuture:
		return upper
	}
	if strings.TrimSpace(optionSide) != "" || strike > 0 {
		return OptionTrackInstrumentTypeOption
	}
	return OptionTrackInstrumentTypeFuture
}

func optionTrackInstrumentType(trade OptionTrackTradeRecord) string {
	return normalizeInstrumentType(trade.InstrumentType, trade.OptionSide, trade.StrikePrice)
}

func hydrateOptionTrackTrade(trade OptionTrackTradeRecord) OptionTrackTradeRecord {
	trade.InstrumentType = optionTrackInstrumentType(trade)
	if trade.InstrumentType == OptionTrackInstrumentTypeFuture {
		trade.OptionSide = ""
		trade.StrikePrice = 0
	}
	return trade
}

func normalizeTradeType(value string) string {
	upper := strings.ToUpper(strings.TrimSpace(value))
	upper = strings.ReplaceAll(upper, "-", "_")
	upper = strings.ReplaceAll(upper, " ", "_")
	switch upper {
	case "买开":
		return OptionTrackTradeTypeBuyOpen
	case "卖开":
		return OptionTrackTradeTypeSellOpen
	case "买平":
		return OptionTrackTradeTypeBuyClose
	case "卖平":
		return OptionTrackTradeTypeSellClose
	default:
		return upper
	}
}

func normalizeOptionSide(value string) string {
	upper := strings.ToUpper(strings.TrimSpace(value))
	switch upper {
	case "认购":
		return "CALL"
	case "认沽":
		return "PUT"
	default:
		return upper
	}
}

func roundMoney(value float64) float64 {
	return math.Round(value*100) / 100
}