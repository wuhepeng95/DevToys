package services

import "testing"

func TestFormatJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		indent  string
		want    string
		wantErr bool
	}{
		{
			name:    "valid json with 2-space indent",
			input:   `{"a":1,"b":2}`,
			indent:  "  ",
			want:    "{\n  \"a\": 1,\n  \"b\": 2\n}",
			wantErr: false,
		},
		{
			name:    "valid json with tab indent",
			input:   `{"a":1}`,
			indent:  "\t",
			want:    "{\n\t\"a\": 1\n}",
			wantErr: false,
		},
		{
			name:    "valid json with empty indent defaults to 2 spaces",
			input:   `{"a":1}`,
			indent:  "",
			want:    "{\n  \"a\": 1\n}",
			wantErr: false,
		},
		{
			name:    "invalid json",
			input:   `{invalid}`,
			indent:  "  ",
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			indent:  "  ",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatJSON(tt.input, tt.indent)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatJSON() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatJSON() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCompressJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "pretty json",
			input:   "{\n  \"a\": 1,\n  \"b\": 2\n}",
			want:    `{"a":1,"b":2}`,
			wantErr: false,
		},
		{
			name:    "already compressed",
			input:   `{"a":1}`,
			want:    `{"a":1}`,
			wantErr: false,
		},
		{
			name:    "invalid json",
			input:   `{invalid}`,
			want:    "",
			wantErr: true,
		},
		{
			name:    "nested json",
			input:   "{\n  \"a\": {\n    \"b\": 2\n  }\n}",
			want:    `{"a":{"b":2}}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompressJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompressJSON() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompressJSON() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestValidateJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		wantErr bool
	}{
		{
			name:    "valid json object",
			input:   `{"a":1}`,
			want:    true,
			wantErr: false,
		},
		{
			name:    "valid json array",
			input:   `[1,2,3]`,
			want:    true,
			wantErr: false,
		},
		{
			name:    "valid json string",
			input:   `"hello"`,
			want:    true,
			wantErr: false,
		},
		{
			name:    "valid json number",
			input:   `42`,
			want:    true,
			wantErr: false,
		},
		{
			name:    "invalid json",
			input:   `{invalid}`,
			want:    false,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJSON() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
