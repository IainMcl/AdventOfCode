package day5

import (
	"strconv"
	"strings"
)

type Seeds struct {
	Seeds []int
}

type SourceDestinationMap struct {
	Source      string
	Destination string
	Ranges      []Range
}

type Range struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func Run(input string, seedsAsRanges bool) string {
	seeds, sourceDestinationMaps := parseInput(input, seedsAsRanges)

	locations := make([]int, len(seeds.Seeds))
	for i, seed := range seeds.Seeds {
		starting := "seed"
		input := seed
		for range locations {
			for _, rangeMap := range sourceDestinationMaps {
				if rangeMap.Source == starting {
					starting = rangeMap.Destination
					for _, r := range rangeMap.Ranges {
						if (input >= r.SourceRangeStart) && (input < r.SourceRangeStart+r.RangeLength) {
							input = r.DestinationRangeStart + (input - r.SourceRangeStart)
							break
						}
					}
					// Input not in any of the ranges so the input stays the same
				}
			}
			break
		}
		locations[i] = input
	}

	min := locations[0]
	for _, location := range locations {
		if location < min {
			min = location
		}
	}
	return strconv.Itoa(min)
}

func parseInput(input string, seedsAsRanges bool) (Seeds, []SourceDestinationMap) {
	sections := strings.Split(input, "\r\n\r\n")
	seedSection := sections[0]
	seedLine := strings.Split(seedSection, "\n")[0]
	seedLine = strings.TrimPrefix(seedLine, "seeds: ")
	seedLine = strings.TrimSuffix(seedLine, "\r")
	seedStrings := strings.Split(seedLine, " ")

	var seeds Seeds
	if seedsAsRanges {
		seeds = parseSeeedsAsRanges(seedStrings)
	} else {
		seeds = parseSeeds(seedStrings)
	}

	sourceDestinationMaps := make([]SourceDestinationMap, len(sections)-1)
	remainingSections := sections[1:]
	for i, section := range remainingSections {
		section = strings.TrimSuffix(section, "\r")
		sectionLines := strings.Split(section, "\n")
		sourceDestinationMaps[i] = parseSourceDestinationMap(sectionLines)
	}
	return seeds, sourceDestinationMaps
}

func parseSeeds(seedStrings []string) Seeds {
	seeds := Seeds{
		Seeds: make([]int, len(seedStrings)),
	}
	for i, seedString := range seedStrings {
		seed, err := strconv.Atoi(seedString)
		if err != nil {
			panic(err)
		}
		seeds.Seeds[i] = seed
	}
	return seeds
}

func parseSeeedsAsRanges(seedStrings []string) Seeds {
	if len(seedStrings)%2 != 0 {
		panic("Seeds must be even length")
	}
	seedsList := make([]int, 0)
	for i := 0; i < len(seedStrings); i += 2 {
		start, err := strconv.Atoi(seedStrings[i])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(seedStrings[i+1])
		if err != nil {
			panic(err)
		}
		for j := 0; j < r; j++ {
			seedsList = append(seedsList, start+j)
		}
	}

	return Seeds{Seeds: seedsList}
}

func parseSourceDestinationMap(sectionLines []string) SourceDestinationMap {
	sourceDestinationMap := SourceDestinationMap{}
	sourceDestinationMap.Ranges = make([]Range, len(sectionLines)-1)

	sourceDestinationString := strings.Split(sectionLines[0], " ")[0]
	sourceDestinationStringSplit := strings.Split(sourceDestinationString, "-")
	sourceDestinationMap.Source = sourceDestinationStringSplit[0]
	sourceDestinationMap.Destination = sourceDestinationStringSplit[2]

	for i, line := range sectionLines[1:] {
		line = strings.TrimSuffix(line, "\r")
		lineSplit := strings.Split(line, " ")
		destRangeStart, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			panic(err)
		}
		sourceRangeStart, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			panic(err)
		}
		rangeLength, err := strconv.Atoi(lineSplit[2])
		if err != nil {
			panic(err)
		}

		sourceDestinationMap.Ranges[i] = Range{
			DestinationRangeStart: destRangeStart,
			SourceRangeStart:      sourceRangeStart,
			RangeLength:           rangeLength,
		}
	}
	return sourceDestinationMap
}
