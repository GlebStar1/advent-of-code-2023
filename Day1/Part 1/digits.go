package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	parentDir := filepath.Dir(currentDir)

	files, err := os.ReadDir(parentDir)
	if err != nil {
		fmt.Println("Something went wrong when scanning folder")
		return
	}

	for _, file := range files {
		if file.Name() == "input.txt" {
			var sum int

			data, err := os.Open(parentDir + `\` + file.Name())
			if err != nil {
				panic(err.Error())
			}

			scanner := bufio.NewScanner(data)
			for scanner.Scan() {
				line := scanner.Text()
				sum += getCurrentLineDigit(line)
			}
			fmt.Printf("Sum is: %v", sum)
		}
	}
}

func getDigit(digitString string) int {
	var digit int
	_, err := fmt.Sscan(digitString, &digit)
	if err != nil {
		fmt.Println("File is incorrect.")
	}

	return digit
}

func getCurrentLineDigit(line string) int {
	var digits []string
	for _, r := range line {
		if unicode.IsDigit(r) {
			digits = append(digits, string(r))
		}
	}
	return getDigit(digits[0] + digits[len(digits)-1])
}
