package cartridge

func getDestinationCode(code byte) bool {
	return code == 0x00
}

// IsJapanese returns "Yes" if the ROM is marked as Japanese, otherwise "No".
func IsJapanese(rom Rom) string {
	s := "No"
	if rom.IsJapanese {
		s = "Yes"
	}
	return s
}
