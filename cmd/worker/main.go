package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"quantengine/internal/worker"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var msg struct {
			Message struct {
				Data []byte `json:"data"`
			} `json:"message"`
		}

		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			log.Println("decode error:", err)
			return
		}

		var task worker.TaskMessage
		if err := json.Unmarshal(msg.Message.Data, &task); err != nil {
			log.Println("unmarshal error:", err)
			return
		}

		log.Println("Running task:", task.Symbol, task.Timeframe)

		err := worker.HandleTask(context.Background(), &task)
		if err != nil {
			log.Println(err)
		}
	})

	log.Println("Worker started on 8080...")
	http.ListenAndServe(":8080", nil)
}
