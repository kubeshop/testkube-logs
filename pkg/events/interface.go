package events

import "github.com/nats-io/nats.go"

// Channel is a generic channel for events of type T
// TODO fix it //go:generate mockgen -destination=./channel_mock.go -package=events "github.com/kubeshop/testkube-logs/pkg/events" Channel[T]

type Channel[T any] interface {
	// Publish publishes event to the bus
	Publish(event T) error
	// Subscribe subscribes to the queue group
	Subscribe(queue string) (chan T, error)
	// Unsuscribe unsubscribes from the queue group
	Unsubscribe(queue string) error
	// Close closes the bus
	Close() error
}

// Eventable is an interface for possible data struct which can be passed through events channel
type Eventable interface {
	Data() any
}

// Bus interface to abstract NATS Connection and other event buses
//
//go:generate mockgen -destination=./bus_mock.go -package=events "github.com/kubeshop/testkube-logs/pkg/events" Bus
type Bus interface {
	Publish(subj string, data []byte) error
	Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error)
	SubscribeSync(subj string) (*nats.Subscription, error)
}
