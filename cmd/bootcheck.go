package cmd

import (
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

func bootChecks(rom cartridge.Rom) (headerChecksum, verificationStatus string, err error) {
	verificationStatus, err = isLogoOfficial(rom)
	if err != nil {
		return "", "", err
	}
	headerChecksum, err = okHeaderChecksum(rom)
	if err != nil {
		return "", "", err
	}
	return verificationStatus, headerChecksum, nil
}
