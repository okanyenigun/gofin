package indicators

import (
	"sync"
	"time"
)

type Calculator struct {
	config Config
}

func NewCalculator(config Config) *Calculator {
	return &Calculator{
		config: config,
	}
}

func (c *Calculator) CalculateAll(highs, lows, closePrices []float64) (Results, time.Duration) {
	start := time.Now()

	var wg sync.WaitGroup
	results := Results{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		results.RSI = CalculateRSI(closePrices, c.config.RSIPeriod)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		results.BollingerBands = CalculateBollingerBands(closePrices, c.config.BBPeriod, c.config.BBStdDev)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		results.SMA10 = CalculateSMA(closePrices, c.config.SMA10Period)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		results.SMA20 = CalculateSMA(closePrices, c.config.SMA20Period)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		results.SMA50 = CalculateSMA(closePrices, c.config.SMA50Period)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		results.MACD = CalculateMACD(closePrices, c.config.MACDFastPeriod, c.config.MACDSlowPeriod, c.config.MACDSignalPeriod)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		results.Stochastic = CalculateStochastic(highs, lows, closePrices, c.config.StochPeriod, c.config.StochSmoothPeriod)
	}()

	wg.Wait()

	duration := time.Since(start)

	return results, duration
}
