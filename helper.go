package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/drpaneas/archimedes/cmd"
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

var found, hasHex, hasInfo bool

func isRomMissing() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	for _, v := range os.Args {
		if v == "--help" || v == "-h" {
			help()
			os.Exit(0)
		}
	}
}

func processUserInput() {
	isRomMissing()

	// Parse checking
	for _, v := range os.Args {
		if v == "--hexdump" || v == "-x" {
			hasHex = true
		}

		if v == "--info" || v == "-i" {
			hasInfo = true
		}

		if v == "--debug" || v == "-d" {
			debugFlag = true
		}
	}

	for k, v := range os.Args {
		if k != 0 && v != "--debug" && v != "--info" && v != "-d" && v != "-i" && v != "--hexdump" && v != "-x" {
			if found {
				debug(fmt.Sprintf("You you have already provided the ROM file '%s'. What is '%s' then?", file, v))
				logError("please provide only one ROM file at a time.")
				fmt.Printf("Run 'archimedes --help' for usage.\n")
				os.Exit(1)
			}
			if !(strings.HasSuffix(v, "gb") || strings.HasSuffix(v, "gbc")) {
				logError("this *.gb or *.gbc only).")
				os.Exit(1)
			}
			file = v
			found = true
			debug(fmt.Sprintf("ROM file: %s", file))
		}
	}

	if file == "" {
		logError("the ROM file is missing.")
		fmt.Printf("Run 'archimedes --help' for usage.\n")
		os.Exit(1)
	}

	if !fileExists(file) {
		logError(fmt.Sprintf("'%s' not found.", file))
		os.Exit(1)
	}

	if strings.HasSuffix(file, "gb") || strings.HasSuffix(file, "gbc") {
		debug("File has proper suffix")
	} else {
		logError("only 'gb' or 'gbc' ROM files are supported.")
		os.Exit(1)
	}

	if hasInfo {
		rom := cartridge.Decode(readRom(file))
		if err := cmd.PrintRomInfo(rom, file); err != nil {
			fmt.Printf("\nProblematic ROM: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if hasHex {
		rom := cartridge.Decode(readRom(file))
		cmd.Hexdump(rom)
		os.Exit(0)
	}
}

func printLogo() {
	fmt.Printf("  _. ._ _ |_  o ._ _   _   _|  _   _ \n (_| | (_ | | | | | | (/_ (_| (/_ _> \n")
	fmt.Printf("archimedes is a GameBoy (Color) emulator by drpaneas, written in Go.\n\n")
	fmt.Printf("  * Dedicated to my grandpa, who bought me my first DMG console.\n")
	fmt.Printf("  > Find more information at: https://github.com/drpaneas/archimedes\n\n")
}

func help() {
	fmt.Printf("Usage:\n  archimedes [rom file] [optional flags]\n\n")
	fmt.Println("Optional flags:")
	fmt.Println("\t-i, --info: displays ROM header info and quits.")
	fmt.Println("\t-d, --debug: enables verbosity output (useful for debugging).")
	fmt.Println("\t-x, --hexdump: displays the ROM in hexdump format.")
	fmt.Println("\t-h, --help: displays this message.")
}

// Exists reports whether the named file or directory exists.
func exists(path string, isDir bool) bool {
	if path == "" {
		fmt.Println("Path is empty")
		return false
	}

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) || os.IsPermission(err) {
			return false
		}
	}

	return isDir == info.IsDir()
}

// fileExists reports whether the provided file exists.
func fileExists(path string) bool {
	return exists(path, false)
}

func debug(s string) {
	if debugFlag {
		fmt.Printf("Debug: %s\n", s)
	}
}

func logError(s string) {
	fmt.Printf("Error: %s\n", s)
}

func readRom(file string) []byte {
	f, err := os.Open(file)

	if err != nil {
		println("Error opening the file. Please open a bug.")
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Error closing the file. Please open a bug.")
			os.Exit(1)
		}
	}(f)

	return getFileInBytes(bufio.NewReader(f))
}

func getFileInBytes(data *bufio.Reader) []byte {
	var fileInBytes []byte
	for {
		buf := make([]byte, 1)
		_, err := data.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Couldn't get file in bytes. Please file a bug")
				fmt.Println(err)
			}
			break // EOF
		}
		fileInBytes = append(fileInBytes, buf[0])
	}

	return fileInBytes
}
