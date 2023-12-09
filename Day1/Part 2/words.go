package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
			count := 0
			for scanner.Scan() {
				count++
				line := scanner.Text()
				fmt.Printf("Line %v", count)
				sum += getCurrentLineDigit(line)
			}
			fmt.Printf("Sum is: %v", sum)
		}
	}
}

func getDigit(digitString string) int {
	var digit int
	digitsMap := make(map[string]string)

	digitsMap["one"] = "1"
	digitsMap["two"] = "2"
	digitsMap["three"] = "3"
	digitsMap["four"] = "4"
	digitsMap["five"] = "5"
	digitsMap["six"] = "6"
	digitsMap["seven"] = "7"
	digitsMap["eight"] = "8"
	digitsMap["nine"] = "9"

	digits := strings.Split(digitString, " ")
	firstDigit := digits[0]
	secondDigit := digits[len(digits)-1]
	var transformedDigits []string

	if len(firstDigit) > 1 {
		transformedDigits = append(transformedDigits, digitsMap[firstDigit])
	} else {
		transformedDigits = append(transformedDigits, firstDigit)
	}

	if len(secondDigit) > 1 {
		transformedDigits = append(transformedDigits, digitsMap[secondDigit])
	} else {
		transformedDigits = append(transformedDigits, secondDigit)
	}
	fmt.Println(transformedDigits)
	_, err := fmt.Sscan(strings.Join(transformedDigits, ""), &digit)
	if err != nil {
		fmt.Println("File is incorrect.")
	}

	return digit
}

func getCurrentLineDigit(line string) int {
	var digits []string

	pattern := "(one|two|three|four|five|six|seven|eight|nine|[1-9])"

	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	matches := re.FindAllString(line, -1)
	if len(matches) < 1 {
		return 0
	}
	for _, r := range matches {
		if len(r) > 0 && r != "" {
			digits = append(digits, r)
		}
	}
	return getDigit(digits[0] + " " + digits[len(digits)-1])
}
