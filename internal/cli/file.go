package cli

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

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

// fileExists reports whether the provided file exists.
func fileExists(path string) bool {
	return exists(path, false)
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
