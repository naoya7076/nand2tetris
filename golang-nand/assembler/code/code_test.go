package code_test

import (
	"assembler/code"
	"testing"
)

func TestDest(t *testing.T) {
	cases := []struct {
		mnemonic string
		want     string
	}{
		{"null", "000"},
		{"M", "001"},
		{"D", "010"},
		{"MD", "011"},
		{"A", "100"},
		{"AM", "101"},
		{"AD", "110"},
		{"AMD", "111"},
	}
	for _, tt := range cases {
		binaryCode := code.Dest(tt.mnemonic)
		if binaryCode != tt.want {
			t.Errorf("want %s, got %s", tt.want, binaryCode)
		}
	}
}
