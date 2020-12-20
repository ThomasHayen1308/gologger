package formatter

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

// JSON is a formatter for outputting json logs
type JSON struct{}

var _ Formatter = &JSON{}

// Format formats the log event data into bytes
func (j *JSON) Format(event *LogEvent) ([]byte, error) {
	data := make(map[string]interface{})
	if label, ok := event.Metadata["label"]; ok {
		if label != "" {
			data["level"] = label
		}
	}
	for k, v := range event.Metadata {
		data[k] = v
	}
	data["msg"] = event.Message
	data["timestamp"] = time.Now().UTC().Format("2006-01-02T15:04:05-0700")
	return jsoniter.Marshal(event)
}
