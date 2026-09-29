package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func writeLatestSnapshot(outputDir string, snapshot MarketSnapshot) error {
	data, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return fmt.Errorf("encode latest snapshot: %w", err)
	}
	return writeFileAtomically(filepath.Join(outputDir, "latest.json"), append(data, '\n'))
}

func writeDailyBasisCSV(path string, records []DailyBasisRecord) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	file, err := os.CreateTemp(filepath.Dir(path), ".daily-basis-*.csv")
	if err != nil {
		return fmt.Errorf("create temporary daily basis file: %w", err)
	}
	temporaryPath := file.Name()
	defer os.Remove(temporaryPath)

	writer := csv.NewWriter(file)
	if err := writer.Write([]string{
		"trade_date", "contract", "expiry_date", "days_to_expiry", "spot_close", "future_close",
		"discount_points", "daily_discount_points", "discount_rate_pct", "annualized_discount_rate_pct",
	}); err != nil {
		file.Close()
		return fmt.Errorf("write daily basis header: %w", err)
	}
	for _, record := range records {
		dailyDiscount := ""
		if record.DailyDiscountPoints != nil {
			dailyDiscount = formatFloat(*record.DailyDiscountPoints)
		}
		annualizedRate := ""
		if record.AnnualizedDiscountRatePct != nil {
			annualizedRate = formatFloat(*record.AnnualizedDiscountRatePct)
		}
		if err := writer.Write([]string{
			record.TradeDate,
			record.Contract,
			record.ExpiryDate,
			strconv.Itoa(record.DaysToExpiry),
			formatFloat(record.SpotClose),
			formatFloat(record.FutureClose),
			formatFloat(record.DiscountPoints),
			dailyDiscount,
			formatFloat(record.DiscountRatePct),
			annualizedRate,
		}); err != nil {
			file.Close()
			return fmt.Errorf("write daily basis row: %w", err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		file.Close()
		return fmt.Errorf("flush daily basis CSV: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close daily basis CSV: %w", err)
	}
	if err := os.Rename(temporaryPath, path); err != nil {
		return fmt.Errorf("replace daily basis CSV: %w", err)
	}
	return nil
}

func writeFileAtomically(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	file, err := os.CreateTemp(filepath.Dir(path), ".latest-*.json")
	if err != nil {
		return fmt.Errorf("create temporary snapshot file: %w", err)
	}
	temporaryPath := file.Name()
	defer os.Remove(temporaryPath)
	if _, err := file.Write(data); err != nil {
		file.Close()
		return fmt.Errorf("write latest snapshot: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close latest snapshot: %w", err)
	}
	if err := os.Rename(temporaryPath, path); err != nil {
		return fmt.Errorf("replace latest snapshot: %w", err)
	}
	return nil
}

func formatFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}