package strategy

import "quantengine/internal/models"

type Strategy interface {
	Evaluate(c models.Candle) SignalBundle
}
