package cmd

import (
	"fmt"
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

func hexdump(rom cartridge.Rom) {
	var count16 int
	var tile string
	for i, v := range rom.Binary {
		count16++

		// we should reset if count16 > 16
		if count16 > 16 {
			count16 = 1
		}

		// Do the stuff you want
		tile += fmt.Sprintf("%02X ", v)

		// Print for specific count
		if count16 == 1 {
			fmt.Printf("%08X  ", i)
		}

		if count16 == 16 {
			fmt.Printf("%s\n", tile)
			tile = "" // reset the tile
		}
	}
}