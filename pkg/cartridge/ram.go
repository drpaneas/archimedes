package cartridge

import "fmt"

const (
	bytesPerKB = 1024              // Number of bytes in a kilobyte
	bytesPerMB = 1024 * bytesPerKB // Bytes in a Megabyte
	noSRAM     = "None"
)

// SRAMInfo represents the save RAM information.
type SRAMInfo struct {
	SizeKB   int // Size in kilobytes
	NumBanks int // Number of banks
}

// getSRAMInfo returns the size of the save RAM in KB and the number of banks.
func getSRAMInfo(rom Rom) (info SRAMInfo) {
	if rom.RAMSize == 0 {
		return SRAMInfo{SizeKB: 0, NumBanks: 0}
	}
	return SRAMInfo{
		SizeKB:   rom.RAMSize / bytesPerKB,
		NumBanks: rom.RAMBanks,
	}
}

// ramSizeToString returns a string representation of the SRAM size in both Kb and KiB, and bank count.
func ramSizeToString(info SRAMInfo) string {
	if info.SizeKB == 0 {
		return noSRAM
	}
	sizeKb := info.SizeKB * 8 // Convert KiB to Kb
	return fmt.Sprintf("%dK (%d KiB) (%d banks)", sizeKb, info.SizeKB, info.NumBanks)
}

// FormatSRAMInfo gets the SRAM info from a Rom and returns its string representation.
func FormatSRAMInfo(rom Rom) string {
	info := getSRAMInfo(rom)
	return ramSizeToString(info)
}
