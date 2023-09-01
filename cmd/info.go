package cmd

import (
	"fmt"
	"hash/crc32"
	"path/filepath"
	"strings"

	"github.com/drpaneas/archimedes/pkg/cartridge"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func PrintRomInfo(rom cartridge.Rom, file string) error {
	console := getConsole(file) // TODO: Να φτιάξω μία γενική function για τις καταλήξεις αρχείων σε Internal Package
	sgbSupport := hasSGBSupport(rom)
	crc := strings.ToUpper(fmt.Sprintf("%08x", crc32.Checksum(rom.Binary, crc32.IEEETable)))
	verificationStatus, headerChecksum, err := bootChecks(rom)

	globalCheckSumStatus := "(Unverified) - but that's OK"
	if rom.IsGlobalChecksumVerified {
		globalCheckSumStatus = "(Verified) ✓"
	}

	fmt.Println("== File Information ==")
	fmt.Println()
	fmt.Printf(" Filename         : %s\n", filepath.Base(file))
	fmt.Printf(" Filesize         : %v KB\n", len(rom.Binary)/1024)
	fmt.Printf(" Checksum (CRC)   : %v\n", crc)
	fmt.Println()

	fmt.Println("== Header Information ==")
	fmt.Println()

	fmt.Printf(" Title            : %s\n", rom.Title)
	fmt.Printf(" Console          : %s\n", console)
	fmt.Printf(" Publisher        : %s\n", rom.Publisher)
	fmt.Printf(" Cartridge Type   : %s\n", rom.CartridgeType)
	fmt.Printf(" Super Gameboy    : %s\n", sgbSupport)
	fmt.Printf(" Color Gameboy    : %s\n", hasCGB(rom))
	fmt.Printf(" Rom Size         : %s\n", getRomInfo(rom))
	fmt.Printf(" Ram Size         : %s\n", getRamInfo(rom))
	fmt.Printf(" Japanese         : %s\n", isJapanese(rom))
	fmt.Printf(" Game Version     : %d\n", rom.Header.MaskRomVersion)
	fmt.Printf(" Entry Point      : %s\n", getEntryPoint(rom))
	fmt.Printf(" Global Checksum  : 0x%s %s\n", rom.GlobalChecksum, globalCheckSumStatus)
	fmt.Println()

	fmt.Println("== Boot Checks ==")
	fmt.Println()
	fmt.Printf(" Header Checksum  : %s\n", headerChecksum)
	fmt.Printf(" Official Logo    : %s\n", verificationStatus)

	return err
}
