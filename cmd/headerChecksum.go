package cmd

import (
	"errors"
	"fmt"
	"github.com/drpaneas/archimedes/pkg/cartridge"
	"strings"
)

func okHeaderChecksum(rom cartridge.Rom) (s string, err error) {
	s = fmt.Sprintf("0x%s (Verified) %s", strings.ToUpper(rom.HeaderChecksum), succeed)
	if !rom.IsChecksumVerified {
		s = fmt.Sprintf("0x%s (Not Verified) %s", strings.ToUpper(rom.HeaderChecksum), failed)
		err = errors.New("the byte 0x014D does not match the boot ROM computation " +
			"across header bytes 0x0134-0x014C. The boot room will lock up and the program won't run ")
	}
	return s, err
}
