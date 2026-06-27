package services

import (
	"sort"
	"strconv"
	"strings"
)

// RemoveDuplicateOptions 文本去重选项
type RemoveDuplicateOptions struct {
	KeepOrder   bool `json:"keepOrder"`
	IgnoreCase  bool `json:"ignoreCase"`
	RemoveEmpty bool `json:"removeEmpty"`
}

// RemoveDuplicate 去除文本中的重复行
func RemoveDuplicate(input string, opts RemoveDuplicateOptions) (string, error) {
	lines := strings.Split(input, "\n")

	if opts.RemoveEmpty {
		filtered := make([]string, 0, len(lines))
		for _, l := range lines {
			if strings.TrimSpace(l) != "" {
				filtered = append(filtered, l)
			}
		}
		lines = filtered
	}

	seen := make(map[string]bool)
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		key := line
		if opts.IgnoreCase {
			key = strings.ToLower(line)
		}
		if !seen[key] {
			seen[key] = true
			result = append(result, line)
		}
	}

	if !opts.KeepOrder {
		sort.Strings(result)
	}

	return strings.Join(result, "\n"), nil
}

// SortLinesOptions 文本排序选项
type SortLinesOptions struct {
	Order      string `json:"order"`
	Numeric    bool   `json:"numeric"`
	IgnoreCase bool   `json:"ignoreCase"`
}

// SortLines 对文本行进行排序
func SortLines(input string, opts SortLinesOptions) (string, error) {
	lines := strings.Split(input, "\n")
	filtered := make([]string, 0, len(lines))
	for _, l := range lines {
		if l != "" {
			filtered = append(filtered, l)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		a, b := filtered[i], filtered[j]
		if opts.IgnoreCase {
			a = strings.ToLower(a)
			b = strings.ToLower(b)
		}
		if opts.Numeric {
			na, _ := strconv.ParseFloat(a, 64)
			nb, _ := strconv.ParseFloat(b, 64)
			if na != 0 || nb != 0 {
				return na < nb
			}
		}
		return a < b
	})

	if opts.Order == "desc" {
		for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
			filtered[i], filtered[j] = filtered[j], filtered[i]
		}
	}

	return strings.Join(filtered, "\n"), nil
}
