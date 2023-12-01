package day1

import (
	"os"
	"testing"
)

// Test day1.Run
// Input:
// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
// Output:
// 142
func TestRun(t *testing.T) {
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	expected := "142"
	actual := Run(input)
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
	expected := "56042"
	actual := Run(input)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
