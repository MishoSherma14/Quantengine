package indicators

import "quantengine/internal/models"

type ADX struct {
	Period int

	trIndicator Indicator
	dmPlus      Indicator
	dmMinus     Indicator
	adxValues   []float64
}

func NewADX(period int) *ADX {
	return &ADX{
		Period:      period,
		trIndicator: NewTR(),
		dmPlus:      NewDMPlus(),
		dmMinus:     NewDMMinus(),
	}
}

func (a *ADX) Reset() {
	a.trIndicator.Reset()
	a.dmPlus.Reset()
	a.dmMinus.Reset()
	a.adxValues = a.adxValues[:0]
}

func (a *ADX) Next(c models.Candle) models.IndicatorResult {
	tr := a.trIndicator.Next(c).Value
	plusDM := a.dmPlus.Next(c).Value
	minusDM := a.dmMinus.Next(c).Value

	if tr == 0 {
		return models.IndicatorResult{Value: 0, Signal: models.SignalNone}
	}

	plusDI := (plusDM / tr) * 100
	minusDI := (minusDM / tr) * 100

	diff := abs(plusDI - minusDI)
	sum := plusDI + minusDI
	if sum == 0 {
		return models.IndicatorResult{Value: 20, Signal: models.SignalNone}
	}

	dx := (diff / sum) * 100

	a.adxValues = append(a.adxValues, dx)
	if len(a.adxValues) > a.Period {
		a.adxValues = a.adxValues[1:]
	}

	if len(a.adxValues) < a.Period {
		return models.IndicatorResult{Value: 20, Signal: models.SignalNone}
	}

	adx := avg(a.adxValues)

	signal := models.SignalNone
	if adx > 25 {
		if plusDI > minusDI {
			signal = models.SignalBuy
		} else {
			signal = models.SignalSell
		}
	}

	return models.IndicatorResult{Value: adx, Signal: signal}
}
