package services

import "testing"

func TestCompareText(t *testing.T) {
	tests := []struct {
		name  string
		left  string
		right string
		want  int
	}{
		{
			name:  "identical texts",
			left:  "a\nb\nc",
			right: "a\nb\nc",
			want:  3,
		},
		{
			name:  "added line",
			left:  "a\nc",
			right: "a\nb\nc",
			want:  3,
		},
		{
			name:  "removed line",
			left:  "a\nb\nc",
			right: "a\nc",
			want:  3,
		},
		{
			name:  "modified line",
			left:  "a\nb\nc",
			right: "a\nB\nc",
			want:  4,
		},
		{
			name:  "completely different",
			left:  "a\nb",
			right: "c\nd",
			want:  4,
		},
		{
			name:  "empty left",
			left:  "",
			right: "a\nb",
			want:  3,
		},
		{
			name:  "empty right",
			left:  "a\nb",
			right: "",
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CompareText(tt.left, tt.right)
			if err != nil {
				t.Errorf("CompareText() error = %v", err)
				return
			}
			if len(result.Lines) != tt.want {
				t.Errorf("CompareText() got %d lines, want %d", len(result.Lines), tt.want)
			}
		})
	}
}

func TestCompareTextContent(t *testing.T) {
	result, err := CompareText("a\nb\nc", "a\nB\nc")
	if err != nil {
		t.Fatal(err)
	}

	sameCount := 0
	addCount := 0
	removeCount := 0
	for _, l := range result.Lines {
		switch l.Type {
		case DiffLineSame:
			sameCount++
		case DiffLineAdd:
			addCount++
		case DiffLineRemove:
			removeCount++
		}
	}

	if sameCount != 2 {
		t.Errorf("expected 2 same lines, got %d", sameCount)
	}
	if addCount != 1 {
		t.Errorf("expected 1 add line, got %d", addCount)
	}
	if removeCount != 1 {
		t.Errorf("expected 1 remove line, got %d", removeCount)
	}
}
