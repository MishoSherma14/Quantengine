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

func RunBacktest(strategy *strategy.Config, symbol string, candles []models.Candle) (*ResultOutput, error) {
	result := backtester.Run(strategy, candles)

	out := &ResultOutput{
		StrategyHash: strategy.Hash(),
		StrategyJSON: strategy.ToJSON(),
		Symbol:       symbol,
		WinRate:      result.WinRate,
		ProfitFactor: result.ProfitFactor,
		MaxDrawdown:  result.MaxDrawdown,
		Sharpe:       result.Sharpe,
		AvgReturn:    result.AvgReturn,
		Score:        result.Score(),  // სპეციალურად ავაგებთ scoring მოდელს Phase 9-ში
		EquityCurve:  result.EquityCurve,
		Timestamp:    time.Now(),
	}

	return out, nil
}
