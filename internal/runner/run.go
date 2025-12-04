package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"quantengine/internal/backtester"
	"quantengine/internal/data"
	"quantengine/internal/strategy"
)

// PubSubMessage — Cloud Run Pub/Sub trigger input
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// Incoming task structure
type Task struct {
	Symbol    string          `json:"symbol"`
	Timeframe string          `json:"timeframe"`
	Strategy  json.RawMessage `json:"strategy"`
}

// Main runner function: triggered per Pub/Sub message
func Run(ctx context.Context, m PubSubMessage) error {
	var task Task
	if err := json.Unmarshal(m.Data, &task); err != nil {
		log.Printf("invalid task json: %v", err)
		return err
	}

	log.Printf("▶ Running backtest for %s", task.Symbol)

	// Parse strategy JSON → Config struct
	cfg, err := strategy.FromJSON(task.Strategy)
	if err != nil {
		log.Printf("failed parsing strategy: %v", err)
		return err
	}

	// Load OHLC candles from GCS (bucket: quantengine-data)
	candles, err := loadCandles(ctx, task.Symbol)
	if err != nil {
		log.Printf("failed loading candles: %v", err)
		return err
	}

	// Run backtest engine
	result := backtester.Run(candles, cfg)

	// Convert to output format for BigQuery
	out := data.ToBQResult(task.Symbol, cfg, result)

	// Save to BigQuery
	if err := data.SaveToBigQuery(ctx, out); err != nil {
		log.Printf("BQ insert failed: %v", err)
		return err
	}

	log.Printf("✔ Done: %s", task.Symbol)
	return nil
}

// loads CSV from GCS bucket: quantengine-data/<symbol>.csv
func loadCandles(ctx context.Context, symbol string) ([]*data.Candle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bkt := client.Bucket("quantengine-data")
	obj := bkt.Object(symbol + ".csv")

	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return data.ParseCSV(reader)
}
