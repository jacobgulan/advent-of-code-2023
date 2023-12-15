package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func getInputs() map[int]map[string][]string {
	// Open input.txt file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Each row number contains a slice for "winningNumbers" and a slice for "myNumbers"
	// Each row is prefixed with "Card #:" where # is the row number and can be any number of digits long
	// The "winningNumbers" and "myNumbers" are separated by a pipe character "|"
	// All the numbers in "winningNumbers" and "myNumbers" are separated by a space character " "
	inputMap := make(map[int]map[string][]string)

	// Read file line by line
	scanner := bufio.NewScanner(file)
	rowNumber := 1
	for scanner.Scan() {
		// Split line by pipe character
		line := scanner.Text()

		// Remove "Card #:" prefix
		// The # can be any number of digits long
		line = strings.Split(line, ":")[1]

		// Split line by pipe character
		lineSplit := strings.Split(line, "|")

		// Get winning numbers
		winningNumbersList := strings.Split(lineSplit[0], " ")
		winningNumbers := []string{}
		for i := range winningNumbersList {
			winningNumbersList[i] = strings.TrimSpace(winningNumbersList[i])
			if winningNumbersList[i] != "" {
				winningNumbers = append(winningNumbers, winningNumbersList[i])
			}
		}

		// Get my numbers
		myNumbersList := strings.Split(lineSplit[1], " ")
		myNumbers := []string{}
		for i := range myNumbersList {
			myNumbersList[i] = strings.TrimSpace(myNumbersList[i])
			if myNumbersList[i] != "" {
				myNumbers = append(myNumbers, myNumbersList[i])
			}
		}

		// Add winning numbers and my numbers to inputMap
		inputMap[rowNumber] = map[string][]string{
			"winningNumbers": winningNumbers,
			"myNumbers":      myNumbers,
		}
		rowNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Return inputs
	return inputMap
}

func FindMyPoints(inputMap map[int]map[string][]string) int {
	// Given a map of winning numbers and my numbers, find the winning numbers for each row
	// A winning number is a number that appears in both the winning numbers and my numbers
	myPoints := 0

	// Loop through each row
	for _, row := range inputMap {
		// Loop through each winning number
		exponent := -1
		for _, winningNumber := range row["winningNumbers"] {
			// Loop through each of my numbers
			for _, myNumber := range row["myNumbers"] {
				// If the winning number matches my number, add it to my winning numbers
				if winningNumber == myNumber {
					exponent++
				}
			}
		}
		if exponent >= 0 {
			myPoints += int(math.Pow(2, float64(exponent)))
		}
	}

	// Return my winning numbers
	return myPoints
}

func FindScratchcards(inputMap map[int]map[string][]string) int {
	// The key is the card number
	// The value is the number copies of the card including the original
	var scratchCardMap map[int][]map[string][]string = make(map[int][]map[string][]string)

	totalScratchCards := 0

	// Store the original scratchcard and its copies in "scratchcards"
	// Convert inputMap to scratchcards
	for i := 1; i <= len(inputMap); i++ {
		scratchCardMap[i] = []map[string][]string{inputMap[i]}
	}

	// Loop through each row
	for i := 1; i <= len(scratchCardMap); i++ {
		// Loop through each scratchcard
		for _, scratchCard := range scratchCardMap[i] {
			totalScratchCards++
			matches := 0

			for _, winningNumber := range scratchCard["winningNumbers"] {
				for _, myNumber := range scratchCard["myNumbers"] {
					// If the winning number matches my number, add it to the matches
					if winningNumber == myNumber {
						matches += 1
					}
				}
			}

			for j := 1; j <= matches; j++ {
				// Add extra scratchcards for as many as there are found
				scratchCardMap[j+i] = append(scratchCardMap[j+i], scratchCardMap[j+i][0])
			}
		}
	}

	return totalScratchCards
}

func Main() {
	main()
}

func main() {
	inputMap := getInputs()
	myPoints := FindMyPoints(inputMap)
	fmt.Println("My points: ", myPoints)
	totalScratchCards := FindScratchcards(inputMap)
	fmt.Println("Total scratchcards: ", totalScratchCards)
}
