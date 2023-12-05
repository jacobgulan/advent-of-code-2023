package main

import "testing"

func TestFindCalibrationValue(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"Test 1", "3nine9fivetwo9twohxhc8", "38"},
		{"Test 2", "7one718onegfqtdbtxfcmd", "71"},
		{"Test 3", "g4", "44"},
		{"Test 4", "pbc19", "19"},
		{"Test 5", "bdnhvtsjmdnklsxbtmnztqjtpnz6fivesevenfourzddgsrfmlq", "64"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := FindCalibrationValue(tc.input)
			if output != tc.expectedOutput {
				t.Errorf("Expected %s, but got %s", tc.expectedOutput, output)
			}
		})
	}
}
