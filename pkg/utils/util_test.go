package utils

import (
	"testing"

	"github.com/zhuyanxi/CarnoFinance/pkg/domain"
)

func Test_setStructFieldValues(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		dr := &domain.StockDailyBeforePrice{}
		err := SetStructFieldValuesByJsonTags(dr, []string{"ts_code", "trade_date", "open", "high", "low", "close", "pre_close", "change", "pct_chg", "vol", "amount"}, []any{"000001.SZ", "20231117", 10.22, 10.25, 10.12, 10.15, 10.24, -0.09, -0.8789, 903362.17, 917616.681})
		if err != nil {
			t.Errorf("err = %v", err)
		}
		if err != nil {
			t.Errorf("err = %v", err)
		}
		if dr.TSCode != "000001.SZ" {
			t.Errorf("dr.TSCode = %v", dr.TSCode)
		}
	})
}
