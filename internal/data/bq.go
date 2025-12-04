package runner

import (
	"cloud.google.com/go/bigquery"
	"context"
)

func SaveToBigQuery(ctx context.Context, r *ResultOutput) error {
	client, err := bigquery.NewClient(ctx, "quantengine")
	if err != nil {
		return err
	}

	u := client.Dataset("quant_results").Table("backtest_results").Uploader()
	return u.Put(ctx, r)
}

