package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type Correspondence struct {
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

type CorrespondenceRange struct {
	destinationRangeStart *int
	sourceRangeStart      *int
	rangeLength           *int
}

type InputMap struct {
	Seeds                 []int
	SeedToSoil            []CorrespondenceRange
	SoilToFertilizer      []CorrespondenceRange
	FertilizerToWater     []CorrespondenceRange
	WaterToLight          []CorrespondenceRange
	LightToTemperature    []CorrespondenceRange
	TemperatureToHumidity []CorrespondenceRange
	HumidityToLocation    []CorrespondenceRange
}

func removeValue(s []string, val string) []string {
	var result []string
	for _, v := range s {
		if v != val {
			result = append(result, v)
		}
	}
	return result
}

func convertToCorrespondenceRange(inputLine string) CorrespondenceRange {
	lineSplit := strings.Split(inputLine, " ")
	correspondenceRange := CorrespondenceRange{}
	filteredLineSplit := removeValue(lineSplit, "")

	for i := range filteredLineSplit {
		lineSplit[i] = strings.TrimSpace(lineSplit[i])
		value, _ := strconv.Atoi(lineSplit[i])
		switch i {
		case 0:
			correspondenceRange.destinationRangeStart = &value
		case 1:
			correspondenceRange.sourceRangeStart = &value
		case 2:
			correspondenceRange.rangeLength = &value
		}
	}

	return correspondenceRange
}

func getInputs() InputMap {
	// Open input.txt file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputMap := InputMap{}

	// Read file line by line
	scanner := bufio.NewScanner(file)
	index := 0
	currentKey := "Seeds"

	for scanner.Scan() {
		// Split line by pipe character
		line := scanner.Text()

		// Read all the seeds from the first line
		if index == 0 {
			line = strings.Split(line, ":")[1]
			seeds := strings.Split(line, " ")
			for i := range seeds {
				seeds[i] = strings.TrimSpace(seeds[i])
				if seeds[i] != "" {
					value, _ := strconv.Atoi(seeds[i])
					inputMap.Seeds = append(inputMap.Seeds, value)
				}
			}
		}

		if line == "seed-to-soil map:" {
			currentKey = "SeedToSoil"
			continue
		}

		if line == "soil-to-fertilizer map:" {
			currentKey = "SoilToFertilizer"
			continue
		}

		if line == "fertilizer-to-water map:" {
			currentKey = "FertilizerToWater"
			continue
		}

		if line == "water-to-light map:" {
			currentKey = "WaterToLight"
			continue
		}

		if line == "light-to-temperature map:" {
			currentKey = "LightToTemperature"
			continue
		}

		if line == "temperature-to-humidity map:" {
			currentKey = "TemperatureToHumidity"
			continue
		}

		if line == "humidity-to-location map:" {
			currentKey = "HumidityToLocation"
			continue
		}

		switch currentKey {
		case "SeedToSoil":
			inputMap.SeedToSoil = append(inputMap.SeedToSoil, convertToCorrespondenceRange(line))
		case "SoilToFertilizer":
			inputMap.SoilToFertilizer = append(inputMap.SoilToFertilizer, convertToCorrespondenceRange(line))
		case "FertilizerToWater":
			inputMap.FertilizerToWater = append(inputMap.FertilizerToWater, convertToCorrespondenceRange(line))
		case "WaterToLight":
			inputMap.WaterToLight = append(inputMap.WaterToLight, convertToCorrespondenceRange(line))
		case "LightToTemperature":
			inputMap.LightToTemperature = append(inputMap.LightToTemperature, convertToCorrespondenceRange(line))
		case "TemperatureToHumidity":
			inputMap.TemperatureToHumidity = append(inputMap.TemperatureToHumidity, convertToCorrespondenceRange(line))
		case "HumidityToLocation":
			inputMap.HumidityToLocation = append(inputMap.HumidityToLocation, convertToCorrespondenceRange(line))
		}

		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Return inputs
	return inputMap
}

func findCorrespondenceMap(destinationRangeStart, sourceRangeStart, rangeLength int) map[int]int {
	correspondenceMap := make(map[int]int)
	for i := sourceRangeStart; i < sourceRangeStart+rangeLength; i++ {
		correspondenceMap[i] = destinationRangeStart + i
	}
	return correspondenceMap
}

func FindSeedMapCorrespondence(inputMap InputMap) map[int]Correspondence {
	seedMapCorrespondence := make(map[int]Correspondence)
	// Maximum seed value
	maxSeed := 0
	for _, seed := range inputMap.Seeds {
		if seed > maxSeed {
			maxSeed = seed

		}
	}

	// Find soil correspondences for all possible seeds
	for _, seed := range inputMap.SeedToSoil {
		correspondenceMap := findCorrespondenceMap(*seed.destinationRangeStart, *seed.sourceRangeStart, *seed.rangeLength)
		for key, value := range correspondenceMap {
			seedMapCorrespondence[key] = Correspondence{Soil: value}
		}
		fmt.Println(seedMapCorrespondence)
	}

	for i := 0; i < maxSeed; i++ {
		if _, ok := seedMapCorrespondence[i]; !ok {
			seedMapCorrespondence[i] = Correspondence{Soil: i}
		}
	}

	// Find fertilizer correspondences for all possible seeds
	for _, soil := range inputMap.SoilToFertilizer {
		correspondenceMap := findCorrespondenceMap(*soil.destinationRangeStart, *soil.sourceRangeStart, *soil.rangeLength)
		for key, value := range correspondenceMap {
			seedMapCorrespondence[key] = Correspondence{Fertilizer: value}
		}
	}

	for i := 0; i < maxSeed; i++ {
		if _, ok := seedMapCorrespondence[i]; !ok {
			seedMapCorrespondence[i] = Correspondence{Fertilizer: i}
		}
	}

	// Find water correspondences for all possible seeds
	for _, fertilizer := range inputMap.FertilizerToWater {
		correspondenceMap := findCorrespondenceMap(*fertilizer.destinationRangeStart, *fertilizer.sourceRangeStart, *fertilizer.rangeLength)
		for key, value := range correspondenceMap {
			seedMapCorrespondence[key] = Correspondence{Water: value}
		}
	}

	for i := 0; i < maxSeed; i++ {
		if _, ok := seedMapCorrespondence[i]; !ok {
			seedMapCorrespondence[i] = Correspondence{Water: i}
		}
	}

	// Find light correspondences for all possible seeds
	for _, water := range inputMap.WaterToLight {
		correspondenceMap := findCorrespondenceMap(*water.destinationRangeStart, *water.sourceRangeStart, *water.rangeLength)
		for key, value := range correspondenceMap {
			seedMapCorrespondence[key] = Correspondence{Light: value}
		}
	}

	for i := 0; i < maxSeed; i++ {
		if _, ok := seedMapCorrespondence[i]; !ok {
			seedMapCorrespondence[i] = Correspondence{Light: i}
		}
	}

	// Find temperature correspondences for all possible seeds
	for _, light := range inputMap.LightToTemperature {
		correspondenceMap := findCorrespondenceMap(*light.destinationRangeStart, *light.sourceRangeStart, *light.rangeLength)
		for key, value := range correspondenceMap {
			seedMapCorrespondence[key] = Correspondence{Temperature: value}
		}
	}

	for i := 0; i < maxSeed; i++ {
		if _, ok := seedMapCorrespondence[i]; !ok {
			seedMapCorrespondence[i] = Correspondence{Temperature: i}
		}
	}

	// Find humidity correspondences for all possible seeds
	for _, temperature := range inputMap.TemperatureToHumidity {
		correspondenceMap := findCorrespondenceMap(*temperature.destinationRangeStart, *temperature.sourceRangeStart, *temperature.rangeLength)
		for key, value := range correspondenceMap {
			seedMapCorrespondence[key] = Correspondence{Humidity: value}
		}
	}

	for i := 0; i < maxSeed; i++ {
		if _, ok := seedMapCorrespondence[i]; !ok {
			seedMapCorrespondence[i] = Correspondence{Humidity: i}
		}
	}

	// Find location correspondences for all possible seeds
	for _, humidity := range inputMap.HumidityToLocation {
		correspondenceMap := findCorrespondenceMap(*humidity.destinationRangeStart, *humidity.sourceRangeStart, *humidity.rangeLength)
		for key, value := range correspondenceMap {
			seedMapCorrespondence[key] = Correspondence{Location: value}
		}
	}

	for i := 0; i < maxSeed; i++ {
		if _, ok := seedMapCorrespondence[i]; !ok {
			seedMapCorrespondence[i] = Correspondence{Location: i}
		}
	}

	return seedMapCorrespondence

}

func Main() {
	main()
}

func main() {
	// Create a channel to receive OS signals.
	sigs := make(chan os.Signal, 1)

	// Register the channel to receive SIGINT and SIGTERM signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine that will do something when a signal is received.
	go func() {
		sig := <-sigs
		fmt.Println("\nReceived signal: ", sig)
		// Do any necessary cleanup here.
		// Then call os.Exit to terminate the program gracefully.
		os.Exit(0)
	}()

	inputMap := getInputs()
	fmt.Println(inputMap)
	seedMapCorrespondence := FindSeedMapCorrespondence(inputMap)

	// Part 1
	// Find the lowest location given from the initial seeds
	lowestLocation := 100000000000000
	for _, seed := range inputMap.Seeds {
		if seedMapCorrespondence[seed].Location < lowestLocation {
			lowestLocation = seedMapCorrespondence[seed].Location
		}
	}

	fmt.Println("Part 1:", lowestLocation)
}
