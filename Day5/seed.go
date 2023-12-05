package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

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
	ch := make(chan int)
	mu := sync.Mutex{}
	var wg sync.WaitGroup
	seeds, sourceDestinationMaps := parseInput(input)
	wg.Add(len(seeds) * 100)

	minLocation := math.MaxInt64

	if seedsAsRanges {
		go getSeedNumbersAsRanges(seeds, ch, &wg)
	} else {
		go getSeedNumbers(seeds, ch, &wg)
	}

	resultCh := make(chan int)
	go func() {
		count := 0
		startTime := time.Now()
		for seed := range ch {
			count++
			if count%1000000 == 0 {
				fmt.Printf("Count %d: Min location %d: Run duration %s\n", count, minLocation, time.Since(startTime))
			}
			dest := getDestination(seed, &sourceDestinationMaps)
			// Lock the minLocation
			mu.Lock()
			if dest < minLocation {
				minLocation = dest
				fmt.Printf("New min location %d\n", minLocation)
			}
			mu.Unlock()
		}
		// return minLocation
		resultCh <- minLocation
	}()
	result := <-resultCh
	return fmt.Sprintf("%d", result)
}

func getDestination(input int, destMaps *[]SourceDestinationMap) int {
	starting := "seed"
	for _, rangeMap := range *destMaps {
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
	return input
}

func parseInput(input string) ([]int, []SourceDestinationMap) {
	sections := strings.Split(input, "\r\n\r\n")
	seedSection := sections[0]
	seedLine := strings.Split(seedSection, "\n")[0]
	seedLine = strings.TrimPrefix(seedLine, "seeds: ")
	seedLine = strings.TrimSuffix(seedLine, "\r")
	seedStrings := strings.Split(seedLine, " ")

	seeds := parseSeeds(seedStrings)

	sourceDestinationMaps := make([]SourceDestinationMap, len(sections)-1)
	remainingSections := sections[1:]
	for i, section := range remainingSections {
		section = strings.TrimSuffix(section, "\r")
		sectionLines := strings.Split(section, "\n")
		sourceDestinationMaps[i] = parseSourceDestinationMap(sectionLines)
	}
	return seeds, sourceDestinationMaps
}

func parseSeeds(seedStrings []string) []int {
	seeds := make([]int, len(seedStrings))
	for i, seedString := range seedStrings {
		seed, err := strconv.Atoi(seedString)
		if err != nil {
			panic(err)
		}
		seeds[i] = seed
	}
	return seeds
}

func getSeedNumbers(seeds []int, ch chan int, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()
	for _, seed := range seeds {
		ch <- seed
	}
}

func getSeedNumbersAsRanges(seedRanges []int, ch chan int, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()
	if len(seedRanges)%2 != 0 {
		panic("Seeds must be even length")
	}

	for i := 0; i < len(seedRanges); i += 2 {
		start := seedRanges[i]
		r := seedRanges[i+1]
		for j := 0; j < r; j++ {
			ch <- start + j
		}
	}
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
