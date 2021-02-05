package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadExcel(t *testing.T) {
	path, _ := os.Getwd()
	filepath := filepath.Join(path, CSI_300_GROWTH_INDEX_EXCEL)
	DownloadExcel(CSI_300_GROWTH_INDEX_URL, filepath)
}
