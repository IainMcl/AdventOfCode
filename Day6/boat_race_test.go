package day6

import (
	"os"
	"testing"
)

// Test day6.Run
// Input:
// Time:      7  15   30
// Distance:  9  40  200
// Output:
// 288
func TestRun(t *testing.T) {
	input := "Time:      7  15   30\r\nDistance:  9  40  200"
	expected := "288"
	actual := Run(input, false)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestRunInput(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	expected := "449820"
	actual := Run(input, false)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestKerningRun(t *testing.T) {
	input := "Time:      7  15   30\r\nDistance:  9  40  200"
	expected := "71503"
	actual := Run(input, true)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestKerningRunInput(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	expected := "42250895"
	actual := Run(input, true)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
