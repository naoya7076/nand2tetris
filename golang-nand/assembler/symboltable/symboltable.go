package symboltable

type SymbolTable struct {
	table      map[string]int
	romAddress int
	ramAddress int
}

var initialTable = map[string]int{
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
}

func New() *SymbolTable {
	return &SymbolTable{table: initialTable, romAddress: 0, ramAddress: 16}
}

func (st *SymbolTable) AddEntry(symbol string, address int) {
	st.table[symbol] = address
}

func (st *SymbolTable) Contains(symbol string) bool {
	_, ok := st.table[symbol]
	return ok
}

func (st *SymbolTable) GetAddress(symbol string) int {
	address := st.table[symbol]
	return address
}

func (st *SymbolTable) IncRomAddress() {
	st.romAddress++
}

func (st *SymbolTable) IncRamAddress() {
	st.ramAddress++
}

func (st *SymbolTable) GetRomAddress() int {
	return st.romAddress
}

func (st *SymbolTable) GetRamAddress() int {
	return st.ramAddress
}
