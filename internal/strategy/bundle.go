package strategy

import "quantengine/internal/models"

type SignalBundle struct {
	EntrySignal        models.Signal
	ConfirmationSignal models.Signal
	BaselineSignal     models.Signal
	VolumeSignal       models.Signal
	ExitSignal         models.Signal
}
