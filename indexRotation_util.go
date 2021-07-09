package main

import "fmt"

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
