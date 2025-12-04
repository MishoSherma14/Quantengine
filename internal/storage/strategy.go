package strategy

import (
	"encoding/json"
)

func FromJSON(data []byte) (*Strategy, error) {
	var s Strategy
	err := json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
