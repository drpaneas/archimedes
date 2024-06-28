package cartridge

import "fmt"

type Rom struct {
	Binary                   []byte // the bytes of the file
	Header                   Header
	Title                    string
	HasSGB                   bool
	Publisher                string
	CartridgeType            string
	RomSize                  int // Size in bytes
	RomBanks                 int // Count number of ROM banks
	RamSize                  int
	RamBanks                 int
	IsJapanese               bool
	OldLicenseeCode          []byte
	HeaderChecksum           string
	IsChecksumVerified       bool
	GlobalChecksum           string
	CGBFlag                  bool
	HasPGB                   bool
	IsLogoVerified           bool
	IsGlobalChecksumVerified bool
}

// NewRom creates and returns a new Rom instance with default values.
func NewRom() Rom {
	return Rom{}
}

// Decode parses the given byte slice and returns a Rom instance.
func Decode(b []byte) Rom {
	return parse(b)
}

// GetRomInfo returns a string describing the ROM size and banking information.
func GetRomInfo(rom Rom) string {
	// Pre-calculate common values
	sizeInKiB := rom.RomSize / bytesPerKB
	sizeInMiB := rom.RomSize / bytesPerMB

	switch {
	case rom.RomBanks == 2:
		return fmt.Sprintf("%d KiB (%d banks - No ROM banking)", sizeInKiB, rom.RomBanks)
	case rom.RomBanks >= 4 && rom.RomBanks <= 32:
		return fmt.Sprintf("%d KiB (%d banks)", sizeInKiB, rom.RomBanks)
	case rom.RomBanks >= 64:
		return fmt.Sprintf("%d MiB (%d banks)", sizeInMiB, rom.RomBanks)
	default:
		return "Rom Size: N/A"
	}
}
