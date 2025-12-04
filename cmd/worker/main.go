package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"quantengine/internal/runner"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data"`
	} `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var m PubSubMessage

	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &m); err != nil {
		log.Println("error unmarshalling pubsub:", err)
		w.WriteHeader(400)
		return
	}

	var payload runner.BacktestPayload
	if err := json.Unmarshal(m.Message.Data, &payload); err != nil {
		log.Println("payload decode error:", err)
		w.WriteHeader(400)
		return
	}

	log.Printf("Received job: %+v\n", payload)

	// Run backtest
	result, err := runner.RunBacktest(payload)
	if err != nil {
		log.Println("backtest error:", err)
		w.WriteHeader(500)
		return
	}

	log.Printf("Backtest finished → score: %.4f", result.Score)

	// Write to BigQuery
	if err := runner.SaveToBigQuery(context.Background(), result); err != nil {
		log.Println("bigquery write error:", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Worker ready, listening for Pub/Sub events on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

