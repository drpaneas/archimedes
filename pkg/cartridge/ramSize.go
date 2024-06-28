package cartridge

// getRamSizeBank returns the size of the RAM and the number of banks based on the RAM size code.
func getRamSizeBank(code byte) (size int, banks int) {
	switch code {
	case 0x00:
		// No RAM when using a MBC2 chip, even though MBC2 includes a built-in RAM of 512 x 4 bits.
		return // Implicitly returns 0, 0
	case 0x01:
		// Unused. Listed in various unofficial docs as 2KB, but a 2 KB RAM chip was never used in a cartridge.
		return // Implicitly returns 0, 0
	case 0x02:
		return 8 * bytesPerKB, 1 // 8 KByte, 1 bank of 8 KB
	case 0x03:
		return 32 * bytesPerKB, 4 // 32 KByte, 4 banks of 8 KB each
	case 0x04:
		return 128 * bytesPerKB, 16 // 128 KByte, 16 banks of 8 KB each
	case 0x05:
		return 64 * bytesPerKB, 8 // 64 KByte, 8 banks of 8 KB each
	default:
		return 0, 0
	}
}
