package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part2() {
	stars, rowKs, colKs := loadInput("input-0.txt", 1000000)
	cnt := len(stars)
	sum := 0
	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			dist := getDist(rowKs, stars[i].Row, stars[j].Row) + getDist(colKs, stars[i].Col, stars[j].Col)
			sum += dist
		}
	}
	fmt.Println("Part-2:", sum)
}

func part1() {
	stars, rowKs, colKs := loadInput("input-0.txt", 2)
	cnt := len(stars)
	sum := 0
	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			dist := getDist(rowKs, stars[i].Row, stars[j].Row) + getDist(colKs, stars[i].Col, stars[j].Col)
			sum += dist
		}
	}
	fmt.Println("Part-1:", sum)
}

func getDist(ks Ints, i1, i2 int) int {
	sum := 0

	if i1 > i2 {
		for i := i2 + 1; i <= i1; i++ {
			sum += ks[i]
		}
	} else {
		for i := i1 + 1; i <= i2; i++ {
			sum += ks[i]
		}
	}
	return sum
}

type Ints []int

type Star struct {
	Row, Col int
}

func loadInput(filename string, expansion int) ([]Star, Ints, Ints) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	stars := make([]Star, 0)

	rowIndex := 0
	colLen := 0
	for scanner.Scan() {
		line := scanner.Text()
		colLen = len(line)

		for colIndex, r := range line {
			if r == '#' {
				stars = append(stars, Star{Row: rowIndex, Col: colIndex})
			}
		}

		rowIndex++
	}

	rowKs := make(Ints, rowIndex)

	for index := range rowKs {
		cnt := 0

		for _, star := range stars {
			if star.Row == index {
				cnt++
			}
		}
		if cnt == 0 {
			rowKs[index] = expansion
		} else {
			rowKs[index] = 1
		}
	}

	colKs := make(Ints, colLen)

	for index := range colKs {
		cnt := 0

		for _, star := range stars {
			if star.Col == index {
				cnt++
			}
		}
		if cnt == 0 {
			colKs[index] = expansion
		} else {
			colKs[index] = 1
		}
	}

	return stars, rowKs, colKs
}
