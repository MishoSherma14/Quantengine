import (
    "context"
    "encoding/json"
    "log"

    "quantengine/internal/data"
    "quantengine/internal/runner"
    "quantengine/internal/backtester"
)

func main() {
    ctx := context.Background()
    log.Println("Worker started...")

    msg, err := runner.PullMessage(ctx)  // Pub/Sub
    if err != nil {
        log.Fatal(err)
    }

    strategyJSON := msg.StrategyJSON
    strategy := runner.ParseStrategy(strategyJSON)

    candles, err := data.LoadCandlesFromGCS(msg.Symbol, msg.Timeframe)
    if err != nil {
        log.Fatal(err)
    }

    engine := backtester.NewEngine(strategy, msg.Symbol, candles)
    engine.Run()

    result := engine.BuildResultOutput(strategyJSON, candles)

    err = runner.SaveToBigQuery(ctx, result)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Result saved:", result.StrategyHash)
}
