package cartridge

func getRomSizeBank(code byte) (size int, banks int) {
	switch code {

	case 0x00:
		return 32*1024, 2 //32 KByte, 2 banks (No Rom banking)
	case 0x01:
		return 64*1024,4 // 64 Byte, 4 banks
	case 0x02:
		return 128*1024,8 // 128 KByte, 8 banks
	case 0x03:
		return 256*1024,16 // 256 KByte, 16 banks
	case 0x04:
		return 512*1024,32 // 512 KByte, 32 banks
	case 0x05:
		return 1 * 1024 * 1024, 64 // 1 MByte, 64 banks
	case 0x06:
		return 2 * 1024 * 1024, 128 // 2 MByte, 128 banks
	case 0x07:
		return 4 * 1024 * 1024, 256 // 4 MByte, 256 banks
	case 0x08:
		return 8 * 1024 * 1024, 512 // 8 MByte, 512 banks
	// -----------------------------------------------------------------
	// Only listed in unofficial docs. No cartridges or ROM files using these sizes are known.
	// As the other ROM sizes are all powers of 2, these are likely inaccurate.
	// The source for these values is unknown.
	case 0x52:
		return 1153433, 72 // 1.1 MByte, 72 banks
	case 0x53:
		return 1258291, 80 // 1.2 MByte, 80 banks
	case 0x54:
		return 1.5 * 1024 * 1024, 96 // 1.5 MByte, 96 banks
	default:
		return 0, 0
	}
}