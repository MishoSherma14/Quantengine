package storage

import (
	"context"
	"cloud.google.com/go/bigquery"
)

type ResultOutput struct {
	StrategyHash string  `bigquery:"strategy_hash"`
	StrategyJSON string  `bigquery:"strategy_json"`
	Symbol       string  `bigquery:"symbol"`
	WinRate      float64 `bigquery:"win_rate"`
	ProfitFactor float64 `bigquery:"profit_factor"`
	MaxDD        float64 `bigquery:"max_drawdown"`
	Sharpe       float64 `bigquery:"sharpe"`
	AvgReturn    float64 `bigquery:"avg_return"`
	Score        float64 `bigquery:"score"`
	EquityCurve  string  `bigquery:"equity_curve"`
	Timestamp    string  `bigquery:"ts"`
}

func SaveToBigQuery(ctx context.Context, r *ResultOutput) error {
	client, err := bigquery.NewClient(ctx, "quantengine")
	if err != nil {
		return err
	}
	defer client.Close()

	u := client.Dataset("quant_results").Table("backtest_results").Uploader()
	return u.Put(ctx, r)
}
