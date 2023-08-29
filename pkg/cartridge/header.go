package cartridge

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
	RamSize          byte
	DestinationCode  byte
	OldLicenseeCode  byte
	MaskRomVersion byte
	Checksum       []byte
	GlobalChecksum []byte // 2 bytes
}

func newHeader(b []byte) Header {
	return Header{
		Entrypoint:       b[0x0100 : 0x0103+1],
		Logo:             b[0x0104 : 0x0133+1],
		NewTitle:         b[0x0134 : 0x013E+1],
		OldTitle:         b[0x0134 : 0x0143+1],
		ManufacturerCode: b[0x013F : 0x0142+1],
		CGBFlag:          b[0x0143],
		NewLicenseeCode:  b[0x0144 : 0x0145+1],
		SGBFlag:          b[0x0146],
		CartridgeType:    b[0x147],
		RomSize:          b[0x0148],
		RamSize:          b[0x0149],
		DestinationCode:  b[0x014A],
		OldLicenseeCode:  b[0x014B],
		MaskRomVersion:   b[0x014c],
		Checksum:         []byte{b[0x014d]},
		GlobalChecksum:   b[0x014e : 0x014f+1],
	}
}
