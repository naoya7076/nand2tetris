package parser_test

import (
	"testing"
	"vmtranslator/parser"
)

func TestHasMoreCommands(t *testing.T) {
	want := true
	p := parser.New("push constant 10")
	got := p.HasMoreCommands()
	if got != want {
		t.Errorf("want %t, got %t", want, got)
	}
	want = false
	p.Advance()
	got = p.HasMoreCommands()
	if got != want {
		t.Errorf("want %t, got %t", want, got)
	}
}
func TestCommandType(t *testing.T) {
	want := "C_PUSH"
	got := parser.New("push constant 10").CommandType()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "C_ARITHMETIC"
	got = parser.New("add").CommandType()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestArg1(t *testing.T) {
	want := "constant"
	got := parser.New("push constant 10").Arg1()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "add"
	got = parser.New("add").Arg1()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestArg2(t *testing.T) {
	want := 10
	got := parser.New("push constant 10").Arg2()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
