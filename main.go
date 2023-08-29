package main

import (
	"fmt"
	"github.com/drpaneas/archimedes/cmd"
	"github.com/drpaneas/archimedes/pkg/cartridge"
	"os"
)

var file string
var debugFlag bool


func init() {
	printLogo()
	processUserInput()
}

func main() {

	rom := cartridge.NewRom()
	rom = cartridge.Decode(readRom(file))
	if err := cmd.PrintRomInfo(rom, file); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//var gb emu.Emulator
	//gb.Run(rom)
}
