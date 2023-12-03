package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Hello Go!")

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Something went wrong when scanning folder")
		return
	}

	for _, file := range files {
		if file.Name() == "input.txt" {
			var sum int
			data, err := os.Open(file.Name())
			if err != nil {
				fmt.Println("Error reading input file.")
				return
			}
			scanner := bufio.NewScanner(data)

			for scanner.Scan() {
				line := scanner.Text()
				sum += getCurrentLineDigit(line)
			}
			fmt.Println(sum)
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
