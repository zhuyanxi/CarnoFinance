package main

import (
	"strconv"
	"strings"
)

// price format is: 1,111.11 or 111.11
func StringToFloat64(price string) float64 {
	pp := strings.Split(price, ",")

	if len(pp) == 2 {

		t, err := strconv.ParseFloat(pp[0], 64)
		if err != nil {
			return 0
		}

		h, err := strconv.ParseFloat(pp[1], 64)
		if err != nil {
			return 0
		}

		ret := t*1000 + h
		return ret
	}

	if len(pp) == 1 {
		h, err := strconv.ParseFloat(pp[0], 64)
		if err != nil {
			return 0
		}

		return h
	}
	return 0
}
