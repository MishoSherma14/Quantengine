package runner

import (
	"context"
	"encoding/json"
	"time"

	"quantengine/internal/backtester"
	"quantengine/internal/strategy"
	"quantengine/internal/data"
)

func RunJob(cfg *strategy.Config, symbol, tf string) (*ResultOutput, error) {
	ctx := context.Background()

	// load candles from GCS
	candles, err := data.LoadCandlesFromGCS(symbol, tf)
	if err != nil {
		return nil, err
	}

	// run backtest
	bt := backtester.NewBacktester(cfg, candles)
	result := bt.Run()

	// strategy JSON
	js, _ := json.Marshal(cfg)

	// compose output
	out := &ResultOutput{
		StrategyHash: cfg.Hash(),
		StrategyJSON: string(js),
		Symbol:       symbol,
		WinRate:      result.WinRate,
		ProfitFactor: result.ProfitFactor,
		MaxDrawdown:  result.MaxDrawdown,
		Score:        result.Score,
		EquityCurve:  result.EquityCurve,
		Timestamp:    time.Now(),
	}

	return out, nil
}
