package main

import (
	"bufio"
	"fmt"
	"os"
)

type Ints [][]int
type Keys map[string]int
type IntsMap [][]Keys
type CheckPoint struct {
	Row, Col int
	Dirs     string
	Heatloss int
}
type Direction struct {
	Row, Col, AddHeat int
	Dir               rune
}

var field Ints
var mins Ints
var maxRowIndex int
var maxColIndex int

func main() {
	field = loadInput("input-0.txt")
	maxRowIndex = len(field) - 1
	maxColIndex = len(field[0]) - 1
	mins = minimums()
	fmt.Println(maxRowIndex+1, "x", maxColIndex+1)

	// printField()
	// printInts(mins)

	approximate := astarUltra()

	fmt.Println("Part-2:", approximate)
}

func astarUltra() int {
	points := []CheckPoint{
		{Row: 0, Col: 0, Heatloss: 0, Dirs: "."},
	}

	visited := make(IntsMap, maxRowIndex+1)
	for irow := range field {
		visited[irow] = make([]Keys, maxColIndex+1)
		for icol := range field[0] {
			visited[irow][icol] = Keys{}
		}
	}

	cnt := 0
	ultraBest := 9 * (maxRowIndex + 1) * maxColIndex

	for len(points) > 0 {
		cnt++

		// first := points[0]
		// points = points[1:]
		first := points[len(points)-1]
		points = points[:len(points)-1]

		loss := field[first.Row][first.Col]
		lineLen := getLineLen(first.Dirs)
		minHeat := mins[first.Row][first.Col]
		fullLoss := first.Heatloss + loss

		if cnt%1000000 == 0 {
			fmt.Println(cnt, "L=", len(points), ultraBest)
		}
		// fmt.Println("L=", len(points), "UB=", ultraBest, "Full=", fullLoss)

		visitedCell := false
		visitedFullLoss := 0

		ll := len(first.Dirs)
		if ll >= 10 {
			dirKey := first.Dirs[ll-lineLen:]
			visitedFullLoss, visitedCell = visited[first.Row][first.Col][dirKey]

			// if visitedCell {
			// 	fmt.Println(first.Dirs)
			// 	fmt.Println(first.Row, first.Col, visited[first.Row][first.Col])
			// 	fmt.Scanln()
			// }
		}

		if (!visitedCell || visitedCell && visitedFullLoss > fullLoss) && (fullLoss < ultraBest) && (fullLoss+minHeat < ultraBest) && (first.Row != maxRowIndex || first.Col != maxColIndex || lineLen >= 4) {

			if ll >= 10 {
				dirKey := first.Dirs[ll-lineLen:]
				visited[first.Row][first.Col][dirKey] = fullLoss
			}

			if first.Row == maxRowIndex && first.Col == maxColIndex && ultraBest > fullLoss {
				ultraBest = fullLoss
			}

			last := rune(first.Dirs[len(first.Dirs)-1])

			// printDir(first.Dirs)
			// fmt.Println("RC=", ultraBest, first.Row, first.Col, first.Dirs, lineLen)

			psbs := LastDirToDirs(last)
			for _, psb := range psbs {

				if lineLen == 0 || lineLen < 4 && psb == last || lineLen >= 4 && lineLen < 10 || lineLen == 10 && psb != last {
					dr, dc := DirToDX(psb)
					if inside(first.Row+dr, first.Col+dc) {
						// fmt.Println("    ", string(psb))
						points = append(
							points, CheckPoint{
								Row:      first.Row + dr,
								Col:      first.Col + dc,
								Dirs:     first.Dirs + string(psb),
								Heatloss: fullLoss,
							},
						)
					}
				}
			}
			// fmt.Scanln()
		}
	}

	return ultraBest - field[0][0]
}

func LastDirToDirs(dir rune) []rune {
	if dir == 'R' {
		return []rune{'R', 'U', 'D'}
	}
	if dir == 'L' {
		return []rune{'L', 'U', 'D'}
	}
	if dir == 'U' {
		return []rune{'U', 'L', 'R'}
	}
	if dir == 'D' {
		return []rune{'D', 'L', 'R'}
	}
	if dir == '.' {
		return []rune{'R', 'D'}
	}
	panic("LastDirToDirs")
}

