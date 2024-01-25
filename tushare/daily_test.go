package tushare

import (
	"fmt"
	"testing"
)

func TestGetDailyPrices(t *testing.T) {
	tushare := NewTushare("3994a21e138271341df8b79335de2c5489d2ec6af490c868cbb72cf1")
	resp, err := tushare.GetDailyPrices(DailyRequest{
		TSCode:    "000001.SZ",
		StartDate: "20231121",
		EndDate:   "20231128",
	})
	if err != nil {
		fmt.Println("get daily price error: ", err.Error())
	}
	fmt.Printf("response: %+v\n", resp)
}
