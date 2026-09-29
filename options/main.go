package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unicode"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	interval := flag.Duration("interval", 5*time.Second, "quote polling interval (minimum 1s)")
	once := flag.Bool("once", false, "fetch one snapshot and exit")
	historyDays := flag.Int("history-days", 0, "download daily basis history for this many calendar days and exit")
	outputDir := flag.String("output-dir", defaultOutputDir(), "directory for latest.json and daily_basis.csv")
	flag.Parse()

	if *interval < time.Second {
		return fmt.Errorf("interval must be at least 1s")
	}
	if *historyDays < 0 {
		return fmt.Errorf("history-days cannot be negative")
	}
	if *historyDays > 0 && *once {
		return fmt.Errorf("-once and -history-days cannot be used together")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	client := NewEastmoneyClient()
	if *historyDays > 0 {
		return runHistory(ctx, client, *historyDays, *outputDir)
	}
	if *once {
		return pollOnce(ctx, client, *outputDir)
	}
	return watch(ctx, client, *outputDir, *interval)
}

func runHistory(ctx context.Context, client *EastmoneyClient, calendarDays int, outputDir string) error {
	records, err := client.DailyHistory(ctx, calendarDays)
	if err != nil {
		return err
	}
	path := filepath.Join(outputDir, "daily_basis.csv")
	if err := writeDailyBasisCSV(path, records); err != nil {
		return err
	}
	fmt.Printf("已写入 %d 条日度贴水记录：%s\n", len(records), path)
	return nil
}

func watch(ctx context.Context, client *EastmoneyClient, outputDir string, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		if err := pollOnce(ctx, client, outputDir); err != nil {
			if ctx.Err() != nil {
				return nil
			}
			fmt.Fprintf(os.Stderr, "行情采集失败：%v\n", err)
		}
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
	}
}

func pollOnce(ctx context.Context, client *EastmoneyClient, outputDir string) error {
	snapshot, err := client.Snapshot(ctx)
	if err != nil {
		return err
	}
	if err := writeLatestSnapshot(outputDir, snapshot); err != nil {
		return err
	}
	printSnapshot(snapshot)
	return nil
}

func printSnapshot(snapshot MarketSnapshot) {
	quoteTime := "未知"
	if snapshot.SpotQuoteTime != nil {
		quoteTime = snapshot.SpotQuoteTime.Format("2006-01-02 15:04:05")
	}
	fmt.Printf("\n采集时间：%s | %s：%.2f | 指数行情时间：%s\n",
		snapshot.RetrievedAt.Format("2006-01-02 15:04:05"), snapshot.SpotName, snapshot.SpotPrice, quoteTime)
	fmt.Print(formatSnapshotTable(snapshot))
}

func formatSnapshotTable(snapshot MarketSnapshot) string {
	rows := [][]string{{"合约", "期货价格", "贴水点数", "每日贴水值", "贴水率", "年化贴水", "到期天数", "到期日"}}
	for _, contract := range snapshot.Contracts {
		dailyDiscount := "N/A"
		if contract.DailyDiscountPoints != nil {
			dailyDiscount = fmt.Sprintf("%.4f", *contract.DailyDiscountPoints)
		}
		annualizedRate := "N/A"
		if contract.AnnualizedDiscountRatePct != nil {
			annualizedRate = fmt.Sprintf("%.2f%%", *contract.AnnualizedDiscountRatePct)
		}
		rows = append(rows, []string{
			contract.Contract,
			fmt.Sprintf("%.2f", contract.FuturePrice),
			fmt.Sprintf("%.2f", contract.DiscountPoints),
			dailyDiscount,
			fmt.Sprintf("%.3f%%", contract.DiscountRatePct),
			annualizedRate,
			fmt.Sprintf("%d", contract.DaysToExpiry),
			contract.ExpiryDate,
		})
	}

	rightAligned := []bool{false, true, true, true, true, true, true, false}
	columnWidths := make([]int, len(rows[0]))
	for _, row := range rows {
		for column, cell := range row {
			if cellWidth := terminalDisplayWidth(cell); cellWidth > columnWidths[column] {
				columnWidths[column] = cellWidth
			}
		}
	}

	var output strings.Builder
	for _, row := range rows {
		for column, cell := range row {
			if column > 0 {
				output.WriteString("  ")
			}
			padding := columnWidths[column] - terminalDisplayWidth(cell)
			if rightAligned[column] {
				output.WriteString(strings.Repeat(" ", padding))
				output.WriteString(cell)
			} else {
				output.WriteString(cell)
				output.WriteString(strings.Repeat(" ", padding))
			}
		}
		output.WriteByte('\n')
	}
	return output.String()
}

func terminalDisplayWidth(value string) int {
	width := 0
	for _, character := range value {
		if unicode.Is(unicode.Han, character) ||
			(character >= '\u3000' && character <= '\u303f') ||
			(character >= '\uff01' && character <= '\uff60') {
			width += 2
		} else {
			width++
		}
	}
	return width
}

func defaultOutputDir() string {
	_, sourceFile, _, ok := runtime.Caller(0)
	if !ok {
		return filepath.Join("options", "data")
	}
	return filepath.Join(filepath.Dir(sourceFile), "data")
}