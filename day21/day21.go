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

	// 65
	// 202300

	res1 := 5801 + 5797 + 5807 + 5811
	res1 += (988 + 981 + 973 + 991) * 202300
	res1 += (6748 + 6751 + 6758 + 6755) * (202300 - 1)
	res1 += 202300*202300*7747 + (202300-1)*(202300-1)*7702

	// 632258017940374
	// 632258017940374

	fmt.Println(res1)

	field, start := loadInput("input-0.txt")

	maxRow := len(field) - 1
	maxCol := len(field[0]) - 1

	cells := map[int]map[int]int{}
	addToCells(cells, start.Row, start.Col)

	width := len(field)
	stop := width * 5
	half := width / 2

	fmt.Println("half", half)

	for i := 0; i < stop; i++ {
		newCells := map[int]map[int]int{}

		for row, rowData := range cells {
			for col := range rowData {
				for _, delta := range deltas {

					newRow := row + delta[0]
					newCol := col + delta[1]

					realRow := remainder(newRow, maxRow+1)
					realCol := remainder(newCol, maxCol+1)

					zero := field[realRow][realCol] == 0

					if zero {
						addToCells(newCells, newRow, newCol)
					}
				}
			}
		}

		if i == 65+width*4-1 {
			test := 0
			for _, v := range newCells {
				test += len(v)
			}
			fmt.Println(i+1, test)

			size := (i/width + 1)
			fmt.Println("size", size, i%width)

			for krow := -1 * size; krow <= size; krow++ {
				for kcol := -1 * size; kcol <= size; kcol++ {

					minRow := width * krow
					maxRow := width*(krow+1) - 1

					minCol := width * kcol
					maxCol := width*(kcol+1) - 1

					// fmt.Println("| ", minRow, "x", minCol, "   ", maxRow, "x", maxCol)
					sum := 0
					for r, vv := range newCells {
						for c := range vv {
							if minRow <= r && r <= maxRow && minCol <= c && c <= maxCol {
								sum++
							}
						}
					}
					fmt.Printf("%5d ", sum)
				}
				fmt.Println()
			}
			fmt.Println()
		}

		cells = newCells
	}

	fmt.Println("Part-2:")
}

func remainder(n1, n2 int) int {
	r := n1 % n2

	if r >= 0 {
		return r
	}
	return r + n2
}

var deltas = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func part1() {
	field, start := loadInput("input-0.txt")

	maxRow := len(field) - 1
	maxCol := len(field[0]) - 1

	cells := map[int]map[int]int{}
	addToCells(cells, start.Row, start.Col)

	stop := 64

	for i := 0; i < stop; i++ {
		// fmt.Println(cells)

		newCells := map[int]map[int]int{}

		for row, rowData := range cells {
			for col := range rowData {
				for _, delta := range deltas {
					newRow := row + delta[0]
					newCol := col + delta[1]
					inside := newRow >= 0 && newRow <= maxRow && newCol >= 0 && newCol <= maxCol
					zero := field[newRow][newCol] == 0

					// fmt.Println(inside, newRow, newCol, zero)

					if inside && zero {
						addToCells(newCells, newRow, newCol)
					}
				}
			}
		}
		print(field, newCells)
		// fmt.Println(newCells)
		// fmt.Println()
		cells = newCells
	}

	sum := 0
	for _, v := range cells {
		sum += len(v)
	}

	fmt.Println("Part-1:", sum)
}

func print(field Field, cells map[int]map[int]int) {

	for row, rowData := range field {
		rowPoints, exRow := cells[row]

		for col := range rowData {
			exCol := false
			if exRow {
				_, exCol = rowPoints[col]
			}

			if exCol {
				fmt.Print("O")
			} else if field[row][col] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func addToCells(cells map[int]map[int]int, row, col int) {
	_, rowEx := cells[row]
	if !rowEx {
		cells[row] = map[int]int{}
	}

	_, colEx := cells[row][col]
	if !colEx {
		cells[row][col] = 1
	}
}

type Field [][]int
type Cell struct {
	Row, Col int
}

func loadInput(fn string) (Field, Cell) {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := make(Field, 0)
	cell := Cell{}
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		temp := []int{}

		for col, b := range line {
			if b == '.' || b == 'S' {
				temp = append(temp, 0)
			} else {
				temp = append(temp, 1)
			}

			if b == 'S' {
				cell.Row = row
				cell.Col = col
			}
		}

		res = append(res, temp)
		row++
	}

	return res, cell
}
