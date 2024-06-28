package cmd

import (
	"fmt"
	"strings"

	"github.com/drpaneas/archimedes/pkg/cartridge"
)

// Hexdump prints the ROM in hexdump format
func Hexdump(rom cartridge.Rom) {
	var sb strings.Builder // Use strings.Builder for efficient string concatenation
	for i, v := range rom.Binary {
		// Use modulus to determine when to print and reset
		if i%16 == 0 {
			if i != 0 {
				fmt.Println(sb.String()) // Print the accumulated line
				sb.Reset()               // Reset the builder for the next line
			}
			fmt.Printf("%08X  ", i) // Print the line prefix
		}

		// Append the current byte in hex format to the builder
		fmt.Fprintf(&sb, "%02X ", v)
	}

	// Print any remaining bytes that didn't fill a complete line
	if sb.Len() > 0 {
		fmt.Println(sb.String())
	}
}
