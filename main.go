package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//begin should after 2010-06-02, end should before 2021-06-28
	begin, err := time.Parse("2006-01-02", "2015-06-05")
	ExitIfErr(err)
	end, err := time.Parse("2006-01-02", "2021-06-25")
	ExitIfErr(err)
	// doHS300AndCYB100(begin, end)
	// doHS300AndZZ500(begin, end)
	doSZ50AndCYB100(begin, end)
}

func doHS300AndCYB100(begin, end time.Time) {
	hs300Inc := CalcInc(GetHS300IndexData(), begin, end)
	cyb100Inc := CalcInc(GetCYB100IndexData(), begin, end)
	fmt.Println("hs300 inc: ", hs300Inc)
	fmt.Println("cyb inc: ", cyb100Inc)

	hs300Inc = CalcProfit(GetHS300IndexData(), begin, end)
	cyb100Inc = CalcProfit(GetCYB100IndexData(), begin, end)
	fmt.Println("hs300 profit: ", hs300Inc)
	fmt.Println("cyb profit: ", cyb100Inc)

	fmt.Println()
	fmt.Println("============================策略开始===============================================")
	fmt.Println()
	index1, index2 := InitData(begin, end, GetHS300IndexData(), GetCYB100IndexData())

	var holdShare *Share
	var totalMoney float64
	// holdShare = Calc2(index1, index2, 20)
	// fmt.Printf("%+v", holdShare)
	// totalMoney := holdShare.HoldMoney + holdShare.Cash
	// fmt.Printf("Total Money:%f\n", totalMoney)
	// fmt.Printf("Profit:%f\n", 100*(totalMoney-OriginMoney)/OriginMoney)

	fmt.Println()
	fmt.Println("============================策略的分界线===============================================")
	fmt.Println()

	holdShare = Calc4(index1, index2, 20)
	fmt.Printf("%+v", holdShare)
	totalMoney = holdShare.HoldMoney + holdShare.Cash
	fmt.Printf("Total Money:%f\n", totalMoney)
	fmt.Printf("Profit:%f\n", 100*(totalMoney-OriginMoney)/OriginMoney)

}

func doSZ50AndCYB100(begin, end time.Time) {
	sz50Inc := CalcInc(GetSZ50IndexData(), begin, end)
	cyb100Inc := CalcInc(GetCYB100IndexData(), begin, end)
	fmt.Println("hs300 inc: ", sz50Inc)
	fmt.Println("cyb inc: ", cyb100Inc)

	sz50Inc = CalcProfit(GetSZ50IndexData(), begin, end)
	cyb100Inc = CalcProfit(GetCYB100IndexData(), begin, end)
	fmt.Println("hs300 profit: ", sz50Inc)
	fmt.Println("cyb profit: ", cyb100Inc)

	fmt.Println()
	fmt.Println("============================策略开始===============================================")
	fmt.Println()
	index1, index2 := InitData(begin, end, GetSZ50IndexData(), GetCYB100IndexData())

	var holdShare *Share
	var totalMoney float64
	// holdShare = Calc2(index1, index2, 20)
	// fmt.Printf("%+v", holdShare)
	// totalMoney := holdShare.HoldMoney + holdShare.Cash
	// fmt.Printf("Total Money:%f\n", totalMoney)
	// fmt.Printf("Profit:%f\n", 100*(totalMoney-OriginMoney)/OriginMoney)

	fmt.Println()
	fmt.Println("============================策略的分界线===============================================")
	fmt.Println()

	holdShare = Calc4(index1, index2, 20)
	fmt.Printf("%+v", holdShare)
	totalMoney = holdShare.HoldMoney + holdShare.Cash
	fmt.Printf("Total Money:%f\n", totalMoney)
	fmt.Printf("Profit:%f\n", 100*(totalMoney-OriginMoney)/OriginMoney)

}

func doHS300AndZZ500(begin, end time.Time) {
	hs300Inc := CalcInc(GetHS300IndexData(), begin, end)
	cyb100Inc := CalcInc(GetZZ500IndexData(), begin, end)
	fmt.Println("hs300 inc: ", hs300Inc)
	fmt.Println("zz500 inc: ", cyb100Inc)

	hs300Inc = CalcProfit(GetHS300IndexData(), begin, end)
	cyb100Inc = CalcProfit(GetZZ500IndexData(), begin, end)
	fmt.Println("hs300 profit: ", hs300Inc)
	fmt.Println("zz500 profit: ", cyb100Inc)

	fmt.Println()
	fmt.Println("============================策略开始===============================================")
	fmt.Println()
	index1, index2 := InitData(begin, end, GetHS300IndexData(), GetZZ500IndexData())

	holdShare := Calc2(index1, index2, 20)
	fmt.Printf("%+v", holdShare)
	totolMoney := holdShare.HoldMoney + holdShare.Cash
	fmt.Printf("Total Money:%f\n", totolMoney)
	fmt.Printf("Profit:%f\n", 100*(totolMoney-OriginMoney)/OriginMoney)

	fmt.Println()
	fmt.Println("============================策略的分界线===============================================")
	fmt.Println()

	holdShare = Calc3(index1, index2, 20)
	fmt.Printf("%+v", holdShare)
	totolMoney = holdShare.HoldMoney + holdShare.Cash
	fmt.Printf("Total Money:%f\n", totolMoney)
	fmt.Printf("Profit:%f\n", 100*(totolMoney-OriginMoney)/OriginMoney)

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
