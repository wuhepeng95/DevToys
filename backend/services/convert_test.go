package services

import (
	"testing"
)

func TestConvertBase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		fromBase int
		want     ConvertNumber
		wantErr  bool
	}{
		{
			name:     "decimal 42",
			input:    "42",
			fromBase: 10,
			want:     ConvertNumber{Binary: "101010", Octal: "52", Decimal: "42", Hexadecimal: "2a"},
		},
		{
			name:     "binary 1010",
			input:    "1010",
			fromBase: 2,
			want:     ConvertNumber{Binary: "1010", Octal: "12", Decimal: "10", Hexadecimal: "a"},
		},
		{
			name:     "hex ff",
			input:    "ff",
			fromBase: 16,
			want:     ConvertNumber{Binary: "11111111", Octal: "377", Decimal: "255", Hexadecimal: "ff"},
		},
		{
			name:     "octal 77",
			input:    "77",
			fromBase: 8,
			want:     ConvertNumber{Binary: "111111", Octal: "77", Decimal: "63", Hexadecimal: "3f"},
		},
		{
			name:     "hex with 0x prefix",
			input:    "0xFF",
			fromBase: 16,
			want:     ConvertNumber{Binary: "11111111", Octal: "377", Decimal: "255", Hexadecimal: "ff"},
		},
		{
			name:     "large number",
			input:    "65535",
			fromBase: 10,
			want:     ConvertNumber{Binary: "1111111111111111", Octal: "177777", Decimal: "65535", Hexadecimal: "ffff"},
		},
		{
			name:     "zero",
			input:    "0",
			fromBase: 10,
			want:     ConvertNumber{Binary: "0", Octal: "0", Decimal: "0", Hexadecimal: "0"},
		},
		{
			name:     "invalid base",
			input:    "123",
			fromBase: 3,
			wantErr:  true,
		},
		{
			name:     "invalid number",
			input:    "XYZ",
			fromBase: 10,
			wantErr:  true,
		},
		{
			name:     "empty input",
			input:    "",
			fromBase: 10,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertBase(tt.input, tt.fromBase)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertBase() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Binary != tt.want.Binary {
					t.Errorf("Binary = %q, want %q", got.Binary, tt.want.Binary)
				}
				if got.Octal != tt.want.Octal {
					t.Errorf("Octal = %q, want %q", got.Octal, tt.want.Octal)
				}
				if got.Decimal != tt.want.Decimal {
					t.Errorf("Decimal = %q, want %q", got.Decimal, tt.want.Decimal)
				}
				if got.Hexadecimal != tt.want.Hexadecimal {
					t.Errorf("Hexadecimal = %q, want %q", got.Hexadecimal, tt.want.Hexadecimal)
				}
			}
		})
	}
}
