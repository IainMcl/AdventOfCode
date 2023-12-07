package day7

import (
	"strconv"
	"strings"
)

func Run(input string, withJokers bool) string {
	hands := parseInput(input, withJokers)
	rankHands(&hands)
	winnings := calcWinnings(&hands)
	return strconv.Itoa(winnings)
}

type handRank int

const (
	HighCard handRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type CamelCard struct {
	Hand     []rune
	Bid      int
	HandRank handRank
}

var cardRankings = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}
var cardRankingsWithJokers = map[rune]int{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func parseInput(input string, withJokers bool) []CamelCard {
	lines := strings.Split(input, "\n")
	hands := make([]CamelCard, len(lines))
	for i, line := range lines {
		line = strings.Replace(line, "\r", "", -1)
		parts := strings.Split(line, " ")
		var err error
		hands[i].Bid, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		hands[i].Hand = []rune(parts[0])

		cardCounts := map[rune]int{}
		for _, card := range hands[i].Hand {
			cardCounts[card]++
		}
		// jokerUsed := false
		if withJokers {
			jokerCount := cardCounts['J']
			cardCounts['J'] = 0
			if jokerCount > 0 {
				// Get the key for the highest value
				maxVal := -1
				var maxKey rune
				for key, val := range cardCounts {
					if val > maxVal && key != 'J' {
						maxVal = val
						maxKey = key
					}
				}
				cardCounts[maxKey] += jokerCount
				// jokerUsed = true
			}
		}

		for _, count := range cardCounts {
			switch count {
			case 5:
				hands[i].HandRank = FiveOfAKind
			case 4:
				hands[i].HandRank = FourOfAKind
			case 3:
				hands[i].HandRank = ThreeOfAKind
				for _, count2 := range cardCounts {
					if count2 == 2 {
						hands[i].HandRank = FullHouse
						break
					}
				}
			case 2:
				hands[i].HandRank = OnePair
				found1 := false
				for _, count2 := range cardCounts {
					if count2 == 2 && found1 {
						hands[i].HandRank = TwoPair
						break
					}
					if count2 == 2 {
						found1 = true
					}
					if count2 == 3 {
						hands[i].HandRank = FullHouse
						break
					}
				}
			}
		}
		// if jokerUsed {
		// 	handRank := ""
		// 	switch hands[i].HandRank {
		// 	case FiveOfAKind:
		// 		handRank = "Five of a kind"
		// 	case FourOfAKind:
		// 		handRank = "Four of a kind"
		// 	case ThreeOfAKind:
		// 		handRank = "Three of a kind"
		// 	case TwoPair:
		// 		handRank = "Two pair"
		// 	case OnePair:
		// 		handRank = "One pair"
		// 	case FullHouse:
		// 		handRank = "Full house"
		// 	case HighCard:
		// 		handRank = "High card"
		// 	}

		// 	hand := ""
		// 	for _, card := range hands[i].Hand {
		// 		hand += string(card)
		// 	}

		// 	fmt.Println(i, " Hand with joker: ", hand, " ", handRank)
		//}
	}
	return hands
}

func rankHands(cards *[]CamelCard) {
	// Quick sort the list using the isHigher function
	interfaceCards := make([]interface{}, len(*cards))
	for i, card := range *cards {
		interfaceCards[i] = card
	}
	quicksort(interfaceCards, 0, len(*cards)-1, isHigher)
	for i, card := range interfaceCards {
		(*cards)[i] = card.(CamelCard)
	}
}

func quicksort(arr []interface{}, low, high int, isHigher func(a, b interface{}) bool) {
	if low < high {
		part := partition(arr, low, high, isHigher)

		quicksort(arr, low, part-1, isHigher)
		quicksort(arr, part+1, high, isHigher)
	}
}

func partition(arr []interface{}, low, high int, isHigher func(a, b interface{}) bool) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if isHigher(arr[j], pivot) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func isHigher(card1, card2 interface{}) bool {
	if card1.(CamelCard).HandRank > card2.(CamelCard).HandRank {
		return false
	}
	if card1.(CamelCard).HandRank < card2.(CamelCard).HandRank {
		return true
	}
	if card1.(CamelCard).HandRank == card2.(CamelCard).HandRank {
		for i, card := range card1.(CamelCard).Hand {
			if cardRankings[card] > cardRankings[card2.(CamelCard).Hand[i]] {
				return false
			}
			if cardRankings[card] < cardRankings[card2.(CamelCard).Hand[i]] {
				return true
			}
		}
	}
	return false
}

func calcWinnings(hands *[]CamelCard) int {
	total := 0
	for i, hand := range *hands {
		total += (i + 1) * hand.Bid
	}
	return total
}
