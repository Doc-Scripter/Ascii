package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"art/ascii"
)

func main() {
	// Check command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <input string>")
		os.Exit(1)
	}

	input := strings.Join(os.Args[1:], " ")
	input = strings.ReplaceAll(input, "\\n", "\n")
	input = strings.ReplaceAll(input, " ", "      ")
	var v string
	if strings.Contains(input, "-th") {
		v = "-th"
	} else if strings.Contains(input, "-sh") {
		v = "-sh"
	} else if strings.Contains(input, "-s") {
		v = "s"
	}
	var file string
	switch v {

	case "-s":
		file = "standard.txt"

	case "-th":
		file = "thinkertoy.txt"

	case "-sh":
		file = "thinkertoy.txt"

	default:
		file = "standard.txt"
	}
	Err := ascii.Checkfiles(file)
	if Err != nil {
		os.Exit(0)
	}
	asciiArtMap, err := ascii.ReadAscii(file)
	if err != nil {
		fmt.Println("Error reading ASCII art file:", err)
		os.Exit(1)
	}

	pattern := regexp.MustCompile(`\S+(\".+\")\S+`)
	input = pattern.ReplaceAllString(input, "$1")

	// Process each argument
	result := strings.Split(input, "\n")
	for i, arg := range result {
		if input == "" {
			fmt.Println("Usage: go run . <input string>")
			os.Exit(1)
		}
		for _, char := range arg {
			if rune(char) < 32 || rune(char) > 126 {
				fmt.Println("input contains non-printable character")
				return
			}
		}
		if i < len(result) && arg == "" {
			fmt.Println()
			continue
		}

		// for _, art := range result {
		art := ascii.GenerateAsciiArt(arg, asciiArtMap)
		fmt.Print(art)
		// }
	}
}
