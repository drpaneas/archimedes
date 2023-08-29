package cmd

import (
	"fmt"

	"github.com/drpaneas/archimedes/pkg/cartridge"
)

func getRamInfo(rom cartridge.Rom) string {
	s := "N/A"
	if rom.RamSize != 0 {
		s = fmt.Sprintf("%d KB (%d banks)", rom.RamSize/1024, rom.RamBanks)
	}
	return s
}
