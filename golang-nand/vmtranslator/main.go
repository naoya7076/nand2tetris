package main

import (
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
