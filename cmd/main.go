package main

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhuyanxi/CarnoFinance/pkg/api"
	"github.com/zhuyanxi/CarnoFinance/pkg/db"
	"github.com/zhuyanxi/CarnoFinance/pkg/domain"
	"github.com/zhuyanxi/CarnoFinance/pkg/helper"
	"github.com/zhuyanxi/CarnoFinance/pkg/xueqiu"
)

// flow:
// init history data: get all stock index and data

func main() {
	a := time.Unix(1709481600, 0).UTC().Format(helper.DateFormat)
	logrus.Infoln(a)
	b := time.Unix(1709481600, 0).Local().Format(helper.DateFormat)
	logrus.Infoln(b)

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
	r.POST("/etf/last", api.SetETFPrice(app))
	r.GET("/rsrslist", api.GetRSRSList(app))

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
