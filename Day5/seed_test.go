package day5

import (
	"os"
	"testing"
)

// Test day5.Run
// Input:
// seeds: 79 14 55 13
//
// seed-to-soil map:
// 50 98 2
// 52 50 48
//
// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15
//
// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4
//
// water-to-light map:
// 88 18 7
// 18 25 70
//
// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13
//
// temperature-to-humidity map:
// 0 69 1
// 1 0 69
//
// humidity-to-location map:
// 60 56 37
// 56 93 442
// Output:
// 35
func TestRun(t *testing.T) {
	input := "seeds: 79 14 55 13\r\n\r\nseed-to-soil map:\n50 98 2\n52 50 48\r\n\r\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\r\n\r\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\r\n\r\nwater-to-light map:\n88 18 7\n18 25 70\r\n\r\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\r\n\r\ntemperature-to-humidity map:\n0 69 1\n1 0 69\r\n\r\nhumidity-to-location map:\n60 56 37\n56 93 442"
	expected := "35"
	actual := Run(input, false)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestInput(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	expected := "324724204"
	actual := Run(input, false)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestSeedsAsRanges(t *testing.T) {
	input := "seeds: 79 14 55 13\r\n\r\nseed-to-soil map:\n50 98 2\n52 50 48\r\n\r\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\r\n\r\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\r\n\r\nwater-to-light map:\n88 18 7\n18 25 70\r\n\r\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\r\n\r\ntemperature-to-humidity map:\n0 69 1\n1 0 69\r\n\r\nhumidity-to-location map:\n60 56 37\n56 93 442"
	expected := "46"
	actual := Run(input, true)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestSeedsAsRangesInput(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	expected := "324724204"
	actual := Run(input, true)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
