package services

import "testing"

func TestRemoveDuplicate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		opts    RemoveDuplicateOptions
		want    string
		wantErr bool
	}{
		{
			name:  "basic dedup keep order",
			input: "a\nb\na\nc",
			opts:  RemoveDuplicateOptions{KeepOrder: true, RemoveEmpty: false, IgnoreCase: false},
			want:  "a\nb\nc",
		},
		{
			name:  "dedup with sort",
			input: "c\na\nb\na",
			opts:  RemoveDuplicateOptions{KeepOrder: false, RemoveEmpty: false, IgnoreCase: false},
			want:  "a\nb\nc",
		},
		{
			name:  "dedup ignore case",
			input: "A\nb\na\nB",
			opts:  RemoveDuplicateOptions{KeepOrder: true, RemoveEmpty: false, IgnoreCase: true},
			want:  "A\nb",
		},
		{
			name:  "dedup remove empty",
			input: "a\n\nb\n\nc",
			opts:  RemoveDuplicateOptions{KeepOrder: true, RemoveEmpty: true, IgnoreCase: false},
			want:  "a\nb\nc",
		},
		{
			name:  "empty input",
			input: "",
			opts:  RemoveDuplicateOptions{KeepOrder: true, RemoveEmpty: false, IgnoreCase: false},
			want:  "",
		},
		{
			name:  "all duplicates",
			input: "x\nx\nx",
			opts:  RemoveDuplicateOptions{KeepOrder: true, RemoveEmpty: false, IgnoreCase: false},
			want:  "x",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveDuplicate(tt.input, tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveDuplicate() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RemoveDuplicate() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSortLines(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		opts    SortLinesOptions
		want    string
		wantErr bool
	}{
		{
			name:  "asc sort",
			input: "c\na\nb",
			opts:  SortLinesOptions{Order: "asc", Numeric: false, IgnoreCase: false},
			want:  "a\nb\nc",
		},
		{
			name:  "desc sort",
			input: "a\nb\nc",
			opts:  SortLinesOptions{Order: "desc", Numeric: false, IgnoreCase: false},
			want:  "c\nb\na",
		},
		{
			name:  "numeric sort",
			input: "10\n2\n1\n20",
			opts:  SortLinesOptions{Order: "asc", Numeric: true, IgnoreCase: false},
			want:  "1\n2\n10\n20",
		},
		{
			name:  "ignore case sort",
			input: "B\na\nC",
			opts:  SortLinesOptions{Order: "asc", Numeric: false, IgnoreCase: true},
			want:  "a\nB\nC",
		},
		{
			name:  "empty input",
			input: "",
			opts:  SortLinesOptions{Order: "asc", Numeric: false, IgnoreCase: false},
			want:  "",
		},
		{
			name:  "mixed strings and numbers numeric off",
			input: "10\n2\nx\n1",
			opts:  SortLinesOptions{Order: "asc", Numeric: false, IgnoreCase: false},
			want:  "1\n10\n2\nx",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SortLines(tt.input, tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("SortLines() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SortLines() = %q, want %q", got, tt.want)
			}
		})
	}
}
