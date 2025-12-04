package generator

import "encoding/json"

func ToJSON(s StrategyTemplate) ([]byte, error) {
	str := map[string]any{
		"entry":        s.Entry,
		"confirmation": s.Confirmation,
		"baseline":     s.Baseline,
		"volume":       s.Volume,
		"exit":         s.Exit,
		"risk": map[string]any{
			"type": "ATR",
		},
	}
	return json.MarshalIndent(str, "", "  ")
}
