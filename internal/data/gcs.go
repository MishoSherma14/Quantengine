package data

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/csv"
	"os"
)

func LoadCandlesFromGCS(symbol, tf string) ([]Candle, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bkt := client.Bucket("quantengine-data")
	obj := bkt.Object(symbol + "_" + tf + ".csv")

	r, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return parseCSV(csv.NewReader(r))
}

