package worker

import (
	"context"
	"encoding/json"
	"log"

	"quantengine/internal/data"
	"quantengine/internal/backtester"
	"quantengine/internal/strategy"
	"quantengine/internal/models"
)

type TaskMessage struct {
	Symbol    string          `json:"symbol"`
	Timeframe string          `json:"timeframe"`
	Config    json.RawMessage `json:"strategy"`
}

func HandleTask(ctx context.Context, m *TaskMessage) error {
	// 1) Load candles from GCS
	candles, err := data.LoadCandlesFromGCS(m.Symbol, m.Timeframe)
	if err != nil {
		log.Println("GCS load error:", err)
		return err
	}

	// 2) Parse strategy config
	cfg, err := strategy.FromJSON(m.Config)
	if err != nil {
		log.Println("Strategy parse error:", err)
		return err
	}

	// 3) Run backtest
	result := backtester.Run(candles, cfg)

	// 4) Convert for BigQuery
	out := models.ResultOutput{
		StrategyHash: cfg.Hash(),
		StrategyJSON: string(m.Config),
		Symbol:       m.Symbol,
		WinRate:      result.WinRate,
		ProfitFactor: result.ProfitFactor,
		MaxDrawdown:  result.MaxDrawdown,
		Score:        result.Score(),
		EquityCurve:  result.EquityJSON(),
	}

	// 5) Save to BigQuery
	err = data.SaveToBigQuery(ctx, &out)
	if err != nil {
		log.Println("BQ save error:", err)
		return err
	}

	return nil
}
