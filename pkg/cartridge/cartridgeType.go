package cartridge

import "fmt"

func getCartridgeType(code byte) string {
	switch code {
	case 0x00:
		return "ROM ONLY"
	case 0x01:
		return "MBC1"
	case 0x02:
		return "MBC1+RAM"
	case 0x03:
		return "MBC1+RAM+BATTERY"
	case 0x05:
		return "MBC2"
	case 0x06:
		return "MBC2+BATTERY"
	case 0x08:
		fmt.Println("BEWARE: This is not a licenced cartridge!")
		return "ROM+RAM 1"
	case 0x09:
		fmt.Println("BEWARE: This is not a licenced cartridge!")
		return "ROM+RAM+BATTERY"
	case 0x0B:
		return "MMM01"
	case 0x0C:
		return "MMM01+RAM"
	case 0x0D:
		return "MMM01+RAM+BATTERY"
	case 0x0F:
		return "MBC3+TIMER+BATTERY"
	case 0x10:
		return "MBC3+TIMER+RAM+BATTERY 2"
	case 0x11:
		return "MBC3"
	case 0x12:
		fmt.Println("Note: 2 MBC3 with RAM size 64 KByte refers to MBC30, used only in Pocket Monsters Crystal Version for Japan")
		return "MBC3+RAM 2"
	case 0x13:
		fmt.Println("Note: 2 MBC3 with RAM size 64 KByte refers to MBC30, used only in Pocket Monsters Crystal Version for Japan")
		return "MBC3+RAM+BATTERY 2"
	case 0x19:
		return "MBC5"
	case 0x1A:
		return "MBC5+RAM"
	case 0x1B:
		return "MBC5+RAM+BATTERY"
	case 0x1C:
		return "MBC5+RUMBLE"
	case 0x1D:
		return "MBC5+RUMBLE+RAM"
	case 0x1E:
		return "MBC5+RUMBLE+RAM+BATTERY"
	case 0x20:
		return "MBC6"
	case 0x22:
		return "MBC7+SENSOR+RUMBLE+RAM+BATTERY"
	case 0xFC:
		return "POCKET CAMERA"
	case 0xFD:
		return "BANDAI TAMA5"
	case 0xFE:
		return "HuC3"
	case 0xFF:
		return "HuC1+RAM+BATTERY"
	default:
		return "N/A"
	}
}
