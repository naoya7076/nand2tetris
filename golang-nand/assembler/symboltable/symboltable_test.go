package symboltable_test

import (
	"assembler/symboltable"
	"testing"
)

func TestNew(t *testing.T) {
	st := symboltable.New()
	if st == nil {
		t.Errorf("symboltable.New() returned nil")
	}
}
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

func TestIncRomAddress(t *testing.T) {
	st := symboltable.New()
	if st.GetRomAddress() != 0 {
		t.Errorf("expected 0, but got %d", st.GetRomAddress())
	}
	st.IncRomAddress()
	if st.GetRomAddress() != 1 {
		t.Errorf("expected 1, but got %d", st.GetRomAddress())
	}
}
