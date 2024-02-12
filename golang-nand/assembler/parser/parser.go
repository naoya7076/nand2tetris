package parser

type Parser struct {
	input string
}

func New(source string) *Parser {
	return &Parser{input: source}
}

func (p *Parser) hasMoreCommands() bool {
	return true
}
