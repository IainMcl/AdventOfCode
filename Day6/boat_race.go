package day6

import (
	"math"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func Run(input string, kerning bool) string {
	var races []Race
	if kerning {
		races = parseInputWithKerning(input)
	} else {
		races = parseInput(input)
	}
	waysToWin := make([]int, len(races))

	for i, race := range races {
		// fmt.Printf("Time: %d, Distance %d\n", race.Time, race.Distance)
		// All of the ways to win will be clumped together around a most efficient solution
		// Since a winner is not greater than or less than a value it makes sense to treat this li
		// like the two crystal balls problem stepping through the array in sqrt(n) steps
		// When a winner is found step backwards a max of sqrt(n) steps to find the start of the clump
		// Then continue sqrt(n) stepping until there is no longer a winner and step back again
		winningMin := -1
		winningMax := -1
		winnerFound := false
		step := math.Floor(math.Sqrt(float64(race.Time)))
		for i := 0; i <= race.Time; i = getNextStep(i, int(step), race.Time) {
			winner := isWinner(i, race.Time, race.Distance)
			if winner && !winnerFound {
				winningMin = i
				winnerFound = true
				// Step backwards until there is no longer a winner
				for j := i - 1; j >= i-int(step); j-- {
					if !isWinner(j, race.Time, race.Distance) {
						winningMin = j + 1
						break
					}
				}
			} else if !winner && winnerFound {
				// Step backwards until there is a winner
				for j := i - 1; j >= i-int(step); j-- {
					if isWinner(j, race.Time, race.Distance) {
						winningMax = j
						break
					}
				}
				break
			}
		}
		if winningMin == -1 {
			winningMin = 0
		}
		if winningMax == -1 {
			winningMax = race.Time
		}
		if winningMin == 0 && winningMax == 0 {
			waysToWin[i] = 0
		} else {
			waysToWin[i] = winningMax - winningMin + 1
		}
	}

	productOfWaysToWin := 1
	for _, ways := range waysToWin {
		productOfWaysToWin *= ways
	}
	return strconv.Itoa(productOfWaysToWin)
}

func getNextStep(i, step, max int) int {
	if i+int(step) > max {
		return max
	}
	return i + step
}

func isWinner(chargeTime, raceTime, winnintDistance int) bool {
	movingTime := raceTime - chargeTime
	dist := movingTime * chargeTime
	return dist > winnintDistance
}

func parseInputWithKerning(input string) []Race {
	lines := strings.Split(input, "\r\n")
	timeLine := lines[0]
	distanceLine := lines[1]

	// Split time on : then remove all white space and read as int
	t := strings.Split(timeLine, ":")[1]
	t = strings.ReplaceAll(t, " ", "")
	time, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}

	// Split distance on : then remove all white space and read as int
	d := strings.Split(distanceLine, ":")[1]
	d = strings.ReplaceAll(d, " ", "")
	distance, err := strconv.Atoi(d)
	if err != nil {
		panic(err)
	}

	race := Race{
		Time:     time,
		Distance: distance,
	}
	return []Race{race}
}

func parseInput(input string) []Race {
	lines := strings.Split(input, "\r\n")
	timeLine := lines[0]
	distanceLine := lines[1]
	var times []int
	var distances []int

	t := strings.Split(timeLine, " ")
	for _, s := range t {
		if s == "" {
			continue
		}
		if num, err := strconv.Atoi(s); err != nil {
			continue
		} else {
			times = append(times, num)
		}
	}

	d := strings.Split(distanceLine, " ")
	for _, s := range d {
		if s == "" {
			continue
		}
		if num, err := strconv.Atoi(s); err != nil {
			continue
		} else {
			distances = append(distances, num)
		}
	}

	races := make([]Race, len(times))
	for i := 0; i < len(times); i++ {
		races[i].Distance = distances[i]
		races[i].Time = times[i]
	}
	return races
}
