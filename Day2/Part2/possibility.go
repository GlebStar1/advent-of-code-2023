package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
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

			allGames := getAllGames(file.Name())

			var sumOfPossPower int = 0

			for _, game := range allGames {

				sumOfPossPower += findFewestPossCubesPower(game.revealedCubes)
			}

			fmt.Println(sumOfPossPower)
		}
	}

}

func findFewestPossCubesPower(cubes []Cube) int {
	var allBlue []int
	var allRed []int
	var allGreen []int

	for _, cube := range cubes {
		if cube.colour == "blue" {
			allBlue = append(allBlue, cube.amount)
		}

		if cube.colour == "red" {
			allRed = append(allRed, cube.amount)
		}

		if cube.colour == "green" {
			allGreen = append(allGreen, cube.amount)
		}
	}

	maxBlue := slices.Max(allBlue[:])
	maxRed := slices.Max(allRed[:])
	maxGreen := slices.Max(allGreen[:])

	return maxBlue * maxRed * maxGreen
}

func getAllGames(fileName string) []Game {
	var games []Game
	data, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()

		games = append(games, mapToGame(line))
	}

	return games
}

func mapToGame(line string) Game {
	splittedLine := strings.Split(line, ":")

	gamePart := splittedLine[0]
	cubeSetPart := splittedLine[1]

	gameId := strings.Replace(gamePart, "Game ", "", 1)
	cubesSet := strings.Split(cubeSetPart, ";")

	game := Game{id: getDigit(gameId)}

	var revealedCubes []Cube

	for _, cubes := range cubesSet {

		revealedCubeSet := strings.Split(cubes, ",")

		for _, cube := range revealedCubeSet {

			cubeInfo := strings.Split(cube, " ")

			revealedCube := Cube{amount: getDigit(cubeInfo[1]), colour: cubeInfo[2]}

			revealedCubes = append(revealedCubes, revealedCube)
		}
	}

	game.revealedCubes = revealedCubes

	return game
}

func getDigit(digitString string) int {
	var digit int
	_, err := fmt.Sscan(digitString, &digit)
	if err != nil {
		fmt.Println("File is incorrect.")
	}

	return digit
}

type Game struct {
	id            int
	revealedCubes []Cube
}

type Cube struct {
	amount int
	colour string
}
