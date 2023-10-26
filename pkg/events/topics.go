package events

import "fmt"

func LogTopic(executionId string) string {
	return fmt.Sprintf("logs.%s", executionId)
}
