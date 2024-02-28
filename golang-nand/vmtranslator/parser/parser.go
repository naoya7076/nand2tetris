package parser

import (
	"strconv"
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
	return p.currentCommandIdx < len(p.commandList)
}

func (p *Parser) Advance() {
	p.currentCommandIdx++
}

func (p *Parser) Arg1() string {
	currentCommand := p.commandList[p.currentCommandIdx]
	if strings.Contains(currentCommand, "push") {
		return strings.Split(currentCommand, " ")[1]
	} else {
		return currentCommand
	}
}

func (p *Parser) Arg2() int {
	currentCommand := p.commandList[p.currentCommandIdx]
	arg2 := strings.Split(currentCommand, " ")[2]
	result, err := strconv.Atoi(arg2)
	if err != nil {
		panic(err)
	}
	return result
}

// C_ARITHMETIC、C_PUSH、C_POP、C_LABEL、C_GOTO、
// C_IF、C_FUNCTION、C_RETURN、C_CALL
func (p *Parser) CommandType() string {
	currentCommand := p.commandList[p.currentCommandIdx]
	if strings.Contains(currentCommand, "push") { // strings.startsWithかも
		return "C_PUSH"
	} else {
		return "C_ARITHMETIC"
	}
}
