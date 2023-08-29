package cartridge

import "C"
import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
)

func parse(b []byte) Rom {
	header := newHeader(b)

	title := string(header.OldTitle)

	// For Game Boy Color exclusive games only:
	// it seems more reasonable to use the new Title format (11 chars)
	if header.CGBFlag == 0xC0 {
		title = string(header.NewTitle)
	}

	// Logoverification
	nintendoExpectedLogo := []byte{
		0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B, 0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D,
		0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E, 0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99,
		0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC, 0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E,
	}

	isLogoVerified := false
	if bytes.Equal(nintendoExpectedLogo, header.Logo) {
		isLogoVerified = true
	}

	// Old cards use the old Licensee Code and new cards use the new one
	// To find which is which: if 0x14B == $33 then the New Licensee should be used
	publisher := getOldLicenseeCode(header.OldLicenseeCode)
	if b[0x14B] == 0x33 {
		publisher = getNewLicenseeCode(header.NewLicenseeCode)
	}

	// The SGB (akaSuper Game Boy) is a peripheral that allows Game Boy cartridges to be played on a SNES console.
	// This means the game will have better colors or audio, if you plug it into a SNES.
	var supportsSGB = false
	if header.SGBFlag == 0x03 {
		supportsSGB = true
	}

	// Specifies the MBC (Memory Bank Controller) and any other external hardware inside the card.
	cartridgeType := getCartridgeType(header.CartridgeType)

	// Specifies the ROM Size of the cartridge.
	romSize, romBanks := getRomSizeBank(header.RomSize)

	// Specifies the size of the external RAM in the cartridge (if any).
	ramSize, ramBanks := getRamSizeBank(header.RamSize)

	// Specifies if this version of the game is supposed to be sold in Japan, or anywhere else.
	isJapanese := getDestinationCode(header.DestinationCode)

	// Contains an 8 bit checksum across the cartridge header bytes $0134-014C. The boot ROM computes x as follows:
	headerChecksum := hex.EncodeToString(header.Checksum)

	// checksum verify
	var x byte
	for i := 0x0134; i <= 0x014C; i++ {
		x = x - b[i] - 1
	}
	isVerified := false
	if hex.EncodeToString([]byte{x}) == headerChecksum {
		isVerified = true
	}

	/*	014E-014F - Global Checksum
		Contains a 16 bit checksum (upper byte first) across the whole cartridge ROM.
		Produced by adding all bytes of the cartridge (except for the two checksum bytes).
		The Game Boy does not verify this checksum.
	*/
	var a uint16
	for k, v := range b {
		if k == 0x014E || k == 0x014F {
			continue
		}
		a += uint16(v)
	}
	globalChecksumImpl := fmt.Sprintf("%X", a)
	globalChecksum := strings.ToUpper(hex.EncodeToString([]byte{header.GlobalChecksum[0], header.GlobalChecksum[1]}))
	isGlobalChecksumVerified := false
	if globalChecksumImpl == globalChecksum {
		isGlobalChecksumVerified = true
	}

	hasPGB := true

	if hasBit(header.CGBFlag, 7) && (hasBit(header.CGBFlag, 2) || hasBit(header.CGBFlag, 3)) {
		hasPGB = false
	}

	return Rom{
		Binary:                   b,
		Header:                   header,
		Title:                    title,
		Publisher:                publisher,
		HasSGB:                   supportsSGB,
		CartridgeType:            cartridgeType,
		RomSize:                  romSize,
		RomBanks:                 romBanks,
		RamBanks:                 ramBanks,
		RamSize:                  ramSize,
		IsJapanese:               isJapanese,
		HeaderChecksum:           headerChecksum,
		IsChecksumVerified:       isVerified,
		GlobalChecksum:           globalChecksum,
		HasPGB:                   hasPGB,
		IsLogoVerified:           isLogoVerified,
		IsGlobalChecksumVerified: isGlobalChecksumVerified,
	}
}

// hasBit returns true if b byte has bit-p set.
func hasBit(b byte, p uint8) bool {
	return b&(1<<p) > 0
}
