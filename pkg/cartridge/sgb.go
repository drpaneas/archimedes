package cartridge

// HasSGBSupport returns "Yes" if the ROM has SGB support, otherwise "No".
func HasSGBSupport(rom Rom) string {
	sgbSupport := "No"
	if rom.HasSGB {
		sgbSupport = "Yes"
	}
	return sgbSupport
}
