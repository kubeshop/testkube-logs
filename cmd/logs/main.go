package main

import (
	"github.com/kubeshop/testkube-logs/pkg/config"
	"github.com/kubeshop/testkube-logs/pkg/events"
	"github.com/kubeshop/testkube-logs/pkg/logger"
)

func main() {
	log := logger.Init().With("sercice", "logs")

	cfg := Must(config.Get())

	// Event bus
	natsConn := Must(events.NewNatsConnection(cfg.NatsURI))
	defer func() {
		log.Infof("closing encoded nats connection")
		natsConn.Close()
	}()

	natsEncodedConn := Must(events.NewEncodedNatsConnection(cfg.NatsURI))
	defer func() {
		log.Infof("closing encoded nats connection")
		natsEncodedConn.Close()
	}()

}

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
