package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part2() {
	data := loadInput("input-0.txt")

	sum := 0
	for i := range data {
		levels := getLevels(data[i])

		for _, level := range levels {
			fmt.Println(level)
		}

		test := 0
		for j := len(levels) - 1; j > 0; j-- {
			test = levels[j-1][0] - test
		}

		sum += test

		fmt.Println("=", test)
		fmt.Println()
	}

	fmt.Println("Part-2:", sum)
}

func part1() {
	data := loadInput("input-0.txt")

	sum := 0
	for i := range data {
		levels := getLevels(data[i])
		test := 0
		for _, level := range levels {
			fmt.Println(level)
			test += level[len(level)-1]
		}

		sum += test

		fmt.Println("=", test)
		fmt.Println()
	}

	fmt.Println("Part-1:", sum)
}

func getLevels(first Ints) []Ints {
	levels := make([]Ints, 0)
	levels = append(levels, first)

	for {
		allZero := true
		next := make([]int, 0)

		lastLevelIndex := len(levels) - 1
		lastLevelElmIndex := len(levels[lastLevelIndex]) - 1

		for index := 1; index <= lastLevelElmIndex; index++ {
			diff := levels[lastLevelIndex][index] - levels[lastLevelIndex][index-1]
			next = append(next, diff)
			if diff != 0 {
				allZero = false
			}
		}
		levels = append(levels, next)

		if allZero {
			break
		}
	}

	return levels
}

type Ints []int

func loadInput(filename string) []Ints {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	res := make([]Ints, 0)

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, parseInts(line))
	}
	scanner.Scan()

	return res
}

func parseInts(line string) []int {
	parts := strings.Split(line, " ")

	res := []int{}
	for _, part := range parts {
		if part != "" {
			n, _ := strconv.Atoi(part)
			res = append(res, n)
		}
	}
	return res
}
