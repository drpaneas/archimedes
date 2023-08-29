package cmd

import (
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

func bootChecks(rom cartridge.Rom) (headerChecksum, verificationStatus string , err error) {
	verificationStatus, err = isLogoOfficial(rom)
	headerChecksum, err = okHeaderChecksum(rom)
	return verificationStatus, headerChecksum, err
}