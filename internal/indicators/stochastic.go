package indicators

import (
	"quantengine/internal/models"
)

type Stochastic struct {
	KPeriod int
	DPeriod int

	highWindow []float64
	lowWindow  []float64
	kValues    []float64
}

func NewStochastic(k, d int) *Stochastic {
	return &Stochastic{
		KPeriod:    k,
		DPeriod:    d,
		highWindow: make([]float64, 0, k),
		lowWindow:  make([]float64, 0, k),
		kValues:    make([]float64, 0, d),
	}
}

func (s *Stochastic) Reset() {
	s.highWindow = s.highWindow[:0]
	s.lowWindow = s.lowWindow[:0]
	s.kValues = s.kValues[:0]
}

func (s *Stochastic) Next(c models.Candle) models.IndicatorResult {
	s.highWindow = append(s.highWindow, c.High)
	s.lowWindow = append(s.lowWindow, c.Low)

	if len(s.highWindow) > s.KPeriod {
		s.highWindow = s.highWindow[1:]
		s.lowWindow = s.lowWindow[1:]
	}

	if len(s.highWindow) < s.KPeriod {
		return models.IndicatorResult{Value: 50, Signal: models.SignalNone}
	}

	highest := maxSlice(s.highWindow)
	lowest := minSlice(s.lowWindow)

	if highest == lowest {
		return models.IndicatorResult{Value: 50, Signal: models.SignalNone}
	}

	k := ((c.Close - lowest) / (highest - lowest)) * 100

	s.kValues = append(s.kValues, k)
	if len(s.kValues) > s.DPeriod {
		s.kValues = s.kValues[1:]
	}

	d := avg(s.kValues)

	signal := models.SignalNone

	if d < 20 {
		signal = models.SignalBuy
	}
	if d > 80 {
		signal = models.SignalSell
	}

	return models.IndicatorResult{
		Value:  d,
		Signal: signal,
	}
}

func maxSlice(v []float64) float64 {
	max := v[0]
	for _, n := range v {
		if n > max {
			max = n
		}
	}
	return max
}

func minSlice(v []float64) float64 {
	min := v[0]
	for _, n := range v {
		if n < min {
			min = n
		}
	}
	return min
}

func avg(v []float64) float64 {
	s := 0.0
	for _, x := range v {
		s += x
	}
	return s / float64(len(v))
}
