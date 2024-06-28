package cartridge

// Header represents the structure of a Game Boy ROM header.
type Header struct {
	Entrypoint       []byte // 4 bytes
	Logo             []byte // 48 bytes
	NewTitle         []byte // 11 bytes
	OldTitle         []byte // 16 bytes
	ManufacturerCode []byte // 4 bytes
	CGBFlag          byte
	NewLicenseeCode  []byte // 2 bytes
	SGBFlag          byte
	CartridgeType    byte
	RomSize          byte
	RAMSize          byte
	DestinationCode  byte
	OldLicenseeCode  byte
	MaskRomVersion   byte
	Checksum         []byte // 1 byte
	GlobalChecksum   []byte // 2 bytes
}

// newHeader creates a new Header from a byte slice.
// It assumes the byte slice has the correct structure and size for a Game Boy ROM header.
func newHeader(b []byte) Header {
	return Header{
		Entrypoint:       b[0x0100:0x0104],
		Logo:             b[0x0104:0x0134],
		NewTitle:         b[0x0134:0x013F],
		OldTitle:         b[0x0134:0x0144],
		ManufacturerCode: b[0x013F:0x0143],
		CGBFlag:          b[0x0143],
		NewLicenseeCode:  b[0x0144:0x0146],
		SGBFlag:          b[0x0146],
		CartridgeType:    b[0x0147],
		RomSize:          b[0x0148],
		RAMSize:          b[0x0149],
		DestinationCode:  b[0x014A],
		OldLicenseeCode:  b[0x014B],
		MaskRomVersion:   b[0x014C],
		Checksum:         []byte{b[0x014D]},
		GlobalChecksum:   b[0x014E:0x0150],
	}
}
