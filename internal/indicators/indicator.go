package indicators

import "quantengine/internal/models"

type Indicator interface {
	Next(c models.Candle) models.IndicatorResult
	Reset()
}
