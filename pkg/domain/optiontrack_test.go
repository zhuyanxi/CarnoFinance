package domain

import (
	"context"
	"database/sql"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func newTestOptionTrackService(t *testing.T) *OptionTrackService {
	t.Helper()
	dbPath := filepath.Join(t.TempDir(), "optiontrack-test.db")
	sqldb, err := sql.Open(sqliteshim.ShimName, dbPath)
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	t.Cleanup(func() {
		_ = sqldb.Close()
	})
	bunDB := bun.NewDB(sqldb, sqlitedialect.New())
	t.Cleanup(func() {
		_ = bunDB.Close()
	})
	service := NewOptionTrackService(context.Background(), bunDB)
	if err := service.Init(); err != nil {
		t.Fatalf("init service: %v", err)
	}
	return service
}

func TestOptionTrackCreatePlanAutoCalculatesAmount(t *testing.T) {
	service := newTestOptionTrackService(t)

	detail, err := service.CreatePlan(OptionTrackPlanCreateRequest{
		Name: "test plan",
		Positions: []OptionTrackTradeUpsertInput{{
			ContractName:       "MO2406-C-5000",
			ExpiryDate:         "2026-06-28",
			StrikePrice:        5000,
			OptionSide:         "CALL",
			ContractMultiplier: 100,
			TradeType:          OptionTrackTradeTypeBuyOpen,
			TradeTime:          "2026-05-09T10:00",
			Price:              12.5,
			Quantity:           2,
		}},
	})
	if err != nil {
		t.Fatalf("create plan: %v", err)
	}

	if len(detail.Trades) != 1 {
		t.Fatalf("expected 1 trade, got %d", len(detail.Trades))
	}
	if detail.Trades[0].Amount != 2500 {
		t.Fatalf("expected auto amount 2500, got %v", detail.Trades[0].Amount)
	}
	if detail.Plan.Status != OptionTrackStatusInProgress {
		t.Fatalf("expected in progress, got %s", detail.Plan.Status)
	}
}

func TestOptionTrackRejectsOverClose(t *testing.T) {
	service := newTestOptionTrackService(t)

	detail, err := service.CreatePlan(OptionTrackPlanCreateRequest{
		Name: "close test",
		Positions: []OptionTrackTradeUpsertInput{{
			ContractName:       "MO2406-C-5000",
			ExpiryDate:         "2026-06-28",
			StrikePrice:        5000,
			OptionSide:         "CALL",
			ContractMultiplier: 100,
			TradeType:          OptionTrackTradeTypeBuyOpen,
			TradeTime:          "2026-05-09T10:00",
			Price:              10,
			Quantity:           1,
		}},
	})
	if err != nil {
		t.Fatalf("create plan: %v", err)
	}

	_, err = service.CreateTrade(detail.Plan.ID, OptionTrackTradeUpsertInput{
		ContractName:       "MO2406-C-5000",
		ExpiryDate:         "2026-06-28",
		StrikePrice:        5000,
		OptionSide:         "CALL",
		ContractMultiplier: 100,
		TradeType:          OptionTrackTradeTypeSellClose,
		TradeTime:          "2026-05-09T11:00",
		Price:              11,
		Quantity:           2,
	})
	if !errors.Is(err, ErrOptionTrackInvalidHistory) {
		t.Fatalf("expected invalid history, got %v", err)
	}
}

func TestOptionTrackCloseAndPreviewPnL(t *testing.T) {
	service := newTestOptionTrackService(t)

	detail, err := service.CreatePlan(OptionTrackPlanCreateRequest{
		Name: "pnl test",
		Positions: []OptionTrackTradeUpsertInput{{
			ContractName:       "MO2406-P-4800",
			ExpiryDate:         "2026-06-28",
			StrikePrice:        4800,
			OptionSide:         "PUT",
			ContractMultiplier: 100,
			TradeType:          OptionTrackTradeTypeBuyOpen,
			TradeTime:          "2026-05-09T09:30",
			Price:              8,
			Quantity:           2,
		}},
	})
	if err != nil {
		t.Fatalf("create plan: %v", err)
	}

	preview, err := service.PreviewPlanPnL(detail.Plan.ID, OptionTrackPnLPreviewRequest{
		Prices: []OptionTrackContractPriceInput{{
			ContractKey:  detail.Positions[0].ContractKey,
			CurrentPrice: 9.5,
		}},
	})
	if err != nil {
		t.Fatalf("preview pnl: %v", err)
	}
	if preview.UnrealizedPnL != 300 {
		t.Fatalf("expected unrealized pnl 300, got %v", preview.UnrealizedPnL)
	}

	_, err = service.CreateTrade(detail.Plan.ID, OptionTrackTradeUpsertInput{
		ContractName:       "MO2406-P-4800",
		ExpiryDate:         "2026-06-28",
		StrikePrice:        4800,
		OptionSide:         "PUT",
		ContractMultiplier: 100,
		TradeType:          OptionTrackTradeTypeSellClose,
		TradeTime:          time.Date(2026, 5, 9, 14, 0, 0, 0, time.Local).Format(time.RFC3339),
		Price:              9,
		Quantity:           2,
	})
	if err != nil {
		t.Fatalf("close position trade: %v", err)
	}

	summary, err := service.ClosePlan(detail.Plan.ID)
	if err != nil {
		t.Fatalf("close plan: %v", err)
	}
	if summary.Status != OptionTrackStatusClosed {
		t.Fatalf("expected closed status, got %s", summary.Status)
	}

	firstTradeID := detail.Trades[0].ID
	_, err = service.UpdateTrade(detail.Plan.ID, firstTradeID, OptionTrackTradeUpsertInput{
		ContractName:       "MO2406-P-4800",
		ExpiryDate:         "2026-06-28",
		StrikePrice:        4800,
		OptionSide:         "PUT",
		ContractMultiplier: 100,
		TradeType:          OptionTrackTradeTypeBuyOpen,
		TradeTime:          "2026-05-09T09:30",
		Price:              8,
		Quantity:           3,
	})
	if !errors.Is(err, ErrOptionTrackClosedEditBreak) {
		t.Fatalf("expected closed edit break, got %v", err)
	}
}

func TestOptionTrackSupportsMixedOptionAndFuturePlan(t *testing.T) {
	service := newTestOptionTrackService(t)

	detail, err := service.CreatePlan(OptionTrackPlanCreateRequest{
		Name: "delta neutral test",
		Positions: []OptionTrackTradeUpsertInput{
			{
				InstrumentType:     OptionTrackInstrumentTypeOption,
				ContractName:       "MO2406-C-5000",
				ContractCode:       "MO2406C5000",
				ExpiryDate:         "2026-06-28",
				StrikePrice:        5000,
				OptionSide:         "CALL",
				ContractMultiplier: 100,
				TradeType:          OptionTrackTradeTypeBuyOpen,
				TradeTime:          "2026-05-10T10:00",
				Price:              12,
				Quantity:           2,
			},
			{
				InstrumentType:     OptionTrackInstrumentTypeFuture,
				ContractName:       "IM2506",
				ContractCode:       "IM2506",
				ExpiryDate:         "2026-06-19",
				ContractMultiplier: 200,
				TradeType:          OptionTrackTradeTypeSellOpen,
				TradeTime:          "2026-05-10T10:01",
				Price:              5230.5,
				Quantity:           1,
			},
		},
	})
	if err != nil {
		t.Fatalf("create mixed plan: %v", err)
	}

	if len(detail.Trades) != 2 {
		t.Fatalf("expected 2 trades, got %d", len(detail.Trades))
	}
	if detail.Trades[1].InstrumentType != OptionTrackInstrumentTypeFuture {
		t.Fatalf("expected future instrument type, got %s", detail.Trades[1].InstrumentType)
	}
	if detail.Trades[1].OptionSide != "" {
		t.Fatalf("expected future option side empty, got %s", detail.Trades[1].OptionSide)
	}
	if detail.Trades[1].StrikePrice != 0 {
		t.Fatalf("expected future strike 0, got %v", detail.Trades[1].StrikePrice)
	}
	if detail.Trades[1].Amount != 1046100 {
		t.Fatalf("expected future amount 1046100, got %v", detail.Trades[1].Amount)
	}

	prices := make([]OptionTrackContractPriceInput, 0, len(detail.Positions))
	for _, position := range detail.Positions {
		currentPrice := 13.0
		if position.InstrumentType == OptionTrackInstrumentTypeFuture {
			currentPrice = 5200.5
		}
		prices = append(prices, OptionTrackContractPriceInput{
			ContractKey:  position.ContractKey,
			CurrentPrice: currentPrice,
		})
	}

	preview, err := service.PreviewPlanPnL(detail.Plan.ID, OptionTrackPnLPreviewRequest{
		Prices: prices,
	})
	if err != nil {
		t.Fatalf("preview mixed pnl: %v", err)
	}
	if len(preview.Contracts) != 2 {
		t.Fatalf("expected 2 pnl contracts, got %d", len(preview.Contracts))
	}
	if preview.TotalPnL != 6200 {
		t.Fatalf("expected total pnl 6200, got %v", preview.TotalPnL)
	}
}
