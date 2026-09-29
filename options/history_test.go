package main

import "testing"

func TestParseDailyKlinesUsesCloseColumn(t *testing.T) {
	closes, err := parseDailyKlines([]string{
		"2026-09-28,7300.0,7290.5,7310.0,7280.0,123,456.0",
		"2026-09-29,7290.0,7302.1,7315.0,7285.0,234,567.0",
	})
	if err != nil {
		t.Fatalf("parse daily klines: %v", err)
	}
	if closes["2026-09-28"] != 7290.5 || closes["2026-09-29"] != 7302.1 {
		t.Fatalf("unexpected daily close map: %#v", closes)
	}
}

func TestParseDailyKlinesRejectsMalformedRows(t *testing.T) {
	if _, err := parseDailyKlines([]string{"2026-09-29,7290.0"}); err == nil {
		t.Fatal("expected malformed daily kline to fail")
	}
}