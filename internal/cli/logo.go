package cli

import "fmt"

// PrintLogo prints the ASCI art logo of the emulator in the terminal.
func PrintLogo() {
	fmt.Printf("  _. ._ _ |_  o ._ _   _   _|  _   _ \n (_| | (_ | | | | | | (/_ (_| (/_ _> \n")
	fmt.Printf("archimedes is a GameBoy emulator by drpaneas, written in Go.\n\n")
	fmt.Printf("  * Dedicated to my grandpa, who bought me my first DMG console.\n")
	fmt.Printf("  > Find more information at: https://github.com/drpaneas/archimedes\n\n")
}
