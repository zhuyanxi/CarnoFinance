package main

import "time"

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

	ConstituentCode string
	ConstituentName string
}
