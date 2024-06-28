# Archimedes

Archimedes is a command-line tool designed for Game Boy developers and enthusiasts. It provides detailed insights into Game Boy ROM files, including hex dumps, licensing information, entry points, and more. Whether you're debugging your latest header for your game or just curious about the internals of classic Game Boy titles, Archimedes is your go-to utility.

## Dedication

This project is lovingly dedicated and named after my grandpa, who bought me my first DMG Game Boy. His gift sparked a lifelong passion for gaming and technology that ultimately led to the creation of Archimedes. Thank you, Grandpa, for setting me on this path.

## Features

- **Hex Dump**: View a hex dump of the ROM file to inspect its raw data.
- **ROM Information**: Extract and display information about the ROM, such as the licensee code, ROM size, RAM size, and entry point.
- **Checksum Verification**: Verify the header checksum and global checksum of the ROM.
- **Logo Verification**: Check if the Nintendo logo in the ROM is correct.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

Clone the repository and build the project:

```sh
git clone https://github.com/yourusername/archimedes.git
cd archimedes
go build -o archimedes ./cmd
```

## Usage

To get information about a ROM:

```shell
  _. ._ _ |_  o ._ _   _   _|  _   _ 
 (_| | (_ | | | | | | (/_ (_| (/_ _> 
archimedes is a GameBoy (Color) emulator by drpaneas, written in Go.

  * Dedicated to my grandpa, who bought me my first DMG console.
  > Find more information at: https://github.com/drpaneas/archimedes


Usage of ./archimedes:
  archimedes [options] rom_file

Options:
  -h    Displays this message
  -i    Displays ROM header info and quits
  -x    Displays the ROM in hexdump format
```

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to improve the tool or add new features.

## License
This project is licensed under the MIT License - see the LICENSE file for details.