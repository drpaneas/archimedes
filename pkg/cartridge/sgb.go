package cartridge

func HasSGBSupport(rom Rom) string {
	sgbSupport := "No"
	if rom.HasSGB {
		sgbSupport = "Yes"
	}
	return sgbSupport
}
