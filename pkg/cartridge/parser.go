package cartridge

import "C"
import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	cgbFlagExclusive        = 0xC0
	sgbFlagSupported        = 0x03
	newLicenseeFlag         = 0x33
	nintendoExpectedLogoStr = "\xCE\xED\x66\x66\xCC\x0D\x00\x0B\x03\x73\x00\x83\x00\x0C\x00\x0D" +
		"\x00\x08\x11\x1F\x88\x89\x00\x0E\xDC\xCC\x6E\xE6\xDD\xDD\xD9\x99" +
		"\xBB\xBB\x67\x63\x6E\x0E\xEC\xCC\xDD\xDC\x99\x9F\xBB\xB9\x33\x3E"
)

// parse takes a byte slice representing a Game Boy cartridge binary and returns a ROM struct.
func parse(b []byte) Rom {
	header := newHeader(b)

	title := string(header.OldTitle)

	// For Game Boy Color exclusive games only:
	// it seems more reasonable to use the new Title format (11 chars)
	if header.CGBFlag == cgbFlagExclusive {
		title = string(header.NewTitle)
	}

	// Logo Verification: Convert back to byte slice when used
	nintendoExpectedLogo := []byte(nintendoExpectedLogoStr)
	isLogoVerified := bytes.Equal(nintendoExpectedLogo, header.Logo)

	// Old cards use the old Licensee Code and new cards use the new one
	// To find which is which: if 0x14B == $33 then the New Licensee should be used
	publisher := getOldLicenseeCode(header.OldLicenseeCode)
	if b[0x14B] == 0x33 {
		publisher = getNewLicenseeCode(header.NewLicenseeCode)
	}

	// The SGB (akaSuper Game Boy) is a peripheral that allows Game Boy cartridges to be played on a SNES console.
	// This means the game will have better colors or audio, if you plug it into a SNES.
	supportsSGB := header.SGBFlag == sgbFlagSupported

	// Specifies the MBC (Memory Bank Controller) and any other external hardware inside the card.
	cartridgeType := getCartridgeType(header.CartridgeType)

	// Specifies the ROM Size of the cartridge.
	romSize, romBanks := getRomSizeBank(header.RomSize)

	// Specifies the size of the external RAM in the cartridge (if any).
	ramSize, ramBanks := getRAMSizeBank(header.RAMSize)

	// Specifies if this version of the game is supposed to be sold in Japan, or anywhere else.
	isJapanese := getDestinationCode(header.DestinationCode)

	// Contains an 8 bit checksum across the cartridge header bytes $0134-014C. The boot ROM computes x as follows:
	headerChecksum := hex.EncodeToString(header.Checksum)

	// checksum verify
	var x byte
	for i := 0x0134; i <= 0x014C; i++ {
		x = x - b[i] - 1
	}
	isVerified := hex.EncodeToString([]byte{x}) == headerChecksum

	/*	014E-014F - Global Checksum
		Contains a 16 bit checksum (upper byte first) across the whole cartridge ROM.
		Produced by adding all bytes of the cartridge (except for the two checksum bytes).
		The Game Boy does not verify this checksum.
	*/
	var a uint16
	for k, v := range b {
		if k != 0x014E && k != 0x014F {
			a += uint16(v)
		}
	}
	globalChecksumImpl := fmt.Sprintf("%X", a)
	globalChecksum := strings.ToUpper(hex.EncodeToString([]byte{header.GlobalChecksum[0], header.GlobalChecksum[1]}))
	isGlobalChecksumVerified := globalChecksumImpl == globalChecksum

	// This mode bypasses the standard CGB functions and leaves the color palettes blank.
	// While the functionality exists, there's no documented purpose from Nintendo.
	// The theory is that this mode could have been used to apply color to DMG games that
	// included their own color data embedded in the ROM at a specific location.
	hasPGB := hasBit(header.CGBFlag, 7) && (hasBit(header.CGBFlag, 2) || hasBit(header.CGBFlag, 3))

	return Rom{
		Binary:                   b,
		Header:                   header,
		Title:                    title,
		Publisher:                publisher,
		HasSGB:                   supportsSGB,
		CartridgeType:            cartridgeType,
		RomSize:                  romSize,
		RomBanks:                 romBanks,
		RAMBanks:                 ramBanks,
		RAMSize:                  ramSize,
		IsJapanese:               isJapanese,
		HeaderChecksum:           headerChecksum,
		IsChecksumVerified:       isVerified,
		GlobalChecksum:           globalChecksum,
		HasPGB:                   hasPGB,
		IsLogoVerified:           isLogoVerified,
		IsGlobalChecksumVerified: isGlobalChecksumVerified,
	}
}

// hasBit returns true if the byte b has the bit at position p set.
func hasBit(b byte, p uint8) bool {
	return b&(1<<p) > 0
}
