package subscriber

import "github.com/kubeshop/testkube-logs/pkg/logs/events"

// Subscriber will listen to log events
type Subscriber interface {

	// Listen will listen to log events for particaular execution id
	Listen(id string, event events.LogChunk) error
}
