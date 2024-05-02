package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhuyanxi/CarnoFinance/pkg/domain"
	"github.com/zhuyanxi/CarnoFinance/pkg/helper"
)

func GetETFClosePrices(app *domain.Domain) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		count, err := helper.ExtractCount(ctx)
		if err != nil {
			logrus.Errorf("extract count error: %v", err)
			return
		}
		dates, list, err := app.GetETFClosePrices(ctx.Param("code"), count)
		if err != nil {
			logrus.Errorf("get etf price list error: %v", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{
			"count":  len(list),
			"dates":  dates,
			"prices": list,
		})
	}
}

func SetETFPrice(app *domain.Domain) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		code := ctx.Query("code")
		if code == "all" {
			codes, err := app.GetETFCodeList()
			if err != nil {
				logrus.Errorf("get etf code list error: %v", err)
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
			for _, code := range codes {
				app.SetETFPrice(code.TSCode, -1)
			}
		} else {
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
		}
		helper.GinOK(ctx)
	}
}

func GetRSRSList(app *domain.Domain) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		period, err := helper.ExtractPeriod(ctx)
		if err != nil {
			logrus.Errorf("set period error: %v", err)
			return
		}
		list, err := app.GetETFRSRSList(period, ctx.DefaultQuery("order", "desc"), ctx.DefaultQuery("date", time.Now().Format("20060102")))
		if err != nil {
			logrus.Errorf("get rsrs list error: %v", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, list)
	}
}
