package main

import (
	"math"
	"testing"
	"time"
)

func TestCalculateBasis(t *testing.T) {
	asOf := time.Date(2026, time.January, 1, 15, 30, 0, 0, time.FixedZone("CST", 8*60*60))
	expiry := time.Date(2026, time.January, 31, 0, 0, 0, 0, time.UTC)

	metrics, err := calculateBasis(5000, 4950, asOf, expiry)
	if err != nil {
		t.Fatalf("calculate basis: %v", err)
	}
	if metrics.DiscountPoints != 50 {
		t.Fatalf("expected 50 discount points, got %v", metrics.DiscountPoints)
	}
	if metrics.DiscountRatePct != 1 {
		t.Fatalf("expected 1%% discount rate, got %v", metrics.DiscountRatePct)
	}
	if metrics.DaysToExpiry != 30 {
		t.Fatalf("expected 30 days to expiry, got %d", metrics.DaysToExpiry)
	}
	if metrics.DailyDiscountPoints == nil || math.Abs(*metrics.DailyDiscountPoints-50.0/30) > 1e-9 {
		t.Fatalf("unexpected daily discount value: %v", metrics.DailyDiscountPoints)
	}
	if metrics.AnnualizedDiscountRatePct == nil || math.Abs(*metrics.AnnualizedDiscountRatePct-365.0/30) > 1e-9 {
		t.Fatalf("unexpected annualized rate: %v", metrics.AnnualizedDiscountRatePct)
	}
}

func TestCalculateBasisAtExpiryHasNoAnnualizedRate(t *testing.T) {
	expiry := time.Date(2026, time.June, 19, 0, 0, 0, 0, time.UTC)
	metrics, err := calculateBasis(5000, 4950, expiry, expiry)
	if err != nil {
		t.Fatalf("calculate basis: %v", err)
	}
	if metrics.DaysToExpiry != 0 {
		t.Fatalf("expected 0 days to expiry, got %d", metrics.DaysToExpiry)
	}
	if metrics.AnnualizedDiscountRatePct != nil {
		t.Fatalf("expected no annualized rate on expiry day, got %v", *metrics.AnnualizedDiscountRatePct)
	}
	if metrics.DailyDiscountPoints != nil {
		t.Fatalf("expected no daily discount value on expiry day, got %v", *metrics.DailyDiscountPoints)
	}
}

func TestContractExpiryUsesThirdFriday(t *testing.T) {
	expiry, err := contractExpiry("IM2606")
	if err != nil {
		t.Fatalf("contract expiry: %v", err)
	}
	expected := time.Date(2026, time.June, 19, 0, 0, 0, 0, time.UTC)
	if !expiry.Equal(expected) {
		t.Fatalf("expected %s, got %s", expected.Format(time.DateOnly), expiry.Format(time.DateOnly))
	}
}

func TestContractExpiryRejectsInvalidCode(t *testing.T) {
	if _, err := contractExpiry("IM2613"); err == nil {
		t.Fatal("expected invalid contract code to fail")
	}
}