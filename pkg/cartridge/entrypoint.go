package cartridge

import (
	"bytes"
	"fmt"
)

// usualEntryPoint is the standard entry point for Game Boy ROMs.
var usualEntryPoint = []byte{0, 195, 80, 1}

// hasUsualEntryPoint checks if the ROM has the usual entry point.
func hasUsualEntryPoint(rom Rom) bool {
	return bytes.Equal(rom.Header.Entrypoint, usualEntryPoint)
}

// GetEntryPoint returns the entry point of the ROM.
// It returns a string with the entry point and a description.
func GetEntryPoint(rom Rom) string {
	if len(rom.Header.Entrypoint) < 4 {
		return "entry point data is incomplete"
	}

	description := "not the most usual instruction"
	if hasUsualEntryPoint(rom) {
		description = "(nop, jp $0150)"
	}

	return fmt.Sprintf("%02X, %02X, %02X, %02X - %s",
		rom.Header.Entrypoint[0], rom.Header.Entrypoint[1],
		rom.Header.Entrypoint[2], rom.Header.Entrypoint[3], description)
}
