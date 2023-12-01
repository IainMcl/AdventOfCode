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

// Test day1.RunWithWords
// Input:
// two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen
// Output:
// 281
func TestRunWithWords(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	expected := "281"
	actual := RunWithWords(input)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestInputWithWords(t *testing.T) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	expected := "55358"
	actual := RunWithWords(input)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
