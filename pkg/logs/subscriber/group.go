package subscriber

import "github.com/kubeshop/testkube-logs/pkg/logs/events"

type SubscribersGroup struct {
	subscribers []Subscriber
}

func (s *SubscribersGroup) Add(sub Subscriber) {
	s.subscribers = append(s.subscribers, sub)
}

func (s *SubscribersGroup) NotifyAll(id string, event events.LogChunk) error {
	for _, sub := range s.subscribers {
		if err := sub.Notify(id, event); err != nil {
			return err
		}
	}
	return nil
}

func (s *SubscribersGroup) StopAll(id string) error {
	for _, sub := range s.subscribers {
		if err := sub.Stop(id); err != nil {
			return err
		}
	}
	return nil
}
