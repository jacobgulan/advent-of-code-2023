package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var symbols []rune = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '-', '=', '[', ']', '{', '}', '|', '\\', ';', ':', '\'', '"', ',', '<', '>', '/', '?'}

func GetNumbersNextToSymbol(puzzleInput [][]rune) map[string]int {
	// Create slice to hold numbers
	var numbers map[string]int = make(map[string]int)

	// Loop through each line and determine if there is a number adjacent to a symbol
	// The number can be on the left, right, top, bottom, or diagonal
	// The number continues until it hits a period or another symbol
	for i, runeRow := range puzzleInput {
		for j, r := range runeRow {
			if isSymbol(r) {
				// Check left
				if j-1 >= 0 {
					if unicode.IsDigit(runeRow[j-1]) {
						startIndex := getStartIndex(runeRow, j-1)
						value, position := getNumberAndPosition(runeRow, startIndex, i)
						if value != -1 {
							numbers[position] = value
						}
					}
				}

				// Check right
				if j+1 < len(runeRow) {
					if unicode.IsDigit(runeRow[j+1]) {
						startIndex := getStartIndex(runeRow, j+1)
						value, position := getNumberAndPosition(runeRow, startIndex, i)
						if value != -1 {
							numbers[position] = value
						}
					}
				}

				// Check top
				if i-1 >= 0 {
					if unicode.IsDigit(puzzleInput[i-1][j]) {
						startIndex := getStartIndex(puzzleInput[i-1], j)
						value, position := getNumberAndPosition(puzzleInput[i-1], startIndex, i-1)
						if value != -1 {
							numbers[position] = value
						}
					}
				}

				// Check bottom
				if i+1 < len(puzzleInput) {
					if unicode.IsDigit(puzzleInput[i+1][j]) {
						startIndex := getStartIndex(puzzleInput[i+1], j)
						value, position := getNumberAndPosition(puzzleInput[i+1], startIndex, i+1)
						if value != -1 {
							numbers[position] = value
						}
					}
				}

				// Check top left
				if i-1 >= 0 && j-1 >= 0 {
					if unicode.IsDigit(puzzleInput[i-1][j-1]) {
						startIndex := getStartIndex(puzzleInput[i-1], j-1)
						value, position := getNumberAndPosition(puzzleInput[i-1], startIndex, i-1)
						if value != -1 {
							numbers[position] = value
						}
					}
				}

				// Check top right
				if i-1 >= 0 && j+1 < len(runeRow) {
					if unicode.IsDigit(puzzleInput[i-1][j+1]) {
						startIndex := getStartIndex(puzzleInput[i-1], j+1)
						value, position := getNumberAndPosition(puzzleInput[i-1], startIndex, i-1)
						if value != -1 {
							numbers[position] = value
						}
					}
				}

				// Check bottom left
				if i+1 < len(puzzleInput) && j-1 >= 0 {
					if unicode.IsDigit(puzzleInput[i+1][j-1]) {
						startIndex := getStartIndex(puzzleInput[i+1], j-1)
						value, position := getNumberAndPosition(puzzleInput[i+1], startIndex, i+1)
						if value != -1 {
							numbers[position] = value
						}
					}
				}

				// Check bottom right
				if i+1 < len(puzzleInput) && j+1 < len(runeRow) {
					if unicode.IsDigit(puzzleInput[i+1][j+1]) {
						startIndex := getStartIndex(puzzleInput[i+1], j+1)
						value, position := getNumberAndPosition(puzzleInput[i+1], startIndex, i+1)
						if value != -1 {
							numbers[position] = value
						}
					}
				}
			}
		}
	}

	return numbers
}

func isSymbol(r rune) bool {
	for _, symbol := range symbols {
		if r == symbol {
			return true
		}
	}
	return false
}

func getInputMatrix() [][]rune {
	// Open input.txt file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file line by line
	var inputMatrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputMatrix = append(inputMatrix, []rune(scanner.Text()))
	}

	// Return slice of strings
	return inputMatrix
}

// Determine start index
func getStartIndex(runeRow []rune, givenIndex int) int {
	// Find furthermost left index that is a number
	var startIndex int = givenIndex
	for i := givenIndex; i >= 0; i-- {
		if unicode.IsDigit(runeRow[i]) {
			startIndex = i
		} else {
			break
		}
	}
	return startIndex
}

// Given a rune row and a start index, return the number and its matrix position
func getNumberAndPosition(runeRow []rune, startIndex int, rowNumber int) (int, string) {
	number := ""
	numberIndicies := ""

	for i := startIndex; i < len(runeRow); i++ {
		if unicode.IsDigit(runeRow[i]) {
			number += string(runeRow[i])
			numberIndicies += fmt.Sprintf("%d,", i)
		} else {
			break
		}
	}

	if number == "" {
		return -1, ""
	}
	value, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}

	return value, fmt.Sprintf("%d:%s", rowNumber, numberIndicies)
}

func main() {
	// Get puzzle input
	inputMatrix := getInputMatrix()

	// Part 1
	// Create slice to hold the sum of each line
	var totalSum int = 0
	numbers := GetNumbersNextToSymbol(inputMatrix)

	// Sum all the numbers
	for _, value := range numbers {
		totalSum += value
	}

	fmt.Println("Part 1:", totalSum)

}
