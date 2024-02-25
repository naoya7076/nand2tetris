package main

import (
	"bytes"
	"os"
	"testing"
)

func TestAssemble(t *testing.T) {
	cases := []struct {
		inPath   string
		outPath  string
		basePath string
	}{
		{"./asm/max/MaxL.asm", "./asm/max/MaxL.hack", "./testdata/MaxL.hack"},
		{"./asm/pong/PongL.asm", "./asm/pong/PongL.hack", "./testdata/PongL.hack"},
		{"./asm/rect/RectL.asm", "./asm/rect/RectL.hack", "./testdata/RectL.hack"},
		{"./asm/add/Add.asm", "./asm/add/Add.hack", "./testdata/Add.hack"},
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
		trimmedText := trimFileText(inFile)
		Assemble(trimmedText, outFile)
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
