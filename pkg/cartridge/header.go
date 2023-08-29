package cartridge

type header struct {
	entryPoint       [4]byte
	logo             [48]byte
	title            [11]byte
	manufacturerCode [4]byte
	cgb              byte
	newLicenseeCode  [2]byte
	sgb              byte
	cartridgeType    byte
	romSize          byte
	ramSize          byte
	destinationCode  byte
	oldLicenseeCode  byte
	maskRomVersion   byte
	headerChecksum   byte
	globalChecksum   [2]byte
}

//func newHeader(b []byte) header {
//	return header{
//		entryPoint:       []b[0x0100:0x0103+1],
//		logo:             [48]byte{},
//		title:            [11]byte{},
//		manufacturerCode: [4]byte{},
//		cgb:              0,
//		newLicenseeCode:  [2]byte{},
//		sgb:              0,
//		cartridgeType:    0,
//		romSize:          0,
//		ramSize:          0,
//		destinationCode:  0,
//		oldLicenseeCode:  0,
//		maskRomVersion:   0,
//		headerChecksum:   0,
//		globalChecksum:   [2]byte{},
//	}
//}
