package indicators

import "quantengine/internal/models"

type SMA struct {
	Length int
	Window []float64
}

func NewSMA(length int) *SMA {
	return &SMA{
		Length: length,
		Window: make([]float64, 0, length),
	}
}

func (s *SMA) Reset() {
	s.Window = s.Window[:0]
}

func (s *SMA) Next(c models.Candle) models.IndicatorResult {
	price := c.Close

	s.Window = append(s.Window, price)
	if len(s.Window) > s.Length {
		s.Window = s.Window[1:]
	}

	sum := 0.0
	for _, v := range s.Window {
		sum += v
	}

	value := sum / float64(len(s.Window))

	return models.IndicatorResult{
		Value:  value,
		Signal: models.SignalNone,
	}
}
