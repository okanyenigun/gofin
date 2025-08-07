package main

import (
	"gofin/internal/display"
	"gofin/pkg/analysis"
	"gofin/pkg/data"
	"gofin/pkg/indicators"
	"log"
)

func main() {
	// Load CSV
	csvData, err := data.LoadCSV("nvda.csv")
	if err != nil {
		log.Fatal("Failed to load CSV:", err)
	}

	closePrices := csvData.GetClosePrices()

	config := indicators.DefaultConfig()

	calculator := indicators.NewCalculator(config)

	results, duration := calculator.CalculateAll(closePrices)

	analyzer := analysis.NewAnalyzer()
	formatter := display.NewFormatter()

	formatter.PrintHeader(csvData.Headers)
	formatter.PrintCalculationTime(duration)
	formatter.PrintFirstRows(csvData, results, config, 5)
	formatter.PrintLastRows(csvData, results, analyzer, 10)

}
