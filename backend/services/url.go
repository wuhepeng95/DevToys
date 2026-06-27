package services

import (
	"fmt"
	"net/url"
)

// URLParts URL 解析结果的各个部分
type URLParts struct {
	Scheme   string `json:"scheme"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Path     string `json:"path"`
	Query    string `json:"query"`
	Fragment string `json:"fragment"`
	Raw      string `json:"raw"`
}

// URLEncode 对整个字符串进行百分号编码
func URLEncode(input string) (string, error) {
	return url.QueryEscape(input), nil
}

// URLDecode 对百分号编码的字符串进行解码
func URLDecode(input string) (string, error) {
	decoded, err := url.QueryUnescape(input)
	if err != nil {
		return "", fmt.Errorf("URL 解码失败: %s", err.Error())
	}
	return decoded, nil
}

// QueryEncode 对 query 参数值进行编码（保留 = & 结构）
func QueryEncode(input string) (string, error) {
	parsed, err := url.ParseQuery(input)
	if err != nil {
		return "", fmt.Errorf("Query 解析失败: %s", err.Error())
	}
	return parsed.Encode(), nil
}

// QueryDecode 解码 URL query 字符串
func QueryDecode(input string) (string, error) {
	decoded, err := url.QueryUnescape(input)
	if err != nil {
		return "", fmt.Errorf("Query 解码失败: %s", err.Error())
	}
	return decoded, nil
}

// ParseURL 解析 URL 的各个组成部分
func ParseURL(rawURL string) (URLParts, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return URLParts{}, fmt.Errorf("URL 解析失败: %s", err.Error())
	}

	host := parsed.Hostname()
	port := parsed.Port()

	result := URLParts{
		Scheme:   parsed.Scheme,
		Host:     host,
		Port:     port,
		Path:     parsed.Path,
		Query:    parsed.RawQuery,
		Fragment: parsed.Fragment,
		Raw:      rawURL,
	}
	return result, nil
}
