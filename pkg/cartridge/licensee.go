package cartridge

import "fmt"

func getNewLicenseeCode(lic []byte) string {
	b1 := lic[0]
	b2 := lic[1]
	code := fmt.Sprintf("%s%s",string(b1),string(b2))
	switch code {
	case "00":
		return "None"
	case "01":
		return "Nintendo R&D1"
	case "08":
		return "Capcom"
	case "13":
		return "Electronic Arts"
	case "18":
		return "Hudson Soft"
	case "19":
		return "b-ai"
	case "20":
		return "kss"
	case "22":
		return "pow"
	case "24":
		return "PCM Complete"
	case "25":
		return "san-x"
	case "28":
		return "Kemco Japan"
	case "29":
		return "seta"
	case "30":
		return "Viacom"
	case "31":
		return "Nintendo"
	case "32":
		return "Bandai"
	case "33":
		return  "Ocean/Acclaim"
	case "34":
		return "Konami"
	case "35":
		return "Hextor"
	case "37":
		return "Taito"
	case "38":
		return "Hudson"
	case "39":
		return "Banpresto"
	case "41":
		return "Ubi Soft"
	case "42":
		return "Atlus"
	case "44":
		return "Malibu"
	case "46":
		return "angel"
	case "47":
		return "Bullet-Proof"
	case "49":
		return "irem"
	case "50":
		return "Absolute"
	case "51":
		return "Acclaim"
	case "52":
		return "Activision"
	case "53":
		return "American sammy"
	case "54":
		return "Konami"
	case "55":
		return "Hi tech entertainment"
	case "56":
		return "LJN"
	case "57":
		return "MMatchbox"
	case "58":
		return "Mattel"
	case "59":
		return "Milton Bradley"
	case "60":
		return "Titus"
	case "61":
		return "Virgin"
	case "64":
		return "LucasArts"
	case "67":
		return "Ocean"
	case "69":
		return "Electronic Arts"
	case "70":
		return "Infogrames"
	case "71":
		return "Interplay"
	case "72":
		return "Broderbund"
	case "73":
		return "sculptured"
	case "75":
		return "sci"
	case "78":
		return "THQ"
	case "79":
		return "Accolade"
	case "80":
		return "misawa"
	case "83":
		return "lozc"
	case "86":
		return "Tokuma Shoten Intermedia"
	case "87":
		return "Tsukuda Original"
	case "91":
		return "Chunsoft"
	case "92":
		return "Video system"
	case "93":
		return "Ocean/Acclaim"
	case "95":
		return "Varie"
	case "96":
		return "Yonezawa/s'pal"
	case "97":
		return "Kaneko"
	case "99":
		return "Pack in soft"
	case "A4":
		return "Konami (Yo-Gi-Oh!)"
	default:
		return "Unknown/Undefined"
	}
}

