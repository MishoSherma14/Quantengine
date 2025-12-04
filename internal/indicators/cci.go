package indicators

import "quantengine/internal/models"

type CCI struct {
	Period int
	window []float64
}

func NewCCI(period int) *CCI {
	return &CCI{
		Period: period,
		window: make([]float64, 0, period),
	}
}

func (c *CCI) Reset() {
	c.window = c.window[:0]
}

func (c *CCI) Next(candle models.Candle) models.IndicatorResult {
	tp := (candle.High + candle.Low + candle.Close) / 3

	c.window = append(c.window, tp)
	if len(c.window) > c.Period {
		c.window = c.window[1:]
	}

	if len(c.window) < c.Period {
		return models.IndicatorResult{Value: 0}
	}

	sum := 0.0
	for _, v := range c.window {
		sum += v
	}
	sma := sum / float64(len(c.window))

	md := 0.0
	for _, v := range c.window {
		md += abs(v - sma)
	}
	md = md / float64(len(c.window))

	cci := (tp - sma) / (0.015 * md)

	return models.IndicatorResult{Value: cci}
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}
