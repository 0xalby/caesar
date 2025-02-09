package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/pflag"
)

var Release string

func main() {
	// Setting flags
	shiftFlag := pflag.IntSliceP("shift", "s", []int{1}, "Shifts by a single value or a range in beetween -12 and 13")
	fullFlag := pflag.BoolP("full", "f", false, "Performs a full shift")
	versionFlag := pflag.BoolP("version", "v", false, "Show version number")
	pflag.Parse()
	// Show version number and exit
	if *versionFlag {
		fmt.Println("caesar version", Release)
		return
	}
	// Collecting inputs
	var input []string
	// Read from files
	if len(pflag.Args()) > 0 {
		for i, f := range pflag.Args() {
			file, err := os.Open(f)
			if err != nil {
				fmt.Println("failed to open file", err)
				continue
			}
			defer file.Close()
			content, err := io.ReadAll(file)
			if err != nil {
				fmt.Println("error reading file", err)
				continue
			}
			if i == 0 {
				input = append(input, string(content))
				continue
			}
			input = append(input, "\n"+string(content))
		}
	}
	// Read from stdin
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("failed to get stdin information", err)
		os.Exit(1)
	}
	// Is stdin is being piped?
	if info.Mode()&os.ModeCharDevice == 0 {
		content, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("failed to read from stdin", err)
			os.Exit(1)
		}
		input = append(input, string(content))
	}
	// Concatenate inputs
	data := strings.Join(input, "")
	if data == "" {
		pflag.Usage()
		return
	}
	// Full shift
	if *fullFlag {
		for i := -12; i < 13; i++ {
			result := Cipher(data, i)
			fmt.Print(result)
		}
		return
	}
	// Two values range
	if len(*shiftFlag) > 2 {
		return
	} else if len(*shiftFlag) == 2 {
		shft := *shiftFlag
		if shft[0] < -12 || shft[1] > 13 {
			fmt.Print("a shift range is between 1 and 26 which is the number of characters in the alphabet")
			return
		}
		var result string
		for sh := shft[0]; sh <= shft[1]; sh++ {
			result = Cipher(data, sh)
			fmt.Print(result)
		}
		// Single shift
	} else {
		shft := *shiftFlag
		result := Cipher(data, shft[0])
		fmt.Print(result)
	}
}

// Ceaser cipher algorithm with the 26 alphabet letters and 10 decimal numbers
func Cipher(text string, shift int) string {
	var result strings.Builder
	letterShift := (shift%26 + 26) % 26
	numShift := (shift%10 + 10) % 10
	for _, char := range text {
		if unicode.IsLower(char) {
			result.WriteRune('a' + (char-'a'+rune(letterShift))%26)
		} else if unicode.IsUpper(char) {
			result.WriteRune('A' + (char-'A'+rune(letterShift))%26)
		} else if unicode.IsDigit(char) {
			result.WriteRune('0' + (char-'0'+rune(numShift))%10)
		} else {
			// Ignoring special characters
			result.WriteRune(char)
		}
	}
	return result.String()
}
