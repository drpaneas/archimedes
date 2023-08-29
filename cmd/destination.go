package cmd

import (
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

func isJapanese(rom cartridge.Rom) string {
	s := "No"
	if rom.IsJapanese {
		s = "Yes"
	}
	return s
}
