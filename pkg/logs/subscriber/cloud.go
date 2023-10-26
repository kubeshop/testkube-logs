package subscriber

import "github.com/kubeshop/testkube-logs/pkg/logs/events"

var _ Subscriber = &CloudSubscriber{}

// NewCloudSubscriber creates new CloudSubscriber which will send data to local MinIO bucket
func NewCloudSubscriber() *CloudSubscriber {
	return &CloudSubscriber{}
}

type CloudSubscriber struct {
	Bucket string
}

func (s *CloudSubscriber) Notify(id string, e events.LogChunk) error {
	panic("not implemented")
}

func (s *CloudSubscriber) Stop(id string) error {
	panic("not implemented")
}

func (s *CloudSubscriber) Name() string {
	return "cloud"
}
