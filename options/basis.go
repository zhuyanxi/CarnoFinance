package main

import (
	"fmt"
	"strconv"
	"time"
)

type BasisMetrics struct {
	DiscountPoints           float64  `json:"discount_points"`
	DiscountRatePct          float64  `json:"discount_rate_pct"`
	DailyDiscountPoints      *float64 `json:"daily_discount_points,omitempty"`
	AnnualizedDiscountRatePct *float64 `json:"annualized_discount_rate_pct"`
	DaysToExpiry              int      `json:"days_to_expiry"`
}

func calculateBasis(spotPrice, futurePrice float64, asOf, expiry time.Time) (BasisMetrics, error) {
	if spotPrice <= 0 {
		return BasisMetrics{}, fmt.Errorf("spot price must be positive")
	}

	quoteDate := dateOnly(asOf)
	expiryDate := dateOnly(expiry)
	daysToExpiry := int(expiryDate.Sub(quoteDate).Hours() / 24)
	discountRatePct := (spotPrice - futurePrice) / spotPrice * 100

	metrics := BasisMetrics{
		DiscountPoints:  spotPrice - futurePrice,
		DiscountRatePct: discountRatePct,
		DaysToExpiry:    daysToExpiry,
	}
	if daysToExpiry > 0 {
		dailyDiscountPoints := metrics.DiscountPoints / float64(daysToExpiry)
		metrics.DailyDiscountPoints = &dailyDiscountPoints
		annualizedRatePct := discountRatePct * 365 / float64(daysToExpiry)
		metrics.AnnualizedDiscountRatePct = &annualizedRatePct
	}
	return metrics, nil
}

func contractExpiry(contract string) (time.Time, error) {
	if len(contract) != 6 || contract[:2] != "IM" {
		return time.Time{}, fmt.Errorf("invalid IM contract code %q", contract)
	}

	year, err := strconv.Atoi(contract[2:4])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year in contract code %q", contract)
	}
	month, err := strconv.Atoi(contract[4:6])
	if err != nil || month < 1 || month > 12 {
		return time.Time{}, fmt.Errorf("invalid month in contract code %q", contract)
	}

	firstDay := time.Date(2000+year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	daysUntilFriday := (int(time.Friday-firstDay.Weekday()) + 7) % 7
	thirdFriday := firstDay.AddDate(0, 0, daysUntilFriday+14)
	return thirdFriday, nil
}

func dateOnly(value time.Time) time.Time {
	return time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, time.UTC)
}