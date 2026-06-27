package services

import "encoding/json"

// FormatJSON 格式化 JSON 字符串，indent 为缩进字符（如 "  " 或 "\t"）
func FormatJSON(input string, indent string) (string, error) {
	if indent == "" {
		indent = "  "
	}
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", indent)
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

// CompressJSON 压缩 JSON 字符串（去除多余空格和换行）
func CompressJSON(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	compressed, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}
	return string(compressed), nil
}

// ValidateJSON 校验 JSON 是否合法
func ValidateJSON(input string) (bool, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return false, err
	}
	return true, nil
}
