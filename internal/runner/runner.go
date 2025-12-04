package runner

import (
	"fmt"
	"quantengine/internal/backtester"
	"quantengine/internal/data"
)

func RunBacktest(p BacktestPayload) (*backtester.Result, error) {
	candles, err := data.LoadCandlesFromGCS(p.Symbol, p.Timeframe)
	if err != nil {
		return nil, fmt.Errorf("load candles: %w", err)
	}

	strategy := backtester.BuildStrategyFromJSON(p.Strategy)

	result := backtester.Run(candles, strategy)

	return result, nil
}

