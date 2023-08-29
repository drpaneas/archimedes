package cartridge

func getRamSizeBank(code byte) (size int, banks int) {
	switch code {

	case 0x00:
		// When using a MBC2 chip, $00 must be specified as the RAM Size,
		// even though the MBC2 includes a built-in RAM of 512 x 4 bits.
		return 0, 0 // No RAM
	case 0x01:
		// Listed in various unofficial docs as 2KB.
		// However, a 2 KB RAM chip was never used in a cartridge. The source for this value is unknown.
		return 0,0 // Unused
	case 0x02:
		return 8 * 1024,1 // 8 KByte, 8 bank of 1 KB
	case 0x03:
		return 32*1024,4 // 32 KByte, 4 banks of 8 KB each
	case 0x04:
		return 128*1024,16 // 128 KByte, 16 banks of 8 KB each
	case 0x05:
		return 64 * 1024, 8 // 64 KByte, 8 banks of 8 KB each
	default:
		return 0, 0
	}
}
