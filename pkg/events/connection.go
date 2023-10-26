package events

import (
	"github.com/kubeshop/testkube-logs/pkg/logger"

	"github.com/nats-io/nats.go"
)

func NewEncodedNatsConnection(natsURI string) (*nats.EncodedConn, error) {
	log := logger.Logger()

	nc, err := nats.Connect(natsURI)
	if err != nil {
		log.Fatalw("error connecting to nats", "error", err)
		return nil, err
	}

	// automatic NATS JSON CODEC
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalw("error connecting to nats", "error", err)
		return nil, err
	}

	return ec, nil
}

func NewNatsConnection(natsURI string) (*nats.Conn, error) {
	log := logger.Logger()

	nc, err := nats.Connect(natsURI)
	if err != nil {
		log.Fatalw("error connecting to nats", "error", err)
		return nil, err
	}

	return nc, nil
}
