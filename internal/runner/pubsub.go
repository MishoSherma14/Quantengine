package runner

import (
    "cloud.google.com/go/pubsub"
    "context"
    "encoding/json"
)

func PullMessage(ctx context.Context) (*TaskMessage, error) {
    client, err := pubsub.NewClient(ctx, "quantengine")
    if err != nil {
        return nil, err
    }

    sub := client.Subscription("strategy-tasks-sub")

    var msg TaskMessage
    err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
        json.Unmarshal(m.Data, &msg)
        m.Ack()
    })
    if err != nil {
        return nil, err
    }

    return &msg, nil
}
