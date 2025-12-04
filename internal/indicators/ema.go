package indicators

import "quantengine/internal/models"

type EMA struct {
	Length  int
	Mult    float64
	Value   float64
	Started bool
}

func NewEMA(length int) *EMA {
	mult := 2.0 / float64(length+1)
	return &EMA{
		Length: length,
		Mult:   mult,
	}
}

func (e *EMA) Reset() {
	e.Started = false
	e.Value = 0
}

func (e *EMA) Next(c models.Candle) models.IndicatorResult {
	price := c.Close

	if !e.Started {
		e.Value = price
		e.Started = true
		return models.IndicatorResult{Value: e.Value, Signal: models.SignalNone}
	}

	e.Value = (price-e.Value)*e.Mult + e.Value

	return models.IndicatorResult{
		Value:  e.Value,
		Signal: models.SignalNone,
	}
}
