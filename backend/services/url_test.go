package services

import "testing"

func TestURLEncode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "simple string",
			input:   "hello world",
			want:    "hello+world",
			wantErr: false,
		},
		{
			name:    "chinese characters",
			input:   "你好",
			want:    "%E4%BD%A0%E5%A5%BD",
			wantErr: false,
		},
		{
			name:    "special chars",
			input:   "a=b&c=d",
			want:    "a%3Db%26c%3Dd",
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "url with spaces",
			input:   "https://example.com/path with spaces",
			want:    "https%3A%2F%2Fexample.com%2Fpath+with+spaces",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := URLEncode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLEncode() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URLEncode() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestURLDecode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "plus as space",
			input:   "hello+world",
			want:    "hello world",
			wantErr: false,
		},
		{
			name:    "percent encoded chinese",
			input:   "%E4%BD%A0%E5%A5%BD",
			want:    "你好",
			wantErr: false,
		},
		{
			name:    "special chars",
			input:   "a%3Db%26c%3Dd",
			want:    "a=b&c=d",
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "invalid percent encoding",
			input:   "%ZZ",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := URLDecode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLDecode() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URLDecode() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestQueryEncode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "simple query",
			input:   "a=1&b=2",
			want:    "a=1&b=2",
			wantErr: false,
		},
		{
			name:    "values with special chars",
			input:   "name=hello world&key=a=b",
			want:    "key=a%3Db&name=hello+world",
			wantErr: false,
		},
		{
			name:    "empty input",
			input:   "",
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryEncode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryEncode() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QueryEncode() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestQueryDecode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "simple query",
			input:   "a=1&b=2",
			want:    "a=1&b=2",
			wantErr: false,
		},
		{
			name:    "encoded values",
			input:   "key=a%3Db&name=hello+world",
			want:    "key=a=b&name=hello world",
			wantErr: false,
		},
		{
			name:    "empty input",
			input:   "",
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryDecode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryDecode() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QueryDecode() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestParseURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    URLParts
		wantErr bool
	}{
		{
			name:  "full url",
			input: "https://example.com:443/path/to/page?name=hello&key=val#section",
			want: URLParts{
				Scheme:   "https",
				Host:     "example.com",
				Port:     "443",
				Path:     "/path/to/page",
				Query:    "name=hello&key=val",
				Fragment: "section",
				Raw:      "https://example.com:443/path/to/page?name=hello&key=val#section",
			},
			wantErr: false,
		},
		{
			name:  "simple url",
			input: "https://example.com",
			want: URLParts{
				Scheme: "https",
				Host:   "example.com",
				Port:   "",
				Path:   "",
				Raw:    "https://example.com",
			},
			wantErr: false,
		},
		{
			name:  "url with query only",
			input: "http://test.com/api?v=1",
			want: URLParts{
				Scheme: "http",
				Host:   "test.com",
				Port:   "",
				Path:   "/api",
				Query:  "v=1",
				Raw:    "http://test.com/api?v=1",
			},
			wantErr: false,
		},
		{
			name:    "invalid url",
			input:   "://invalid",
			want:    URLParts{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseURL() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got.Scheme != tt.want.Scheme {
					t.Errorf("ParseURL() Scheme = %q, want %q", got.Scheme, tt.want.Scheme)
				}
				if got.Host != tt.want.Host {
					t.Errorf("ParseURL() Host = %q, want %q", got.Host, tt.want.Host)
				}
				if got.Path != tt.want.Path {
					t.Errorf("ParseURL() Path = %q, want %q", got.Path, tt.want.Path)
				}
				if got.Query != tt.want.Query {
					t.Errorf("ParseURL() Query = %q, want %q", got.Query, tt.want.Query)
				}
			}
		})
	}
}
