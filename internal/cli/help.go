package cli

import "fmt"

func help() {
	fmt.Printf("Usage:\n  archimedes [rom file] [optional flags]\n\n")
	fmt.Println("Optional flags:")
	fmt.Println("\t-i, --info:    displays ROM header information and quits.")
	fmt.Println("\t-x, --hexdump: prints the ROM in hexdump format and quits.")
	fmt.Println("\t-h, --help:    displays this message.")
}
