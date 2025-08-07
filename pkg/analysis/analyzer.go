package analysis

import "gofin/pkg/indicators"

type TrendType string

const (
	StrongUpTrend   TrendType = "Strong Uptrend"
	UpTrend         TrendType = "Uptrend"
	Sideways        TrendType = "Sideways"
	Downtrend       TrendType = "Downtrend"
	StrongDowntrend TrendType = "Strong Downtrend"
	Unknown         TrendType = "Unknown"
)

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) DetermineTrend(sma10, sma20, sma50 float64) TrendType {
	if sma10 == 0 || sma20 == 0 || sma50 == 0 {
		return Unknown
	}

	if sma10 > sma20 && sma20 > sma50 {
		return StrongUpTrend
	} else if sma10 > sma20 {
		return UpTrend
	} else if sma10 < sma20 && sma20 < sma50 {
		return StrongDowntrend
	} else if sma10 < sma20 {
		return Downtrend
	}
	return Sideways
}

func (a *Analyzer) AnalyzeBollingerPosition(price, upper, lower float64) string {
	if upper == 0 || lower == 0 {
		return "N/A"
	}

	if price > upper {
		return "Above Upper Band"
	} else if price < lower {
		return "Below Lower Band"
	} else {
		return "Within Bands"
	}
}

func (a *Analyzer) GetRSISignal(rsi float64) string {
	if rsi == 0 {
		return "N/A"
	}

	if rsi > 70 {
		return "Overbought"
	} else if rsi < 30 {
		return "Oversold"
	}

	return "Neutral"
}

type AnalysisResult struct {
	Trend             TrendType
	BollingerPosition string
	RSISignal         string
}

func (a *Analyzer) AnalyzePoint(results indicators.Results, index int, price float64) AnalysisResult {
	if index >= len(results.SMA10) || index >= len(results.SMA20) || index >= len(results.SMA50) {
		return AnalysisResult{
			Trend:             Unknown,
			BollingerPosition: "N/A",
			RSISignal:         "N/A",
		}
	}

	trend := a.DetermineTrend(results.SMA10[index], results.SMA20[index], results.SMA50[index])

	var bollingerPos string
	if index < len(results.BollingerBands.Upper) {
		bollingerPos = a.AnalyzeBollingerPosition(price, results.BollingerBands.Upper[index], results.BollingerBands.Lower[index])
	} else {
		bollingerPos = "N/A"
	}

	var rsiSignal string
	if index < len(results.RSI) {
		rsiSignal = a.GetRSISignal(results.RSI[index])
	} else {
		rsiSignal = "N/A"
	}

	return AnalysisResult{
		Trend:             trend,
		BollingerPosition: bollingerPos,
		RSISignal:         rsiSignal,
	}
}
