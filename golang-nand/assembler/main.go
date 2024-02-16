package main

import (
	"assembler/code"
	"assembler/parser"
	"bufio"
	"fmt"
	"os"
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
	// 入力ファイルのテキストからコメントアウトと空白行を削除
	trimmedSrc := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		// ~コメントアウトと空白行をfileから削除
		if len(text) == 0 || text[0] == '/' || text[0] == ' ' {
			continue
		}
		trimmedSrc += text + "\n"
	}
	// 出力ファイルを作成
	outputFileName := fileName[:len(fileName)-3] + ".hack"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("File not found")
	}
	defer outputFile.Close()
	// 入力ファイルをパース
	// パースした結果を出力ファイルに書き込む
	p := parser.New(trimmedSrc)
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case "A_COMMAND":
			outputFile.WriteString(p.Symbol() + "\n")
		case "L_COMMAND":
			outputFile.WriteString(p.Symbol() + "\n")
		case "C_COMMAND":
			dest := code.Dest(p.Dest())
			comp := code.Comp(p.Comp())
			jump := code.Jump(p.Jump())
			outputFile.WriteString("111" + comp + dest + jump + "\n")
		default:
			fmt.Println("Command type not found")
		}
	}
	fmt.Println(outputFile)
}
