package main

import (
	"fmt"
	"os"

	"github.com/IainMcl/AdventOfCode2023/day5"
)

// func main() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		panic(err)
// 	}
// 	input := utils.ReadInput('1')
// 	out := day1.Run(input)
// 	fmt.Println(out)
// }

func main() {
	data, err := os.ReadFile("Day5/input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	actual := day5.Run(input, true)
	fmt.Println(actual)
}
