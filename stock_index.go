package main

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	CSI_300_GROWTH_INDEX_URL = "http://www.csindex.com.cn/uploads/file/autofile/cons/000918cons.xls?t=1612515850"
	CSI_300_VALUE_INDEX_URL  = "http://www.csindex.com.cn/uploads/file/autofile/cons/000919cons.xls?t=1612515836"

	CSI_300_GROWTH_INDEX_EXCEL = "csi_300_growth_index.xls"
	CSI_300_VALUE_INDEX_EXCEL  = "csi_300_value_index.xls"
)

// StockIndex :
type StockIndex struct {
	Code    string
	Date    time.Time
	Name    string
	EngName string

	Stocks []Stock
}

type StockIndexScraper struct {
	URL       string
	ExcelPath string
}

func NewStockIndexScraper(url, excelPath string) *StockIndexScraper {
	return &StockIndexScraper{url, excelPath}
}

func (s *StockIndexScraper) Download() error {
	err := DownloadExcel(s.URL, s.ExcelPath)
	return err
}

func (s *StockIndexScraper) Read() (StockIndex, error) {
	f, err := excelize.OpenFile(s.ExcelPath)
	if err != nil {
		fmt.Println(err)
		return StockIndex{}, err
	}

	var stockIndex StockIndex

	sheetName := f.GetSheetName(0)
	rows := f.GetRows(sheetName)
	for i := range rows {
		for j := range rows[i] {
			fmt.Print(rows[i][j], "\t")
		}
		fmt.Println()
	}

	return stockIndex, nil
}
