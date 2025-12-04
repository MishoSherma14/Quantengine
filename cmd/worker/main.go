package main

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"quantengine/internal/runner"
	"quantengine/internal/strategy"
)

type JobMessage struct {
	Strategy json.RawMessage `json:"strategy"`
	Symbol   string          `json:"symbol"`
	Timeframe string         `json:"timeframe"`
}

func main() {
	ctx := context.Background()

	// Pub/Sub subscriber client
	client, err := pubsub.NewClient(ctx, "quantengine")
	if err != nil {
		log.Fatal(err)
	}

	sub := client.Subscription("strategy-tasks-sub")

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()

		var m JobMessage
		if err := json.Unmarshal(msg.Data, &m); err != nil {
			log.Println("decode err:", err)
			return
		}

		// parse strategy
		cfg, err := strategy.FromJSON(m.Strategy)
		if err != nil {
			log.Println("strategy parse err:", err)
			return
		}

		// run backtest
		out, err := runner.RunJob(cfg, m.Symbol, m.Timeframe)
		if err != nil {
			log.Println("run err:", err)
			return
		}

		// save to BigQuery
		if err := runner.SaveToBigQuery(ctx, out); err != nil {
			log.Println("bq err:", err)
			return
		}

		log.Println("DONE:", out.StrategyHash, m.Symbol)
	})

	if err != nil {
		log.Fatal(err)
	}
}
