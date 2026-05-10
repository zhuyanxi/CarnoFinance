package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhuyanxi/CarnoFinance/pkg/domain"
)

func CreateOptionTrackPlan(app *domain.OptionTrackService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.OptionTrackPlanCreateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		plan, err := app.CreatePlan(req)
		if err != nil {
			handleOptionTrackError(ctx, err)
			return
		}
		ctx.JSON(http.StatusCreated, plan)
	}
}

func ListOptionTrackPlans(app *domain.OptionTrackService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
		resp, err := app.ListPlans(page, pageSize)
		if err != nil {
			handleOptionTrackError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func GetOptionTrackPlanDetail(app *domain.OptionTrackService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planID, err := parseOptionTrackID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid plan id"})
			return
		}
		detail, err := app.GetPlanDetail(planID)
		if err != nil {
			handleOptionTrackError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, detail)
	}
}

func CreateOptionTrackTrade(app *domain.OptionTrackService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planID, err := parseOptionTrackID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid plan id"})
			return
		}
		var req domain.OptionTrackTradeUpsertInput
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		detail, err := app.CreateTrade(planID, req)
		if err != nil {
			handleOptionTrackError(ctx, err)
			return
		}
		ctx.JSON(http.StatusCreated, detail)
	}
}

func UpdateOptionTrackTrade(app *domain.OptionTrackService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planID, err := parseOptionTrackID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid plan id"})
			return
		}
		tradeID, err := parseOptionTrackID(ctx.Param("tradeId"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid trade id"})
			return
		}
		var req domain.OptionTrackTradeUpsertInput
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		detail, err := app.UpdateTrade(planID, tradeID, req)
		if err != nil {
			handleOptionTrackError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, detail)
	}
}

func DeleteOptionTrackTrade(app *domain.OptionTrackService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planID, err := parseOptionTrackID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid plan id"})
			return
		}
		tradeID, err := parseOptionTrackID(ctx.Param("tradeId"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid trade id"})
			return
		}
		detail, err := app.DeleteTrade(planID, tradeID)
		if err != nil {
			handleOptionTrackError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, detail)
	}
}

func CloseOptionTrackPlan(app *domain.OptionTrackService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planID, err := parseOptionTrackID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid plan id"})
			return
		}
		summary, err := app.ClosePlan(planID)
		if err != nil {
			handleOptionTrackError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, summary)
	}
}

func GetOptionTrackPnL(app *domain.OptionTrackService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planID, err := parseOptionTrackID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid plan id"})
			return
		}
		resp, err := app.GetPlanPnL(planID)
		if err != nil {
			handleOptionTrackError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func PreviewOptionTrackPnL(app *domain.OptionTrackService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		planID, err := parseOptionTrackID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid plan id"})
			return
		}
		var req domain.OptionTrackPnLPreviewRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, err := app.PreviewPlanPnL(planID, req)
		if err != nil {
			handleOptionTrackError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func parseOptionTrackID(raw string) (int64, error) {
	return strconv.ParseInt(raw, 10, 64)
}

func handleOptionTrackError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrOptionTrackInvalidPayload), errors.Is(err, domain.ErrOptionTrackInvalidHistory):
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrOptionTrackNotFound), errors.Is(err, domain.ErrOptionTrackTradeNotFound):
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrOptionTrackClosedPlan), errors.Is(err, domain.ErrOptionTrackOpenPositions), errors.Is(err, domain.ErrOptionTrackClosedEditBreak):
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}