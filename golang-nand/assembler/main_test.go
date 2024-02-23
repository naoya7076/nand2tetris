package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// MaxL.asmを変換してMaxL.hackと比較する
func TestAssemble(t *testing.T) {
	// テストデータのパスを取得
	inPath := filepath.Join("./asm/max/MaxL.asm")
	inFile, err := os.Open(inPath)
	if err != nil {
		t.Fatal(err)
	}
	defer inFile.Close()
	// 出力ファイルのパスを取得
	outPath := filepath.Join("./asm/max/dist/MaxL.hack")
	outFile, err := os.Create(outPath)
	if err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()
	// Assemble関数を実行
	Assemble(inFile, outFile)
	comp, err := os.ReadFile("./asm/max/MaxL.hack")
	if err != nil {
		t.Fatal(err)
	}
	out, err := os.ReadFile("./asm/max/dist/MaxL.hack")
	if err != nil {
		t.Fatal(err)
	}
	// 出力ファイルと期待するファイルを比較
	if !bytes.Equal(comp, out) {
		t.Fatalf("expected %s, but got %s", comp, out)
	}
}
