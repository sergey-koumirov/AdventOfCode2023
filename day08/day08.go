package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	instructions, nodes := loadInput("input-0.txt")

	// fmt.Println(instructions)
	// fmt.Println(nodes)

	res := []int{}

	for _, n := range nodes {
		if n.Name[2] == 'A' {
			test := findForName(nodes, instructions, n.Name)
			res = append(res, test)
		}
	}

	lcm := res[0]
	for i := 1; i < len(res); i++ {
		lcm = LCM(lcm, res[i])
	}

	fmt.Println("Part-2:", res, lcm)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func findForName(nodes Nodes, instructions, current string) int {
	iLen := len(instructions)

	step := 0

	for {
		node := nodes[current]
		dir := instructions[step%iLen]

		if dir == 'L' && node.Left[2] == 'Z' || dir == 'R' && node.Right[2] == 'Z' {
			break
		}

		if dir == 'L' {
			current = node.Left
		}

		if dir == 'R' {
			current = node.Right
		}

		step++
	}

	return step + 1
}

func part1() {
	instructions, nodes := loadInput("input-0.txt")
	iLen := len(instructions)

	// fmt.Println(instructions)
	// fmt.Println(nodes)

	step := 0

	current := "AAA"

	for {
		node := nodes[current]
		dir := instructions[step%iLen]

		if dir == 'L' && node.Left == "ZZZ" || dir == 'R' && node.Right == "ZZZ" {
			break
		}

		if dir == 'L' {
			current = node.Left
		}

		if dir == 'R' {
			current = node.Right
		}

		step++
	}

	fmt.Println("Part-1:", step+1)
}

type Node struct {
	Name, Left, Right string
}

type Nodes map[string]Node

func loadInput(filename string) (string, Nodes) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	res := make(Nodes)

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		parts1 := strings.Split(line, " = ")
		parts2 := strings.Split(parts1[1][1:9], ", ")

		temp := Node{
			Name:  parts1[0],
			Left:  parts2[0],
			Right: parts2[1],
		}

		res[temp.Name] = temp
	}
	scanner.Scan()

	return instructions, res
}
