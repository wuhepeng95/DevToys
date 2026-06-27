package main

import (
	"context"

	"devtoys/backend/services"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// FormatJSON 格式化 JSON 字符串
func (a *App) FormatJSON(input string, indent string) services.Result {
	output, err := services.FormatJSON(input, indent)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// CompressJSON 压缩 JSON 字符串
func (a *App) CompressJSON(input string) services.Result {
	output, err := services.CompressJSON(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// ValidateJSON 校验 JSON 合法性
func (a *App) ValidateJSON(input string) services.Result {
	valid, err := services.ValidateJSON(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(valid)
}

// RemoveDuplicate 文本去重
func (a *App) RemoveDuplicate(input string, opts services.RemoveDuplicateOptions) services.Result {
	output, err := services.RemoveDuplicate(input, opts)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// SortLines 文本排序
func (a *App) SortLines(input string, opts services.SortLinesOptions) services.Result {
	output, err := services.SortLines(input, opts)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// EncodeBase64 Base64 编码
func (a *App) EncodeBase64(input string, urlSafe bool) services.Result {
	output, err := services.EncodeBase64(input, urlSafe)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// DecodeBase64 Base64 解码
func (a *App) DecodeBase64(input string, urlSafe bool) services.Result {
	output, err := services.DecodeBase64(input, urlSafe)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// CompareText 文本 Diff
func (a *App) CompareText(left string, right string) services.Result {
	result, err := services.CompareText(left, right)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(result)
}

// UnixToDate 时间戳转日期
func (a *App) UnixToDate(timestamp string, unit string, useUTC bool) services.Result {
	result, err := services.UnixToDate(timestamp, unit, useUTC)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(result)
}

// DateToUnix 日期转时间戳
func (a *App) DateToUnix(dateStr string) services.Result {
	result, err := services.DateToUnix(dateStr)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(result)
}

// URLEncode URL 编码
func (a *App) URLEncode(input string) services.Result {
	output, err := services.URLEncode(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// URLDecode URL 解码
func (a *App) URLDecode(input string) services.Result {
	output, err := services.URLDecode(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// QueryEncode Query 参数编码
func (a *App) QueryEncode(input string) services.Result {
	output, err := services.QueryEncode(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// QueryDecode Query 参数解码
func (a *App) QueryDecode(input string) services.Result {
	output, err := services.QueryDecode(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// ParseURL 解析 URL
func (a *App) ParseURL(rawURL string) services.Result {
	result, err := services.ParseURL(rawURL)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(result)
}

// GeneratePasswords 生成随机密码
func (a *App) GeneratePasswords(opts services.PasswordOptions) services.Result {
	result, err := services.GeneratePasswords(opts)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(result)
}

// JSONToYAML JSON 转 YAML
func (a *App) JSONToYAML(input string) services.Result {
	output, err := services.JSONToYAML(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// YAMLToJSON YAML 转格式化 JSON
func (a *App) YAMLToJSON(input string) services.Result {
	output, err := services.YAMLToJSON(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// YAMLToJSONCompact YAML 转压缩 JSON
func (a *App) YAMLToJSONCompact(input string) services.Result {
	output, err := services.YAMLToJSONCompact(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// FormatSQL 格式化 SQL 语句
func (a *App) FormatSQL(input string) services.Result {
	output, err := services.FormatSQL(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// MinifySQL 压缩 SQL 语句
func (a *App) MinifySQL(input string) services.Result {
	output, err := services.MinifySQL(input)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(output)
}

// ConvertBase 进制转换
func (a *App) ConvertBase(input string, fromBase int) services.Result {
	result, err := services.ConvertBase(input, fromBase)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(result)
}

// CalculateSubnet 计算子网信息
func (a *App) CalculateSubnet(cidr string) services.Result {
	result, err := services.CalculateSubnet(cidr)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(result)
}

// SplitSubnet 拆分 CIDR 子网
func (a *App) SplitSubnet(cidr string, targetSize int) services.Result {
	result, err := services.SplitSubnet(cidr, targetSize)
	if err != nil {
		return services.Error(err.Error())
	}
	return services.Success(result)
}
