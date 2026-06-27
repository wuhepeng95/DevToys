package services

import (
	"fmt"
	"math/big"
	"strings"
)

// ConvertNumber 将数字在进制间转换，返回各进制表示
type ConvertNumber struct {
	Binary      string `json:"binary"`
	Octal       string `json:"octal"`
	Decimal     string `json:"decimal"`
	Hexadecimal string `json:"hexadecimal"`
}

func isValidBase(base int) bool {
	return base == 2 || base == 8 || base == 10 || base == 16
}

// ConvertBase 将输入字符串从指定进制转换为其他进制
func ConvertBase(input string, fromBase int) (ConvertNumber, error) {
	if !isValidBase(fromBase) {
		return ConvertNumber{}, fmt.Errorf("不支持的进制: %d，仅支持 2/8/10/16", fromBase)
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return ConvertNumber{}, fmt.Errorf("输入为空")
	}

	// 去除 0x/0b 前缀
	if fromBase == 16 && len(input) > 2 && strings.ToLower(input[:2]) == "0x" {
		input = input[2:]
	}
	if fromBase == 2 && len(input) > 2 && strings.ToLower(input[:2]) == "0b" {
		input = input[2:]
	}

	n := new(big.Int)
	n, ok := n.SetString(input, fromBase)
	if !ok {
		return ConvertNumber{}, fmt.Errorf("无效的 %d 进制数: %s", fromBase, input)
	}

	return ConvertNumber{
		Binary:      n.Text(2),
		Octal:       n.Text(8),
		Decimal:     n.Text(10),
		Hexadecimal: strings.ToLower(n.Text(16)),
	}, nil
}
