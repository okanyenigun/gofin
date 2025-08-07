package indicators

import "math"

func CalculateBollingerBands(closePrices []float64, period int, stdDevMultiplier float64) BollingerBands {
	bands := BollingerBands{
		Upper:  make([]float64, len(closePrices)),
		Middle: make([]float64, len(closePrices)),
		Lower:  make([]float64, len(closePrices)),
	}

	if len(closePrices) < period {
		return bands
	}

	for i := period - 1; i < len(closePrices); i++ {
		sum := 0.0
		for j := i - period + 1; j <= i; j++ {
			sum += closePrices[j]
		}

		sma := sum / float64(period)
		bands.Middle[i] = sma

		sumSquaredDiff := 0.0
		for j := i - period + 1; j <= i; j++ {
			diff := closePrices[j] - sma
			sumSquaredDiff += diff * diff
		}
		stdDev := math.Sqrt(sumSquaredDiff / float64(period))

		bands.Upper[i] = sma + (stdDevMultiplier * stdDev)
		bands.Lower[i] = sma - (stdDevMultiplier * stdDev)
	}

	return bands

}
