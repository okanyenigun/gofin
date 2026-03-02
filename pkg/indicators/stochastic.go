package indicators

// CalculateStochastic computes the Stochastic Oscillator.
//
// %K = (Close - LowestLow(period)) / (HighestHigh(period) - LowestLow(period)) * 100
// %D = SMA of %K over smoothPeriod (typically 3)
//
// Values are aligned to the full input slice; indices before period-1 are zero.
func CalculateStochastic(highs, lows, closes []float64, period, smoothPeriod int) Stochastic {
	n := len(closes)
	k := make([]float64, n)
	d := make([]float64, n)

	if n < period || len(highs) != n || len(lows) != n {
		return Stochastic{K: k, D: d}
	}
	for i := period - 1; i < n; i++ {
		highestHigh := highs[i]
		lowestLow := lows[i]
		for j := i - (period - 1); j <= i; j++ {
			if highs[j] > highestHigh {
				highestHigh = highs[j]
			}
			if lows[j] < lowestLow {
				lowestLow = lows[j]
			}
		}

		diff := highestHigh - lowestLow
		if diff == 0 {
			k[i] = 50
		} else {
			k[i] = (closes[i] - lowestLow) / diff * 100
		}
	}

	for i := period - 1 + smoothPeriod - 1; i < n; i++ {
		var sum float64
		for j := i - (smoothPeriod - 1); j <= i; j++ {
			sum += k[j]
		}
		d[i] = sum / float64(smoothPeriod)
	}

	return Stochastic{K: k, D: d}
}
