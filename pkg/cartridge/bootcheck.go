package cartridge

// BootChecks performs checks on the ROM's boot logo and header checksum.
// It returns the header checksum, the verification status of the official logo, and any error encountered.
func BootChecks(rom Rom) (headerChecksum, verificationStatus string) {
	// Check if the ROM's logo matches the official logo.
	verificationStatus, _ = isLogoOfficial(rom) // Do not take any action if an error is encountered.

	// Validate the ROM's header checksum.
	headerChecksum, _ = okHeaderChecksum(rom) // Do not take any action if an error is encountered.

	// Why I don't care about the error returned by isLogoOfficial and okHeaderChecksum?
	// Because I still want to see all rest of the info, even if there are errors.

	// Return the header checksum and verification status if no errors were encountered.
	return headerChecksum, verificationStatus
}
