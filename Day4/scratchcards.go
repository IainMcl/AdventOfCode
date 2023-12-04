package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Scratchcard struct {
	CardNumber     int
	WinningNumbers []int
	CardNumbers    []int
}

func Run(input string) string {
	lines := strings.Split(input, "\n")
	// Remove any trailing return \r
	cards := make([]Scratchcard, len(lines))
	points := 0.0
	for i, line := range lines {
		lines[i] = strings.TrimSuffix(line, "\r")
		cards[i] = parseScratchcard(line)
		winningNumbers := getWinningNumbers(cards[i])
		if len(winningNumbers) > 0 {
			points += math.Pow(2, float64(len(winningNumbers)-1))
		}
	}
	return fmt.Sprintf("%.0f", points)
}

type CardDetails struct {
	CardNumber      int
	CardCoppies     int
	NumberOfWinners int
}

func WinningMoreCards(input string) string {
	lines := strings.Split(input, "\n")
	// Remove any trailing return \r
	cards := make([]Scratchcard, len(lines))
	cardDetails := make([]CardDetails, len(lines))
	// Set a value in cardCoppies for each line number
	for i, line := range lines {
		lines[i] = strings.TrimSuffix(line, "\r")
		cards[i] = parseScratchcard(line)
		winningNumbers := getWinningNumbers(cards[i])
		cardDetails[i] = CardDetails{
			CardNumber:      cards[i].CardNumber,
			CardCoppies:     1,
			NumberOfWinners: len(winningNumbers),
		}
	}

	totalCards := 0
	for i, cardDetail := range cardDetails {
		if cardDetail.NumberOfWinners > 0 {
			maxIndex := i + 1 + cardDetail.NumberOfWinners
			if maxIndex > len(cardDetails) {
				maxIndex = len(cardDetails)
			}
			for j := i + 1; j < maxIndex; j++ {
				cardDetails[j].CardCoppies += cardDetail.CardCoppies
			}
		}
	}

	for _, cardDetail := range cardDetails {
		totalCards += cardDetail.CardCoppies
	}

	return fmt.Sprintf("%d", totalCards)
}

// Parse a scratchcard from a string
// Example input:
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Gives:
//
//	Scratchcard{
//	  CardNumber: 1,
//	  WinningNumbers: [41 48 83 86 17],
//	  CardNumbers: [83 86  6 31 17  9 48 53],
//	}
func parseScratchcard(input string) Scratchcard {
	idNumbersSplit := strings.Split(input, ":")
	cardNumberSection := strings.Split(idNumbersSplit[0], " ")
	cardNumber := cardNumberSection[len(cardNumberSection)-1]
	cardNumbersSplit := strings.Split(idNumbersSplit[1], "|")
	winningNumbers := strings.Split(strings.TrimSpace(cardNumbersSplit[0]), " ")
	cardNumbers := strings.Split(strings.TrimSpace(cardNumbersSplit[1]), " ")

	cardNumberInt, err := strconv.Atoi(cardNumber)
	if err != nil {
		panic(err)
	}
	winningNumbersInt := []int{}
	for _, winningNumber := range winningNumbers {
		if winningNumber == "" {
			continue
		}
		winningNumberInt, err := strconv.Atoi(winningNumber)
		if err != nil {
			panic(err)
		}
		winningNumbersInt = append(winningNumbersInt, winningNumberInt)
	}
	CardNumbersInt := []int{}
	for _, cardNumber := range cardNumbers {
		if cardNumber == "" {
			continue
		}
		cardNumberInt, err := strconv.Atoi(cardNumber)
		if err != nil {
			panic(err)
		}
		CardNumbersInt = append(CardNumbersInt, cardNumberInt)
	}
	if len(winningNumbers) > 0 && len(cardNumbers) > 0 {
		return Scratchcard{
			CardNumber:     cardNumberInt,
			WinningNumbers: winningNumbersInt,
			CardNumbers:    CardNumbersInt,
		}
	}
	return Scratchcard{}
}

func getWinningNumbers(card Scratchcard) []int {
	winningNumbers := []int{}
	for _, winningNumber := range card.WinningNumbers {
		for _, cardNumber := range card.CardNumbers {
			if winningNumber == cardNumber {
				winningNumbers = append(winningNumbers, winningNumber)
			}
		}
	}
	return winningNumbers
}
