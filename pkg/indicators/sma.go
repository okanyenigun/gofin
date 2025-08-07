package indicators

func CalculateSMA(prices []float64, period int) []float64 {
	if len(prices) < period {
		return nil
	}

	smaValues := make([]float64, len(prices))

	for i := period - 1; i < len(prices); i++ {
		sum := 0.0
		for j := i - period + 1; j <= i; j++ {
			sum += prices[j]
		}
		smaValues[i] = sum / float64(period)
	}

	return smaValues
}
