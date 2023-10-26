package subscriber

import "github.com/kubeshop/testkube-logs/pkg/logs/events"

// Subscriber will listen to log chunks, and handle them based on log id (execution Id)
type Subscriber interface {
	// Notify will send data log events for particaular execution id
	Notify(id string, event events.LogChunk) error
	// Stop will stop listening subscriber from sending data
	Stop(id string) error
	// Name subscriber name
	Name() string
}
