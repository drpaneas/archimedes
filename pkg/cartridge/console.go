package cartridge

import "strings"

func GetConsole(file string) (console string) {
	if strings.HasSuffix(file, "gb") || strings.HasSuffix(file, "GB") {
		console = "Game Boy"
	}

	if strings.HasSuffix(file, "gbc") || strings.HasSuffix(file, "GBC") {
		console = "Game Boy Color"
	}

	return console
}
