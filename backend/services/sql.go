package services

import (
	"fmt"
	"strings"
	"unicode"
)

var sqlKeywords = map[string]bool{
	"SELECT": true, "FROM": true, "WHERE": true, "AND": true, "OR": true,
	"INSERT": true, "INTO": true, "VALUES": true, "UPDATE": true, "SET": true,
	"DELETE": true, "CREATE": true, "TABLE": true, "DROP": true, "ALTER": true,
	"ADD": true, "COLUMN": true, "INDEX": true, "VIEW": true, "JOIN": true,
	"LEFT": true, "RIGHT": true, "INNER": true, "OUTER": true, "FULL": true,
	"ON": true, "AS": true, "IN": true, "IS": true, "NULL": true,
	"NOT": true, "LIKE": true, "BETWEEN": true, "EXISTS": true, "CASE": true,
	"WHEN": true, "THEN": true, "ELSE": true, "END": true, "ORDER": true,
	"BY": true, "GROUP": true, "HAVING": true, "LIMIT": true, "OFFSET": true,
	"UNION": true, "ALL": true, "DISTINCT": true, "TOP": true, "WITH": true,
	"RECURSIVE": true, "CAST": true, "COALESCE": true, "IF": true,
	"PRIMARY": true, "KEY": true, "FOREIGN": true, "REFERENCES": true,
	"CONSTRAINT": true, "DEFAULT": true, "CHECK": true, "UNIQUE": true,
	"AUTO_INCREMENT": true, "CASCADE": true, "GRANT": true, "REVOKE": true,
	"BEGIN": true, "COMMIT": true, "ROLLBACK": true, "TRANSACTION": true,
	"EXEC": true, "EXECUTE": true, "PROCEDURE": true, "FUNCTION": true,
	"RETURN": true, "DECLARE": true, "CURSOR": true, "FETCH": true,
	"CLOSE": true, "PRINT": true, "RAISE": true, "EXCEPTION": true,
	"TRIGGER": true, "EVENT": true, "SCHEMA": true,
	"DATABASE": true, "USE": true, "SHOW": true, "DESCRIBE": true, "EXPLAIN": true,
	"REPLACE": true, "TRUNCATE": true, "RENAME": true, "CROSS": true, "NATURAL": true,
	"USING": true, "ANY": true, "SOME": true, "TRUE": true, "FALSE": true,
	"INTERSECT": true, "EXCEPT": true, "MINUS": true,
}

var newlineBeforeKeywords = map[string]bool{
	"SELECT": true, "FROM": true, "WHERE": true, "ORDER": true,
	"GROUP": true, "HAVING": true, "LIMIT": true, "OFFSET": true,
	"UNION": true, "INSERT": true, "INTO": true, "VALUES": true,
	"UPDATE": true, "SET": true, "DELETE": true, "CREATE": true,
	"ALTER": true, "DROP": true, "JOIN": true, "INNER": true,
	"LEFT": true, "RIGHT": true, "FULL": true, "OUTER": true,
	"CROSS": true, "NATURAL": true, "ON": true, "AND": true, "OR": true,
	"BY": true,
}

// FormatSQL 格式化 SQL 语句
func FormatSQL(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("SQL 内容为空")
	}
	tokens := tokenize(input)
	return formatTokens(tokens), nil
}

// MinifySQL 压缩 SQL 语句（去除多余空白）
func MinifySQL(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("SQL 内容为空")
	}
	tokens := tokenize(input)
	var out strings.Builder
	lastWasSpace := false
	for _, t := range tokens {
		if t == " " {
			if !lastWasSpace {
				out.WriteString(" ")
				lastWasSpace = true
			}
		} else {
			out.WriteString(t)
			lastWasSpace = false
		}
	}
	return strings.TrimSpace(out.String()), nil
}

type tokenType int

const (
	tokKeyword tokenType = iota
	tokIdentifier
	tokString
	tokNumber
	tokOperator
	tokParenOpen
	tokParenClose
	tokComma
	tokSemicolon
	tokWhitespace
)

type token struct {
	text string
	typ  tokenType
}

func tokenize(input string) []string {
	var tokens []string
	current := strings.Builder{}
	flush := func() {
		if current.Len() > 0 {
			tokens = append(tokens, current.String())
			current.Reset()
		}
	}

	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		switch {
		case ch == '\'' || ch == '"':
			flush()
			current.WriteRune(ch)
			quote := ch
			for i++; i < len(runes); i++ {
				current.WriteRune(runes[i])
				if runes[i] == quote && (i+1 >= len(runes) || runes[i+1] != quote) {
					break
				}
			}
			flush()

		case ch == '-' && i+1 < len(runes) && runes[i+1] == '-':
			flush()
			current.WriteString("--")
			for i += 2; i < len(runes); i++ {
				if runes[i] == '\n' {
					i--
					break
				}
				current.WriteRune(runes[i])
			}
			flush()

		case ch == '/' && i+1 < len(runes) && runes[i+1] == '*':
			flush()
			current.WriteString("/*")
			for i += 2; i < len(runes); i++ {
				current.WriteRune(runes[i])
				if runes[i] == '*' && i+1 < len(runes) && runes[i+1] == '/' {
					current.WriteRune('/')
					i++
					break
				}
			}
			flush()

		case ch == '(':
			flush()
			tokens = append(tokens, "(")
		case ch == ')':
			flush()
			tokens = append(tokens, ")")
		case ch == ',':
			flush()
			tokens = append(tokens, ",")
		case ch == ';':
			flush()
			tokens = append(tokens, ";")
		case unicode.IsSpace(ch):
			flush()
			tokens = append(tokens, " ")
		default:
			current.WriteRune(ch)
		}
	}
	flush()
	return tokens
}

func formatTokens(tokens []string) string {
	var out strings.Builder
	indent := 0
	parenDepth := 0
	needNewline := false
	prevWasKeyword := false
	prevWasComma := false

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t == " " {
			continue
		}

		upper := strings.ToUpper(t)
		isKeyword := sqlKeywords[upper]
		isNewlineKW := newlineBeforeKeywords[upper]

		if t == "(" {
			parenDepth++
		}
		if t == ")" {
			parenDepth--
		}

		if isNewlineKW && !prevWasKeyword {
			needNewline = true
		}
		if t == "," {
			prevWasComma = true
		}

		if needNewline {
			out.WriteString("\n")
			if prevWasComma {
				out.WriteString(strings.Repeat("    ", indent))
			} else {
				out.WriteString(strings.Repeat("    ", indent))
			}
			needNewline = false
		}

		if t == "(" && prevWasKeyword {
			out.WriteString(" ")
		}
		out.WriteString(t)

		if t == "(" {
			indent++
			if i+1 < len(tokens) && tokens[i+1] != ")" {
				out.WriteString("\n")
				out.WriteString(strings.Repeat("    ", indent))
			}
		}
		if t == ")" {
			indent--
			if indent < 0 {
				indent = 0
			}
		}

		prevWasKeyword = isKeyword
		if t != "," {
			prevWasComma = false
		}

		if isKeyword && upper != "AS" {
			out.WriteString(" ")
		} else if t == "," {
			out.WriteString(" ")
		}
	}

	result := strings.TrimSpace(out.String())
	return result
}
