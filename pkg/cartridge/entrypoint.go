package cartridge

import (
	"bytes"
	"fmt"
)

func hasUsualEntryPoint(rom Rom) bool {
	usualEntryPoint := []byte{0, 195, 80, 1}
	isUsualEntryPoint := false
	if bytes.Equal(rom.Header.Entrypoint, usualEntryPoint) {
		isUsualEntryPoint = true
	}
	return isUsualEntryPoint
}

func GetEntryPoint(rom Rom) string {
	s := fmt.Sprintf("%02X, %02X, %02X, %02X - not the most usual instruction",
		rom.Header.Entrypoint[0], rom.Header.Entrypoint[1], rom.Header.Entrypoint[2], rom.Header.Entrypoint[3])
	if hasUsualEntryPoint(rom) {
		s = fmt.Sprintf("%02X, %02X, %02X, %02X - (nop, jp $0150)", rom.Header.Entrypoint[0],
			rom.Header.Entrypoint[1], rom.Header.Entrypoint[2], rom.Header.Entrypoint[3])
	}
	return s
}
