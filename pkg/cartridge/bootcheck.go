package cartridge

func BootChecks(rom Rom) (headerChecksum, verificationStatus string, err error) {
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
