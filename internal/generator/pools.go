package generator

// Entry Indicators
var EntryIndicators = []string{
	"RSI",
	"Stochastic",
	"Momentum",
	"CCI",
}

// Confirmation Indicators
var ConfirmationIndicators = []string{
	"ADX",
	"OBV",
	"VolumeSpike",
}

// Baseline Indicators
var BaselineIndicators = []string{
	"EMA",
	"SMA",
	"SuperTrend",
}

// Volume Indicators
var VolumeIndicators = []string{
	"OBV",
	"VolumeSpike",
}

// Exit Indicators (simple approach)
var ExitIndicators = []string{
	"RSI",
	"CCI",
	"Stochastic",
}

// ========== PARAMETER POOLS ==========

// RSI
var RSI_Periods = []int{7, 10, 14, 21, 28}

// Stochastic
var Stoch_K = []int{9, 12, 14}
var Stoch_D = []int{3, 5}

// Momentum
var Momentum_P = []int{5, 10, 14, 20}

// CCI
var CCI_Periods = []int{10, 14, 20}

// EMA / SMA
var EMA_Lengths = []int{20, 50, 100, 200}
var SMA_Lengths = []int{20, 50, 100, 200}

// SuperTrend
var SuperTrend_Periods = []int{7, 10, 14}
var SuperTrend_Mult = []float64{1.5, 2.0, 3.0}

// ADX
var ADX_Periods = []int{10, 14, 20}

// Volume Spike
var VolumeSpike_Periods = []int{10, 20}
