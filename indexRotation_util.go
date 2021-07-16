package main

import (
	"fmt"
)

type Share struct {
	HoldType  int     // 持有类型0, 1 or 2(其中0表示空仓)
	HoldHand  int     // 持有手数
	HoldMoney float64 //持有金额
	Cash      float64 // 现金

	CurProfitRatio float64 //当前收益百分比
	MaxHoldMoney   float64 //最大持仓市值
	MaxMoney       float64 //最大金额
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
