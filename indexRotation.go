package main

import (
	"encoding/csv"
	"io/ioutil"
	"strings"
)

type IndexData struct {
	Date   string
	Open   string
	High   string
	Low    string
	Close  string
	Change string //  e.g.: 0.1%
	Volume string
}

func GetIndexData(file string) []IndexData {
	ret := []IndexData{}

	f, err := ioutil.ReadFile(file)
	if err != nil {
		return ret
	}

	hs300Data, err := csv.NewReader(strings.NewReader(string(f))).ReadAll()
	if err != nil {
		return ret
	}

	curRow := 1
	for _, line := range hs300Data {
		if curRow == 1 {
			curRow++
			continue
		}
		var d IndexData
		d.Date = line[0]
		d.Close = line[1]
		d.Open = line[2]
		d.High = line[3]
		d.Low = line[4]
		d.Volume = line[5]
		d.Change = line[6]

		ret = append(ret, d)
	}
	return ret
}
