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
