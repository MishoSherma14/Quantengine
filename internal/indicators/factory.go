package indicators

import (
	"encoding/json"
	"fmt"
)

func CreateIndicator(name string, params json.RawMessage) (Indicator, error) {

	switch name {

	case "EMA":
		var p struct{ Length int }
		json.Unmarshal(params, &p)
		return NewEMA(p.Length), nil

	case "SMA":
		var p struct{ Length int }
		json.Unmarshal(params, &p)
		return NewSMA(p.Length), nil

	case "RSI":
		var p struct{ Period int }
		json.Unmarshal(params, &p)
		return NewRSI(p.Period), nil

	case "Momentum":
		var p struct{ Period int }
		json.Unmarshal(params, &p)
		return NewMomentum(p.Period), nil

	case "CCI":
		var p struct{ Period int }
		json.Unmarshal(params, &p)
		return NewCCI(p.Period), nil

	// SUPER TREND, ADX, STOCHASTIC, OBV, VOLUME SPIKE HERE LATER

	default:
		return nil, fmt.Errorf("unknown indicator: %s", name)
	}
}
