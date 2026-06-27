package services

import (
	"encoding/json"
	"testing"
)

func TestJSONToYAML(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		check   func(t *testing.T, got string)
		wantErr bool
	}{
		{
			name:  "simple object",
			input: `{"name":"test","value":42}`,
			check: func(t *testing.T, got string) {
				if got != "name: test\nvalue: 42\n" && got != "value: 42\nname: test\n" {
					t.Errorf("unexpected yaml output: %q", got)
				}
			},
		},
		{
			name:  "nested object",
			input: `{"a":{"b":1}}`,
			check: func(t *testing.T, got string) {
				if got != "a:\n    b: 1\n" {
					t.Errorf("unexpected yaml output: %q", got)
				}
			},
		},
		{
			name:  "array",
			input: `{"items":["a","b"]}`,
			check: func(t *testing.T, got string) {
				if got != "items:\n    - a\n    - b\n" {
					t.Errorf("unexpected yaml output: %q", got)
				}
			},
		},
		{
			name:    "invalid json",
			input:   `{invalid}`,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSONToYAML(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONToYAML() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.check != nil {
				tt.check(t, got)
			}
		})
	}
}

func TestYAMLToJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "simple object",
			input: "name: test\nvalue: 42\n",
			want:  "{\n  \"name\": \"test\",\n  \"value\": 42\n}",
		},
		{
			name:  "nested object",
			input: "a:\n  b: 1\n",
			want:  "{\n  \"a\": {\n    \"b\": 1\n  }\n}",
		},
		{
			name:  "array",
			input: "items:\n  - a\n  - b\n",
			want:  "{\n  \"items\": [\n    \"a\",\n    \"b\"\n  ]\n}",
		},
		{
			name:  "boolean",
			input: "flag: true\n",
			want:  "{\n  \"flag\": true\n}",
		},
		{
			name:  "empty string",
			input: "",
			want:  "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YAMLToJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("YAMLToJSON() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !jsonEqual(got, tt.want) {
				t.Errorf("YAMLToJSON() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestJSONToYAMLRoundTrip(t *testing.T) {
	input := `{"name":"test","nested":{"a":1,"b":2},"items":["x","y"],"flag":true}`
	yaml, err := JSONToYAML(input)
	if err != nil {
		t.Fatal(err)
	}
	jsonOut, err := YAMLToJSON(yaml)
	if err != nil {
		t.Fatal(err)
	}
	if !jsonEqual(jsonOut, input) {
		t.Errorf("round trip failed:\ninput:  %s\nyaml:   %s\noutput: %s", input, yaml, jsonOut)
	}
}

func TestYAMLToJSONCompact(t *testing.T) {
	input := "name: test\nvalue: 42\n"
	got, err := YAMLToJSONCompact(input)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"name":"test","value":42}`
	if got != want {
		t.Errorf("YAMLToJSONCompact() = %q, want %q", got, want)
	}
}

func jsonEqual(a, b string) bool {
	var va, vb interface{}
	if err := json.Unmarshal([]byte(a), &va); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &vb); err != nil {
		return false
	}
	ba, _ := json.Marshal(va)
	bb, _ := json.Marshal(vb)
	return string(ba) == string(bb)
}
