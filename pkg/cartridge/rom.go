package cartridge

type Rom struct {
	Binary                   []byte // the bytes of the file
	Header                   Header
	Title                    string
	HasSGB                   bool
	Publisher                string
	CartridgeType            string
	RomSize                  int
	RomBanks                 int
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

func NewRom() Rom {
	return Rom{}
}

func Decode(b []byte) Rom {
	return parse(b)
}
