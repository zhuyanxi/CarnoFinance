package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhuyanxi/CarnoFinance/pkg/domain"
	"github.com/zhuyanxi/CarnoFinance/pkg/helper"
)

func SetETFPrice(app *domain.Domain) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		code := ctx.Query("code")
		count, err := helper.ExtractCount(ctx)
		if err != nil {
			logrus.Errorf("extract count error: %v", err)
			return
		}
		if err := app.SetETFPrice(code, count); err != nil {
			logrus.Errorf("set etf price error: %v", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"code": code, "count": count})
	}
}

func GetRSRSList(app *domain.Domain) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		period, err := helper.ExtractPeriod(ctx)
		if err != nil {
			logrus.Errorf("set period error: %v", err)
			return
		}
		list, err := app.GetRSRSList(period, ctx.DefaultQuery("order", "desc"))
		if err != nil {
			logrus.Errorf("get rsrs list error: %v", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, list)
	}
}
