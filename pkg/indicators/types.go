package indicators

type BollingerBands struct {
	Upper  []float64
	Middle []float64
	Lower  []float64
}

type Results struct {
	RSI            []float64
	BollingerBands BollingerBands
	SMA10          []float64
	SMA20          []float64
	SMA50          []float64
}

type Config struct {
	RSIPeriod   int
	BBPeriod    int
	BBStdDev    float64
	SMA10Period int
	SMA20Period int
	SMA50Period int
}

func DefaultConfig() Config {
	return Config{
		RSIPeriod:   14,
		BBPeriod:    20,
		BBStdDev:    2.0,
		SMA10Period: 10,
		SMA20Period: 20,
		SMA50Period: 50,
	}
}
