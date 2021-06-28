package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
)

type IndexData struct {
	Date       string
	Open       string
	High       string
	Low        string
	Close      string
	Increase   string
	Change     string //  e.g.: 0.1%
	Volume     string
	VolumeYang string
}

func GetHS300() []IndexData {
	ret := []IndexData{}

	f, err := ioutil.ReadFile("HS300.csv")
	if err != nil {
		return ret
	}

	hs300Data, err := csv.NewReader(strings.NewReader(string(f))).ReadAll()
	if err != nil {
		return ret
	}
	fmt.Println("HS300 rows:", len(hs300Data))

	curRow := 1
	for _, line := range hs300Data {
		if curRow == 1 {
			curRow++
			continue
		}
		var d IndexData
		d.Date = line[0]
		d.Open = line[6]
		d.High = line[4]
		d.Low = line[5]
		d.Close = line[3]
		d.Increase = line[8]
		d.Change = line[9]
		d.Volume = line[10]
		d.VolumeYang = line[11]

		ret = append(ret, d)
	}
	return ret
}
