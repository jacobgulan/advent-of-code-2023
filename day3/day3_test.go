package main

import (
	"testing"
)

var inputString = [][]rune{
	[]rune("...................548............456#..815....$...@.....357.........619......*.............@.................-.......................15...."),
	[]rune(".865../....404..994.354........=.#738....535....+..=..........134.........................=..862.......215.939.......821*....#.............."),
	[]rune("......*......&....284................887.183....#....446*220.749...........732....538...................327.......356...793.....703........."),
	[]rune("...*................130...792........353.....*.795..........................383..690..53.....*.......385.&677..598*.......=....500.34./....."),
}

var inputString2 = [][]rune{
	[]rune("467..114.."),
	[]rune("...*......"),
	[]rune("..35..633."),
	[]rune("......#..."),
	[]rune("617*......"),
	[]rune(".....+.58."),
	[]rune("..592....."),
	[]rune("......755."),
	[]rune("...$.*...."),
	[]rune(".664.598.."),
}

func TestGetNumbersNextToSymbol(t *testing.T) {

	expectedNumbers := [][]int{
		{456},
		{404, 738, 862, 939, 821},
		{446, 220, 327, 356, 793},
		{795, 677, 598},
	}
	sumExpected := 0
	for _, row := range expectedNumbers {
		for _, value := range row {
			sumExpected += value
		}
	}

	actual := GetNumbersNextToSymbol(inputString)
	sumActual := 0
	for _, value := range actual {
		sumActual += value
	}

	// Assert expected sum equals actual sum
	if sumExpected != sumActual {
		t.Errorf("Expected sum of %d, got %d", sumExpected, sumActual)
	}
}

func TestSumPart2Numbers(t *testing.T) {
	numbersAndAsterisks := GetNumbersNextToSymbolAndAsterisks(inputString)
	expectedSum := (821 * 793) + (446 * 220) + (356 * 598)
	actualSum := SumPart2Numbers(inputString, numbersAndAsterisks)

	if expectedSum != actualSum {
		t.Errorf("Part 2: Expected sum of %d, got %d", expectedSum, actualSum)
	}

	expectedSum2 := 467835
	actualSum2 := SumPart2Numbers(inputString2, GetNumbersNextToSymbolAndAsterisks(inputString2))

	if expectedSum2 != actualSum2 {
		t.Errorf("Part 2: Expected sum of %d, got %d", expectedSum2, actualSum2)
	}
}
