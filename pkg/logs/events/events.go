package events

import "time"

// Generic event like log-start log-end
type Trigger struct {
	Id       string
	Type     string
	Metadata map[string]string
}

// log line/chunk data
type LogChunk struct {
	Time     time.Time         `json:"ts"`
	Content  string            `json:"content"`
	Metadata map[string]string `json:"metadata"`
}
