package main

import (
	"fmt"
	"time"
)

type Share struct {
	HoldType  int     // 持有类型0, 1 or 2(其中0表示空仓)
	HoldHand  int     // 持有手数
	HoldMoney float64 //持有金额
	Cash      float64 // 现金

	// CurProfitRatio float64 //当前收益百分比
	// MaxHoldMoney   float64 //最大持仓市值
}

func (s *Share) Buy(holdType int, price float64) {

}

var ShowLog = true

func LogPrintln(log ...interface{}) {
	if ShowLog {
		fmt.Println(log...)
	}
}

// 计算最大回撤
func CalcMaxRe(money []float64) {
	var top float64
	var maxRet float64
	top = money[0]
	maxRet = 0
	for i := 1; i < len(money); i++ {
		if money[i] > top {
			top = money[i]
		} else {
			tmpRet := (money[i] - top) / money[i]
			if tmpRet < maxRet {
				maxRet = tmpRet
			}
		}
	}
	fmt.Println("max ret:", maxRet)
}

func DoDayByDay(begin, end time.Time, calcFunc func(orderedIndex1, orderedIndex2 []IndexData, n int) *Share) {
	beginDateExcludeWeekend := []time.Time{}
	for {
		begin = begin.Add(24 * time.Hour)
		if begin.Add(40 * 24 * time.Hour).After(end) {
			break
		}
		exludeTime1, err := time.Parse("2006-01-02", "2015-03-31")
		ExitIfErr(err)
		exludeTime2, err := time.Parse("2006-01-02", "2015-07-31")
		ExitIfErr(err)
		if begin.After(exludeTime1) && begin.Before(exludeTime2) {
			fmt.Println(begin)
			continue
		}
		if begin.Weekday() == time.Saturday || begin.Weekday() == time.Sunday {
			continue
		}
		beginDateExcludeWeekend = append(beginDateExcludeWeekend, begin)
	}

	totalDays := len(beginDateExcludeWeekend)
	type profitDay struct {
		DD     time.Time
		profit float64
	}
	lossArr := []profitDay{}
	profitArr := []profitDay{}
	for _, v := range beginDateExcludeWeekend {
		fmt.Println(v)

		// profit := doHS300AndCYB100(v, end)
		endV := v.Add(2 * 365 * 24 * time.Hour)
		if endV.After(end) {
			endV = end
		}
		profit := doHS300AndCYB100(v, endV, calcFunc)
		profitArr = append(profitArr, profitDay{
			DD:     v,
			profit: profit,
		})
	}

	maxProfit := profitArr[0]
	minProfit := profitArr[0]
	for _, v := range profitArr {
		if v.profit < 0 {
			lossArr = append(lossArr, v)
		}
		if v.profit > maxProfit.profit {
			maxProfit = v
		}
		if v.profit < minProfit.profit {
			minProfit = v
		}
	}

	fmt.Println("max profit: ", maxProfit)
	fmt.Println("min profit: ", minProfit)

	for _, v := range lossArr {
		fmt.Println("loss day: ", v.DD, " | loss profit: ", v.profit)
	}

	fmt.Println("total days:", totalDays)
	fmt.Println("profit day", len(profitArr))
	fmt.Println("loss days", len(lossArr))
}
