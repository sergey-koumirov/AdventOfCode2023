package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

type Lens struct {
	Label string
	FL    int
}

func part2() {
	lines := loadInput("input-0.txt")

	boxes := [256][]Lens{}

	for _, line := range lines {
		label, operation, num := parse(line)

		index := hash(label)

		if operation == '-' {
			for i := 0; i < len(boxes[index]); i++ {
				if boxes[index][i].Label == label {
					boxes[index] = append(boxes[index][:i], boxes[index][i+1:]...)
					break
				}
			}
		}

		if operation == '=' {
			found := false
			for i := 0; i < len(boxes[index]); i++ {
				if boxes[index][i].Label == label {
					found = true
					boxes[index][i].Label = label
					boxes[index][i].FL = num
				}
			}
			if !found {
				boxes[index] = append(boxes[index], Lens{Label: label, FL: num})
			}
		}

		print(boxes, line)
	}

	sum := 0
	for boxIndex, box := range boxes {
		for lensIndex, lens := range box {
			sum += (boxIndex + 1) * (lensIndex + 1) * lens.FL
		}
	}

	fmt.Println("Part-2:", sum)
}

func print(boxes [256][]Lens, cmd string) {
	fmt.Println("After", cmd)
	for i, box := range boxes {

		if len(box) > 0 {
			fmt.Print("Box", i, " ")
			for _, lens := range box {
				fmt.Print(lens, " ")
			}
			fmt.Println()
		}
	}
	fmt.Println()
}

func parse(line string) (string, rune, int) {
	if strings.IndexRune(line, '=') > -1 {
		parts := strings.Split(line, "=")
		num, _ := strconv.Atoi(parts[1])
		return parts[0], '=', num
	}
	parts := strings.Split(line, "-")
	return parts[0], '-', -1
}

func part1() {
	lines := loadInput("input-0.txt")

	sum := 0

	for _, line := range lines {
		h := hash(line)
		sum += h
	}

	fmt.Println("Part-1:", sum)
}

func hash(line string) int {
	current := 0
	for _, code := range line {
		current += int(code)
		current *= 17
		current = current % 256
	}

	return current
}

func loadInput(fn string) []string {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	res := strings.Split(line, ",")

	return res
}
