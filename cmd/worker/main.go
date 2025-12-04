package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"quantengine/internal/data"
	"quantengine/internal/runner"
	"quantengine/internal/strategy"
)

type TaskMessage struct {
	Strategy  json.RawMessage `json:"strategy"`
	Symbol    string          `json:"symbol"`
	Timeframe string          `json:"timeframe"`
}

func main() {
	fmt.Println("Worker Booted...")

	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		panic(err)
	}

	sub := pubsubClient.Subscription(os.Getenv("PUBSUB_SUBSCRIPTION"))

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered:", r)
			}
		}()

		var t TaskMessage
		if err := json.Unmarshal(msg.Data, &t); err != nil {
			log.Println("JSON decode error:", err)
			msg.Nack()
			return
		}

		// Load candles from GCS
		candles, err := data.LoadCandlesFromGCS(t.Symbol, t.Timeframe)
		if err != nil {
			log.Println("Load candles error:", err)
			msg.Nack()
			return
		}

		// Parse strategy config
		conf, err := strategy.FromJSON(t.Strategy)
		if err != nil {
			log.Println("Strategy decode error:", err)
			msg.Nack()
			return
		}

		// Execute backtest
		result := runner.Run(conf, candles)

		// Save results
		err = runner.SaveToBigQuery(ctx, result)
		if err != nil {
			log.Println("BQ Save error:", err)
			msg.Nack()
			return
		}

		msg.Ack()
	})

	if err != nil {
		log.Fatal(err)
	}
}
