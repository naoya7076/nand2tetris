package symboltable_test

import (
	"assembler/symboltable"
	"testing"
)

func TestAddEntry(t *testing.T) {
	st := symboltable.New()
	symbol := "test"
	st.AddEntry(symbol, 1)
	if !st.Contains(symbol) {
		t.Errorf("%s not found", symbol)
	}
}

func TestContains(t *testing.T) {
	st := symboltable.New()
	symbol := "test"
	if st.Contains(symbol) {
		t.Errorf("%s found", symbol)
	}
}

func TestGetAddress(t *testing.T) {
	st := symboltable.New()
	symbol := "test"
	st.AddEntry(symbol, 1)
	if st.GetAddress(symbol) != 1 {
		t.Errorf("expected 1, but got %d", st.GetAddress(symbol))
	}
}
