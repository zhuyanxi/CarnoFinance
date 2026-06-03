package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"time"
)

const (
	inputFileName  = "csi1000_daily.csv"
	dateLayout     = "2006-01-02"
	bucketSizeDays = 10
)

type thresholdConfig struct {
	Label     string
	Value     float64
	OutputDir string
}

type dailyBar struct {
	TradeDate time.Time
	High      float64
	Low       float64
	Close     float64
}

type moveResult struct {
	NaturalDays int
	BreachDate  time.Time
	BreachSide  string
	BreachHigh  float64
	BreachLow   float64
	HasBreach   bool
}

type bucketRow struct {
	Label      string
	StartDay   int
	EndDay     int
	Count      int
	Percentage float64
}

func main() {
	baseDir, err := resolveBaseDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "resolve base dir: %v\n", err)
		os.Exit(1)
	}

	series, err := readDailyBars(filepath.Join(baseDir, "total", inputFileName))
	if err != nil {
		fmt.Fprintf(os.Stderr, "read daily bars csv: %v\n", err)
		os.Exit(1)
	}

	configs := []thresholdConfig{
		{Label: "5pct", Value: 0.05, OutputDir: "5pct-move"},
		{Label: "10pct", Value: 0.10, OutputDir: "10pct-move"},
	}

	for _, config := range configs {
		results := computeFirstHighLowMoveDays(series, config.Value)
		outputDir := filepath.Join(baseDir, config.OutputDir)

		detailPath := filepath.Join(outputDir, fmt.Sprintf("csi1000_high_low_move_%s_natural_days.csv", config.Label))
		if err := writeResults(detailPath, series, results, config); err != nil {
			fmt.Fprintf(os.Stderr, "write %s detail csv: %v\n", config.Label, err)
			os.Exit(1)
		}

		summaryPath := filepath.Join(outputDir, fmt.Sprintf("csi1000_high_low_move_%s_summary.md", config.Label))
		if err := writeSummary(summaryPath, results, config); err != nil {
			fmt.Fprintf(os.Stderr, "write %s summary file: %v\n", config.Label, err)
			os.Exit(1)
		}

		bucketPath := filepath.Join(outputDir, fmt.Sprintf("csi1000_high_low_move_%s_bucket_distribution.csv", config.Label))
		if err := writeBucketDistribution(bucketPath, results); err != nil {
			fmt.Fprintf(os.Stderr, "write %s bucket file: %v\n", config.Label, err)
			os.Exit(1)
		}

		fmt.Printf("saved %s high/low results to %s, %s, and %s\n", config.Label, detailPath, summaryPath, bucketPath)
	}
}

func resolveBaseDir() (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("cannot resolve current file path")
	}
	return filepath.Dir(currentFile), nil
}

func readDailyBars(path string) ([]dailyBar, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	series := make([]dailyBar, 0, 4096)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) < 5 {
			return nil, fmt.Errorf("unexpected record length: %d", len(record))
		}

		tradeDate, err := time.Parse(dateLayout, record[0])
		if err != nil {
			return nil, fmt.Errorf("parse trade date %q: %w", record[0], err)
		}
		highPrice, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, fmt.Errorf("parse high %q: %w", record[2], err)
		}
		lowPrice, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, fmt.Errorf("parse low %q: %w", record[3], err)
		}
		closePrice, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return nil, fmt.Errorf("parse close %q: %w", record[4], err)
		}

		series = append(series, dailyBar{
			TradeDate: tradeDate,
			High:      highPrice,
			Low:       lowPrice,
			Close:     closePrice,
		})
	}

	return series, nil
}

func computeFirstHighLowMoveDays(series []dailyBar, threshold float64) []moveResult {
	results := make([]moveResult, len(series))
	for i := range results {
		upperBound := series[i].Close * (1 + threshold)
		lowerBound := series[i].Close * (1 - threshold)

		for j := i + 1; j < len(series); j++ {
			upBreach := series[j].High > upperBound
			downBreach := series[j].Low < lowerBound
			if !upBreach && !downBreach {
				continue
			}

			side := "up"
			if downBreach {
				side = "down"
			}
			if upBreach && downBreach {
				side = "both"
			}

			results[i] = moveResult{
				NaturalDays: int(series[j].TradeDate.Sub(series[i].TradeDate).Hours()/24) + 1,
				BreachDate:  series[j].TradeDate,
				BreachSide:  side,
				BreachHigh:  series[j].High,
				BreachLow:   series[j].Low,
				HasBreach:   true,
			}
			break
		}
	}
	return results
}

