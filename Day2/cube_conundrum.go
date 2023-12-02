package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Id      int
	Results []Pull
}

type Pull struct {
	Balls []Balls
}

type Balls struct {
	Number int
	Color  string
}

var MaxNums = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Run(input string) string {
	gameStrings := strings.Split(input, "\n")
	games := make([]Game, len(gameStrings))
	for i, game := range gameStrings {
		parseGame(&games[i], game)
	}

	// Find the max number of each color from each game
	possibleGameIds := make([]int, 0)
	for _, game := range games {
		exceeded := false
		for _, pull := range game.Results {
			for _, ball := range pull.Balls {
				if MaxNums[ball.Color] < ball.Number {
					exceeded = true
					break
				}
			}
			if exceeded {
				break
			}
		}
		if !exceeded {
			possibleGameIds = append(possibleGameIds, game.Id)
		}
	}

	// Sum the possible game ids
	sum := 0
	for _, id := range possibleGameIds {
		sum += id
	}

	return fmt.Sprintf("%d", sum)
}

func Fewest(input string) string {
	gameStrings := strings.Split(input, "\n")
	games := make([]Game, len(gameStrings))
	for i, game := range gameStrings {
		parseGame(&games[i], game)
	}

	// Find the max number of each color from each game
	gamePowers := make([]int, 0)
	for _, game := range games {
		colorMaxes := make(map[string]int)
		for _, pull := range game.Results {
			for _, ball := range pull.Balls {
				if colorMaxes[ball.Color] < ball.Number {
					colorMaxes[ball.Color] = ball.Number
				}
			}
		}
		power := 1
		for _, max := range colorMaxes {
			power *= max
		}
		gamePowers = append(gamePowers, power)
	}

	// Sum the possible game ids
	sum := 0
	for _, id := range gamePowers {
		sum += id
	}

	return fmt.Sprintf("%d", sum)
}

func parseGame(game *Game, line string) {
	// Get last char before ':'
	split := strings.Split(line, ":")
	gameId := strings.Split(split[0], " ")[1]
	var err error
	game.Id, err = strconv.Atoi(gameId)
	if err != nil {
		panic(err)
	}
	results := split[1]
	pulls := strings.Split(results, ";")
	game.Results = make([]Pull, len(pulls))
	for i, pull := range pulls {
		parsePull(&game.Results[i], pull)
	}
}

func parsePull(pull *Pull, line string) {
	balls := strings.Split(line, ",")
	pull.Balls = make([]Balls, len(balls))
	for i, ball := range balls {
		parseBall(&pull.Balls[i], ball)
	}
}

func parseBall(ball *Balls, line string) {
	split := strings.Split(line, " ")
	// Remove slice indexex that only contain whitespace
	noWhiteSpace := make([]string, 0)
	for _, s := range split {
		if s != "" {
			noWhiteSpace = append(noWhiteSpace, s)
		}
	}
	// Convert noWhiteSpace[0] to int
	var err error
	ball.Number, err = strconv.Atoi(noWhiteSpace[0])
	if err != nil {
		panic(err)
	}
	//
	ball.Color = strings.Replace(noWhiteSpace[1], "\r", "", -1)
}
