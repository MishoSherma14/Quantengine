package indicators

import "quantengine/internal/models"

type OBV struct {
	prevClose float64
	value     float64
	started   bool
}

func NewOBV() *OBV {
	return &OBV{}
}

func (o *OBV) Reset() {
	o.started = false
	o.value = 0
}

func (o *OBV) Next(c models.Candle) models.IndicatorResult {
	if !o.started {
		o.prevClose = c.Close
		o.started = true
		return models.IndicatorResult{Value: 0, Signal: models.SignalNone}
	}

	if c.Close > o.prevClose {
		o.value += c.Volume
	} else if c.Close < o.prevClose {
		o.value -= c.Volume
	}

	o.prevClose = c.Close

	signal := models.SignalNone
	if o.value > 0 {
		signal = models.SignalBuy
	}
	if o.value < 0 {
		signal = models.SignalSell
	}

	return models.IndicatorResult{Value: o.value, Signal: signal}
}
