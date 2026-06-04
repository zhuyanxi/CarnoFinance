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
	inputFileName   = "csi1000_daily.csv"
	outputFileName  = "csi1000_close_move_5pct_after_5_quiet_days_natural_days.csv"
	summaryFileName = "csi1000_close_move_5pct_after_5_quiet_days_summary.md"
	bucketFileName  = "csi1000_close_move_5pct_after_5_quiet_days_bucket_distribution.csv"
	dateLayout      = "2006-01-02"
	moveThreshold   = 0.05
	quietDays       = 5
	bucketSizeDays  = 10
)

type dailyClose struct {
	TradeDate time.Time
	Close     float64
}

type quietMoveResult struct {
	StartDate            time.Time
	AnchorDate           time.Time
	StartClose           float64
	AnchorClose          float64
	NaturalDaysFromAnchor int
	BreachDate           time.Time
	BreachClose          float64
	HasBreach            bool
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

	series, err := readDailyClose(filepath.Join(baseDir, "..", "total", inputFileName))
	if err != nil {
		fmt.Fprintf(os.Stderr, "read daily close csv: %v\n", err)
		os.Exit(1)
	}

	results := computeQuietFiveDayMoveDays(series)
	outputPath := filepath.Join(baseDir, outputFileName)
	if err := writeResults(outputPath, results); err != nil {
		fmt.Fprintf(os.Stderr, "write result csv: %v\n", err)
		os.Exit(1)
	}

	summaryPath := filepath.Join(baseDir, summaryFileName)
	if err := writeSummary(summaryPath, results); err != nil {
		fmt.Fprintf(os.Stderr, "write summary file: %v\n", err)
		os.Exit(1)
	}

	bucketPath := filepath.Join(baseDir, bucketFileName)
	if err := writeBucketDistribution(bucketPath, results); err != nil {
		fmt.Fprintf(os.Stderr, "write bucket file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("saved %d rows to %s, %s, and %s\n", len(results), outputPath, summaryPath, bucketPath)
}

func resolveBaseDir() (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("cannot resolve current file path")
	}
	return filepath.Dir(currentFile), nil
}

func readDailyClose(path string) ([]dailyClose, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	series := make([]dailyClose, 0, 4096)
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
		closePrice, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return nil, fmt.Errorf("parse close %q: %w", record[4], err)
		}

		series = append(series, dailyClose{TradeDate: tradeDate, Close: closePrice})
	}

	return series, nil
}

func computeQuietFiveDayMoveDays(series []dailyClose) []quietMoveResult {
	results := make([]quietMoveResult, 0, len(series))
	for i := 0; i+quietDays < len(series); i++ {
		upperBound := series[i].Close * (1 + moveThreshold)
		lowerBound := series[i].Close * (1 - moveThreshold)
		quiet := true
		for j := i + 1; j <= i+quietDays; j++ {
			if series[j].Close > upperBound || series[j].Close < lowerBound {
				quiet = false
				break
			}
		}
		if !quiet {
			continue
		}

		result := quietMoveResult{
			StartDate:   series[i].TradeDate,
			AnchorDate:  series[i+quietDays].TradeDate,
			StartClose:  series[i].Close,
			AnchorClose: series[i+quietDays].Close,
		}

		for j := i + quietDays + 1; j < len(series); j++ {
			if series[j].Close > upperBound || series[j].Close < lowerBound {
				result.NaturalDaysFromAnchor = int(series[j].TradeDate.Sub(series[i+quietDays].TradeDate).Hours()/24) + 1
				result.BreachDate = series[j].TradeDate
				result.BreachClose = series[j].Close
				result.HasBreach = true
				break
			}
		}

		results = append(results, result)
	}
	return results
}

func writeResults(path string, results []quietMoveResult) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"trade_date", "start_close", "anchor_date", "anchor_close", "natural_days_from_day5_to_5pct_close_move", "breach_date", "breach_close"}); err != nil {
		return err
	}

	for _, result := range results {
		naturalDays := ""
		breachDate := ""
		breachClose := ""
		if result.HasBreach {
			naturalDays = strconv.Itoa(result.NaturalDaysFromAnchor)
			breachDate = result.BreachDate.Format(dateLayout)
			breachClose = strconv.FormatFloat(result.BreachClose, 'f', 2, 64)
		}
		if err := writer.Write([]string{
			result.StartDate.Format(dateLayout),
			strconv.FormatFloat(result.StartClose, 'f', 2, 64),
			result.AnchorDate.Format(dateLayout),
			strconv.FormatFloat(result.AnchorClose, 'f', 2, 64),
			naturalDays,
			breachDate,
			breachClose,
		}); err != nil {
			return err
		}
	}

	return writer.Error()
}

func writeSummary(path string, results []quietMoveResult) error {
	values := make([]float64, 0, len(results))
	missingCount := 0
	for _, result := range results {
		if !result.HasBreach {
			missingCount++
			continue
		}
		values = append(values, float64(result.NaturalDaysFromAnchor))
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
	lines := []string{
		"# CSI1000 Close 5% Move After 5 Quiet Days Summary",
		"",
		"Basis: next 5 closes after start must all stay within +/-5% of start close; count natural days from the 5th observation day to first breach.",
		"",
		fmt.Sprintf("- qualified_start_dates: %d", len(results)),
		fmt.Sprintf("- sample_with_breach: %d", len(values)),
		fmt.Sprintf("- sample_without_breach: %d", missingCount),
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

func writeBucketDistribution(path string, results []quietMoveResult) error {
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
		if err := writer.Write([]string{bucket.Label, startDay, endDay, strconv.Itoa(bucket.Count), fmt.Sprintf("%.4f", bucket.Percentage)}); err != nil {
			return err
		}
	}

	return writer.Error()
}

func buildBucketDistribution(results []quietMoveResult) []bucketRow {
	maxDays := 0
	missingCount := 0
	for _, result := range results {
		if !result.HasBreach {
			missingCount++
			continue
		}
		if result.NaturalDaysFromAnchor > maxDays {
			maxDays = result.NaturalDaysFromAnchor
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
		buckets = append(buckets, bucketRow{Label: fmt.Sprintf("%d-%d", startDay, endDay), StartDay: startDay, EndDay: endDay})
	}

	for _, result := range results {
		if !result.HasBreach {
			continue
		}
		bucketIndex := (result.NaturalDaysFromAnchor - 1) / bucketSizeDays
		buckets[bucketIndex].Count++
	}

	total := float64(len(results))
	for idx := range buckets {
		buckets[idx].Percentage = float64(buckets[idx].Count) / total * 100
	}

	if missingCount > 0 {
		buckets = append(buckets, bucketRow{Label: "no_breach_yet", Count: missingCount, Percentage: float64(missingCount) / total * 100})
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