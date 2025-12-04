package strategy

import "encoding/json"

type IndicatorConfig struct {
	Name   string          `json:"name"`
	Params json.RawMessage `json:"params"`
}

type StrategyConfig struct {
	Entry        IndicatorConfig `json:"entry"`
	Confirmation IndicatorConfig `json:"confirmation"`
	Baseline     IndicatorConfig `json:"baseline"`
	Volume       IndicatorConfig `json:"volume"`
	Exit         IndicatorConfig `json:"exit"`
	Risk         json.RawMessage `json:"risk"`
}
