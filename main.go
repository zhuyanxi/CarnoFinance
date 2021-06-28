package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	hs300 := GetIndexData("沪深300指数历史数据.csv")
	cyb100 := GetIndexData("创业板指数历史数据.csv")

	sort.Slice(hs300, func(i, j int) bool {
		return hs300[i].Date < hs300[j].Date
	})
	sort.Slice(cyb100, func(i, j int) bool {
		return cyb100[i].Date < cyb100[j].Date
	})

	// doGrowthIndex()
	// doValueIndex()
}

func doGrowthIndex() {
	path, _ := os.Getwd()
	filepath := filepath.Join(path, CSI_300_GROWTH_INDEX_EXCEL)
	growthScraper := NewStockIndexScraper(CSI_300_GROWTH_INDEX_URL, filepath)
	err := growthScraper.Download()
	if err != nil {
		panic(err)
	}
	growthScraper.Read()
}

func doValueIndex() {
	path, _ := os.Getwd()
	filepath := filepath.Join(path, CSI_300_VALUE_INDEX_EXCEL)
	growthScraper := NewStockIndexScraper(CSI_300_VALUE_INDEX_URL, filepath)
	err := growthScraper.Download()
	if err != nil {
		panic(err)
	}
	growthScraper.Read()
}

func doRequestXueqiu() {
	// Make request
	response, err := http.Get("https://xueqiu.com/S/SZ300910")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(".quote-info").Each(func(i int, si *goquery.Selection) {
		si.Find("tr").Find("td").Each(func(j int, sj *goquery.Selection) {
			span := sj.Find("span").Text()
			fmt.Println(span)
		})
	})
}
