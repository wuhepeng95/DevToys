package services

import "strings"

// DiffLineType 表示一行在 diff 中的类型
type DiffLineType string

const (
	DiffLineSame   DiffLineType = "same"
	DiffLineAdd    DiffLineType = "add"
	DiffLineRemove DiffLineType = "remove"
)

// DiffLine 表示 diff 中的一行
type DiffLine struct {
	Type DiffLineType `json:"type"`
	Text string       `json:"text"`
}

// DiffResult 表示 diff 的完整结果
type DiffResult struct {
	Lines []DiffLine `json:"lines"`
	Text  string     `json:"text"`
}

// CompareText 对比两段文本，返回逐行 diff 结果
func CompareText(left, right string) (DiffResult, error) {
	leftLines := strings.Split(left, "\n")
	rightLines := strings.Split(right, "\n")

	lcs := computeLCS(leftLines, rightLines)
	lines := buildDiff(leftLines, rightLines, lcs)

	var sb strings.Builder
	for _, l := range lines {
		switch l.Type {
		case DiffLineAdd:
			sb.WriteString("+ ")
		case DiffLineRemove:
			sb.WriteString("- ")
		default:
			sb.WriteString("  ")
		}
		sb.WriteString(l.Text)
		sb.WriteString("\n")
	}

	return DiffResult{
		Lines: lines,
		Text:  strings.TrimRight(sb.String(), "\n"),
	}, nil
}

// computeLCS 计算最长公共子序列（用于 diff）
func computeLCS(a, b []string) [][]int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp
}

// buildDiff 根据 LCS 表构建 diff 行
func buildDiff(a, b []string, dp [][]int) []DiffLine {
	var lines []DiffLine
	i, j := len(a), len(b)
	var temp []DiffLine

	for i > 0 || j > 0 {
		if i > 0 && j > 0 && a[i-1] == b[j-1] {
			temp = append(temp, DiffLine{Type: DiffLineSame, Text: a[i-1]})
			i--
			j--
		} else if j > 0 && (i == 0 || dp[i][j-1] >= dp[i-1][j]) {
			temp = append(temp, DiffLine{Type: DiffLineAdd, Text: b[j-1]})
			j--
		} else if i > 0 {
			temp = append(temp, DiffLine{Type: DiffLineRemove, Text: a[i-1]})
			i--
		}
	}

	for k := len(temp) - 1; k >= 0; k-- {
		lines = append(lines, temp[k])
	}
	return lines
}
