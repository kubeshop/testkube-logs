package logs

import (
	"context"
	"encoding/json"

	"github.com/kubeshop/testkube-logs/pkg/logs/events"
	"github.com/kubeshop/testkube-logs/pkg/logs/subscriber"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

const (
	StreamName = "LOGS"
)

func NewLogsService(nats *nats.EncodedConn, js jetstream.JetStream) *LogsService {
	return &LogsService{
		nats:        nats,
		subscribers: []subscriber.Subscriber{},
		js:          js,
	}
}

type LogsService struct {
	log         *zap.SugaredLogger
	nats        *nats.EncodedConn
	js          jetstream.JetStream
	subscribers []subscriber.Subscriber
}

func (l *LogsService) AddSubscriber(s subscriber.Subscriber) {
	l.subscribers = append(l.subscribers, s)
}

func (l *LogsService) Run(ctx context.Context) error {

	// TODO refactor abstract NATS logic from here?
	// TODO consider using durable topics for queue with Ack / Nack

	l.nats.QueueSubscribe("events.logs.stop", "startevents", func(event events.Trigger) {
		// TODO stop all consumers from consuming data for given execution id
	})

	// this one will must a queue group each pod will get it's own
	l.nats.QueueSubscribe("events.logs.start", "startevents", func(event events.Trigger) {
		log := l.log.With("id", event.Id)

		// create stream for incoming logs
		streamName := StreamName + event.Id
		_, err := l.js.CreateStream(ctx, jetstream.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamName + ".*"},
			// MaxAge:   time.Minute,
			Storage: jetstream.FileStorage, // durable stream
		})

		if err != nil {
			l.log.Errorw("error creating stream", "error", err, "id", event.Id)
			return
		}

		// for each consumer create nats consumer and consume stream from it e.g. cloud s3 or others
		for _, sub := range l.subscribers {
			name := "lc" + event.Id + "." + sub.Name()
			c, err := l.js.CreateOrUpdateConsumer(ctx, StreamName, jetstream.ConsumerConfig{
				Name:          name,
				Durable:       name,
				FilterSubject: streamName,
				DeliverPolicy: jetstream.DeliverAllPolicy,
			})

			if err != nil {
				log.Errorw("error creating consumer", "error", err)
			}

			cons, err := c.Consume(func(msg jetstream.Msg) {
				// deliver to subscriber
				logChunk := events.LogChunk{}
				json.Unmarshal(msg.Data(), &logChunk)
				err := sub.Notify(event.Id, logChunk)

				if err != nil {
					if err := msg.Nak(); err != nil {
						log.Errorw("error nacking message", "error", err)
						return
					}
					return
				}

				if err := msg.Ack(); err != nil {
					log.Errorw("error acking message", "error", err)
				}
			})

			// TODO add `cons` and stop it on stop event
			var _ = cons

			if err != nil {
				log.Errorw("error consuming", "error", err, "consumer", c.CachedInfo())
			}
		}
	})

	// TODO
	// assuming this one will be scaled to multiple instances
	// how to handle pod issues here?
	// how to know that there is topic which is not handled by any subscriber?
	// we woudl need to check pending log topics and handle them

	// block

	return nil
}
