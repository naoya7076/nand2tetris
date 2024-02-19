package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	currentCommandIdx int
	commandList       []string
}

func New(source string) *Parser {
	commandList := strings.Split(source, "\n")
	if commandList[len(commandList)-1] == "" {
		commandList = commandList[:len(commandList)-1]
	}
	return &Parser{currentCommandIdx: 0, commandList: commandList}
}

func (p *Parser) HasMoreCommands() bool {
	// 入力にまだコマンドが存在するか？
	return p.currentCommandIdx < len(p.commandList)
}

func (p *Parser) Advance() {
	p.currentCommandIdx++
}

func (p *Parser) CommandType() string {
	// 現在のコマンドの種類を返す
	currentCommand := p.commandList[p.currentCommandIdx]
	if strings.Contains(currentCommand, "@") {
		return "A_COMMAND"
	} else if strings.Contains(currentCommand, "(") {
		return "L_COMMAND"
	} else {
		return "C_COMMAND"
	}
}

func (p *Parser) Symbol() string {
	// 現在のコマンドのシンボルまたは10進数を2進数に変換した値を返す
	currentCommand := p.commandList[p.currentCommandIdx]
	if p.CommandType() == "A_COMMAND" {
		sym := strings.TrimPrefix(currentCommand, "@")
		return parseDecimalToBinary(sym)
	} else if p.CommandType() == "L_COMMAND" {
		sym := strings.TrimSuffix(strings.TrimPrefix(currentCommand, "("), ")")
		return sym
	} else {
		return ""
	}
}

func parseDecimalToBinary(decim string) string {
	// string to int
	// int to binary
	i, err := strconv.Atoi(decim)
	if err != nil {
		panic(err)
	}
	binaryString := fmt.Sprintf("0"+"%015b", i)
	return binaryString
}

func (p *Parser) Dest() string {
	currentCommand := p.commandList[p.currentCommandIdx]
	if !strings.Contains(currentCommand, "=") {
		return "null"
	} else {
		return strings.Split(currentCommand, "=")[0]
	}
}

func (p *Parser) Comp() string {
	currentCommand := p.commandList[p.currentCommandIdx]
	if strings.Contains(currentCommand, "=") {
		return strings.Split(currentCommand, "=")[1]
	} else {
		return strings.Split(currentCommand, ";")[0]
	}
}

func (p *Parser) Jump() string {
	currentCommand := p.commandList[p.currentCommandIdx]
	if !strings.Contains(currentCommand, ";") {
		return "null"
	} else {
		return strings.Split(p.commandList[p.currentCommandIdx], ";")[1]
	}
}
