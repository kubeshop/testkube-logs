package events

import "time"

// Generic event like log-start log-end
type Trigger struct {
	Id       string            `json:"id,omitempty"`
	Type     string            `json:"type,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// log line/chunk data
type LogChunk struct {
	Time     time.Time         `json:"ts"`
	Content  string            `json:"content"`
	Metadata map[string]string `json:"metadata"`
}
