type ResultOutput struct {
    StrategyHash string  `bigquery:"strategy_hash"`
    StrategyJSON string  `bigquery:"strategy_json"`
    Symbol       string  `bigquery:"symbol"`
    WinRate      float64 `bigquery:"win_rate"`
    ProfitFactor float64 `bigquery:"profit_factor"`
    MaxDD        float64 `bigquery:"max_drawdown"`
    Sharpe       float64 `bigquery:"sharpe"`
    AvgReturn    float64 `bigquery:"avg_return"`
    Score        float64 `bigquery:"score"`
    EquityCurve  string  `bigquery:"equity_curve"`
    TS           int64   `bigquery:"ts"`
}
