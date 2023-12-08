package main

import (
	"testing"
)

func TestGetNumbersNextToSymbol(t *testing.T) {
	inputString := [][]rune{
		[]rune("...................548............456#..815....$...@.....357.........619......*.............@.................-.......................15...."),
		[]rune(".865../....404..994.354........=.#738....535....+..=..........134.........................=..862.......215.939.......821*....#.............."),
		[]rune("......*......&....284................887.183....#....446*220.749...........732....538...................327.......356...793.....703........."),
		[]rune("...*................130...792........353.....*.795..........................383..690..53.....*.......385.&677..598*.......=....500.34./....."),
	}

	expected := [][]int{
		{456},
		{404, 738, 862, 939, 821},
		{446, 220, 327, 356, 793},
		{795, 677, 598},
	}
	sumExpected := 0
	for _, row := range expected {
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
