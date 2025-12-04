package backtester

import (
	"quantengine/internal/models"
	"quantengine/internal/money"
	"quantengine/internal/strategy"
)

type Engine struct {
	Strat      strategy.Strategy
	ATR        indicators.Indicator
	ATRManager *money.ATRManager

	Position *Position
	Equity   float64
	Curve    []float64
}

func NewEngine(strat strategy.Strategy, initialEquity float64, atr indicators.Indicator) *Engine {
	return &Engine{
		Strat:      strat,
		ATR:        atr,
		ATRManager: money.NewATRManager(initialEquity),
		Equity:     initialEquity,
		Curve:      []float64{initialEquity},
	}
}

func (e *Engine) Run(candles []models.Candle) models.BacktestResult {
	for _, c := range candles {
		e.tick(c)
	}

	return models.BacktestResult{
		EquityCurve: e.Curve,
	}
}

func (e *Engine) tick(c models.Candle) {
	// 1. Calculate ATR
	atrOut := e.ATR.Next(c)
	atr := atrOut.Value

	sig := e.Strat.Evaluate(c)

	// Exit logic first
	if e.Position != nil && e.Position.IsOpen {
		// Stop Loss
		if e.Position.CheckStop(c.Close) {
			pnl := e.closePos(c.Close)
			e.updateEquity(pnl)
			return
		}

		// TP1 → partial close
		if e.Position.CheckTarget1(c.Close) {
			pnl := e.partialClose(c.Close)
			e.updateEquity(pnl)
			e.Position.ActivateTrailing(atr, c.Close)
		}

		// Trailing update
		e.Position.UpdateTrailing(atr, c.Close)

		// Trailing Exit
		if e.Position.CheckTrailingHit(c.Close) {
			pnl := e.closePos(c.Close)
			e.updateEquity(pnl)
			return
		}
	}

	// If it's open do not open again
	if e.Position != nil && e.Position.IsOpen {
		e.updateEquity(0)
		return
	}

	// ENTRY logic
	if sig.EntrySignal == models.SignalBuy &&
		sig.ConfirmationSignal == models.SignalBuy &&
		sig.BaselineSignal == models.SignalBuy &&
		sig.VolumeSignal == models.SignalBuy {

		e.openPos(true, atr, c.Close)
	}

	if sig.EntrySignal == models.SignalSell &&
		sig.ConfirmationSignal == models.SignalSell &&
		sig.BaselineSignal == models.SignalSell &&
		sig.VolumeSignal == models.SignalSell {

		e.openPos(false, atr, c.Close)
	}

	e.updateEquity(0)
}

func (e *Engine) openPos(isLong bool, atr float64, price float64) {
	size := e.ATRManager.PositionSize(atr)
	stop := e.ATRManager.StopLoss(price, atr, isLong)
	tp1 := e.ATRManager.Target1(price, atr, isLong)

	e.Position = &Position{
		IsOpen:     true,
		IsLong:     isLong,
		EntryPrice: price,
		Size:       size,
		StopLoss:   stop,
		Target1:    tp1,
	}
}

func (e *Engine) partialClose(price float64) float64 {
	e.Position.PartialClosed = true
	half := e.Position.Size * 0.5

	if e.Position.IsLong {
		return (price - e.Position.EntryPrice) * half
	}
	return (e.Position.EntryPrice - price) * half
}

func (e *Engine) closePos(price float64) float64 {
	full := e.Position.Size * 0.5

	if e.Position.IsLong {
		pnl := (price - e.Position.EntryPrice) * full
		e.Position.IsOpen = false
		return pnl
	}

	pnl := (e.Position.EntryPrice - price) * full
	e.Position.IsOpen = false
	return pnl
}

func (e *Engine) updateEquity(pnl float64) {
	e.Equity += pnl
	e.Curve = append(e.Curve, e.Equity)
}
