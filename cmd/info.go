package cmd

import (
	"fmt"
	"hash/crc32"
	"path/filepath"
	"strings"

	"github.com/drpaneas/archimedes/pkg/cartridge"
)

// PrintRomInfo prints detailed information about the ROM.
func PrintRomInfo(rom cartridge.Rom, file string) {
	console := cartridge.GetConsole(file)
	sgbSupport := cartridge.HasSGBSupport(rom)
	crc := strings.ToUpper(fmt.Sprintf("%08x", crc32.ChecksumIEEE(rom.Binary)))
	headerChecksum, verificationStatus := cartridge.BootChecks(rom)

	globalCheckSumStatus := "(Unverified) - but that's OK"
	if rom.IsGlobalChecksumVerified {
		globalCheckSumStatus = "(Verified) ✓"
	}

	fmt.Println("== File Information ==")
	fmt.Printf(" Filename          : %s\n", filepath.Base(file))
	fmt.Printf(" Filesize          : %v KB\n", len(rom.Binary)/1024)
	fmt.Printf(" Checksum (CRC)    : %v\n", crc)
	fmt.Println("\n== Header Information ==")
	fmt.Printf(" Title             : %s\n", rom.Title)
	fmt.Printf(" Console           : %s\n", console)
	fmt.Printf(" Publisher         : %s\n", rom.Publisher)
	fmt.Printf(" Cartridge Type    : %s\n", rom.CartridgeType)
	fmt.Printf(" Super Gameboy     : %s\n", sgbSupport)
	fmt.Printf(" Color Gameboy     : %s\n", cartridge.HasCGB(rom))
	fmt.Printf(" Unofficial Colors : %s\n", cartridge.HasPGBSupport(rom))
	fmt.Printf(" Rom Size          : %s\n", cartridge.GetRomInfo(rom))
	fmt.Printf(" SRAM Save Size    : %s\n", cartridge.FormatSRAMInfo(rom))
	fmt.Printf(" Is Japanese       : %s\n", cartridge.IsJapanese(rom))
	fmt.Printf(" Game Version      : %d\n", rom.Header.MaskRomVersion)
	fmt.Printf(" Entry Point       : %s\n", cartridge.GetEntryPoint(rom))
	fmt.Printf(" Global Checksum   : 0x%s %s\n", rom.GlobalChecksum, globalCheckSumStatus)
	fmt.Println("\n== Boot Checks ==")
	fmt.Printf(" Header Checksum   : %s\n", headerChecksum)
	fmt.Printf(" Official Logo     : %s\n", verificationStatus)
}
