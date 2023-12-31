package day3

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func Main() {
	f, _ := os.Open("./input_red.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)
	scheme := [][]rune{}
	sum := 0
	for sc.Scan() {
		scheme = append(scheme, []rune(sc.Text()))
	}

	for i := 0; i < len(scheme); i++ {
		for j := 0; j < len(scheme[i]); j++ {
			if scheme[i][j] != '.' && !unicode.IsDigit(scheme[i][j]) {
				sum += getNums(scheme, i, j)
			}
		}

	}

	println(sum)
}

func getNums(sh [][]rune, i, j int) int {

	sum := 0
	if i != 0 {
		for _, v := range getHoriNums(sh, i-1, j) {
			sum += v
		}
	}
	for _, v := range getHoriNums(sh, i, j) {
		sum += v
	}
	if i+1 < len(sh) {
		for _, v := range getHoriNums(sh, i+1, j) {
			sum += v
		}
	}
	return sum
}

func getHoriNums(sh [][]rune, i int, j int) []int {
	ln, rn := "", ""

	for k := j + 1; k < len(sh[i]); k++ {
		if !unicode.IsDigit(sh[i][k]) {
			break
		}
		rn += string(sh[i][k])
	}
	for k := j - 1; k >= 0; k-- {
		if !unicode.IsDigit(sh[i][k]) {
			break
		}
		ln = string(sh[i][k]) + ln
	}

	if unicode.IsDigit(sh[i][j]) {
		n, _ := strconv.Atoi(ln + string(sh[i][j]) + rn)
		return []int{n}
	}
	num := []int{}
	lnn, _ := strconv.Atoi(ln)
	rnn, _ := strconv.Atoi(rn)
	if lnn != 0 {
		num = append(num, lnn)
	}
	if rnn != 0 {
		num = append(num, rnn)
	}
	return num
}
