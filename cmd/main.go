package main

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhuyanxi/CarnoFinance/pkg/db"
	"github.com/zhuyanxi/CarnoFinance/pkg/domain"
	"github.com/zhuyanxi/CarnoFinance/pkg/helper"
	"github.com/zhuyanxi/CarnoFinance/pkg/xueqiu"
)

// flow:
// init history data: get all stock index and data

func main() {
	a := time.Unix(1709222400, 0).Format(helper.DateFormat)
	logrus.Infoln(a)

	envs := os.Environ()
	logrus.Infof("%+v", envs)

	ctx := context.Background()
	sqlite := db.NewSqlite(ctx)
	xqc := xueqiu.New()
	app := domain.NewDomain(ctx, sqlite, xqc)
	app.Init()

	r := gin.Default()
	r.Static("/front", "../public")
	r.GET("/healthz", func(ctx *gin.Context) {
		helper.GinOK(ctx)
	})

	//SZ159915
	r.POST("/etf/last", func(ctx *gin.Context) {

		// var price domain.ETFDailyPrice
		// if err := ctx.ShouldBindJSON(&price); err != nil {
		// 	logrus.Errorf("bind json error: %v", err)
		// 	return
		// }
		// if price.TradeDate == "" {
		// 	price.TradeDate = time.Now().Format("20060102")
		// }
		// app.InsertETFDailyPrice(price)

		code := ctx.Query("code")
		count, err := helper.ExtractCount(ctx)
		if err != nil {
			logrus.Errorf("extract count error: %v", err)
			return
		}
		app.SetETFPrice(code, count)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
