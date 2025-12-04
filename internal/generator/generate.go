package generator

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomChoice[T any](arr []T) T {
	return arr[rand.Intn(len(arr))]
}

func GenerateRandomStrategy() StrategyTemplate {

	entry := RandomEntry()
	confirmation := RandomConfirmation()
	baseline := RandomBaseline()
	volume := RandomVolume()
	exit := RandomExit()

	return StrategyTemplate{
		Entry:        entry,
		Confirmation: confirmation,
		Baseline:     baseline,
		Volume:       volume,
		Exit:         exit,
	}
}

func RandomEntry() IndicatorSpec {
	switch RandomChoice(EntryIndicators) {

	case "RSI":
		return IndicatorSpec{"RSI", map[string]any{"period": RandomChoice(RSI_Periods)}}

	case "Stochastic":
		return IndicatorSpec{"Stochastic", map[string]any{
			"k": RandomChoice(Stoch_K),
			"d": RandomChoice(Stoch_D),
		}}

	case "Momentum":
		return IndicatorSpec{"Momentum", map[string]any{
			"period": RandomChoice(Momentum_P),
		}}

	case "CCI":
		return IndicatorSpec{"CCI", map[string]any{"period": RandomChoice([]int{10, 14, 20})}}
	}

	return IndicatorSpec{}
}

func RandomConfirmation() IndicatorSpec {
	switch RandomChoice(ConfirmationIndicators) {

	case "ADX":
		return IndicatorSpec{"ADX", map[string]any{"period": RandomChoice(ADX_Periods)}}

	case "OBV":
		return IndicatorSpec{"OBV", map[string]any{}}

	case "VolumeSpike":
		return IndicatorSpec{"VolumeSpike", map[string]any{
			"period": RandomChoice(VolumeSpike_Periods),
		}}
	}

	return IndicatorSpec{}
}

func RandomBaseline() IndicatorSpec {
	switch RandomChoice(BaselineIndicators) {

	case "EMA":
		return IndicatorSpec{"EMA", map[string]any{"length": RandomChoice(EMA_Lengths)}}

	case "SMA":
		return IndicatorSpec{"SMA", map[string]any{"length": RandomChoice(SMA_Lengths)}}

	case "SuperTrend":
		return IndicatorSpec{"SuperTrend", map[string]any{
			"period": RandomChoice(SuperTrend_Periods),
			"mult":   RandomChoice(SuperTrend_Mult),
		}}
	}

	return IndicatorSpec{}
}

func RandomVolume() IndicatorSpec {
	switch RandomChoice(VolumeIndicators) {

	case "OBV":
		return IndicatorSpec{"OBV", map[string]any{}}

	case "VolumeSpike":
		return IndicatorSpec{"VolumeSpike", map[string]any{
			"period": RandomChoice(VolumeSpike_Periods),
		}}
	}

	return IndicatorSpec{}
}

func RandomExit() IndicatorSpec {
	return IndicatorSpec{"RSI", map[string]any{"period": RandomChoice([]int{7, 14})}}
}
