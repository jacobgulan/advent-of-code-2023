package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Read input.txt file and return a slice of strings
func getInputs() []string {
	// Open input.txt file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file line by line
	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	// Return slice of strings
	return inputs
}

func parseGameString(gameString string) map[string]int {
	// Store the color-count pairs for each game in a map
	maximumColorCount := map[string]int{"red": 0, "blue": 0, "green": 0}

	// Remove everything up to the first colon
	gameString = strings.Split(gameString, ": ")[1]

	// Split the string into rounds
	rounds := strings.Split(gameString, "; ")

	for _, round := range rounds {
		// Split the round into its colors
		colors := strings.Split(round, ", ")

		// Split each color into its color and count
		for _, color := range colors {
			colorCount := strings.Split(color, " ")
			count, colorName := colorCount[0], colorCount[1]

			// Convert count to an int
			countInt, err := strconv.Atoi(count)
			if err != nil {
				log.Fatal(err)
			}

			// Determine if count exceeds current maximum
			if countInt > maximumColorCount[colorName] {
				maximumColorCount[colorName] = countInt
			}
		}

	}

	return maximumColorCount
}

// Given a map of color-count pairs, determine if a game is possible
// A game is possible if the values of each color does not exceed the maxiumum value
// passed into the function
func isGamePossible(colorMap map[string]int, maxRed int, maxBlue int, maxGreen int) bool {
	return maxRed >= colorMap["red"] && maxBlue >= colorMap["blue"] && maxGreen >= colorMap["green"]
}

func main() {
	// Get inputs from input.txt file
	inputs := getInputs()

	// Part 1
	gameResults := make(map[int]map[string]int)
	gameNumber := 0
	gameIDSum := 0
	for _, gameString := range inputs {
		// Increment game number
		gameNumber++
		colorMap := parseGameString(gameString)

		// Add color map to game results
		gameResults[gameNumber] = colorMap

		// Determine if game is possible
		if isGamePossible(colorMap, 12, 14, 13) {
			fmt.Println("Game", gameNumber, "is possible")
			gameIDSum += gameNumber
		}
	}

	fmt.Println(gameIDSum)

}
