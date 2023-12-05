package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

var digits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Return inputs
	return inputs
}

func FindCalibrationValue(calibrationDocument string) string {
	// Given a calibration document, find the first and last digit
	// The digit can either be a numeric value or spelled out

	// Track first and last digit
	var firstDigit int
	var firstDigitIndex int = len(calibrationDocument) + 1
	var lastDigit int
	var lastDigitIndex int = -1

	// Loop through the digits and see if they are in the calibration document
	for _, digit := range digits {
		if strings.Contains(calibrationDocument, digit) {
			firstIndexOccurrence := strings.Index(calibrationDocument, digit)
			lastIndexOccurrence := strings.LastIndex(calibrationDocument, digit)
			if firstIndexOccurrence < firstDigitIndex {
				firstDigitIndex = firstIndexOccurrence
				firstDigit = digitMap[digit]
			}
			if lastIndexOccurrence > lastDigitIndex {
				lastDigitIndex = lastIndexOccurrence
				lastDigit = digitMap[digit]
			}
		}
	}

	// Return first and last digit
	return fmt.Sprintf("%d%d", firstDigit, lastDigit)
}

func main() {
	// List of strings containg calibration documents
	var calibrationDocuments []string = getInputs()

	// Track sum of all calibration values
	var sumCalibrationValues int = 0

	// Retrieve the first digit in the string and the last digit from the calibration document
	for _, calibrationDocument := range calibrationDocuments {
		// Calculate calibration value by combining first and last digit into a string
		calibrationValue := FindCalibrationValue(calibrationDocument)
		fmt.Println("Calibration value: ", calibrationValue)

		// Add calibration value to sum of all calibration values
		value, err := strconv.Atoi(calibrationValue)
		if err != nil {
			fmt.Println(err)
			return
		}
		sumCalibrationValues += value
	}

	// Print sum of all calibration values
	fmt.Println("Sum of all calibration values: ", sumCalibrationValues)
}
