// Coded By : 0xdy44 and sf
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func GetAscii(char int, fileContent string) string {
	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	asciiOffset := char - 32
	lines := strings.Split(fileContent, "\n\n")
	return lines[asciiOffset]
}

func writeToFile(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func VirInput(input string) []string {
	for _, char := range input {
		if char < 32 || char > 126 {
			fmt.Println("Please provide a valid string ._. !")
			os.Exit(1)
		}
	}
	lines := strings.Split(input, "\\n")
	return lines
}

func isEmpty(input []string) bool {
	for _, char := range input {
		if char != "" {
			return false
		}
	}
	return true
}

func printH(lines [][]string) string {
	var result strings.Builder
	for i := 0; i < 8; i++ {
		for j := 0; j < len(lines); j++ {
			if i < len(lines[j]) {
				result.WriteString(lines[j][i])
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func printTxt(input, fileContent string) string {
	var nStr [][]string
	lines := VirInput(input)
	var result strings.Builder
	for i, line := range lines {
		input = string(line)
		if input == "" {
			if i != 0 || !isEmpty(lines) {
				result.WriteString("\n")
			}
			continue
		}
		for _, char := range line {
			asciiStr := GetAscii(int(char), fileContent)
			if asciiStr != "" {
				lines := strings.Split(asciiStr, "\n")
				nStr = append(nStr, lines)
			} else {
				nStr = append(nStr, []string{})
			}
		}
		result.WriteString(printH(nStr))
	}
	return result.String()
}

func wrintInFile(file_r string) {
	var outputFilename string
	inputUser := os.Args[2]
	styleBanner := file_r
	if strings.Contains(styleBanner, ".txt") {
		styleBanner = string(styleBanner)
	} else {
		styleBanner = string(styleBanner + ".txt")
	}
	content, err := os.ReadFile(styleBanner)
	if err != nil {
		fmt.Printf("Error reading file: %v \n", err)
		return
	}
	content = content[1:]
	flag.StringVar(&outputFilename, "output", "", "Output file name")
	flag.Parse()
	output := printTxt(inputUser, string(content))
	writeToFile(outputFilename, output)
}

func main() {
	args := os.Args
	inputLen := len(args) - 1
	if inputLen >= 4 || inputLen <= 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	} else if inputLen <= 2 {
		if strings.Contains(os.Args[1], "--output=") {
			wrintInFile("standard")
			return
		}
		styleBanner := ""
		if len(os.Args) == 3 {
			styleBanner = os.Args[2]
		} else {
			styleBanner = "standard"
		}
		if strings.Contains(styleBanner, ".txt") {
			styleBanner = string(styleBanner)
		} else {
			styleBanner = string(styleBanner + ".txt")
		}
		content, err := os.ReadFile(styleBanner)
		if err != nil {
			fmt.Printf("Error reading file: %v \n", err)
			return
		}
		content = content[1:]
		inputUser := os.Args[1]
		fmt.Print(printTxt(inputUser, string(content)))

	} else {
		wrintInFile(os.Args[3])
	}
}
