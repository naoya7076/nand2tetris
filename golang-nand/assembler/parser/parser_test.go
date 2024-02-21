package parser_test

import (
	"assembler/parser"
	"testing"
)

func TestCommandType(t *testing.T) {
	want := "A_COMMAND"
	got := parser.New("@100").CommandType()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "L_COMMAND"
	got = parser.New("(LOOP)").CommandType()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "C_COMMAND"
	got = parser.New("D=M").CommandType()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestHasMoreCommands(t *testing.T) {
	p := parser.New("@100")
	if !p.HasMoreCommands() {
		t.Errorf("want true, got false")
	}
	p.Advance()
	if p.HasMoreCommands() {
		t.Errorf("want false, got true")
	}
}

func TestSymbol(t *testing.T) {
	want := "100"
	got := parser.New("@100").Symbol()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "LOOP"
	got = parser.New("(LOOP)").Symbol()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestDest(t *testing.T) {
	want := "D"
	got := parser.New("D=M").Dest()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "null"
	got = parser.New("M;JMP").Dest()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestJump(t *testing.T) {
	want := "JMP"
	got := parser.New("D;JMP").Jump()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "null"
	got = parser.New("D=M").Jump()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestComp(t *testing.T) {
	want := "C"
	got := parser.New("D=C").Comp()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "C"
	got = parser.New("C;JMP").Comp()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "M"
	got = parser.New("D=M").Comp()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
	want = "D+1"
	got = parser.New("D=D+1").Comp()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
