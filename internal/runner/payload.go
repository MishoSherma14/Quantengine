package runner

type BacktestPayload struct {
	Strategy  map[string]interface{} `json:"strategy"`
	Symbol    string                 `json:"symbol"`
	Timeframe string                 `json:"timeframe"`
}

