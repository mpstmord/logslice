package parser

import (
	"encoding/json"
	"fmt"
	"time"
)

// LogLine represents a single parsed log entry.
type LogLine struct {
	Timestamp time.Time
	Raw       string
	Fields    map[string]interface{}
}

// CommonTimeFormats lists the time formats attempted during parsing.
var CommonTimeFormats = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02T15:04:05.000Z0700",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04:05.000",
}

// ParseJSON attempts to parse a raw log line as a JSON object.
// It looks for a timestamp field ("time", "ts", "timestamp", "@timestamp")
// and promotes it to LogLine.Timestamp.
func ParseJSON(raw string) (*LogLine, error) {
	var fields map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &fields); err != nil {
		return nil, fmt.Errorf("not valid JSON: %w", err)
	}

	line := &LogLine{
		Raw:    raw,
		Fields: fields,
	}

	for _, key := range []string{"time", "ts", "timestamp", "@timestamp"} {
		if v, ok := fields[key]; ok {
			if t, err := parseTimestampValue(v); err == nil {
				line.Timestamp = t
				break
			}
		}
	}

	return line, nil
}

// parseTimestampValue converts an interface{} value (string or float64 unix epoch)
// into a time.Time.
func parseTimestampValue(v interface{}) (time.Time, error) {
	switch val := v.(type) {
	case string:
		for _, layout := range CommonTimeFormats {
			if t, err := time.Parse(layout, val); err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("unrecognised time string: %s", val)
	case float64:
		return time.Unix(int64(val), 0).UTC(), nil
	}
	return time.Time{}, fmt.Errorf("unsupported timestamp type %T", v)
}
