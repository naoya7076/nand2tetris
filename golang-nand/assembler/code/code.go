package code

func Dest(mnemonic string) string {
	mnemonics := map[string]string{
		"null": "000",
		"M":    "001",
		"D":    "010",
		"MD":   "011",
		"A":    "100",
		"AM":   "101",
		"AD":   "110",
		"AMD":  "111",
	}
	return mnemonics[mnemonic]
}

