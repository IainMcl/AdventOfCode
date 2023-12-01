package main

import (
	"fmt"

	"github.com/IainMcl/AdventOfCode2023/day1"
	"github.com/IainMcl/AdventOfCode2023/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	input := utils.ReadInput('1')
	out := day1.Run(input)
	fmt.Println(out)
}
