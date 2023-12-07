package day7

import (
	"os"
	"testing"
)

// Test day7.Run
// Input:
// 32T3K 765
// T55J5 684
// KK677 28
// KTJJT 220
// QQQJA 483
// output:
// 6440
func TestRun(t *testing.T) {
	input := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"
	expected := "6440"
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
	expected := "246424613"
	actual := Run(input, false)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

// Test day7.Run with jokers
// Input:
// 32T3K 765
// T55J5 684
// KK677 28
// KTJJT 220
// QQQJA 483
// output:
// 6440
func TestRunWithJokers(t *testing.T) {
	input := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"
	expected := "5905"
	actual := Run(input, true)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestInputWithJokers(t *testing.T) {
	data, err := os.ReadFile("input2.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	expected := "248256639"
	actual := Run(input, true)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
