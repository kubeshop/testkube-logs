package logs

import (
	"github.com/kubeshop/testkube-logs/pkg/logs/events"
	"github.com/kubeshop/testkube-logs/pkg/logs/subscriber"
	"github.com/nats-io/nats.go"
)

func NewLogsService(nats *nats.EncodedConn) *LogsService {
	return &LogsService{
		nats:        nats,
		subscribers: []subscriber.Subscriber{},
	}
}

type LogsService struct {
	nats        *nats.EncodedConn
	subscribers []subscriber.Subscriber
}

func (l *LogsService) AddSubscriber(s subscriber.Subscriber) {
	l.subscribers = append(l.subscribers, s)
}

func (l *LogsService) Run() error {
	// this one will must a queue group each pod will get it's own
	l.nats.QueueSubscribe("events.logs.start", "startevents", func(event events.Trigger) {
		// handle subscribers

		for _, s := range l.subscribers {

		}
	})

	// listen for "events.logs" topic

	// for each topic like "events.logs.<id>" create subscriber

	// assuming this one will be scaled to multiple instances
	// how to handle pod issues here?
	// how to know that there is topic which is not handled by any subscriber?

	// block

	return nil
}
