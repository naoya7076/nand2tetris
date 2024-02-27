package main

import (
	"assembler/code"
	"assembler/parser"
	"assembler/symboltable"
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

	trimmedText := trimFileText(file)
	Assemble(trimmedText, outputFile)
}

func Assemble(in string, out *os.File) {
	// 入力ファイルのテキストからコメントアウトと空白行を削除
	// 入力ファイルをパース
	// パースした結果を出力ファイルに書き込む
	p, p2 := parser.New(in), parser.New(in)
	st := firstPass(p)
	for p2.HasMoreCommands() {
		switch p2.CommandType() {
		case "A_COMMAND":
			sym, err := p2.Symbol()
			if err != nil {
				// symbolの場合
				if st.Contains(sym) {
					out.WriteString("0" + fmt.Sprintf("%015b", st.GetAddress(sym)) + "\n")
				} else {
					out.WriteString("0" + fmt.Sprintf("%015b", st.GetRamAddress()) + "\n")
					st.AddEntry(sym, st.GetRamAddress())
					st.IncRamAddress()
				}
			} else {
				// 10進数の場合
				out.WriteString(sym + "\n")
			}
		case "L_COMMAND":
			sym, err := p2.Symbol()
			if err != nil {
				panic(err)
			}
			if !st.Contains(sym) {
				out.WriteString("0" + fmt.Sprintf("%015b", st.GetRamAddress()) + "\n")
				st.AddEntry(sym, st.GetRamAddress())
				st.IncRamAddress()
			}
		case "C_COMMAND":
			dest := code.Dest(p2.Dest())
			comp := code.Comp(p2.Comp())
			jump := code.Jump(p2.Jump())
			out.WriteString("111" + comp + dest + jump + "\n")
		default:
			fmt.Println("Command type not found")
		}
		p2.Advance()
	}
}

func firstPass(p *parser.Parser) *symboltable.SymbolTable {
	st := symboltable.New()
	// 1回目のパスでラベルを探し、シンボルテーブルに追加する
	for p.HasMoreCommands() {
		switch p.CommandType() {
		case "A_COMMAND":
			st.IncRomAddress()
		case "L_COMMAND":
			symbol, err := p.Symbol()
			if err != nil {
				panic(err)
			}
			st.AddEntry(symbol, st.GetRomAddress())
		case "C_COMMAND":
			st.IncRomAddress()
		default:
			fmt.Println("Command type not found")
		}
		p.Advance()
	}
	return st
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
