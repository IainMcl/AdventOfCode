package day1

import (
	"fmt"
	"strings"
)

func Run(input string) string {
	sum := 0

	// Split the input into lines
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var first, last rune

		for i := 0; i < len(line); i++ {
			if isNum(line[i]) {
				first = rune(line[i]) - '0'
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if isNum(line[i]) {
				last = rune(line[i]) - '0'
				break
			}
		}

		// Join the numbers together in the form of a string then add to sum
		// '1' + '2' = '12'
		sum += int(first*10 + last)
	}
	return fmt.Sprintf("%d", sum)
}

func isNum(char byte) bool {
	return char >= '0' && char <= '9'
}

func RunWithWords(input string) string {
	sum := 0

	// Split the input into lines
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var first, last rune

		for i := 0; i < len(line); i++ {
			if isNum(line[i]) {
				first = rune(line[i]) - '0'
				break
			}
			if isWordNumber(line[i:], false) {
				num, _ := getWordNumber(line[i:], false)
				first = num
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if isNum(line[i]) {
				last = rune(line[i]) - '0'
				break
			}
			if isWordNumber(line[:i+1], true) {
				num, _ := getWordNumber(line[:i+1], true)
				last = num
				break
			}
		}

		// Join the numbers together in the form of a string then add to sum
		// '1' + '2' = '12'
		sum += int(first*10 + last)
	}
	return fmt.Sprintf("%d", sum)
}

func isWordNumber(remaining string, rev bool) bool {
	wordNumbers := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for _, word := range wordNumbers {
		if rev {
			if strings.HasSuffix(remaining, word) {
				return true
			}
		} else {
			if strings.HasPrefix(remaining, word) {
				return true
			}
		}
	}
	return false
}

func getWordNumber(remaining string, rev bool) (rune, error) {
	wordNumbers := map[string]rune{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for word, num := range wordNumbers {
		if rev {
			if strings.HasSuffix(remaining, word) {
				return num, nil
			}
		} else {

			if strings.HasPrefix(remaining, word) {
				return num, nil
			}
		}
	}
	return 0, fmt.Errorf("no word number found")
}
