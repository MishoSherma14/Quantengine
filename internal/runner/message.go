package runner

type TaskMessage struct {
    StrategyJSON string `json:"strategy"`
    Symbol       string `json:"symbol"`
    Timeframe    string `json:"timeframe"`
}
