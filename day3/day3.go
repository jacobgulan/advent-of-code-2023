package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func GetNumbersNextToSymbolAndAsterisks(puzzleInput [][]rune) map[string]int {
	// Create slice to hold numbers
	var numbersAndAsterisks map[string]int = make(map[string]int)

	// Loop through each line and determine if there is a number adjacent to a symbol
	// The number can be on the left, right, top, bottom, or diagonal
	// The number continues until it hits a period or another symbol
	for i, runeRow := range puzzleInput {
		for j, r := range runeRow {
			if isSymbol(r) {
				// Store the presence of an asterisk with the key being the row and column
				if r == '*' {
					numbersAndAsterisks[fmt.Sprintf("%d:%d", i, j)] = '*'
				}

				// Check left
				if j-1 >= 0 {
					if unicode.IsDigit(runeRow[j-1]) {
						startIndex := getStartIndex(runeRow, j-1)
						value, position := getNumberAndPosition(runeRow, startIndex, i)
						if value != -1 {
							numbersAndAsterisks[position] = value
						}
					}
				}

				// Check right
				if j+1 < len(runeRow) {
					if unicode.IsDigit(runeRow[j+1]) {
						startIndex := getStartIndex(runeRow, j+1)
						value, position := getNumberAndPosition(runeRow, startIndex, i)
						if value != -1 {
							numbersAndAsterisks[position] = value
						}
					}
				}

				// Check top
				if i-1 >= 0 {
					if unicode.IsDigit(puzzleInput[i-1][j]) {
						startIndex := getStartIndex(puzzleInput[i-1], j)
						value, position := getNumberAndPosition(puzzleInput[i-1], startIndex, i-1)
						if value != -1 {
							numbersAndAsterisks[position] = value
						}
					}
				}

				// Check bottom
				if i+1 < len(puzzleInput) {
					if unicode.IsDigit(puzzleInput[i+1][j]) {
						startIndex := getStartIndex(puzzleInput[i+1], j)
						value, position := getNumberAndPosition(puzzleInput[i+1], startIndex, i+1)
						if value != -1 {
							numbersAndAsterisks[position] = value
						}
					}
				}

				// Check top left
				if i-1 >= 0 && j-1 >= 0 {
					if unicode.IsDigit(puzzleInput[i-1][j-1]) {
						startIndex := getStartIndex(puzzleInput[i-1], j-1)
						value, position := getNumberAndPosition(puzzleInput[i-1], startIndex, i-1)
						if value != -1 {
							numbersAndAsterisks[position] = value
						}
					}
				}

				// Check top right
				if i-1 >= 0 && j+1 < len(runeRow) {
					if unicode.IsDigit(puzzleInput[i-1][j+1]) {
						startIndex := getStartIndex(puzzleInput[i-1], j+1)
						value, position := getNumberAndPosition(puzzleInput[i-1], startIndex, i-1)
						if value != -1 {
							numbersAndAsterisks[position] = value
						}
					}
				}

				// Check bottom left
				if i+1 < len(puzzleInput) && j-1 >= 0 {
					if unicode.IsDigit(puzzleInput[i+1][j-1]) {
						startIndex := getStartIndex(puzzleInput[i+1], j-1)
						value, position := getNumberAndPosition(puzzleInput[i+1], startIndex, i+1)
						if value != -1 {
							numbersAndAsterisks[position] = value
						}
					}
				}

				// Check bottom right
				if i+1 < len(puzzleInput) && j+1 < len(runeRow) {
					if unicode.IsDigit(puzzleInput[i+1][j+1]) {
						startIndex := getStartIndex(puzzleInput[i+1], j+1)
						value, position := getNumberAndPosition(puzzleInput[i+1], startIndex, i+1)
						if value != -1 {
							numbersAndAsterisks[position] = value
						}
					}
				}
			}
		}
	}

	return numbersAndAsterisks
}

