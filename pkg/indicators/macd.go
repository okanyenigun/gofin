package indicators

// calculateEMA computes the Exponential Moving Average for a given period.
// Indices 0..period-2 are zero (undefined). Index period-1 is seeded with SMA.
func calculateEMA(prices []float64, period int) []float64 {
	if len(prices) < period {
		return nil
	}

	ema := make([]float64, len(prices))
	multiplier := 2.0 / float64(period+1)

	var sum float64
	for i := 0; i < period; i++ {
		sum += prices[i]
	}

	ema[period-1] = sum / float64(period)

	for i := period; i < len(prices); i++ {
		ema[i] = prices[i]*multiplier + ema[i-1]*(1-multiplier)
	}

	return ema
}

// CalculateMACD computes the MACD line, signal line, and histogram.
//
// MACD line  = EMA(fast) − EMA(slow)          valid from index slowPeriod-1
// Signal     = EMA(signalPeriod) of MACD line  valid from index slowPeriod+signalPeriod-2
// Histogram  = MACD line − Signal
func CalculateMACD(closePrices []float64, fastPeriod, slowPeriod, signalPeriod int) MACD {
	n := len(closePrices)
	if n < slowPeriod+signalPeriod {
		return MACD{
			Line:      make([]float64, n),
			Signal:    make([]float64, n),
			Histogram: make([]float64, n),
		}
	}

	fastEMA := calculateEMA(closePrices, fastPeriod)
	slowEMA := calculateEMA(closePrices, slowPeriod)

	macdLine := make([]float64, n)
	for i := slowPeriod - 1; i < n; i++ {
		macdLine[i] = fastEMA[i] - slowEMA[i]
	}

	macdSlice := macdLine[slowPeriod-1:]
	signalSlice := calculateEMA(macdSlice, signalPeriod)

	signal := make([]float64, n)
	histogram := make([]float64, n)

	signalStart := slowPeriod - 1 + signalPeriod - 1
	for i := signalStart; i < n; i++ {
		signal[i] = signalSlice[i-(slowPeriod-1)]
		histogram[i] = macdLine[i] - signal[i]
	}

	return MACD{
		Line:      macdLine,
		Signal:    signal,
		Histogram: histogram,
	}
}
