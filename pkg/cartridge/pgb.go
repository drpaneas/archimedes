package cartridge

// HasPGBSupport returns "Yes" if the ROM has PGB support, otherwise "No".
func HasPGBSupport(rom Rom) string {
	pgbSupport := "No"
	if rom.HasPGB {
		pgbSupport = "Yes"
	}
	return pgbSupport
}
