package main

import (
	"context"

	"github.com/kubeshop/testkube-logs/pkg/config"
	"github.com/kubeshop/testkube-logs/pkg/events"
	"github.com/kubeshop/testkube-logs/pkg/logger"
	"github.com/kubeshop/testkube-logs/pkg/logs"
	"github.com/kubeshop/testkube-logs/pkg/logs/consumer"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	log := logger.Init().With("sercice", "logs")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := Must(config.Get())

	// Event bus
	natsConn := Must(events.NewNatsConnection(cfg.NatsURI))
	defer func() {
		log.Infof("closing nats connection")
		natsConn.Close()
	}()

	natsEncodedConn := Must(events.NewEncodedNatsConnection(cfg.NatsURI))
	defer func() {
		log.Infof("closing encoded nats connection")
		natsEncodedConn.Close()
	}()

	js := Must(jetstream.New(natsConn))

	svc := logs.NewLogsService(natsEncodedConn, js)
	svc.AddSubscriber(consumer.NewDummyConsumer())
	svc.Run(ctx)
}

// Must helper function to panic on error
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
