package cartridge

import (
	"errors"
	"fmt"
)

func isLogoOfficial(rom Rom) (s string, err error) {
	if rom.IsLogoVerified {
		return fmt.Sprintf("Verified %s", succeed), nil
	}

	return fmt.Sprintf("Unverified %s", failed), errors.New("boot procedure cannot verify the content and gets locked up")
}
