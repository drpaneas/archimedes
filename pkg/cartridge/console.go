package cartridge

import "strings"

// GetConsole identifies the console type based on the file extension.
// It supports Game Boy and Game Boy Color files.
func GetConsole(file string) string {
	file = strings.ToLower(file) // Normalize the file name to lower case for comparison

	switch {
	case strings.HasSuffix(file, "gb"):
		return "Game Boy"
	case strings.HasSuffix(file, "gbc"):
		return "Game Boy Color"
	default:
		return "Unknown"
	}
}
