package cmd

import (
	"errors"
	"fmt"
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

func isLogoOfficial(rom cartridge.Rom) (s string, err error) {
	s = "N/A"
	if rom.IsLogoVerified {
		s = fmt.Sprintf("Verified %s",succeed)
	} else {
		s = fmt.Sprintf("Unverified %s", failed)
		err = errors.New("boot procedure cannot verify the content and gets locked up")
		if err != nil {
			return s, err
		}
	}
	return s, err
}
