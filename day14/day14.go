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
	lines := loadInput("input-0.txt")

	memory := map[string]int{}

	step := 1
	firstStep := -1
	for {
		if step%1000 == 0 {
			fmt.Println(step)
		}
		cycle(lines)
		key := makeKey(lines)

		memStep, exists := memory[key]
		if exists {
			firstStep = memStep
			break
		}
		memory[key] = step
		step++
	}

	// sum := calcLoad(lines)
	fmt.Println()

	rest := (1_000_000_000 - firstStep) % (step - firstStep)

	for i := 1; i <= rest; i++ {
		cycle(lines)
	}
	sum := calcLoad(lines)

	fmt.Println("Part-2:", firstStep, step, sum)
}

func makeKey(lines [][]byte) string {
	res := ""
	for _, line := range lines {
		res += string(line)
	}
	return res
}

func part1() {
	lines := loadInput("input-1.txt")

	tiltNorth(lines)

	sum := calcLoad(lines)
	fmt.Println()

	fmt.Println("Part-1:", sum)
}

func cycle(lines [][]byte) {
	tiltNorth(lines)
	tiltWest(lines)
	tiltSouth(lines)
	tiltEast(lines)
}

func tiltEast(lines [][]byte) {
	lastColIndex := len(lines[0]) - 1
	lastRowIndex := len(lines) - 1

	for colIndex := lastColIndex - 1; colIndex >= 0; colIndex-- {
		for rowIndex := 0; rowIndex <= lastRowIndex; rowIndex++ {
			element := lines[rowIndex][colIndex]
			if element == 'O' {
				for moveIndex := colIndex + 1; moveIndex <= lastColIndex; moveIndex++ {
					if lines[rowIndex][moveIndex] == '.' {
						lines[rowIndex][moveIndex] = 'O'
						lines[rowIndex][moveIndex-1] = '.'
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltWest(lines [][]byte) {
	lastColIndex := len(lines[0]) - 1
	lastRowIndex := len(lines) - 1

	for colIndex := 1; colIndex <= lastColIndex; colIndex++ {
		for rowIndex := 0; rowIndex <= lastRowIndex; rowIndex++ {
			element := lines[rowIndex][colIndex]
			if element == 'O' {
				for moveIndex := colIndex - 1; moveIndex >= 0; moveIndex-- {
					if lines[rowIndex][moveIndex] == '.' {
						lines[rowIndex][moveIndex] = 'O'
						lines[rowIndex][moveIndex+1] = '.'
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltSouth(lines [][]byte) {
	lastIndex := len(lines) - 1

	for rowIndex := lastIndex - 1; rowIndex >= 0; rowIndex-- {
		for colIndex, element := range lines[rowIndex] {
			if element == 'O' {
				for index := rowIndex + 1; index <= lastIndex; index++ {
					if lines[index][colIndex] == '.' {
						lines[index][colIndex] = 'O'
						lines[index-1][colIndex] = '.'
					} else {
						break
					}
				}
			}
		}
	}
}

func tiltNorth(lines [][]byte) {
	for rowIndex := 1; rowIndex < len(lines); rowIndex++ {
		for colIndex, element := range lines[rowIndex] {
			if element == 'O' {
				for index := rowIndex - 1; index >= 0; index-- {
					if lines[index][colIndex] == '.' {
						lines[index][colIndex] = 'O'
						lines[index+1][colIndex] = '.'
					} else {
						break
					}
				}
			}
		}
	}
}

func calcLoad(lines [][]byte) int {
	res := 0
	for rowIndex, row := range lines {
		for _, element := range row {
			if element == 'O' {
				res += (len(lines) - rowIndex)
			}
		}
	}
	return res
}

func print(lines [][]byte) {
	for _, line := range lines {
		fmt.Println(string(line))
	}
}

func loadInput(fn string) [][]byte {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := [][]byte{}

	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		res = append(res, bytes)
	}

	return res
}
