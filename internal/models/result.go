package models

type BacktestResult struct {
	TotalTrades  int
	WinRate      float64
	MaxDrawdown  float64
	ProfitFactor float64
	Expectancy   float64
	EquityCurve  []float64
}
