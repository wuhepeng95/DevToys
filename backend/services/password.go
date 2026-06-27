package services

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	upperChars     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars     = "abcdefghijklmnopqrstuvwxyz"
	digitChars     = "0123456789"
	reducedSpecial = "!@#$%^&*()_+-="
	ambiguous      = "0O1lI"
)

// PasswordOptions 密码生成选项
type PasswordOptions struct {
	Length           int  `json:"length"`
	IncludeUpper     bool `json:"includeUpper"`
	IncludeLower     bool `json:"includeLower"`
	IncludeDigits    bool `json:"includeDigits"`
	IncludeSpecial   bool `json:"includeSpecial"`
	ExcludeAmbiguous bool `json:"excludeAmbiguous"`
	Count            int  `json:"count"`
}

// PasswordResult 密码生成结果
type PasswordResult struct {
	Passwords []string `json:"passwords"`
}

// GeneratePasswords 生成随机密码
func GeneratePasswords(opts PasswordOptions) (PasswordResult, error) {
	if opts.Length < 4 {
		return PasswordResult{}, errors.New("密码长度不能少于 4 位")
	}
	if opts.Count < 1 {
		opts.Count = 1
	}
	if opts.Count > 100 {
		return PasswordResult{}, errors.New("一次最多生成 100 个密码")
	}

	charset := buildCharset(opts)
	if charset == "" {
		return PasswordResult{}, errors.New("至少选择一种字符类型")
	}

	passwords := make([]string, opts.Count)
	for i := 0; i < opts.Count; i++ {
		pwd, err := generateOnePassword(opts.Length, charset)
		if err != nil {
			return PasswordResult{}, err
		}
		passwords[i] = pwd
	}

	return PasswordResult{Passwords: passwords}, nil
}

func buildCharset(opts PasswordOptions) string {
	var cs string
	if opts.IncludeUpper {
		cs += upperChars
	}
	if opts.IncludeLower {
		cs += lowerChars
	}
	if opts.IncludeDigits {
		cs += digitChars
	}
	if opts.IncludeSpecial {
		cs += reducedSpecial
	}
	if opts.ExcludeAmbiguous {
		cs = filterChars(cs, ambiguous)
	}
	return cs
}

func filterChars(charset, exclude string) string {
	excludeSet := make(map[rune]bool)
	for _, r := range exclude {
		excludeSet[r] = true
	}
	var result []rune
	for _, r := range charset {
		if !excludeSet[r] {
			result = append(result, r)
		}
	}
	return string(result)
}

func generateOnePassword(length int, charset string) (string, error) {
	charsetLen := big.NewInt(int64(len(charset)))
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		password[i] = charset[idx.Int64()]
	}

	return string(password), nil
}