func getOldLicenseeCode(code byte) string {
	switch code {
	case 0x00:
		return "None"
	case 0x01:
		return "Nintendo"
	case 0x08:
		return "Capcom"
	case 0x09:
		return "Hot-B"
	case 0x0A:
		return "Jaleco"
	case 0x0B:
		return "Coconuts"
	case 0x0C:
		return "Elite Systems"
	case 0x13:
		return "Electronic Arts"
	case 0x18:
		return "Hudsonsoft"
	case 0x19:
		return "ITC Entertainment"
	case 0x1A:
		return "Yanoman"
	case 0x1D:
		return "Clary"
	case 0x1F:
		return "Virgin"
	case 0x24:
		return "PCM Complete"
	case 0x25:
		return "San-X"
	case 0x28:
		return  "Kotobuki Systems"
	case 0x29:
		return "Seta"
	case 0x30:
		return "Infogrames"
	case 0x31:
		return "Nintendo"
	case 0x32:
		return "Bandai"
	case 0x33:
		return "See Above"
	case 0x34:
		return "Konami"
	case 0x35:
		return "Hector"
	case 0x38:
		return "Capcom"
	case 0x39:
		return "Banpresto"
	case 0x3C:
		return "*Entertainment I"
	case 0x3E:
		return "Gremlin"
	case 0x41:
		return "Ubi Soft"
	case 0x42:
		return "Atlus"
	case 0x44:
		return "Malibu"
	case 0x46:
		return "Angel"
	case 0x47:
		return "Spectrum Holoby"
	case 0x49:
		return "Irem"
	case 0x4A:
		return "Virgin"
	case 0x4D:
		return "Malibu"
	case 0x4F:
		return "U.S. Gold"
	case 0x50:
		return "Absolute"
	case 0x51:
		return "Acclaim"
	case 0x52:
		return "Activision"
	case 0x53:
		return "American Sammy"
	case 0x54:
		return "Gametek"
	case 0x55:
		return "Park Place"
	case 0x56:
		return "LJN"
	case 0x57:
		return "Matchbox"
	case 0x59:
		return "Milton Bradley"
	case 0x5A:
		return "Mindscape"
	case 0x5B:
		return "Romstar"
	case 0x5C:
		return "Naxat Soft"
	case 0x5D:
		return "Tradewest"
	case 0x60:
		return "Titus"
	case 0x61:
		return "Virgin"
	case 0x67:
		return "Ocean"
	case 0x69:
		return "Electronic Arts"
	case 0x6E:
		return "Elite Systems"
	case 0x6F:
		return "Electro Brain"
	case 0x70:
		return "Infogrames"
	case 0x71:
		return "Interplay"
	case 0x72:
		return "Broderbund"
	case 0x73:
		return "Sculptered Soft"
	case 0x75:
		return "The Sales Curbe"
	case 0x78:
		return "T*HQ"
	case 0x79:
		return "Accolade"
	case 0x7A:
		return "Triffix Entertainment"
	case 0x7C:
		return "Microprose"
	case 0x7F:
		return "Kemco"
	case 0x83:
		return "Lozc"
	case 0x86:
		return "*Tokuma Shoten I"
	case 0x8B:
		return "Bullet-Proof Software"
	case 0x8C:
		return "Vic Tokai"
	case 0x8E:
		return "Ape"
	case 0x8F:
		return "iMax"
	case 0x91:
		return "Chun Soft"
	case 0x92:
		return "Video System"
	case 0x93:
		return "Tsuburava"
	case 0x95:
		return "Varie"
	case 0x96:
		return "Yonezawa's Pal"
	case 0x97:
		return "Kaneko"
	case 0x99:
		return "Arc"
	case 0x9a:
		return "Nihon Bussan"
	case 0x9B:
		return "Tecmo"
	case 0x9C:
		return "Imagineer"
	case 0x9D:
		return "Banpresto"
	case 0x9F:
		return "Nova"
	case 0xA1:
		return "Hori Electric"
	case 0xA2:
		return "Bandai"
	case 0xA4:
		return "Konami"
	case 0xA6:
		return "Kawada"
	case 0xA7:
		return "Takara"
	case 0xA9:
		return "Technos Japan"
	case 0xAA:
		return "Broderbund"
	case 0xAC:
		return "Toei Animation"
	case 0xAD:
		return "ToHo"
	case 0xAF:
		return "Namco"
	case 0xB0:
		return "Acclaim"
	case 0xB1:
		return "Nexoft"
	case 0xB2:
		return "Bandai"
	case 0xB4:
		return "Enix"
	case 0xB6:
		return "Hal"
	case 0xB7:
		return "SNK"
	case 0xB9:
		return "Pony Canyon"
	case 0xBA:
		return "*Culture Brain O"
	case 0xBB:
		return "SunSoft"
	case 0xBD:
		return "Sony ImageSoft"
	case 0xBF:
		return "Sammy"
	case 0xC0:
		return "Taito"
	case 0xC2:
		return "Kemco"
	case 0xC3:
		return "SquareSoft"
	case 0xC4:
		return "*Tokuma Shoten I"
	case 0xC5:
		return "Data East"
	case 0xC6:
		return "Tonkin House"
	case 0xC8:
		return "Koei"
	case 0xC9:
		return "UFL"
	case 0xCA:
		return "Ultra"
	case 0xCB:
		return "VAP"
	case 0xCC:
		return "Use"
	case 0xCD:
		return "Meldac"
	case 0xCE:
		return "*Pony Canyon Or"
	case 0xCF:
		return "Angel"
	case 0xD0:
		return "Taito"
	case 0xD1:
		return "Sofel"
	case 0xD2:
		return "Quest"
	case 0xD3:
		return "Sigma Enterprises"
	case 0xD4:
		return "Ask Kodansha"
	case 0xD6:
		return "Naxat Soft"
	case 0xD7:
		return "Copya Systems"
	case 0xD9:
		return "Banpresto"
	case 0xDA:
		return "Tomy"
	case 0xDB:
		return "LJN"
	case 0xDD:
		return "NCS"
	case 0xDE:
		return "Human"
	case 0xDF:
		return "Altron"
	case 0xE0:
		return "Jaleco"
	case 0xE1:
		return "Towachiki"
	case 0xE2:
		return "Uutaka"
	case 0xE3:
		return "Varie"
	case 0xE5:
		return "Epoch"
	case 0xE7:
		return "Athena"
	case 0xE8:
		return "Asmik"
	case 0xE9:
		return "Natsume"
	case 0xEA:
		return "King Records"
	case 0xEB:
		return "Atlus"
	case 0xEC:
		return "Epic/Sony Records"
	case 0xEE:
		return "IGS"
	case 0xF0:
		return "A Wave"
	case 0xF3:
		return "Extreme Entertainment"
	case 0xFF:
		return "LJN"
	default:
		return "Unknown/Undefined"
	}
}