func DirToDX(dir rune) (int, int) {
	if dir == 'R' {
		return 0, 1
	}
	if dir == 'L' {
		return 0, -1
	}
	if dir == 'U' {
		return -1, 0
	}
	if dir == 'D' {
		return 1, 0
	}
	if dir == '.' {
		return 0, 0
	}
	panic("DirToDX")
}

func inside(row, col int) bool {
	return row >= 0 && row <= maxRowIndex && col >= 0 && col <= maxColIndex
}

func loadInput(fn string) Ints {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := Ints{}

	for scanner.Scan() {
		line := scanner.Text()

		temp := []int{}

		for _, r := range line {
			temp = append(temp, int(r)-48)
		}
		res = append(res, temp)
	}

	return res
}

func printField() {
	fmt.Println()
	for _, row := range field {
		for _, n := range row {
			fmt.Printf("%2d", n)
		}

		fmt.Println()
	}
	fmt.Println()
}

func printDir(dir string) {
	row := 0
	col := 0
	path := map[int]int{}
	for _, r := range dir {
		dr, dc := DirToDX(r)
		row += dr
		col += dc

		key := row*1000 + col
		path[key] = 1
	}

	fmt.Println()
	for irow := range field {
		for icol := range field[0] {
			key := irow*1000 + icol
			_, ex := path[key]
			if ex {
				fmt.Print(" *")
			} else {
				fmt.Printf("%2d", field[irow][icol])
			}

		}
		fmt.Println()
	}
	// fmt.Println()
}

func getLineLen(s string) int {
	last := s[len(s)-1]

	if last == '.' {
		return 0
	}

	cnt := 1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] != last {
			return cnt
		}
		cnt++
	}
	return cnt
}

func minimums() Ints {
	kk := astarMin(maxRowIndex, maxColIndex, 0, 0)
	res := make(Ints, maxRowIndex+1)
	for r := 0; r <= maxRowIndex; r++ {
		res[r] = make([]int, maxColIndex+1)
		for c := 0; c <= maxColIndex; c++ {
			res[r][c] = kk[r][c] - field[r][c]
		}
	}
	return res
}

func astarMin(row, col, destRow, destCol int) Ints {
	fmt.Println("astarMin")

	res := make(Ints, maxRowIndex+1)
	for rowIdx := range field {
		res[rowIdx] = make([]int, maxColIndex+1)
		for colIdx := range field[0] {
			res[rowIdx][colIdx] = -1
		}
	}

	points := []CheckPoint{
		{Row: row, Col: col},
	}

	currentBest := -1

	dirs := []Direction{
		{Row: 0, Col: 1, Dir: 'R'},
		{Row: 1, Col: 0, Dir: 'D'},
		{Row: 0, Col: -1, Dir: 'L'},
		{Row: -1, Col: 0, Dir: 'U'},
	}

	for len(points) > 0 {
		first := points[0]
		points = points[1:]

		loss := field[first.Row][first.Col]
		best := res[first.Row][first.Col]

		if (best == -1 || first.Heatloss+loss < best) && (currentBest == -1 || first.Heatloss+loss < currentBest) {
			res[first.Row][first.Col] = first.Heatloss + loss

			if destRow == first.Row && destCol == first.Col {
				if currentBest == -1 || currentBest > res[first.Row][first.Col] {
					currentBest = res[first.Row][first.Col]
				}
			}

			for _, d := range dirs {
				if inside(first.Row+d.Row, first.Col+d.Col) {
					next := CheckPoint{
						Row:      first.Row + d.Row,
						Col:      first.Col + d.Col,
						Heatloss: first.Heatloss + loss,
					}
					points = append(points, next)
				}
			}
			// printInts(res)
			// fmt.Scanln()
		}

	}

	// printInts(res)

	return res
}

func printInts(ii Ints) {
	fmt.Println()
	for _, row := range ii {
		for _, n := range row {
			fmt.Printf("%3d", n)
		}

		fmt.Println()
	}
	fmt.Println()
}
