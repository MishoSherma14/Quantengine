package runner

import "time"

type ResultOutput struct {
	StrategyHash   string    `bigquery:"strategy_hash"`
	StrategyJSON   string    `bigquery:"strategy_json"`
	Symbol         string    `bigquery:"symbol"`
	WinRate        float64   `bigquery:"win_rate"`
	ProfitFactor   float64   `bigquery:"profit_factor"`
	MaxDrawdown    float64   `bigquery:"max_drawdown"`
	Sharpe         float64   `bigquery:"sharpe"`
	AvgReturn      float64   `bigquery:"avg_return"`
	Score          float64   `bigquery:"score"`
	EquityCurve    []float64 `bigquery:"equity_curve"`
	Timestamp      time.Time `bigquery:"ts"`
}
