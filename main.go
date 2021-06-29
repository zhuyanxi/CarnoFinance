package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
	"gonum.org/v1/plot/plotter"
)

func main() {
	points := make(map[string]plotter.XYs)
	points["First"] = PlotRandomPoints(15)
	PlotLine("test", "X", "Y", "random.png", points)

	indexes := make(map[string]plotter.XYs)
	indexes["HS300"] = PlotPoints(HS300ClosePoints())
	PlotLine("indexes", "date", "point", "index-trend.png", indexes)

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
