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
	_, fieldArr, row, col := loadInput("input-0.txt")

	curRow := row
	curCol := col
	step := 0
	dir := possiblePrevDirs[fieldArr[row][col].Symbol][0]
	for {
		curRow, curCol, dir = nextPipe(&fieldArr, curRow, curCol, dir)
		fieldArr[curRow][curCol].Tag = 1
		if col == curCol && row == curRow {
			break
		}
		step++
	}

	print(fieldArr, row, col, 1)

	sum := 0

	for _, vals := range fieldArr {
		inside := false
		for _, node := range vals {

			if node.Tag == 0 && inside {
				sum++
			}

			if node.Tag == 1 && (node.Symbol == '|' || node.Symbol == '7' || node.Symbol == 'F') {
				inside = !inside
			}
		}
	}

	fmt.Println("Part-2:", sum)
}

var possiblePrevDirs = map[rune][]rune{
	'|': {'U', 'D'},
	'F': {'L', 'U'},
	'L': {'L', 'D'},
	'J': {'R', 'D'},
	'7': {'R', 'U'},
	'-': {'L', 'R'},
}

func part1() {
	_, fieldArr, row, col := loadInput("input-3.txt")

	print(fieldArr, row, col, 0)

	curRow := row
	curCol := col
	step := 0
	dir := possiblePrevDirs[fieldArr[row][col].Symbol][0]
	for {
		curRow, curCol, dir = nextPipe(&fieldArr, curRow, curCol, dir)

		fieldArr[curRow][curCol].Tag = 1
		// print(fieldArr, curRow, curCol)

		if col == curCol && row == curRow {
			break
		}
		step++
	}

	print(fieldArr, row, col, 1)

	fmt.Println("Part-1:", (step+1)/2)
}

func nextPipe(fieldRef *NodeArr, curRow, curCol int, prevDir rune) (int, int, rune) {
	sym := (*fieldRef)[curRow][curCol].Symbol

	// fmt.Println(string(sym), string(prevDir))

	if sym == '|' && prevDir == 'U' {
		return curRow - 1, curCol, 'U'
	}
	if sym == '|' && prevDir == 'D' {
		return curRow + 1, curCol, 'D'
	}

	if sym == '-' && prevDir == 'L' {
		return curRow, curCol - 1, 'L'
	}
	if sym == '-' && prevDir == 'R' {
		return curRow, curCol + 1, 'R'
	}

	if sym == 'F' && prevDir == 'U' {
		return curRow, curCol + 1, 'R'
	}
	if sym == 'F' && prevDir == 'L' {
		return curRow + 1, curCol, 'D'
	}

	if sym == 'L' && prevDir == 'L' {
		return curRow - 1, curCol, 'U'
	}
	if sym == 'L' && prevDir == 'D' {
		return curRow, curCol + 1, 'R'
	}

	if sym == 'J' && prevDir == 'R' {
		return curRow - 1, curCol, 'U'
	}
	if sym == 'J' && prevDir == 'D' {
		return curRow, curCol - 1, 'L'
	}

	if sym == '7' && prevDir == 'U' {
		return curRow, curCol - 1, 'L'
	}
	if sym == '7' && prevDir == 'R' {
		return curRow + 1, curCol, 'D'
	}

	return -1, -1, '.'
}

type Node struct {
	Symbol   rune
	Key      string
	Row, Col int
	Children []*Node
	Tag      int
}

var SforFN = map[string]rune{
	"input-0.txt": 'J',
	"input-1.txt": 'F',
	"input-2.txt": 'F',
	"input-3.txt": '7',
	"input-4.txt": 'F',
}

var RuneToIcon = map[rune]string{
	'|': "│",
	'L': "╰",
	'-': "─",
	'F': "╭",
	'J': "╯",
	'7': "╮",
	'.': " ",
}

type NodeArr [][]Node

func print(field NodeArr, curRow, curCol, tag int) {
	fmt.Println()
	for row, vals := range field {
		for col, node := range vals {
			if curCol == col && curRow == row {
				fmt.Print("*")
			} else if node.Tag != tag {
				fmt.Print(".")
			} else {
				fmt.Print(RuneToIcon[node.Symbol])
			}
		}
		fmt.Println()
	}
}

func loadInput(filename string) (map[string]Node, NodeArr, int, int) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	resArr := make(NodeArr, 0)

	startRow := -1
	startCol := -1

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		// fmt.Println(line)

		temp := make([]Node, 0)

		for col, r := range line {
			corrected := r
			if r == 'S' {
				startRow = row
				startCol = col
				corrected = SforFN[filename]
			}

			temp = append(
				temp,
				Node{
					Symbol:   corrected,
					Key:      fmt.Sprintf("%d-%d", row, col),
					Row:      row,
					Col:      col,
					Children: []*Node{},
				},
			)
		}

		resArr = append(resArr, temp)
		row++
	}

	// makeChildren(&resArr)

	resMap := make(map[string]Node)
	for _, row := range resArr {
		for _, node := range row {
			test := node
			resMap[node.Key] = test
		}
	}

	return resMap, resArr, startRow, startCol
}
