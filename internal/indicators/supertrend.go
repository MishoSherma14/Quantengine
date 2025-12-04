package indicators

import "quantengine/internal/models"

type SuperTrend struct {
	Period     int
	Multiplier float64

	atr       Indicator
	prevUpper float64
	prevLower float64
	trendUp   bool
	started   bool
}

func NewSuperTrend(period int, mult float64) *SuperTrend {
	return &SuperTrend{
		Period:     period,
		Multiplier: mult,
		atr:        NewATR(period),
	}
}

func (st *SuperTrend) Reset() {
	st.started = false
	st.prevUpper = 0
	st.prevLower = 0
	st.trendUp = true
	st.atr.Reset()
}

func (st *SuperTrend) Next(c models.Candle) models.IndicatorResult {
	atrOut := st.atr.Next(c)
	atr := atrOut.Value

	hl2 := (c.High + c.Low) / 2

	upper := hl2 + st.Multiplier*atr
	lower := hl2 - st.Multiplier*atr

	if !st.started {
		st.prevUpper = upper
		st.prevLower = lower
		st.started = true
		return models.IndicatorResult{Value: hl2, Signal: models.SignalNone}
	}

	if upper < st.prevUpper {
		st.prevUpper = upper
	}
	if lower > st.prevLower {
		st.prevLower = lower
	}

	var signal models.Signal

	if c.Close > st.prevUpper {
		st.trendUp = true
		signal = models.SignalBuy
		st.prevLower = lower
	} else if c.Close < st.prevLower {
		st.trendUp = false
		signal = models.SignalSell
		st.prevUpper = upper
	}

	return models.IndicatorResult{
		Value:  hl2,
		Signal: signal,
	}
}
