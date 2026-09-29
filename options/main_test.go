package main

import (
	"strings"
	"testing"
)

func TestFormatSnapshotTableAlignsColumnsAndShowsDailyDiscount(t *testing.T) {
	dailyDiscount := 50.0 / 30
	annualizedRate := 365.0 / 30
	table := formatSnapshotTable(MarketSnapshot{
		Contracts: []ContractSnapshot{{
			Contract:    "IM2610",
			FuturePrice: 4950,
			ExpiryDate:  "2026-10-16",
			BasisMetrics: BasisMetrics{
				DiscountPoints:            50,
				DailyDiscountPoints:       &dailyDiscount,
				DiscountRatePct:           1,
				AnnualizedDiscountRatePct: &annualizedRate,
				DaysToExpiry:               30,
			},
		}},
	})
	lines := strings.Split(strings.TrimSpace(table), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected header and one data row, got %d lines", len(lines))
	}
	if !strings.Contains(lines[0], "每日贴水值") || !strings.Contains(lines[1], "1.6667") {
		t.Fatalf("table is missing daily discount value:\n%s", table)
	}

	headerSpans := tableColumnSpans(lines[0])
	rowSpans := tableColumnSpans(lines[1])
	if len(headerSpans) != len(rowSpans) {
		t.Fatalf("header and row have different column counts: %d != %d", len(headerSpans), len(rowSpans))
	}
	rightAligned := []bool{false, true, true, true, true, true, true, false}
	for column := range headerSpans {
		if rightAligned[column] && headerSpans[column][1] != rowSpans[column][1] {
			t.Errorf("column %d right edges differ: %d != %d", column, headerSpans[column][1], rowSpans[column][1])
		}
		if !rightAligned[column] && headerSpans[column][0] != rowSpans[column][0] {
			t.Errorf("column %d left edges differ: %d != %d", column, headerSpans[column][0], rowSpans[column][0])
		}
	}
}

func tableColumnSpans(line string) [][2]int {
	fields := strings.Fields(line)
	spans := make([][2]int, 0, len(fields))
	byteOffset := 0
	for _, field := range fields {
		fieldOffset := strings.Index(line[byteOffset:], field)
		if fieldOffset < 0 {
			return nil
		}
		start := byteOffset + fieldOffset
		end := start + len(field)
		spans = append(spans, [2]int{
			terminalDisplayWidth(line[:start]),
			terminalDisplayWidth(line[:end]),
		})
		byteOffset = end
	}
	return spans
}