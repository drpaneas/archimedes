package cartridge

import (
	"errors"
	"fmt"
)

const (
	verifiedMessage   = "Verified"
	unverifiedMessage = "Fake (Unverified) - the boot procedure cannot verify the content and gets locked up"
)

// isLogoOfficial checks if the ROM's logo is officially verified.
// It returns a descriptive string and an error if the logo is not verified.
func isLogoOfficial(rom Rom) (string, error) {
	if rom.IsLogoVerified {
		return fmt.Sprintf("%s %s", verifiedMessage, succeed), nil
	}

	return fmt.Sprintf("%s %s", unverifiedMessage, failed), errors.New("logo verification failed")
}
