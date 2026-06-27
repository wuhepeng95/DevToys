package services

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// JSONToYAML 将 JSON 字符串转换为 YAML
func JSONToYAML(input string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return "", err
	}
	out, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// YAMLToJSON 将 YAML 字符串转换为格式化 JSON
func YAMLToJSON(input string) (string, error) {
	var data interface{}
	if err := yaml.Unmarshal([]byte(input), &data); err != nil {
		return "", err
	}
	data = normalizeYAML(data)
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// YAMLToJSONCompact 将 YAML 转换为压缩 JSON
func YAMLToJSONCompact(input string) (string, error) {
	var data interface{}
	if err := yaml.Unmarshal([]byte(input), &data); err != nil {
		return "", err
	}
	data = normalizeYAML(data)
	out, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// normalizeYAML 将 YAML 解析后的 map[interface{}]interface{} 转为 JSON 兼容类型
func normalizeYAML(v interface{}) interface{} {
	switch x := v.(type) {
	case map[string]interface{}:
		for k, val := range x {
			x[k] = normalizeYAML(val)
		}
		return x
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for k, val := range x {
			key, _ := k.(string)
			m[key] = normalizeYAML(val)
		}
		return m
	case []interface{}:
		for i, item := range x {
			x[i] = normalizeYAML(item)
		}
	}
	return v
}
