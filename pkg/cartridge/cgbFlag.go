package cartridge

import "fmt"

// HasCGB determines if the ROM supports Color Game Boy (CGB) functions.
// It returns a descriptive string based on the CGB flag in the ROM header.
func HasCGB(rom Rom) string {
	switch rom.Header.CGBFlag {
	case 0x80:
		return fmt.Sprintf("0x%02X - Game supports CGB functions, but also works on old Game Boys.", rom.Header.CGBFlag)
	case 0xC0:
		return fmt.Sprintf("0x%02X - Game works on CGB only (physically the same as $80).", rom.Header.CGBFlag)
	default:
		return "No"
	}
}
