package services

import (
	"strings"
	"testing"
)

func TestFormatSQL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		check   func(t *testing.T, got string)
		wantErr bool
	}{
		{
			name:  "simple select",
			input: "SELECT * FROM users WHERE id = 1",
			check: func(t *testing.T, got string) {
				if !containsSubseq(got, "SELECT", "FROM", "WHERE") {
					t.Errorf("formatted SQL missing keywords: %q", got)
				}
				if strings.Count(got, "\n") < 2 {
					t.Errorf("expected at least 2 newlines, got %q", got)
				}
			},
		},
		{
			name:  "select with join",
			input: "SELECT u.name, o.total FROM users u INNER JOIN orders o ON u.id = o.user_id WHERE u.active = 1",
			check: func(t *testing.T, got string) {
				if len(got) < 50 {
					t.Errorf("formatted SQL too short: %q", got)
				}
			},
		},
		{
			name:  "insert statement",
			input: "INSERT INTO users (name, email) VALUES ('Alice', 'alice@test.com')",
			check: func(t *testing.T, got string) {
				if len(got) < 30 {
					t.Errorf("formatted SQL too short: %q", got)
				}
			},
		},
		{
			name:  "minify complex",
			input: "SELECT * FROM (SELECT id, name FROM users) AS sub WHERE sub.id > 10",
			check: func(t *testing.T, got string) {
				if len(got) < 30 {
					t.Errorf("formatted SQL too short: %q", got)
				}
			},
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:    "whitespace only",
			input:   "   \n  \t  ",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatSQL(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatSQL() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.check != nil {
				tt.check(t, got)
			}
		})
	}
}

func TestMinifySQL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "simple select",
			input: "SELECT * FROM users WHERE id = 1",
			want:  "SELECT * FROM users WHERE id = 1",
		},
		{
			name:  "extra whitespace",
			input: "SELECT   *   FROM   users",
			want:  "SELECT * FROM users",
		},
		{
			name:  "multi line",
			input: "SELECT\n  *\nFROM\n  users",
			want:  "SELECT * FROM users",
		},
		{
			name:  "string literals preserved",
			input: "SELECT 'hello  world' AS msg",
			want:  "SELECT 'hello  world' AS msg",
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinifySQL(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinifySQL() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("MinifySQL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func containsSubseq(s string, subs ...string) bool {
	pos := 0
	for _, sub := range subs {
		idx := strings.Index(s[pos:], sub)
		if idx < 0 {
			return false
		}
		pos += idx + len(sub)
	}
	return true
}
