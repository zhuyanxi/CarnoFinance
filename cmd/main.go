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
)

// flow:
// init history data: get all stock index and data

func main() {
	envs := os.Environ()
	logrus.Infof("%+v", envs)

	ctx := context.Background()
	sqlite := db.NewSqlite(ctx)
	app := domain.NewDomain(ctx, sqlite)
	app.Init()

	r := gin.Default()
	r.Static("/front", "../public")
	r.GET("/healthz", func(ctx *gin.Context) {
		helper.GinOK(ctx)
	})
	r.POST("/etf", func(ctx *gin.Context) {
		var price domain.ETFDailyPrice
		if err := ctx.ShouldBindJSON(&price); err != nil {
			logrus.Errorf("bind json error: %v", err)
			return
		}
		if price.TradeDate == "" {
			price.TradeDate = time.Now().Format("20060102")
		}
		app.InsertETFDailyPrice(price)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
