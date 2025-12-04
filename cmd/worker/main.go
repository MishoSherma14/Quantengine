package main

import (
	"context"
	"encoding/json"
	"log"

	"quantengine/internal/data"
	"quantengine/internal/runner"
	"quantengine/internal/strategy"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data"`
	} `json:"message"`
}

type TaskInput struct {
	Strategy  json.RawMessage `json:"strategy"`
	Symbol    string          `json:"symbol"`
	Timeframe string          `json:"timeframe"`
}

func main() {
	log.Println("Worker started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		var m PubSubMessage
		json.NewDecoder(r.Body).Decode(&m)

		var task TaskInput
		json.Unmarshal(m.Message.Data, &task)

		// Load candles
		candles, err := data.LoadCandlesFromGCS(task.Symbol, task.Timeframe)
		if err != nil {
			log.Println("LoadCandles error:", err)
			return
		}

		// Build strategy
		conf, err := strategy.FromJSON(task.Strategy)
		if err != nil {
			log.Println("strategy parse error:", err)
			return
		}

		// Run backtest
		out, err := runner.RunBacktest(conf, task.Symbol, candles)
		if err != nil {
			log.Println("backtest error:", err)
			return
		}

		// Save result to BQ
		err = data.SaveToBigQuery(ctx, out)
		if err != nil {
			log.Println("BQ save error:", err)
			return
		}

		log.Println("Backtest OK:", task.Symbol)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
