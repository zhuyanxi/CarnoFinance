package main

import (
	"encoding/csv"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

const IndexDateLayout = "2006年1月2日"

type IndexData struct {
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Change string //  e.g.: 0.1%
	Volume string
}

func GetIndexData(file string) []IndexData {
	ret := []IndexData{}

	f, err := ioutil.ReadFile(file)
	if err != nil {
		return ret
	}

	idxData, err := csv.NewReader(strings.NewReader(string(f))).ReadAll()
	if err != nil {
		return ret
	}

	curRow := 1
	for _, line := range idxData {
		if curRow == 1 {
			curRow++
			continue
		}
		var d IndexData
		d.Date, _ = time.Parse("2006年1月2日", line[0])
		// fmt.Println(d.Date)
		d.Close = StringToFloat64(line[1])
		d.Open = StringToFloat64(line[2])
		d.High = StringToFloat64(line[3])
		d.Low = StringToFloat64(line[4])
		d.Volume = line[5]
		d.Change = line[6]

		ret = append(ret, d)
	}
	return ret
}

func GetHS300IndexData() []IndexData {
	return GetIndexData("沪深300指数历史数据.csv")
}

func GetCYB100IndexData() []IndexData {
	return GetIndexData("创业板指数历史数据.csv")
}

func HS300ClosePoints() ([]float64, []float64) {

	hs300Data := GetHS300IndexData()
	sort.Slice(hs300Data, func(i, j int) bool {
		return hs300Data[i].Date.Before(hs300Data[j].Date)
	})

	dataX := make([]float64, len(hs300Data))
	dataY := make([]float64, len(hs300Data))
	for i := 0; i < len(hs300Data); i++ {
		dataX[i], _ = strconv.ParseFloat(hs300Data[i].Date.Format("20060102"), 64)
		dataY[i] = hs300Data[i].Close
	}

	return dataX, dataY
}
