package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const IndexDateLayout = "2006年1月2日"
const OriginMoney = 10000000

type IndexData struct {
	Date       time.Time
	DateString string
	Open       float64
	High       float64
	Low        float64
	Close      float64
	Change     string //  e.g.: 0.1%
	Volume     string
}

func GetIndexData(file string) []IndexData {
	ret := []IndexData{}

	f, err := ioutil.ReadFile(file)
	if err != nil {
		return ret
	}

	reader := csv.NewReader(strings.NewReader(string(f)))
	reader.LazyQuotes = true
	idxData, err := reader.ReadAll()

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
		d.DateString = d.Date.Format("2006年01月02日")
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

func GetZZ500IndexData() []IndexData {
	return GetIndexData("中证500指数历史数据.csv")
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

// get the data between begin and end,
// begin should after 2010-06-02, end should before 2021-06-28
func InitData(begin, end time.Time, org1, org2 []IndexData) (orderedIdx1, orderedIdx2 []IndexData) {
	// if end.Sub(begin).Hours()/24 < 40 {
	// 	fmt.Println("end time should bigger than begin time plus 40days")
	// 	os.Exit(1)
	// }
	minT, err := time.Parse("2006-01-02", "2010-06-02")
	ExitIfErr(err)
	maxT, err := time.Parse("2006-01-02", "2021-06-28")
	ExitIfErr(err)
	if begin.Before(minT) || end.After(maxT) {
		fmt.Println("begin should after 2010-06-02, end should before 2021-06-28")
		os.Exit(1)
	}

	orderedIdx1 = []IndexData{}
	orderedIdx2 = []IndexData{}

	for _, val := range org1 {
		// fmt.Println("begin:", begin, " | end:", end, " | curDay:", val.Date)
		if val.Date.After(begin) && val.Date.Before(end) {
			orderedIdx1 = append(orderedIdx1, val)
		}
	}
	sort.Slice(orderedIdx1, func(i, j int) bool {
		return orderedIdx1[i].Date.Before(orderedIdx1[j].Date)
	})

	for _, val := range org2 {
		if val.Date.After(begin) && val.Date.Before(end) {
			orderedIdx2 = append(orderedIdx2, val)
		}
	}
	sort.Slice(orderedIdx2, func(i, j int) bool {
		return orderedIdx2[i].Date.Before(orderedIdx2[j].Date)
	})

	// for _,val1:=range orderedHS300{
	// 	for _,val2:=range orderedCYB100{
	// 		if val1.Date.Equal()
	// 	}
	// }
	len1 := len(orderedIdx1)
	len2 := len(orderedIdx2)
	if len1 != len2 {
		fmt.Printf("len1=%d not equal to len2=%d", len1, len2)
		os.Exit(1)
	}
	if len1 < 21 {
		fmt.Println("exchange days should be larger than 20")
		os.Exit(1)
	}

	return
}

// 增长百分比
func CalcInc(idx []IndexData, begin, end time.Time) float64 {
	retIdx := []IndexData{}
	for _, val := range idx {
		if val.Date.After(begin) && val.Date.Before(end) {
			retIdx = append(retIdx, val)
		}
	}
	sort.Slice(retIdx, func(i, j int) bool {
		return retIdx[i].Date.Before(retIdx[j].Date)
	})
	return 100 * (retIdx[len(retIdx)-1].Close - retIdx[0].Close) / retIdx[0].Close
}

// 计算单个指数收益
func CalcProfit(idx []IndexData, begin, end time.Time) float64 {
	retIdx := []IndexData{}
	for _, val := range idx {
		if val.Date.After(begin) && val.Date.Before(end) {
			retIdx = append(retIdx, val)
		}
	}
	sort.Slice(retIdx, func(i, j int) bool {
		return retIdx[i].Date.Before(retIdx[j].Date)
	})

	holdF := OriginMoney / retIdx[0].Open
	holdHand := int(math.Floor(holdF))

	CurMoney := float64(holdHand) * retIdx[len(retIdx)-1].Close
	profit := 100 * (CurMoney - OriginMoney) / OriginMoney
	return profit
}

// 计算前20日(交易日)两个指数谁的涨幅大，判断当天是否前20日涨幅大的指数涨幅依然更大，返回是的天数和否的天数
func Calc1(orderedIndex1, orderedIndex2 []IndexData) (int, int) {
	if len(orderedIndex1) != len(orderedIndex2) {
		fmt.Println("length not the same")
		return 0, 0
	}

	same := 0
	notSame := 0
	length := len(orderedIndex1)
	for i := 20; i < length; i++ {
		inc20Idx1 := (orderedIndex1[i-1].Close - orderedIndex1[i-20].Close) / orderedIndex1[i-20].Close
		inc20Idx2 := (orderedIndex2[i-1].Close - orderedIndex2[i-20].Close) / orderedIndex2[i-20].Close

		incCurIdx1 := (orderedIndex1[i].Close - orderedIndex1[i-1].Close) / orderedIndex1[i-1].Close
		incCurIdx2 := (orderedIndex2[i].Close - orderedIndex2[i-1].Close) / orderedIndex2[i-1].Close

		if inc20Idx1 >= inc20Idx2 && incCurIdx1 >= incCurIdx2 {
			same++
		} else if inc20Idx1 < inc20Idx2 && incCurIdx1 < incCurIdx2 {
			same++
		} else if inc20Idx1 >= inc20Idx2 && incCurIdx1 < incCurIdx2 {
			notSame++
		} else if inc20Idx1 < inc20Idx2 && incCurIdx1 >= incCurIdx2 {
			notSame++
		}

		fmt.Printf("percentage recent 20 days(index1[%f],index2[%f]) | percentage curDay(index1[%f],index2[%f])\n", inc20Idx1*100, inc20Idx2*100, incCurIdx1*100, incCurIdx2*100)
		// time.Sleep(500 * time.Millisecond)
	}
	return same, notSame
}

type Share struct {
	HoldType  int     // 持有类型0, 1 or 2(其中0表示空仓)
	HoldHand  int     // 持有手数
	HoldMoney float64 //持有金额
	Cash      float64 // 现金

	MaxHoldMoney float64 //最大持仓市值
	MaxMoney     float64 //最大金额
	MaxRe        float64 //最大回撤
}

// 计算前n(20,30,60等等)日(交易日)两个指数,谁的涨幅大，第一次当天以开盘价买入谁，之后以开盘价保持或换仓涨幅大的
func Calc2(orderedIndex1, orderedIndex2 []IndexData, n int) *Share {
	if len(orderedIndex1) != len(orderedIndex2) {
		fmt.Println("length not the same")
		os.Exit(1)
	}
	length := len(orderedIndex1)
	if length <= n {
		fmt.Println("length should bigger than n")
		os.Exit(1)
	}

	holdShare := new(Share)
	holdShare.HoldType = 0
	holdShare.HoldHand = 0
	holdShare.HoldMoney = 0
	holdShare.Cash = OriginMoney
	holdShare.MaxHoldMoney = 0
	holdShare.MaxMoney = OriginMoney
	holdShare.MaxRe = 0

	changeCount := 0

	for i := n; i < length; i++ {
		inc20Idx1 := (orderedIndex1[i-1].Close - orderedIndex1[i-n].Close) / orderedIndex1[i-n].Close
		inc20Idx2 := (orderedIndex2[i-1].Close - orderedIndex2[i-n].Close) / orderedIndex2[i-n].Close

		if inc20Idx1 >= inc20Idx2 {
			if holdShare.HoldType == 0 {
				buyPrice := orderedIndex1[i].Open
				holdF := holdShare.Cash / buyPrice
				holdHand := int(math.Floor(holdF))
				holdShare.HoldType = 1
				holdShare.HoldHand = holdHand
				holdShare.HoldMoney = float64(holdHand) * buyPrice
				holdShare.Cash = holdShare.Cash - holdShare.HoldMoney
				fmt.Println("Buy hold index1 date:", orderedIndex1[i].Date)
			}
			if holdShare.HoldType == 2 {
				sellPrice := orderedIndex2[i-1].Close
				sellMoney := sellPrice * float64(holdShare.HoldHand)
				totalMoney := sellMoney + holdShare.Cash

				buyPrice := orderedIndex1[i].Open
				holdF := totalMoney / buyPrice
				holdHand := int(math.Floor(holdF))
				holdShare.HoldType = 1
				holdShare.HoldHand = holdHand
				holdShare.HoldMoney = float64(holdHand) * buyPrice
				holdShare.Cash = totalMoney - holdShare.HoldMoney

				fmt.Println("Change hold from index2 to index1 date:", orderedIndex1[i].Date)
				changeCount++
			}
		}
		if inc20Idx1 < inc20Idx2 {
			if holdShare.HoldType == 0 {
				buyPrice := orderedIndex2[i].Open
				holdF := holdShare.Cash / buyPrice
				holdHand := int(math.Floor(holdF))
				holdShare.HoldType = 2
				holdShare.HoldHand = holdHand
				holdShare.HoldMoney = float64(holdHand) * buyPrice
				holdShare.Cash = holdShare.Cash - holdShare.HoldMoney
				fmt.Println("Buy hold index2 date:", orderedIndex2[i].Date)
			}
			if holdShare.HoldType == 1 {
				sellPrice := orderedIndex1[i-1].Close
				sellMoney := sellPrice * float64(holdShare.HoldHand)
				totalMoney := sellMoney + holdShare.Cash

				buyPrice := orderedIndex2[i].Open
				holdF := totalMoney / buyPrice
				holdHand := int(math.Floor(holdF))
				holdShare.HoldType = 2
				holdShare.HoldHand = holdHand
				holdShare.HoldMoney = float64(holdHand) * buyPrice
				holdShare.Cash = totalMoney - holdShare.HoldMoney

				fmt.Println("Change hold from index1 to index2 date:", orderedIndex1[i].Date)
				changeCount++
			}
		}
	}
	fmt.Println("change count:", changeCount)

	var closePrice float64
	if holdShare.HoldType == 1 {
		closePrice = orderedIndex1[length-1].Close
	} else {
		closePrice = orderedIndex2[length-1].Close
	}
	holdShare.HoldMoney = closePrice * float64(holdShare.HoldHand)

	return holdShare
}

// 计算前n(20,30,60等等)日(交易日)两个指数,谁的涨幅大，第一次当天以开盘价买入谁，之后以开盘价保持或换仓涨幅大的
// 若两个涨幅均小于0，则空仓
func Calc3(orderedIndex1, orderedIndex2 []IndexData, n int) *Share {
	if len(orderedIndex1) != len(orderedIndex2) {
		fmt.Println("length not the same")
		os.Exit(1)
	}
	length := len(orderedIndex1)
	if length <= n {
		fmt.Println("length should bigger than n")
		os.Exit(1)
	}

	holdShare := new(Share)
	holdShare.HoldType = 0
	holdShare.HoldHand = 0
	holdShare.HoldMoney = 0
	holdShare.Cash = OriginMoney
	holdShare.MaxHoldMoney = 0
	holdShare.MaxMoney = OriginMoney
	holdShare.MaxRe = 0

	changeCount := 0

	for i := n; i < length; i++ {
		inc20Idx1 := (orderedIndex1[i-1].Close - orderedIndex1[i-n].Close) / orderedIndex1[i-n].Close
		inc20Idx2 := (orderedIndex2[i-1].Close - orderedIndex2[i-n].Close) / orderedIndex2[i-n].Close

		if inc20Idx1 < 0 && inc20Idx2 < 0 {
			if holdShare.HoldType == 0 {
				continue
			}

			var sellPrice float64
			if holdShare.HoldType == 1 {
				fmt.Println("Sell hold index1 date:", orderedIndex1[i].Date)
				sellPrice = orderedIndex1[i-1].Close
			}

			if holdShare.HoldType == 2 {
				fmt.Println("Sell hold index2 date:", orderedIndex1[i].Date)
				sellPrice = orderedIndex2[i-1].Close
			}

			sellMoney := sellPrice * float64(holdShare.HoldHand)
			totalMoney := sellMoney + holdShare.Cash

			holdShare.HoldType = 0
			holdShare.HoldHand = 0
			holdShare.HoldMoney = 0
			holdShare.Cash = totalMoney

			continue
		}

		if inc20Idx1 >= inc20Idx2 {
			if holdShare.HoldType == 0 {
				buyPrice := orderedIndex1[i].Open
				holdF := holdShare.Cash / buyPrice
				holdHand := int(math.Floor(holdF))
				holdShare.HoldType = 1
				holdShare.HoldHand = holdHand
				holdShare.HoldMoney = float64(holdHand) * buyPrice
				holdShare.Cash = holdShare.Cash - holdShare.HoldMoney
				fmt.Println("Buy hold index1 date:", orderedIndex1[i].Date)
			}
			if holdShare.HoldType == 2 {
				sellPrice := orderedIndex2[i-1].Close
				sellMoney := sellPrice * float64(holdShare.HoldHand)
				totalMoney := sellMoney + holdShare.Cash

				buyPrice := orderedIndex1[i].Open
				holdF := totalMoney / buyPrice
				holdHand := int(math.Floor(holdF))
				holdShare.HoldType = 1
				holdShare.HoldHand = holdHand
				holdShare.HoldMoney = float64(holdHand) * buyPrice
				holdShare.Cash = totalMoney - holdShare.HoldMoney

				fmt.Println("Change hold from index2 to index1 date:", orderedIndex1[i].Date)
				changeCount++
			}
		}
		if inc20Idx1 < inc20Idx2 {
			if holdShare.HoldType == 0 {
				buyPrice := orderedIndex2[i].Open
				holdF := holdShare.Cash / buyPrice
				holdHand := int(math.Floor(holdF))
				holdShare.HoldType = 2
				holdShare.HoldHand = holdHand
				holdShare.HoldMoney = float64(holdHand) * buyPrice
				holdShare.Cash = holdShare.Cash - holdShare.HoldMoney
				fmt.Println("Buy hold index2 date:", orderedIndex2[i].Date)
			}
			if holdShare.HoldType == 1 {
				sellPrice := orderedIndex1[i-1].Close
				sellMoney := sellPrice * float64(holdShare.HoldHand)
				totalMoney := sellMoney + holdShare.Cash

				buyPrice := orderedIndex2[i].Open
				holdF := totalMoney / buyPrice
				holdHand := int(math.Floor(holdF))
				holdShare.HoldType = 2
				holdShare.HoldHand = holdHand
				holdShare.HoldMoney = float64(holdHand) * buyPrice
				holdShare.Cash = totalMoney - holdShare.HoldMoney

				fmt.Println("Change hold from index1 to index2 date:", orderedIndex1[i].Date)
				changeCount++
			}
		}
	}
	fmt.Println("change count:", changeCount)

	if holdShare.HoldType != 0 {
		var closePrice float64
		if holdShare.HoldType == 1 {
			closePrice = orderedIndex1[length-1].Close
		} else {
			closePrice = orderedIndex2[length-1].Close
		}
		holdShare.HoldMoney = closePrice * float64(holdShare.HoldHand)
	}

	return holdShare
}
