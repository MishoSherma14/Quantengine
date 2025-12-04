package indicators

import "quantengine/internal/models"

type Momentum struct {
	Period int
	window []float64
}

func NewMomentum(period int) *Momentum {
	return &Momentum{
		Period: period,
		window: make([]float64, 0, period),
	}
}

func (m *Momentum) Reset() {
	m.window = m.window[:0]
}

func (m *Momentum) Next(c models.Candle) models.IndicatorResult {
	price := c.Close

	m.window = append(m.window, price)
	if len(m.window) > m.Period {
		m.window = m.window[1:]
	}

	if len(m.window) < m.Period {
		return models.IndicatorResult{Value: 0}
	}

	return models.IndicatorResult{
		Value: price - m.window[0],
	}
}
