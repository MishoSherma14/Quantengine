for _, strat := range strategies {
    for _, symbol := range runner.Markets {
        msg := map[string]any{
            "strategy": json.RawMessage(strat),
            "symbol":   symbol,
            "timeframe": runner.Timeframe,
        }

        data, _ := json.Marshal(msg)
        topic.Publish(ctx, &pubsub.Message{Data: data})
    }
}
