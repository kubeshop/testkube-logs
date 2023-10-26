package events

import "github.com/nats-io/nats.go"

// Bus interface to abstract NATS Connection and other event buses
//
//go:generate mockgen -destination=./bus_mock.go -package=events "github.com/kubeshop/testkube-logs/pkg/events" Bus
type Bus interface {
	Publish(subj string, data []byte) error
	Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error)
	SubscribeSync(subj string) (*nats.Subscription, error)
}
