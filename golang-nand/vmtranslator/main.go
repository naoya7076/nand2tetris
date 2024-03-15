package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vmtranslator/parser"
)

func main() {
	// コマンドラインの引数から入力ファイルの名前を受け取る
	// 入力ファイルを開く
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
	}
	defer file.Close()
	// 出力ファイルを作成
	outputFileName := fileName[:len(fileName)-3] + "asm"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("File not found")
	}
	defer outputFile.Close()

	trimmedText := trimFileText(file)
	Translate(trimmedText, outputFile)
}

func Translate(in string, out *os.File) {
	// 入力ファイルを1行ずつ読み込み、Parserに渡す
	// Parserから得た結果をCodeWriterに渡す
	parser := parser.New(in)
	// codeWriter := New(out)
	for parser.HasMoreCommands() {
		parser.Advance()
		switch parser.CommandType() {
		case "C_ARITHMETIC":
			// codeWriter.WriteArithmetic(parser.Arg1())
		case "C_PUSH":
			// codeWriter.WritePushPop("push", parser.Arg1(), parser.Arg2())
		case "C_POP":
			// codeWriter.WritePushPop("pop", parser.Arg1(), parser.Arg2())
		}
	}
}

func trimFileText(input *os.File) string {
	out := ""
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		// コメントアウトと空行を削除
		if len(text) == 0 || text[0] == '/' {
			continue
		}
		out += text + "\n"
	}
	return out
}
