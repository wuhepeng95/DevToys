package services

import (
	"strings"
	"testing"
)

func TestGeneratePasswords(t *testing.T) {
	tests := []struct {
		name    string
		opts    PasswordOptions
		wantErr bool
		check   func(t *testing.T, result PasswordResult)
	}{
		{
			name: "basic password with all types",
			opts: PasswordOptions{
				Length:         12,
				IncludeUpper:   true,
				IncludeLower:   true,
				IncludeDigits:  true,
				IncludeSpecial: true,
				Count:          1,
			},
			wantErr: false,
			check: func(t *testing.T, r PasswordResult) {
				if len(r.Passwords) != 1 {
					t.Fatalf("expected 1 password, got %d", len(r.Passwords))
				}
				p := r.Passwords[0]
				if len(p) != 12 {
					t.Errorf("expected length 12, got %d", len(p))
				}
				if !containsAny(p, upperChars) {
					t.Error("password missing uppercase")
				}
				if !containsAny(p, lowerChars) {
					t.Error("password missing lowercase")
				}
				if !containsAny(p, digitChars) {
					t.Error("password missing digit")
				}
				if !containsAny(p, reducedSpecial) {
					t.Error("password missing special char")
				}
			},
		},
		{
			name: "exclude ambiguous chars",
			opts: PasswordOptions{
				Length:           16,
				IncludeUpper:     true,
				IncludeLower:     true,
				IncludeDigits:    true,
				ExcludeAmbiguous: true,
				Count:            1,
			},
			wantErr: false,
			check: func(t *testing.T, r PasswordResult) {
				p := r.Passwords[0]
				for _, c := range ambiguous {
					if strings.ContainsRune(p, c) {
						t.Errorf("password contains ambiguous char %c", c)
					}
				}
			},
		},
		{
			name: "generate multiple passwords",
			opts: PasswordOptions{
				Length:       8,
				IncludeUpper: true,
				IncludeLower: true,
				Count:        5,
			},
			wantErr: false,
			check: func(t *testing.T, r PasswordResult) {
				if len(r.Passwords) != 5 {
					t.Fatalf("expected 5 passwords, got %d", len(r.Passwords))
				}
				seen := make(map[string]bool)
				for _, p := range r.Passwords {
					if seen[p] {
						t.Error("duplicate password generated")
					}
					seen[p] = true
				}
			},
		},
		{
			name: "lowercase only",
			opts: PasswordOptions{
				Length:       10,
				IncludeLower: true,
				Count:        1,
			},
			wantErr: false,
			check: func(t *testing.T, r PasswordResult) {
				p := r.Passwords[0]
				for _, c := range p {
					if c < 'a' || c > 'z' {
						t.Errorf("unexpected char %c in lowercase-only password", c)
					}
				}
			},
		},
		{
			name: "length too short",
			opts: PasswordOptions{
				Length: 2,
			},
			wantErr: true,
		},
		{
			name: "no charset selected",
			opts: PasswordOptions{
				Length: 8,
				Count:  1,
			},
			wantErr: true,
		},
		{
			name: "too many passwords",
			opts: PasswordOptions{
				Length:       8,
				IncludeLower: true,
				Count:        200,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GeneratePasswords(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePasswords() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.check != nil {
				tt.check(t, result)
			}
		})
	}
}

func containsAny(s, chars string) bool {
	for _, c := range chars {
		if strings.ContainsRune(s, c) {
			return true
		}
	}
	return false
}
