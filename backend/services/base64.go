package services

import (
	"encoding/base64"
)

// EncodeBase64 对输入进行 Base64 编码
func EncodeBase64(input string, urlSafe bool) (string, error) {
	data := []byte(input)
	if urlSafe {
		return base64.RawURLEncoding.EncodeToString(data), nil
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// DecodeBase64 对输入进行 Base64 解码
func DecodeBase64(input string, urlSafe bool) (string, error) {
	var decoded []byte
	var err error

	if urlSafe {
		decoded, err = base64.RawURLEncoding.DecodeString(input)
	} else {
		decoded, err = base64.StdEncoding.DecodeString(input)
	}
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
