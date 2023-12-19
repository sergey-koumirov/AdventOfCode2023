package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	// part2()
}

func part2() {
	field := loadInput("input-0.txt")

	lastRowIndex := len(field) - 1
	lastColIndex := len(field[0]) - 1

	max := -1

	for indexRow := 0; indexRow <= lastRowIndex; indexRow++ {
		init := []Coords{{Row: indexRow, Col: 0, From: 'R'}}
		field[indexRow][0].UsedDirs['R'] = 1
		emulateLight(field, init)

		test := getEnergy(field)
		if test > max || max == -1 {
			max = test
		}

		reset(field)
	}

	for indexRow := 0; indexRow <= lastRowIndex; indexRow++ {
		init := []Coords{{Row: indexRow, Col: lastColIndex, From: 'L'}}
		field[indexRow][lastColIndex].UsedDirs['L'] = 1
		emulateLight(field, init)

		test := getEnergy(field)
		if test > max || max == -1 {
			max = test
		}

		reset(field)
	}

	for indexCol := 0; indexCol <= lastColIndex; indexCol++ {
		init := []Coords{{Row: 0, Col: indexCol, From: 'D'}}
		field[0][indexCol].UsedDirs['D'] = 1
		emulateLight(field, init)

		test := getEnergy(field)
		if test > max || max == -1 {
			max = test
		}

		reset(field)
	}

	for indexCol := 0; indexCol <= lastColIndex; indexCol++ {
		init := []Coords{{Row: lastRowIndex, Col: indexCol, From: 'U'}}
		field[lastRowIndex][indexCol].UsedDirs['U'] = 1
		emulateLight(field, init)

		test := getEnergy(field)
		if test > max || max == -1 {
			max = test
		}

		reset(field)
	}

	fmt.Println("Part-2:", max)
}

func part1() {
	field := loadInput("input-2.txt")

	init := []Coords{{Row: 0, Col: 0, From: 'R'}}
	field[0][0].UsedDirs['R'] = 1

	emulateLight(field, init)

	print(field)

	fmt.Println("Part-1:", getEnergy(field))
}

func reset(field Field) {
	for ir, row := range field {
		for ic, _ := range row {
			field[ir][ic].UsedDirs = map[rune]int{}
		}
	}
}

func getEnergy(field Field) int {
	sum := 0

	for _, row := range field {
		for _, cell := range row {
			if len(cell.UsedDirs) > 0 {
				sum += 1
			}
		}
	}
	return sum
}

func emulateLight(field Field, queue []Coords) {

	maxRow := len(field) - 1
	maxCol := len(field[0]) - 1

	for len(queue) > 0 {
		coords := queue[0]
		queue = queue[1:]

		current := field[coords.Row][coords.Col]

		isDot := current.Symbol == '.'
		isPassHor := current.Symbol == '-' && (coords.From == 'R' || coords.From == 'L')
		isPassVer := current.Symbol == '|' && (coords.From == 'U' || coords.From == 'D')
		if isDot || isPassHor || isPassVer {
			dr, dc := fromToNextDot(coords.From)

			if inside(field, coords, dr, dc, maxRow, maxCol) {
				queue = apply(field, queue, coords.Row+dr, coords.Col+dc, coords.From)
			}
		}

		if current.Symbol == '/' {
			dr, dc, from := fromToNextUp(coords.From)

			if inside(field, coords, dr, dc, maxRow, maxCol) {
				queue = apply(field, queue, coords.Row+dr, coords.Col+dc, from)
			}
		}

		if current.Symbol == '\\' {
			dr, dc, from := fromToNextDown(coords.From)
			if inside(field, coords, dr, dc, maxRow, maxCol) {
				queue = apply(field, queue, coords.Row+dr, coords.Col+dc, from)
			}
		}

		if current.Symbol == '-' && (coords.From == 'U' || coords.From == 'D') {
			if inside(field, coords, 0, -1, maxRow, maxCol) {
				queue = apply(field, queue, coords.Row, coords.Col-1, 'L')
			}
			if inside(field, coords, 0, 1, maxRow, maxCol) {
				queue = apply(field, queue, coords.Row, coords.Col+1, 'R')

			}
		}

		if current.Symbol == '|' && (coords.From == 'R' || coords.From == 'L') {
			if inside(field, coords, -1, 0, maxRow, maxCol) {
				queue = apply(field, queue, coords.Row-1, coords.Col, 'U')
			}
			if inside(field, coords, 1, 0, maxRow, maxCol) {
				queue = apply(field, queue, coords.Row+1, coords.Col, 'D')
			}
		}

		print(field)
		fmt.Scanln()
	}
}

func apply(f Field, q []Coords, row, col int, from rune) []Coords {
	f[row][col].UsedDirs[from] = 1
	return append(q, Coords{Row: row, Col: col, From: from})
}

// \\
func fromToNextDown(r rune) (int, int, rune) {
	if r == 'R' {
		return 1, 0, 'D'
	}
	if r == 'L' {
		return -1, 0, 'U'
	}
	if r == 'U' {
		return 0, -1, 'L'
	}
	if r == 'D' {
		return 0, 1, 'R'
	}
	panic("fromToNextDown !!!")
}

// /
func fromToNextUp(r rune) (int, int, rune) {
	if r == 'R' {
		return -1, 0, 'U'
	}
	if r == 'L' {
		return 1, 0, 'D'
	}
	if r == 'U' {
		return 0, 1, 'R'
	}
	if r == 'D' {
		return 0, -1, 'L'
	}
	panic("fromToNextUp !!!")
}

func inside(f Field, coords Coords, dr, dc, maxRow, maxCol int) bool {
	inField := coords.Row+dr >= 0 && coords.Row+dr <= maxRow && coords.Col+dc >= 0 && coords.Col+dc <= maxCol

	if inField {
		dir := deltaToDir(dr, dc)
		_, ex := f[coords.Row+dr][coords.Col+dc].UsedDirs[dir]
		return !ex
	}

	return false
}

func deltaToDir(dr, dc int) rune {
	if dr == 0 && dc == 1 {
		return 'R'
	}
	if dr == 0 && dc == -1 {
		return 'L'
	}
	if dr == 1 && dc == 0 {
		return 'D'
	}
	if dr == -1 && dc == 0 {
		return 'U'
	}
	panic("deltaToDir !!!")
}

func fromToNextDot(r rune) (int, int) {
	if r == 'R' {
		return 0, 1
	}
	if r == 'L' {
		return 0, -1
	}
	if r == 'U' {
		return -1, 0
	}
	if r == 'D' {
		return 1, 0
	}
	panic("fromToNextDot !!!")
}

func print(f Field) {

	for _, row := range f {
		for _, cell := range row {
			fmt.Print(string(cell.Symbol))
		}
		fmt.Print("       ")
		for _, cell := range row {
			if len(cell.UsedDirs) > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type Cell struct {
	Symbol   rune
	UsedDirs map[rune]int
}

type Coords struct {
	Row, Col int
	From     rune
}

type Field [][]Cell

func loadInput(fn string) Field {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := Field{}

	for scanner.Scan() {
		line := scanner.Text()

		temp := []Cell{}

		for _, r := range line {
			temp = append(temp, Cell{Symbol: r, UsedDirs: map[rune]int{}})
		}
		res = append(res, temp)
	}

	return res
}