func SumPart2Numbers(inputMatrix [][]rune, numbersAndAsterisks map[string]int) int {
	part2TotalSum := 0

	// Loop through each number and multiply adjacent numbers by each other if an asterisk is found as their adjacent symbol
	// Delete the asterisk and the number from the map after multiplying
	// The number can be on the left, right, top, bottom, or diagonal
	// The number is up to 3 digits long
	for key, value := range numbersAndAsterisks {
		if value != '*' {
			continue
		}

		// Get row and column
		rowAndColumn := key
		rowAndColumnSplit := strings.Split(rowAndColumn, ":")
		row, err := strconv.Atoi(rowAndColumnSplit[0])
		if err != nil {
			log.Fatal(err)
		}
		column, err := strconv.Atoi(rowAndColumnSplit[1])
		if err != nil {
			log.Fatal(err)
		}

		// Make a map to hold the multiplying numbers and their keys
		multipylingNumbers := make(map[string]int)

		// Find the starting indicies of numbers adjacent to the asterisk
		// Check left
		if column-1 >= 0 {
			if unicode.IsDigit(inputMatrix[row][column-1]) {
				startIndex := getStartIndex(inputMatrix[row], column-1)
				value, position := getNumberAndPosition(inputMatrix[row], startIndex, row)
				if value != -1 {
					multipylingNumbers[position] = value
				}
			}
		}

		// Check right
		if column+1 < len(inputMatrix[row]) {
			if unicode.IsDigit(inputMatrix[row][column+1]) {
				startIndex := getStartIndex(inputMatrix[row], column+1)
				value, position := getNumberAndPosition(inputMatrix[row], startIndex, row)
				if value != -1 {
					multipylingNumbers[position] = value
				}
			}
		}

		// Check top
		if row-1 >= 0 {
			if unicode.IsDigit(inputMatrix[row-1][column]) {
				startIndex := getStartIndex(inputMatrix[row-1], column)
				value, position := getNumberAndPosition(inputMatrix[row-1], startIndex, row-1)
				if value != -1 {
					multipylingNumbers[position] = value
				}
			}
		}

		// Check bottom
		if row+1 < len(inputMatrix) {
			if unicode.IsDigit(inputMatrix[row+1][column]) {
				startIndex := getStartIndex(inputMatrix[row+1], column)
				value, position := getNumberAndPosition(inputMatrix[row+1], startIndex, row+1)
				if value != -1 {
					multipylingNumbers[position] = value
				}
			}
		}

		// Check top left
		if row-1 >= 0 && column-1 >= 0 {
			if unicode.IsDigit(inputMatrix[row-1][column-1]) {
				startIndex := getStartIndex(inputMatrix[row-1], column-1)
				value, position := getNumberAndPosition(inputMatrix[row-1], startIndex, row-1)
				if value != -1 {
					multipylingNumbers[position] = value
				}
			}
		}

		// Check top right
		if row-1 >= 0 && column+1 < len(inputMatrix[row]) {
			if unicode.IsDigit(inputMatrix[row-1][column+1]) {
				startIndex := getStartIndex(inputMatrix[row-1], column+1)
				value, position := getNumberAndPosition(inputMatrix[row-1], startIndex, row-1)
				if value != -1 {
					multipylingNumbers[position] = value
				}
			}
		}

		// Check bottom left
		if row+1 < len(inputMatrix) && column-1 >= 0 {
			if unicode.IsDigit(inputMatrix[row+1][column-1]) {
				startIndex := getStartIndex(inputMatrix[row+1], column-1)
				value, position := getNumberAndPosition(inputMatrix[row+1], startIndex, row+1)
				if value != -1 {
					multipylingNumbers[position] = value
				}
			}
		}

		// Check bottom right
		if row+1 < len(inputMatrix) && column+1 < len(inputMatrix[row]) {
			if unicode.IsDigit(inputMatrix[row+1][column+1]) {
				startIndex := getStartIndex(inputMatrix[row+1], column+1)
				value, position := getNumberAndPosition(inputMatrix[row+1], startIndex, row+1)
				if value != -1 {
					multipylingNumbers[position] = value
				}
			}
		}

		if len(multipylingNumbers) == 0 || len(multipylingNumbers) == 1 {
			delete(numbersAndAsterisks, key)
			continue
		}

		// Multiply the numbers together
		product := 1
		for _, value := range multipylingNumbers {
			product *= value
		}

		// Delete the asterisk and the number from the map
		delete(numbersAndAsterisks, key)
		for multiplyKey := range multipylingNumbers {
			delete(numbersAndAsterisks, multiplyKey)
		}
		part2TotalSum += product

	}

	return part2TotalSum
}

func isSymbol(r rune) bool {
	for _, symbol := range symbols {
		if r == symbol {
			return true
		}
	}
	return false
}

// "3:144": 42
// "2:144": 356
// "3:111": 598

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

	for i := startIndex; i < len(runeRow); i++ {
		if unicode.IsDigit(runeRow[i]) {
			number += string(runeRow[i])
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

	return value, fmt.Sprintf("%d:%d", rowNumber, startIndex)
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

	var part2TotalSum int = 0
	numbersAndAsterisks := GetNumbersNextToSymbolAndAsterisks(inputMatrix)
	part2TotalSum = SumPart2Numbers(inputMatrix, numbersAndAsterisks)
	fmt.Println("Part 2:", part2TotalSum)
}
