package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func findFirstDigit(s string) int {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return int(c - '0')
		}
	}
	return -1
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// List of strings containg calibration documents
	var calibrationDocuments []string = getInputs()

	// Track sum of all calibration values
	var sumCalibrationValues int = 0

	// Retrieve the first digit in the string and the last digit from the calibration document
	for _, calibrationDocument := range calibrationDocuments {
		// Get first and last digit from calibraiton document
		firstDigit := findFirstDigit(calibrationDocument)
		lastDigit := findFirstDigit(reverseString(calibrationDocument))

		// Calculate calibration value by combining first and last digit into a string
		calibrationValue := fmt.Sprintf("%d%d", firstDigit, lastDigit)
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
