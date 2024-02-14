package parser

import (
	"strings"
)

type Parser struct {
	currentCommandIdx int
	commandList       []string
}

func New(source string) *Parser {
	return &Parser{currentCommandIdx: 0, commandList: strings.Split(source, "\n")}
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
	// 現在のコマンドのシンボルまたは10進数の値を返す
	currentCommand := p.commandList[p.currentCommandIdx]
	if p.CommandType() == "A_COMMAND" {
		sym := strings.TrimPrefix(currentCommand, "@")
		return sym
	} else if p.CommandType() == "L_COMMAND" {
		sym := strings.TrimSuffix(strings.TrimPrefix(currentCommand, "("), ")")
		return sym
	} else {
		return ""
	}
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
