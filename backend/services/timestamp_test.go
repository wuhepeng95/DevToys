package services

import (
	"strconv"
	"strings"
	"testing"
)

func TestUnixToDate(t *testing.T) {
	tests := []struct {
		name      string
		timestamp string
		unit      string
		useUTC    bool
		wantErr   bool
	}{
		{
			name:      "seconds",
			timestamp: "1700000000",
			unit:      "s",
			useUTC:    false,
			wantErr:   false,
		},
		{
			name:      "milliseconds",
			timestamp: "1700000000000",
			unit:      "ms",
			useUTC:    false,
			wantErr:   false,
		},
		{
			name:      "microseconds",
			timestamp: "1700000000000000",
			unit:      "us",
			useUTC:    false,
			wantErr:   false,
		},
		{
			name:      "nanoseconds",
			timestamp: "1700000000000000000",
			unit:      "ns",
			useUTC:    false,
			wantErr:   false,
		},
		{
			name:      "UTC mode",
			timestamp: "1700000000",
			unit:      "s",
			useUTC:    true,
			wantErr:   false,
		},
		{
			name:      "invalid timestamp",
			timestamp: "abc",
			unit:      "s",
			useUTC:    false,
			wantErr:   true,
		},
		{
			name:      "invalid unit",
			timestamp: "1700000000",
			unit:      "xyz",
			useUTC:    false,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := UnixToDate(tt.timestamp, tt.unit, tt.useUTC)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnixToDate() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result.Local == "" {
					t.Error("UnixToDate() returned empty local date")
				}
				if result.UTC == "" {
					t.Error("UnixToDate() returned empty UTC date")
				}
			}
		})
	}
}

func TestDateToUnix(t *testing.T) {
	tests := []struct {
		name    string
		dateStr string
		wantErr bool
	}{
		{
			name:    "RFC3339 format",
			dateStr: "2023-11-14T22:13:20Z",
			wantErr: false,
		},
		{
			name:    "standard datetime",
			dateStr: "2023-11-14 22:13:20",
			wantErr: false,
		},
		{
			name:    "date only",
			dateStr: "2023-11-14",
			wantErr: false,
		},
		{
			name:    "invalid date",
			dateStr: "not-a-date",
			wantErr: true,
		},
		{
			name:    "empty string",
			dateStr: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DateToUnix(tt.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateToUnix() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result.Seconds <= 0 {
					t.Error("DateToUnix() returned non-positive seconds")
				}
				if result.Milliseconds <= 0 {
					t.Error("DateToUnix() returned non-positive milliseconds")
				}
			}
		})
	}
}

func TestDateToUnixConsistency(t *testing.T) {
	// 验证秒/毫秒/微秒/纳秒的倍数关系
	result, err := DateToUnix("2023-11-14 22:13:20")
	if err != nil {
		t.Fatal(err)
	}
	if result.Milliseconds != result.Seconds*1000 {
		t.Errorf("milliseconds should be seconds*1000, got %d vs %d", result.Milliseconds, result.Seconds*1000)
	}
}

func TestUnixToDateRoundTrip(t *testing.T) {
	dateResult, err := DateToUnix("2023-11-14T22:13:20Z")
	if err != nil {
		t.Fatal(err)
	}

	ts := dateResult.Seconds
	timeResult, err := UnixToDate(strconv.FormatInt(ts, 10), "s", true)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(timeResult.UTC, "2023-11-14") {
		t.Errorf("round trip failed, expected date containing 2023-11-14, got %s", timeResult.UTC)
	}
}
