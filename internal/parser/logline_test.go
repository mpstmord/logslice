package parser

import (
	"testing"
	"time"
)

func TestParseJSON_ValidWithRFC3339(t *testing.T) {
	raw := `{"time":"2024-03-15T12:00:00Z","level":"info","msg":"started"}`
	line, err := ParseJSON(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if line.Timestamp.IsZero() {
		t.Error("expected non-zero timestamp")
	}
	want := time.Date(2024, 3, 15, 12, 0, 0, 0, time.UTC)
	if !line.Timestamp.Equal(want) {
		t.Errorf("timestamp = %v, want %v", line.Timestamp, want)
	}
	if line.Fields["level"] != "info" {
		t.Errorf("expected level=info, got %v", line.Fields["level"])
	}
}

func TestParseJSON_UnixEpoch(t *testing.T) {
	raw := `{"ts":1710504000,"msg":"tick"}`
	line, err := ParseJSON(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if line.Timestamp.IsZero() {
		t.Error("expected non-zero timestamp from unix epoch")
	}
}

func TestParseJSON_NoTimestamp(t *testing.T) {
	raw := `{"level":"warn","msg":"no time field"}`
	line, err := ParseJSON(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !line.Timestamp.IsZero() {
		t.Errorf("expected zero timestamp, got %v", line.Timestamp)
	}
}

func TestParseJSON_InvalidJSON(t *testing.T) {
	raw := `not json at all`
	_, err := ParseJSON(raw)
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestParseJSON_AtTimestampKey(t *testing.T) {
	raw := `{"@timestamp":"2024-06-01T08:30:00Z","service":"api"}`
	line, err := ParseJSON(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := time.Date(2024, 6, 1, 8, 30, 0, 0, time.UTC)
	if !line.Timestamp.Equal(want) {
		t.Errorf("timestamp = %v, want %v", line.Timestamp, want)
	}
}
