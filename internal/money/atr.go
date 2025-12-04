package money

type ATRManager struct {
	Balance         float64
	RiskPercent     float64 // e.g. 0.02
	ATRMultiplierSL float64 // 1.5
	ATRMultiplierTP float64 // 3.0
}

func NewATRManager(balance float64) *ATRManager {
	return &ATRManager{
		Balance:         balance,
		RiskPercent:     0.02,
		ATRMultiplierSL: 1.5,
		ATRMultiplierTP: 3.0,
	}
}

func (a *ATRManager) PositionSize(atr float64) float64 {
	possibleLose := a.Balance * a.RiskPercent
	stopDistance := atr * a.ATRMultiplierSL
	return possibleLose / stopDistance // quantity
}

func (a *ATRManager) StopLoss(entry float64, atr float64, isLong bool) float64 {
	if isLong {
		return entry - atr*a.ATRMultiplierSL
	}
	return entry + atr*a.ATRMultiplierSL
}

func (a *ATRManager) Target1(entry float64, atr float64, isLong bool) float64 {
	if isLong {
		return entry + atr*a.ATRMultiplierTP
	}
	return entry - atr*a.ATRMultiplierTP
}
