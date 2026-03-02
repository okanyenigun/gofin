package indicators

type BollingerBands struct {
	Upper  []float64
	Middle []float64
	Lower  []float64
}

type MACD struct {
	Line      []float64
	Signal    []float64
	Histogram []float64
}

type Stochastic struct {
	K []float64
	D []float64
}

type Results struct {
	RSI            []float64
	BollingerBands BollingerBands
	SMA10          []float64
	SMA20          []float64
	SMA50          []float64
	MACD           MACD
	Stochastic     Stochastic
}

type Config struct {
	RSIPeriod         int
	BBPeriod          int
	BBStdDev          float64
	SMA10Period       int
	SMA20Period       int
	SMA50Period       int
	MACDFastPeriod    int
	MACDSlowPeriod    int
	MACDSignalPeriod  int
	StochPeriod       int
	StochSmoothPeriod int
}

func DefaultConfig() Config {
	return Config{
		RSIPeriod:         14,
		BBPeriod:          20,
		BBStdDev:          2.0,
		SMA10Period:       10,
		SMA20Period:       20,
		SMA50Period:       50,
		MACDFastPeriod:    12,
		MACDSlowPeriod:    26,
		MACDSignalPeriod:  9,
		StochPeriod:       14,
		StochSmoothPeriod: 3,
	}
}
