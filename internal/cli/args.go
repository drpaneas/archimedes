package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/drpaneas/archimedes/cmd"
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

const (
	MAX_ARGS = 3 // archimedes arg[1] romfile arg[1] --flag [arg[2]]
)

// ProcessInput is responsible for handling command line flags and arguments and calls the appropriate functions.
// Finally, if no flags are given and the ROM file exists, it loads it in memory.
func ProcessInput() {
	if len(os.Args) > MAX_ARGS {
		LogError("too many arguments.")
		fmt.Println("Run 'archimedes --help' for usage.")
		os.Exit(1)
	}

	// Handle 1st argument, that is program name itself, without any arguments.
	if len(os.Args) == 1 {
		// If the user has typed "archimedes" then we should display the help message and exit gracefully.
		help()
		os.Exit(0)
	}

	// Handle the 2nd argument, if it's the help flag.
	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		// If the user has typed "archimedes --help" or "archimedes -h" then we should display the help message and exit gracefully.
		help()
		os.Exit(0)
	}

	// Loop through all the args, if youd find any that is a flag and it's an unsupported flag,
	// then display a message that the flag is not supported.
	for _, v := range os.Args[1:] { // skip the first argument, which is the program name (archimedes)
		if v == "--info" || v == "-i" || v == "--hexdump" || v == "-x" || v == "--help" || v == "-h" {
			continue
		}
		if strings.HasPrefix(v, "-") {
			LogError(fmt.Sprintf("'%s' is not a valid flag.\nRun 'archimedes --help' for usage.\n", v))
			os.Exit(1)
		}
	}

	// Handle the 2nd argument, it should be the ROM file.
	file := os.Args[1]

	// However, if instead of file, the user has typed a flag, then we should inform them that ROM file is missing.
	if strings.HasPrefix(file, "-") {
		LogError("the ROM file either missing or it's not provided in the expected syntax order.")
		os.Exit(1)
	}

	if !(strings.HasSuffix(file, "gb") || strings.HasSuffix(file, "gbc")) {
		LogError("only 'gb' ROM files are supported.")
		os.Exit(1)
	}

	if file == "" {
		LogError("the ROM file is missing.")
		fmt.Printf("Run 'archimedes --help' for usage.\n")
		os.Exit(1)
	}

	if !fileExists(file) {
		LogError(fmt.Sprintf("'%s'.", file))
		os.Exit(1)
	}

	rom := cartridge.Decode(readRom(file))

	for _, arg := range os.Args {
		switch arg {
		case "--info", "-i":
			if err := cmd.PrintRomInfo(rom, file); err != nil {
				fmt.Printf("\nProblematic ROM: %v\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		case "--hexdump", "-x":
			cmd.Hexdump(rom)
			os.Exit(0)
		case "--help", "-h":
			help()
			os.Exit(0)
		}
	}
}
