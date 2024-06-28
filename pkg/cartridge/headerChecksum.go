package cartridge

import (
	"errors"
	"fmt"
	"strings"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// okHeaderChecksum checks if the ROM's header checksum is verified.
// It returns a descriptive string and an error if the checksum is not verified.
func okHeaderChecksum(rom Rom) (string, error) {
	if rom.IsChecksumVerified {
		return fmt.Sprintf("0x%s (Verified) %s", strings.ToUpper(rom.HeaderChecksum), succeed), nil
	}

	errorMsg := fmt.Sprintf("0x%s (Unverified) %s - the byte 0x014D does not match the boot ROM computation across header bytes 0x0134-0x014C. The boot ROM will lock up, and the program won't run", strings.ToUpper(rom.HeaderChecksum), failed)
	return errorMsg, errors.New("checksum verification failed")
}
