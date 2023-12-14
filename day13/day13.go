package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// part1()
	part2()
}

func part2() {
	fields := loadInput("input-0.txt")

	sum := 0
	for _, field := range fields {

		// print(field)

		col, row := fixSmudge(field)

		if col > -1 {
			sum += col
		}
		if row > -1 {
			sum += 100 * row
		}

		fmt.Println("Col:", col, "Row:", row)

		// print(field)
	}

	fmt.Println("Part-2:", sum)
}

func fixSmudge(field Field) (int, int) {
	oldCol := checkVer(field, -1)
	oldRow := checkHor(field, -1)

	// fmt.Println("Col:", oldCol, "Row:", oldRow)

	for rowIndex, line := range field {
		for colIndex := range line {
			if field[rowIndex][colIndex] == 0 {
				field[rowIndex][colIndex] = 1
			} else {
				field[rowIndex][colIndex] = 0
			}

			col := checkVer(field, oldCol-1)
			row := checkHor(field, oldRow-1)

			// if rowIndex == 0 && colIndex == 4 {
			// 	fmt.Println("    ", col, row)
			// 	print(field)
			// }

			if (col > -1 || row > -1) && (oldCol != col || oldRow != row) {

				if oldCol == col {
					col = -1
				}

				if oldRow == row {
					row = -1
				}

				return col, row
			}

			if field[rowIndex][colIndex] == 0 {
				field[rowIndex][colIndex] = 1
			} else {
				field[rowIndex][colIndex] = 0
			}
		}
	}
	return -1, -1
}

func part1() {
	fields := loadInput("input-3.txt")

	sum := 0
	for _, field := range fields {
		col := checkVer(field, -1)
		row := checkHor(field, -1)

		if col > -1 {
			sum += col
		}
		if row > -1 {
			sum += 100 * row
		}

		fmt.Println(col, row)
		print(field)
	}

	fmt.Println("Part-1:", sum)
	fmt.Println()
	fmt.Println()
}

func checkHor(field Field, forbiddenRow int) int {
	res := -1
	height := len(field)
	width := len(field[0])

	for row := 0; row < height-1; row++ {

		checkHeight := row + 1
		if checkHeight > height-row-1 {
			checkHeight = height - row - 1
		}

		sym := true
		for colIndex := 0; colIndex < width; colIndex++ {
			if !symmetricH(field, row, colIndex, checkHeight) {
				sym = false
			}
		}

		if sym && (forbiddenRow == -1 || forbiddenRow != row) {
			res = row
			// fmt.Println("H-Sym at ", row)
		}
	}

	if res > -1 {
		res += 1
	}

	return res
}

func symmetricH(field Field, row, colIndex int, checkHeight int) bool {

	for i := 0; i < checkHeight; i++ {

		index1 := row - i
		index2 := row + i + 1
		if field[index1][colIndex] != field[index2][colIndex] {
			return false
		}
	}

	return true
}

func checkVer(field Field, forbiddenCol int) int {
	res := -1

	width := len(field[0])

	for col := 0; col < width-1; col++ {

		checkWidth := col + 1
		if checkWidth > width-col-1 {
			checkWidth = width - col - 1
		}

		sym := true

		for _, row := range field {
			if !symmetric(row, col, checkWidth) {
				sym = false
			}
		}

		if sym && (forbiddenCol == -1 || forbiddenCol != col) {
			res = col
			// fmt.Println("V-Sym at ", col)
		}
	}

	if res > -1 {
		res += 1
	}

	return res
}

func symmetric(row []int, col int, checkWidth int) bool {
	for i := 0; i < checkWidth; i++ {
		index1 := col - i
		index2 := col + i + 1
		if row[index1] != row[index2] {
			return false
		}
	}

	return true
}

func print(field Field) {
	for _, row := range field {

		for _, val := range row {
			if val == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type Field [][]int

func loadInput(fn string) []Field {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := []Field{}

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			res = append(res, toField(lines))
			lines = []string{}
		} else {
			lines = append(lines, line)
		}

	}

	res = append(res, toField(lines))

	return res
}

func toField(lines []string) Field {
	res := Field{}
	for _, row := range lines {
		temp := []int{}
		for _, col := range row {
			if col == '#' {
				temp = append(temp, 1)
			} else {
				temp = append(temp, 0)
			}
		}
		res = append(res, temp)
	}

	return res
}
