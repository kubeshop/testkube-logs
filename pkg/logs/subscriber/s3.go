package subscriber

import "github.com/kubeshop/testkube-logs/pkg/logs/events"

var _ Subscriber = &S3Subscriber{}

// NewS3Subscriber creates new S3Subscriber which will send data to local MinIO bucket
func NewS3Subscriber() *S3Subscriber {
	return &S3Subscriber{}
}

type S3Subscriber struct {
	Bucket string
}

func (s *S3Subscriber) Notify(id string, e events.LogChunk) error {
	panic("not implemented")
}

func (s *S3Subscriber) Stop(id string) error {
	panic("not implemented")
}

func (s *S3Subscriber) Name() string {
	return "s3"
}
