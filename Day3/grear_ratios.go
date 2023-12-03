package day3

import (
	"strconv"
	"strings"
)

func Run(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\r", "")
	}
	sum := 0
	for i, line := range lines {
		// Check in line
		numBuffer := make([]byte, 0)
		for j, char := range line {
			if isNumber(byte(char)) {
				numBuffer = append(numBuffer, byte(char))
			}
			if len(numBuffer) > 0 && (j == len(line)-1 || !isNumber(line[j+1])) {
				if hasSurroundingSymbol(len(numBuffer), i, j, &lines) {
					num, err := strconv.Atoi(string(numBuffer))
					if err != nil {
						panic(err)
					}
					sum += num
				}
				numBuffer = make([]byte, 0)
			}
		}
	}
	return strconv.Itoa(sum)
}

func RunGearRatios(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\r", "")
	}
	sum := 0
	for i, line := range lines {
		// Check in line
		// numBuffer := make([]byte, 0)
		for j, char := range line {
			if isAsterix(byte(char)) {
				nums := getSurroundingNumbers(i, j, &lines)
				// If between exactly two parts then are geers so get product
				if len(nums) == 2 {
					product := 1
					for _, num := range nums {
						product *= num
					}
					sum += product
				}
			}
		}
	}
	return strconv.Itoa(sum)
}

func isAsterix(char byte) bool {
	return char == '*'
}

// Checks around a given position for any number characters. If there are number
// characters, they are returned as a slice of ints.
func getSurroundingNumbers(i int, j int, grid *[]string) []int {
	nums := make([]int, 0)
	hStart := j - 1
	hEnd := j + 1
	vStart := i - 1
	vEnd := i + 1
	if hStart < 0 {
		hStart = 0
	}
	if hEnd >= len((*grid)[0]) {
		hEnd = len((*grid)[0]) - 1
	}
	if vStart < 0 {
		vStart = 0
	}
	if vEnd >= len(*grid) {
		vEnd = len(*grid) - 1
	}

	for v := vStart; v <= vEnd; v++ {
		for h := hStart; h <= hEnd; h++ {
			if isNumber((*grid)[v][h]) {
				// Find the range of the number left and right of the current position
				numStart := h
				numEnd := h
				for numStart > 0 && isNumber((*grid)[v][numStart-1]) {
					numStart--
				}
				for numEnd < len((*grid)[0])-1 && isNumber((*grid)[v][numEnd+1]) {
					numEnd++
				}

				num, err := strconv.Atoi(string((*grid)[v][numStart : numEnd+1]))
				if err != nil {
					panic(err)
				}
				nums = append(nums, num)
				h = numEnd
			}
		}
	}
	return nums
}

func hasSurroundingSymbol(numLength int, i int, j int, grid *[]string) bool {
	hStart := j - numLength + 1
	hEnd := j + 1
	vStart := i - 1
	vEnd := i + 1
	if hStart > 0 {
		hStart--
	}
	if hEnd >= len((*grid)[0]) {
		hEnd = len((*grid)[0]) - 1
	}
	if vStart < 0 {
		vStart = 0
	}
	if vEnd >= len(*grid) {
		vEnd = len(*grid) - 1
	}

	for v := vStart; v <= vEnd; v++ {
		for h := hStart; h <= hEnd; h++ {
			if isSymbol((*grid)[v][h]) {
				return true
			}
		}
	}
	return false
}

func isSymbol(char byte) bool {
	// Return true for any symbol that is not a '.' or a number
	return char != '.' && !isNumber(char)
}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}
