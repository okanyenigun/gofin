package indicators

import "math"

func CalculateRSI(closePrices []float64, period int) []float64 {
	if len(closePrices) < period+1 {
		return nil
	}

	rsiValues := make([]float64, len(closePrices))

	priceChanges := make([]float64, len(closePrices)-1)
	for i := 1; i < len(closePrices); i++ {
		priceChanges[i-1] = closePrices[i] - closePrices[i-1]
	}

	var avgGain, avgLoss float64
	for i := 0; i < period; i++ {
		if priceChanges[i] > 0 {
			avgGain += priceChanges[i]
		} else {
			avgLoss += math.Abs(priceChanges[i])
		}
	}

	avgGain /= float64(period)
	avgLoss /= float64(period)

	if avgLoss == 0 {
		rsiValues[period] = 100
	} else {
		rs := avgGain / avgLoss
		rsiValues[period] = 100 - (100 / (1 + rs))
	}

	for i := period + 1; i < len(closePrices); i++ {
		change := priceChanges[i-1]
		var gain, loss float64

		if change > 0 {
			gain = change
		} else {
			loss = math.Abs(change)
		}

		avgGain = ((avgGain * float64(period-1)) + gain) / float64(period)
		avgLoss = ((avgLoss * float64(period-1)) + loss) / float64(period)

		if avgLoss == 0 {
			rsiValues[i] = 100
		} else {
			rs := avgGain / avgLoss
			rsiValues[i] = 100 - (100 / (1 + rs))
		}
	}

	return rsiValues

}
