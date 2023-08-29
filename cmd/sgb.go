package cmd

import "github.com/drpaneas/archimedes/pkg/cartridge"

func hasSGBSupport(rom cartridge.Rom) string {
	sgbSupport := "No"
	if rom.HasSGB {
		sgbSupport = "Yes"
	}
	return sgbSupport
}
