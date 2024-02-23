package main

import (
	"assembler/code"
	"assembler/parser"
	"bufio"
	"fmt"
	"os"
	"strings"
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
	outputFileName := fileName[:len(fileName)-3] + "sample.hack"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("File not found")
	}
	defer outputFile.Close()

	Assemble(file, outputFile)
}

func Assemble(in *os.File, out *os.File) {
	// 入力ファイルのテキストからコメントアウトと空白行を削除
	trimmedSrc := ""
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		text := scanner.Text()
		// ~コメントアウトと空白行をfileから削除
		if len(text) == 0 || text[0] == '/' || text[0] == ' ' {
			continue
		}
		trimmedSrc += strings.ReplaceAll(text, " ", "") + "\n"
	}
	// 入力ファイルをパース
	// パースした結果を出力ファイルに書き込む
	p := parser.New(trimmedSrc)
	for p.HasMoreCommands() {
		switch p.CommandType() {
		case "A_COMMAND":
			out.WriteString(p.Symbol() + "\n")
		case "L_COMMAND":
			out.WriteString(p.Symbol() + "\n")
		case "C_COMMAND":
			dest := code.Dest(p.Dest())
			comp := code.Comp(p.Comp())
			jump := code.Jump(p.Jump())
			out.WriteString("111" + comp + dest + jump + "\n")
		default:
			fmt.Println("Command type not found")
		}
		p.Advance()
	}
}
