package indicators

import "quantengine/internal/models"

type VolumeSpike struct {
	Period int
	window []float64
}

func NewVolumeSpike(period int) *VolumeSpike {
	return &VolumeSpike{
		Period: period,
		window: make([]float64, 0, period),
	}
}

func (v *VolumeSpike) Reset() {
	v.window = v.window[:0]
}

func (v *VolumeSpike) Next(c models.Candle) models.IndicatorResult {
	v.window = append(v.window, c.Volume)
	if len(v.window) > v.Period {
		v.window = v.window[1:]
	}

	avgVol := avg(v.window)

	signal := models.SignalNone
	if c.Volume > avgVol*1.5 {
		signal = models.SignalBuy
	}

	return models.IndicatorResult{
		Value:  c.Volume,
		Signal: signal,
	}
}
