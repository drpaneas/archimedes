![logo.png](logo.png)

# Archimedes

Archimedes is a command-line tool designed for Game Boy developers and enthusiasts. It provides detailed insights into Game Boy ROM files, including hex dumps, licensing information, entry points, and more. Whether you're debugging your latest header for your game or just curious about the internals of classic Game Boy titles, Archimedes is your go-to utility.

```shell
# For example:
Processing ROM file: roms/mario.gb

== File Information ==
 Filename          : mario.gb
 Filesize          : 64 KB
 Checksum (CRC)    : 2C27EC70

== Header Information ==
 Title             : SUPER MARIOLAND
 Console           : Game Boy
 Publisher         : Nintendo
 Cartridge Type    : MBC1
 Super Gameboy     : No
 Color Gameboy     : No
 Unofficial Colors : No
 Rom Size          : 64 KiB (4 banks)
 SRAM Save Size    : None
 Is Japanese       : Yes
 Game Version      : 1
 Entry Point       : 00, C3, 50, 01 - (nop, jp $0150)
 Global Checksum   : 0x5ECF (Verified) ✓

== Boot Checks ==
 Header Checksum   : 0x9D (Verified) ✓
 Official Logo     : Verified ✓
```

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
go build
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

## Dedication

*This project is lovingly dedicated and named after my grandpa (Αμοιρίδης Αρχιμήδης), who bought me my first DMG Game Boy. His gift sparked a lifelong passion for gaming and technology that ultimately led to the creation of Archimedes. Thank you, Grandpa, for setting me on this path.*