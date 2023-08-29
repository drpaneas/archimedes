package cmd

import (
	"fmt"
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

func getRomInfo(rom cartridge.Rom) string {
	s := fmt.Sprintf("Rom Size         : N/A\n")
	if rom.RomBanks == 2 {
		s = fmt.Sprintf("%d KB (%d banks - No ROM banking)", rom.RomSize/1024, rom.RomBanks)
	}
	if rom.RomBanks >= 4 && rom.RomBanks <= 32 {
		s = fmt.Sprintf("%d KB (%d banks)", rom.RomSize/1024, rom.RomBanks)
	}
	if rom.RomBanks >= 64 {
		s = fmt.Sprintf("%d MB (%d banks)", rom.RomSize/1024/1024, rom.RomBanks)
	}
	return s
}
