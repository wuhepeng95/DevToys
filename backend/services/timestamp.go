package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// UnixToDateResult 时间戳转日期的结果
type UnixToDateResult struct {
	Local   string `json:"local"`
	UTC     string `json:"utc"`
	ISO8601 string `json:"iso8601"`
	RFC3339 string `json:"rfc3339"`
}

// DateToUnixResult 日期转时间戳的结果
type DateToUnixResult struct {
	Seconds      int64 `json:"seconds"`
	Milliseconds int64 `json:"milliseconds"`
	Microseconds int64 `json:"microseconds"`
	Nanoseconds  int64 `json:"nanoseconds"`
}

// UnixToDate 将时间戳转换为日期字符串
func UnixToDate(timestamp string, unit string, useUTC bool) (UnixToDateResult, error) {
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return UnixToDateResult{}, fmt.Errorf("无效的时间戳: %s", timestamp)
	}

	var t time.Time
	switch strings.ToLower(unit) {
	case "s":
		t = time.Unix(ts, 0)
	case "ms":
		t = time.UnixMilli(ts)
	case "us":
		t = time.UnixMicro(ts)
	case "ns":
		t = time.Unix(0, ts)
	default:
		return UnixToDateResult{}, fmt.Errorf("不支持的时间单位: %s", unit)
	}

	if useUTC {
		t = t.UTC()
	}

	local := t.Format("2006-01-02 15:04:05")
	utc := t.UTC().Format("2006-01-02 15:04:05")

	result := UnixToDateResult{
		Local:   local,
		UTC:     utc,
		ISO8601: t.Format(time.RFC3339),
		RFC3339: t.Format(time.RFC3339),
	}
	return result, nil
}

// DateToUnix 将日期字符串转换为时间戳
func DateToUnix(dateStr string) (DateToUnixResult, error) {
	formats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02",
		"2006/01/02 15:04:05",
		"2006/01/02",
		"01/02/2006",
		"2006-01-02T15:04:05Z",
	}

	dateStr = strings.TrimSpace(dateStr)
	var t time.Time
	var err error

	for _, f := range formats {
		t, err = time.Parse(f, dateStr)
		if err == nil {
			break
		}
	}

	if err != nil {
		return DateToUnixResult{}, fmt.Errorf("无法解析日期: %s，支持的格式: YYYY-MM-DD HH:mm:ss、RFC3339、ISO8601", dateStr)
	}

	ms := t.UnixMilli()
	result := DateToUnixResult{
		Seconds:      t.Unix(),
		Milliseconds: ms,
		Microseconds: ms * 1000,
		Nanoseconds:  ms * 1000000,
	}
	return result, nil
}