func writeResults(path string, series []dailyBar, results []moveResult, config thresholdConfig) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"trade_date",
		"start_close",
		fmt.Sprintf("natural_days_to_%s_high_low_move", config.Label),
		"breach_date",
		"breach_side",
		"breach_high",
		"breach_low",
		"upper_bound",
		"lower_bound",
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	for i, item := range series {
		upperBound := item.Close * (1 + config.Value)
		lowerBound := item.Close * (1 - config.Value)
		record := []string{
			item.TradeDate.Format(dateLayout),
			formatPrice(item.Close),
			"",
			"",
			"",
			"",
			"",
			formatPrice(upperBound),
			formatPrice(lowerBound),
		}
		if results[i].HasBreach {
			record[2] = strconv.Itoa(results[i].NaturalDays)
			record[3] = results[i].BreachDate.Format(dateLayout)
			record[4] = results[i].BreachSide
			record[5] = formatPrice(results[i].BreachHigh)
			record[6] = formatPrice(results[i].BreachLow)
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return writer.Error()
}

func writeSummary(path string, results []moveResult, config thresholdConfig) error {
	values := make([]float64, 0, len(results))
	missingCount := 0
	upCount := 0
	downCount := 0
	bothCount := 0
	for _, result := range results {
		if !result.HasBreach {
			missingCount++
			continue
		}
		values = append(values, float64(result.NaturalDays))
		switch result.BreachSide {
		case "up":
			upCount++
		case "down":
			downCount++
		case "both":
			bothCount++
		}
	}

	sort.Float64s(values)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if len(values) == 0 {
		_, err = fmt.Fprintln(file, "No breach rows found.")
		return err
	}

	buckets := buildBucketDistribution(results)
	titlePct := int(config.Value * 100)
	lines := []string{
		fmt.Sprintf("# CSI1000 High/Low %d%% Move Summary", titlePct),
		"",
		"Basis: start from each date's close, then scan later daily high/low for first threshold breach.",
		"",
		fmt.Sprintf("- sample_with_breach: %d", len(values)),
		fmt.Sprintf("- sample_without_breach: %d", missingCount),
		fmt.Sprintf("- up_breach: %d", upCount),
		fmt.Sprintf("- down_breach: %d", downCount),
		fmt.Sprintf("- both_side_same_day_breach: %d", bothCount),
		fmt.Sprintf("- mean_days: %.2f", mean(values)),
		fmt.Sprintf("- median_days: %.2f", percentile(values, 0.50)),
		fmt.Sprintf("- p10_days: %.2f", percentile(values, 0.10)),
		fmt.Sprintf("- p25_days: %.2f", percentile(values, 0.25)),
		fmt.Sprintf("- p75_days: %.2f", percentile(values, 0.75)),
		fmt.Sprintf("- p90_days: %.2f", percentile(values, 0.90)),
		fmt.Sprintf("- p95_days: %.2f", percentile(values, 0.95)),
		fmt.Sprintf("- min_days: %.0f", values[0]),
		fmt.Sprintf("- max_days: %.0f", values[len(values)-1]),
		"",
		"## Bucket Distribution",
		"",
		"| bucket | start_day | end_day | count | percentage |",
		"| --- | ---: | ---: | ---: | ---: |",
	}

	for _, bucket := range buckets {
		startDay := ""
		endDay := ""
		if bucket.StartDay > 0 {
			startDay = strconv.Itoa(bucket.StartDay)
		}
		if bucket.EndDay > 0 {
			endDay = strconv.Itoa(bucket.EndDay)
		}
		lines = append(lines, fmt.Sprintf("| %s | %s | %s | %d | %.2f%% |", bucket.Label, startDay, endDay, bucket.Count, bucket.Percentage))
	}

	for _, line := range lines {
		if _, err := fmt.Fprintln(file, line); err != nil {
			return err
		}
	}

	return nil
}

func writeBucketDistribution(path string, results []moveResult) error {
	buckets := buildBucketDistribution(results)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"bucket", "start_day", "end_day", "count", "percentage"}); err != nil {
		return err
	}

	for _, bucket := range buckets {
		startDay := ""
		endDay := ""
		if bucket.StartDay > 0 {
			startDay = strconv.Itoa(bucket.StartDay)
		}
		if bucket.EndDay > 0 {
			endDay = strconv.Itoa(bucket.EndDay)
		}
		if err := writer.Write([]string{
			bucket.Label,
			startDay,
			endDay,
			strconv.Itoa(bucket.Count),
			fmt.Sprintf("%.4f", bucket.Percentage),
		}); err != nil {
			return err
		}
	}

	return writer.Error()
}

func buildBucketDistribution(results []moveResult) []bucketRow {
	maxDays := 0
	missingCount := 0
	for _, result := range results {
		if !result.HasBreach {
			missingCount++
			continue
		}
		if result.NaturalDays > maxDays {
			maxDays = result.NaturalDays
		}
	}

	bucketCount := 0
	if maxDays > 0 {
		bucketCount = (maxDays + bucketSizeDays - 1) / bucketSizeDays
	}
	buckets := make([]bucketRow, 0, bucketCount+1)
	for idx := 0; idx < bucketCount; idx++ {
		startDay := idx*bucketSizeDays + 1
		endDay := (idx + 1) * bucketSizeDays
		buckets = append(buckets, bucketRow{
			Label:    fmt.Sprintf("%d-%d", startDay, endDay),
			StartDay: startDay,
			EndDay:   endDay,
		})
	}

	for _, result := range results {
		if !result.HasBreach {
			continue
		}
		bucketIndex := (result.NaturalDays - 1) / bucketSizeDays
		buckets[bucketIndex].Count++
	}

	total := float64(len(results))
	for idx := range buckets {
		buckets[idx].Percentage = float64(buckets[idx].Count) / total * 100
	}

	if missingCount > 0 {
		buckets = append(buckets, bucketRow{
			Label:      "no_breach_yet",
			Count:      missingCount,
			Percentage: float64(missingCount) / total * 100,
		})
	}

	return buckets
}

func mean(values []float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

func percentile(values []float64, p float64) float64 {
	if len(values) == 1 {
		return values[0]
	}
	position := p * float64(len(values)-1)
	lower := int(math.Floor(position))
	upper := int(math.Ceil(position))
	if lower == upper {
		return values[lower]
	}
	weight := position - float64(lower)
	return values[lower] + (values[upper]-values[lower])*weight
}

func formatPrice(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64)
}