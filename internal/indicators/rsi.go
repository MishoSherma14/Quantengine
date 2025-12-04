package indicators

import "quantengine/internal/models"

type RSI struct {
	Period      int
	gains       []float64
	losses      []float64
	prevClose   float64
	initialized bool
}

func NewRSI(period int) *RSI {
	return &RSI{
		Period: period,
		gains:  make([]float64, 0, period),
		losses: make([]float64, 0, period),
	}
}

func (r *RSI) Reset() {
	r.initialized = false
	r.gains = r.gains[:0]
	r.losses = r.losses[:0]
}

func (r *RSI) Next(c models.Candle) models.IndicatorResult {
	if !r.initialized {
		r.prevClose = c.Close
		r.initialized = true
		return models.IndicatorResult{Value: 50}
	}

	change := c.Close - r.prevClose
	r.prevClose = c.Close

	if change >= 0 {
		r.gains = append(r.gains, change)
		r.losses = append(r.losses, 0)
	} else {
		r.gains = append(r.gains, 0)
		r.losses = append(r.losses, -change)
	}

	if len(r.gains) > r.Period {
		r.gains = r.gains[1:]
		r.losses = r.losses[1:]
	}

	if len(r.gains) < r.Period {
		return models.IndicatorResult{Value: 50}
	}

	avgGain := sum(r.gains) / float64(r.Period)
	avgLoss := sum(r.losses) / float64(r.Period)

	var rsi float64
	if avgLoss == 0 {
		rsi = 100
	} else {
		rs := avgGain / avgLoss
		rsi = 100 - (100 / (1 + rs))
	}

	return models.IndicatorResult{Value: rsi}
}

func sum(v []float64) float64 {
	total := 0.0
	for _, n := range v {
		total += n
	}
	return total
}
