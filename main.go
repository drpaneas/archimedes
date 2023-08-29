package main

import (
	"fmt"
	"os"

	"github.com/drpaneas/archimedes/cmd"
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

var file string
var debugFlag bool
var hexdumpFlag bool
var rom cartridge.Rom

func init() {
	printLogo()
	processUserInput()
}

func main() {

	rom = cartridge.Decode(readRom(file))
	if err := cmd.PrintRomInfo(rom, file); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if hexdumpFlag {
		cmd.Hexdump(rom)
	}

	//var gb emu.Emulator
	//gb.Run(rom)
}
