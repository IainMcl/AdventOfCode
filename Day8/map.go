package day8

import (
	"strconv"
	"strings"
)

type Node struct {
	Name  string
	Left  string
	Right string
}

var nodesMap map[string]Node = make(map[string]Node)

func Run(input string) string {
	inst := parseInput(input)
	start := nodesMap["AAA"]
	end := nodesMap["ZZZ"]
	current := start

	step := 0
	for current != end {
		if inst[step%len(inst)] == 'L' {
			current = nodesMap[current.Left]
		} else {
			current = nodesMap[current.Right]
		}
		step++
	}
	return strconv.Itoa(step)
}

func parseInput(input string) string {
	input = strings.ReplaceAll(input, "\r", "")
	in := strings.Split(input, "\n\n")
	instructions := strings.TrimSuffix(in[0], "\n")
	instructions = strings.TrimSuffix(instructions, "\r")

	// Parse nodes <name> = (<left>, <right>)
	nodes := strings.Split(in[1], "\n")
	for _, node := range nodes {
		name, n := parseNode(node)
		nodesMap[name] = n
	}
	return instructions
}

func parseNode(node string) (string, Node) {
	node = strings.TrimSuffix(node, "\r")
	node = strings.ReplaceAll(node, " ", "")
	node = strings.ReplaceAll(node, "(", "")
	node = strings.ReplaceAll(node, ")", "")
	parts := strings.Split(node, "=")
	name := parts[0]
	lr := strings.Split(parts[1], ",")
	left := lr[0]
	right := lr[1]
	return name, Node{
		Name:  name,
		Left:  left,
		Right: right,
	}
}
