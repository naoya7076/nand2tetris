package main

import (
	"bytes"
	"os"
	"testing"
)

// MaxL.asmを変換してMaxL.hackと比較する
func TestAssemble(t *testing.T) {
	cases := []struct {
		inPath   string
		outPath  string
		basePath string
	}{
		{"./asm/max/MaxL.asm", "./asm/max/MaxL.hack", "./testdata/MaxL.hack"},
	}
	for _, tt := range cases {
		inFile, err := os.Open(tt.inPath)
		if err != nil {
			t.Fatal(err)
		}
		defer inFile.Close()
		outFile, err := os.Create(tt.outPath)
		if err != nil {
			t.Fatal(err)
		}
		defer outFile.Close()
		Assemble(inFile, outFile)
		// 出力ファイルと期待するファイルを比較
		base, err := os.ReadFile(tt.basePath)
		if err != nil {
			t.Fatal(err)
		}
		out, err := os.ReadFile(tt.outPath)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(base, out) {
			t.Fatalf("expected %s, but got %s", base, out)
		}
	}
}
