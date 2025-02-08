# Caesar
Rotate anything with the Caesar cipher using a single shift or a range of them

> This was inspired by Lukas Kyle researches in Silo

## Installation
* Download the release
* Install with go
```sh
go install github.com/0xalby/caesar@latest
```
* Compile from source
```sh
git clone https://github.com/0xalby/caesar
cd caesar
go mod tidy
make build
```
## Usage
### From files
```sh
caesar rotate -s 1,10 rotated.txt
```
### From standard input
```sh
echo "tupjdjtn" | caesar rotate -f
```
### Help
```
Usage of ./bin/caesar_linux_amd64:
  -f, --full         Performs a full shift
  -s, --shift ints   Shifts by a single value or a range in beetween 1 and 26 (default [1])
  -v, --version      Show version number
```