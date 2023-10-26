package subscriber

import (
	"fmt"

	"github.com/kubeshop/testkube-logs/pkg/logs/events"
)

var _ Subscriber = &DummySubscriber{}

// NewS3Subscriber creates new DummySubscriber which will send data to local MinIO bucket
func NewDummySubscriber() *DummySubscriber {
	return &DummySubscriber{}
}

type DummySubscriber struct {
	Bucket string
}

func (s *DummySubscriber) Notify(id string, e events.LogChunk) error {
	fmt.Printf("%s %+v\n", id, e)
	return nil
}

func (s *DummySubscriber) Stop(id string) error {
	fmt.Printf("stopping %s \n", id)
	return nil
}

func (s *DummySubscriber) Name() string {
	return "dummy"
}
