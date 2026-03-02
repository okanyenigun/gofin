package display

import (
	"fmt"
	"gofin/pkg/analysis"
	"gofin/pkg/data"
	"gofin/pkg/indicators"
	"time"
)

type Formatter struct{}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) PrintHeader(headers []string) {
	fmt.Println(headers)
}

func (f *Formatter) PrintCalculationTime(duration time.Duration) {
	fmt.Printf("\nCalculating technical indicators concurrently...\n")
	fmt.Printf("All indicators calculated successfully in %v!\n", duration)
}

func (f *Formatter) PrintFirstRows(csvData *data.CSVData, results indicators.Results, config indicators.Config, count int) {
	fmt.Printf("\nFirst %d rows with RSI (period=%d), Bollinger Bands (period=%d, std=%.1f), SMAs, MACD (%d/%d/%d), and Stochastic (%d/%d):\n",
		count, config.RSIPeriod, config.BBPeriod, config.BBStdDev,
		config.MACDFastPeriod, config.MACDSlowPeriod, config.MACDSignalPeriod,
		config.StochPeriod, config.StochSmoothPeriod)
	fmt.Printf("Datetime, Close, RSI, BB_Upper, BB_Middle, BB_Lower, SMA10, SMA20, SMA50, MACD, MACD_Signal, MACD_Hist, Stoch_K, Stoch_D")

	macdStart := config.MACDSlowPeriod + config.MACDSignalPeriod - 2
	stochStart := config.StochPeriod + config.StochSmoothPeriod - 2

	for i := 0; i < count && i < len(csvData.Records); i++ {
		record := csvData.Records[i]

		rsiStr := f.formatValue(results.RSI, i+config.RSIPeriod)
		bbUpperStr := f.formatValue(results.BollingerBands.Upper, i+config.BBPeriod-1)
		bbMiddleStr := f.formatValue(results.BollingerBands.Middle, i+config.BBPeriod-1)
		bbLowerStr := f.formatValue(results.BollingerBands.Lower, i+config.BBPeriod-1)
		sma10Str := f.formatValue(results.SMA10, i+config.SMA10Period-1)
		sma20Str := f.formatValue(results.SMA20, i+config.SMA20Period-1)
		sma50Str := f.formatValue(results.SMA50, i+config.SMA50Period-1)
		macdLineStr := f.formatValue(results.MACD.Line, i+macdStart)
		macdSigStr := f.formatValue(results.MACD.Signal, i+macdStart)
		macdHistStr := f.formatValue(results.MACD.Histogram, i+macdStart)
		stochKStr := f.formatValue(results.Stochastic.K, i+stochStart)
		stochDStr := f.formatValue(results.Stochastic.D, i+stochStart)

		fmt.Printf("%s, %.2f, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s\n",
			record.Datetime, record.Close,
			rsiStr, bbUpperStr, bbMiddleStr, bbLowerStr,
			sma10Str, sma20Str, sma50Str,
			macdLineStr, macdSigStr, macdHistStr,
			stochKStr, stochDStr)
	}
}

func (f *Formatter) PrintLastRows(csvData *data.CSVData, results indicators.Results, analyzer *analysis.Analyzer, count int) {
	fmt.Printf("\nLast %d values with all indicators:\n", count)
	fmt.Println("Date, Close, RSI, BB_Upper, BB_Middle, BB_Lower, SMA10, SMA20, SMA50, MACD, MACD_Signal, MACD_Hist, Stoch_K, Stoch_D, Trend, RSI_Signal, MACD_Signal, Stoch_Signal")

	start := len(results.RSI) - count
	if start < 0 {
		start = 0
	}

	for i := start; i < len(results.RSI) && i < len(csvData.Records); i++ {
		if results.RSI[i] != 0 && i < len(results.BollingerBands.Upper) && results.BollingerBands.Upper[i] != 0 {
			record := csvData.Records[i]
			analysisResult := analyzer.AnalyzePoint(results, i, record.Close)

			macdLine := results.MACD.Line[i]
			macdSig := results.MACD.Signal[i]
			macdHist := results.MACD.Histogram[i]
			stochK := results.Stochastic.K[i]
			stochD := results.Stochastic.D[i]

			fmt.Printf("%s, %.2f, %.2f, %.2f, %.2f, %.2f, %.2f, %.2f, %.2f, %.4f, %.4f, %.4f, %.2f, %.2f, %s, %s, %s, %s\n",
				record.Datetime, record.Close, results.RSI[i],
				results.BollingerBands.Upper[i], results.BollingerBands.Middle[i], results.BollingerBands.Lower[i],
				results.SMA10[i], results.SMA20[i], results.SMA50[i],
				macdLine, macdSig, macdHist,
				stochK, stochD,
				analysisResult.Trend, analysisResult.RSISignal, analysisResult.MACDSignal, analysisResult.StochasticSignal)
		}
	}
}

func (f *Formatter) formatValue(values []float64, index int) string {
	if index >= 0 && index < len(values) && values[index] != 0 {
		return fmt.Sprintf("%.2f", values[index])
	}
	return "N/A"
}
