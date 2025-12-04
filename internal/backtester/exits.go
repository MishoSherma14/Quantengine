package backtester

func (p *Position) CheckStop(c float64) bool {
	if !p.IsOpen {
		return false
	}
	if p.IsLong && c <= p.StopLoss {
		return true
	}
	if !p.IsLong && c >= p.StopLoss {
		return true
	}
	return false
}

func (p *Position) CheckTarget1(c float64) bool {
	if p.PartialClosed {
		return false
	}

	if p.IsLong && c >= p.Target1 {
		return true
	}
	if !p.IsLong && c <= p.Target1 {
		return true
	}
	return false
}

func (p *Position) ActivateTrailing(atr float64, close float64) {
	if !p.TrailingOn {
		p.TrailingOn = true
		if p.IsLong {
			p.TrailingStop = close - atr*1.5
		} else {
			p.TrailingStop = close + atr*1.5
		}
	}
}

func (p *Position) UpdateTrailing(atr float64, close float64) {
	if !p.TrailingOn {
		return
	}

	if p.IsLong {
		newTS := close - atr*1.5
		if newTS > p.TrailingStop {
			p.TrailingStop = newTS
		}
	} else {
		newTS := close + atr*1.5
		if newTS < p.TrailingStop {
			p.TrailingStop = newTS
		}
	}
}

func (p *Position) CheckTrailingHit(c float64) bool {
	if !p.TrailingOn {
		return false
	}

	if p.IsLong && c <= p.TrailingStop {
		return true
	}
	if !p.IsLong && c >= p.TrailingStop {
		return true
	}
	return false
}
