package services

import "testing"

func TestEncodeBase64(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		urlSafe bool
		want    string
		wantErr bool
	}{
		{
			name:    "standard encode",
			input:   "Hello, World!",
			urlSafe: false,
			want:    "SGVsbG8sIFdvcmxkIQ==",
		},
		{
			name:    "url safe encode",
			input:   "a?b=c&d",
			urlSafe: true,
			want:    "YT9iPWMmZA",
		},
		{
			name:    "empty string",
			input:   "",
			urlSafe: false,
			want:    "",
		},
		{
			name:    "unicode",
			input:   "你好",
			urlSafe: false,
			want:    "5L2g5aW9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeBase64(tt.input, tt.urlSafe)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeBase64() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeBase64() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDecodeBase64(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		urlSafe bool
		want    string
		wantErr bool
	}{
		{
			name:    "standard decode",
			input:   "SGVsbG8sIFdvcmxkIQ==",
			urlSafe: false,
			want:    "Hello, World!",
		},
		{
			name:    "url safe decode",
			input:   "YT9iPWMmZA",
			urlSafe: true,
			want:    "a?b=c&d",
		},
		{
			name:    "empty string",
			input:   "",
			urlSafe: false,
			want:    "",
		},
		{
			name:    "invalid input",
			input:   "!!!not-base64!!!",
			urlSafe: false,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeBase64(tt.input, tt.urlSafe)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBase64() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeBase64() = %q, want %q", got, tt.want)
			}
		})
	}
}
