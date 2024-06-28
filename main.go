package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/drpaneas/archimedes/cmd"
	"github.com/drpaneas/archimedes/pkg/cartridge"
)

const (
	errRomFileMissing  = "the ROM file is missing\nRun 'archimedes --help' for usage"
	errUnsupportedFile = "only 'gb' or 'gbc' ROM files are supported"
	errFileNotExist    = "'%s' does not exist"
)

func init() {
	// Customize flag.Usage to display a more helpful message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\nUsage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  archimedes [options] rom_file\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\nOptions:\n")
		flag.PrintDefaults()
	}
}

func main() {
	printLogo()
	processInput()
}

func printLogo() {
	fmt.Printf("  _. ._ _ |_  o ._ _   _   _|  _   _ \n (_| | (_ | | | | | | (/_ (_| (/_ _> \n")
	fmt.Printf("archimedes is a GameBoy (Color) emulator by drpaneas, written in Go.\n\n")
	fmt.Printf("  * Dedicated to my grandpa, who bought me my first DMG console.\n")
	fmt.Printf("  > Find more information at: https://github.com/drpaneas/archimedes\n\n")
}

func processInput() {
	infoFlag := flag.Bool("i", false, "Displays ROM header info and quits")
	hexdumpFlag := flag.Bool("x", false, "Displays the ROM in hexdump format")
	helpFlag := flag.Bool("h", false, "Displays this message")

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		fmt.Println("Error:", errRomFileMissing)
		os.Exit(1)
	}

	romFile, err := validateAndReturnFile(flag.Arg(0))
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	romData := cartridge.Decode(readRom(romFile))

	// Error handling for Decode remains as a TODO.
	fmt.Printf("Processing ROM file: %s\n\n", romFile)

	if *infoFlag {
		cmd.PrintRomInfo(romData, romFile)
		os.Exit(0)
	}

	if *hexdumpFlag {
		cmd.Hexdump(romData)
		os.Exit(0)
	}

}

// readRom reads the ROM file and returns it as a byte slice.
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

// getFileInBytes reads the file and returns it as a byte slice.
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

func validateAndReturnFile(file string) (string, error) {
	if !(strings.HasSuffix(file, "gb") || strings.HasSuffix(file, "gbc")) {
		return "", errors.New(errUnsupportedFile)
	}

	if !fileExists(file) {
		return "", fmt.Errorf(errFileNotExist, file)
	}

	return file, nil
}

// fileExists reports whether the provided file exists.
func fileExists(path string) bool {
	return exists(path, false)
}

// exists reports whether the named file or directory exists.
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
