package main

import (
    "context"
    "encoding/json"
    "log"

    "cloud.google.com/go/pubsub"
    "quantengine/internal/backtester"
    "quantengine/internal/strategy"
    "quantengine/internal/data"
)

type TaskMessage struct {
    Strategy  json.RawMessage `json:"strategy"`
    Symbol    string          `json:"symbol"`
    Timeframe string          `json:"timeframe"`
}

func main() {
    ctx := context.Background()
    log.Println("Worker Started")

    // Pub/Sub subscriber
    client, err := pubsub.NewClient(ctx, "quantengine")
    if err != nil {
        log.Fatal(err)
    }

    sub := client.Subscription("strategy-tasks-sub")

    sub.ReceiveSettings.MaxOutstandingMessages = 10
    sub.ReceiveSettings.NumGoroutines = 5

    err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
        var task TaskMessage
        if err := json.Unmarshal(msg.Data, &task); err != nil {
            log.Println("Bad msg:", err)
            msg.Nack()
            return
        }

        // Load candles
        candles, err := data.LoadCandles(task.Symbol)
        if err != nil {
            log.Println("load error:", err)
            msg.Nack()
            return
        }

        // Parse strategy config
        strat, err := strategy.FromJSON(task.Strategy)
        if err != nil {
            log.Println("strat error:", err)
            msg.Nack()
            return
        }

        // Run backtest
        result := backtester.Run(candles, strat)

        // TODO: write result to BigQuery

        msg.Ack()
    })

    if err != nil {
        log.Fatal(err)
    }
}
