package strategy

import (
	"quantengine/internal/indicators"
	"quantengine/internal/models"
)

// SimpleStrategy — strategy that combines entry + confirmation + baseline + volume + exit signals
type SimpleStrategy struct {
	entry        indicators.Indicator
	confirmation indicators.Indicator
	baseline     indicators.Indicator
	volume       indicators.Indicator
	exitInd      indicators.Indicator
}

// BuildStrategy — creates indicator instances from StrategyConfig JSON
func BuildStrategy(cfg StrategyConfig) (SimpleStrategy, error) {

	entry, err := indicators.CreateIndicator(cfg.Entry.Name, cfg.Entry.Params)
	if err != nil {
		return SimpleStrategy{}, err
	}

	confirmation, err := indicators.CreateIndicator(cfg.Confirmation.Name, cfg.Confirmation.Params)
	if err != nil {
		return SimpleStrategy{}, err
	}

	baseline, err := indicators.CreateIndicator(cfg.Baseline.Name, cfg.Baseline.Params)
	if err != nil {
		return SimpleStrategy{}, err
	}

	volume, err := indicators.CreateIndicator(cfg.Volume.Name, cfg.Volume.Params)
	if err != nil {
		return SimpleStrategy{}, err
	}

	exitInd, err := indicators.CreateIndicator(cfg.Exit.Name, cfg.Exit.Params)
	if err != nil {
		return SimpleStrategy{}, err
	}

	return SimpleStrategy{
		entry:        entry,
		confirmation: confirmation,
		baseline:     baseline,
		volume:       volume,
		exitInd:      exitInd,
	}, nil
}

// Evaluate — runs all indicators and returns combined signal bundle
func (s SimpleStrategy) Evaluate(c models.Candle) SignalBundle {

	entryRes := s.entry.Next(c)
	confirmationRes := s.confirmation.Next(c)
	baselineRes := s.baseline.Next(c)
	volumeRes := s.volume.Next(c)
	exitRes := s.exitInd.Next(c)

	return SignalBundle{
		EntrySignal:        entryRes.Signal,
		ConfirmationSignal: confirmationRes.Signal,
		BaselineSignal:     baselineRes.Signal,
		VolumeSignal:       volumeRes.Signal,
		ExitSignal:         exitRes.Signal,
	}
}
