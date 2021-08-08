package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	ETFPrefix  = "https://fundf10.eastmoney.com/jjjz_"
	ETF_513050 = ETFPrefix + "513050.html" //中概互联
	ETF_515000 = ETFPrefix + "515000.html" //科技
	ETF_510230 = ETFPrefix + "510230.html" //金融
	ETF_512690 = ETFPrefix + "512690.html" //酒
	ETF_512480 = ETFPrefix + "512480.html" //半导体
	ETF_516160 = ETFPrefix + "516160.html" //新能源
	ETF_515120 = ETFPrefix + "515120.html" //创新药
	ETF_512170 = ETFPrefix + "512170.html" //医疗
	ETF_516780 = ETFPrefix + "516780.html" //稀土
	ETF_515050 = ETFPrefix + "515050.html" //5G
)

func doScrapeETF(url string) []IndexData {
	idx := make([]IndexData, 20)

	// Make request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find("#jztable").Find("table").Each(func(i int, s *goquery.Selection) {
		aa := s.Text()
		fmt.Println(aa)
	})

	doc.Find("#jztable").Find("table").Find("tbody").Find("tr").Each(func(i int, si *goquery.Selection) {
		// dd := IndexData{}
		si.Find("td").Each(func(j int, sj *goquery.Selection) {
			span := sj.Text()
			fmt.Println(span)
		})
	})

	return idx
}
