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

var field Ints

// var mins Ints
var maxRowIndex int
var maxColIndex int

func part2() {
	field = loadInput("input-1.txt")
	maxRowIndex = len(field) - 1
	maxColIndex = len(field[0]) - 1
	fmt.Println(maxRowIndex+1, "x", maxColIndex+1)

	printField()

	approximate := astarUltra()

	fmt.Println("Part-2:", approximate)
}

func astarUltra() int {
	res := make(IntsMap, maxRowIndex+1)
	for rowIdx := range field {
		res[rowIdx] = make([]Keys, maxColIndex+1)

		for colIdx := range field[0] {
			res[rowIdx][colIdx] = Keys{}
		}
	}

	points := []CheckPoint{
		{Row: 1, Col: 0, Heatloss: 0, Dirs: ".........D"},
		{Row: 0, Col: 1, Heatloss: 0, Dirs: ".........R"},
	}

	cnt := 0
	ultraBest := -1
	ultraPath := ""
	for len(points) > 0 {
		// fmt.Println(points)
		cnt++
		if cnt%1000000 == 0 {
			fmt.Println(cnt, "L=", len(points), ultraBest)
		}

		first := points[0]
		points = points[1:]

		loss := field[first.Row][first.Col]

		last10 := first.Dirs[len(first.Dirs)-10:]
		best, bestExist := res[first.Row][first.Col][last10]

		// k := mins[first.Row][first.Col]

		lineLen := getLineLen(first.Dirs)

		if (!bestExist || first.Heatloss+loss < best) && !(first.Row == maxRowIndex && first.Col == maxColIndex && lineLen < 4) {
			// fmt.Println("res", first.Row, first.Col, res[first.Row][first.Col])

			if first.Row == maxRowIndex && first.Col == maxColIndex {
				// fmt.Println(first.Dirs)
				if ultraBest == -1 || ultraBest > first.Heatloss+loss {
					ultraBest = first.Heatloss + loss
					ultraPath = first.Dirs
				}
			}

			res[first.Row][first.Col][last10] = first.Heatloss + loss

			dirsLen := len(first.Dirs)
			lastRune := rune(first.Dirs[dirsLen-1])
			dirs := fromToDeltas(lastRune)

			for _, d := range dirs {
				if inside(first.Row+d.DR, first.Col+d.DC) && (d.Dir == lastRune && lineLen < 4 || lineLen >= 4 && lineLen < 10) {
					next := CheckPoint{
						Row:      first.Row + d.DR,
						Col:      first.Col + d.DC,
						Heatloss: first.Heatloss + loss,
						Dirs:     first.Dirs + string(d.Dir),
					}
					points = append(points, next)
				}
			}
		}
		// fmt.Println(first.Row, first.Col, first.Heatloss, first.Dirs[2:], first.Heatloss+loss)
		// print(res, field)
		// fmt.Scanln()
	}

	fmt.Println(ultraPath)
	fmt.Println(res[maxRowIndex][maxColIndex])

	min := -1
	for _, v := range res[maxRowIndex][maxColIndex] {
		if min == -1 || min > v {
			min = v
		}
	}

	return min
}

func getLineLen(s string) int {
	last := s[len(s)-1]

	cnt := 1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] != last {
			return cnt
		}
		cnt++
	}
	return cnt
}

func part1() {
	field = loadInput("input-0.txt")
	maxRowIndex = len(field) - 1
	maxColIndex = len(field[0]) - 1

	// fmt.Println("minimums ...")
	// mins = minimums()

	if maxColIndex < 20 {
		printField()
		// printInts(mins)
	}

	approximate := astar()
	// print(lossmap, field)

	fmt.Println("approximate", approximate)

	// deep(0, 0, "...", 0, &approximate, 0)

	fmt.Println("Part-1:", approximate)
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

	// return res
}

func astarMin(row, col, destRow, destCol int) Ints {
	res := make(Ints, len(field))
	for rowIdx := range field {
		res[rowIdx] = make([]int, len(field[0]))
		for colIdx := range field {
			res[rowIdx][colIdx] = -1
		}
	}

	points := []CheckPoint{
		{Row: row, Col: col},
	}

	currentBest := -1

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

			dirs := fromToDeltas('.')

			for _, d := range dirs {
				if inside(first.Row+d.DR, first.Col+d.DC) {
					next := CheckPoint{
						Row:      first.Row + d.DR,
						Col:      first.Col + d.DC,
						Heatloss: first.Heatloss + loss,
					}
					points = append(points, next)
				}
			}
		}
	}

	// printInts(res)

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

func astar() int {
	res := make(IntsMap, maxRowIndex+1)
	for rowIdx := range field {
		res[rowIdx] = make([]Keys, maxColIndex+1)
		for colIdx := range field {
			res[rowIdx][colIdx] = Keys{}
		}
	}

	points := []CheckPoint{
		{Row: 1, Col: 0, Heatloss: 0, Dirs: "..D"},
		{Row: 0, Col: 1, Heatloss: 0, Dirs: "..R"},
	}

	cnt := 0
	for len(points) > 0 {
		// fmt.Println(points)
		cnt++
		if cnt%1000000 == 0 {
			fmt.Println(cnt, "L=", len(points))
		}

		first := points[0]
		points = points[1:]

		loss := field[first.Row][first.Col]

		last3 := first.Dirs[len(first.Dirs)-3:]
		best, bestExist := res[first.Row][first.Col][last3]

		// k := mins[first.Row][first.Col]

		if !bestExist || first.Heatloss+loss < best {
			res[first.Row][first.Col][last3] = first.Heatloss + loss

			dirsLen := len(first.Dirs)

			oneRune := rune(first.Dirs[dirsLen-3])
			twoRune := rune(first.Dirs[dirsLen-2])
			lastRune := rune(first.Dirs[dirsLen-1])

			dirs := fromToDeltas(lastRune)

			for _, d := range dirs {
				if inside(first.Row+d.DR, first.Col+d.DC) && (d.Dir != oneRune || d.Dir != twoRune || d.Dir != lastRune) {
					next := CheckPoint{
						Row:      first.Row + d.DR,
						Col:      first.Col + d.DC,
						Heatloss: first.Heatloss + loss,
						Dirs:     first.Dirs + string(d.Dir),
					}
					points = append(points, next)
				}
			}
		}
		// fmt.Println(first.Row, first.Col, first.Heatloss, first.Dirs[2:], first.Heatloss+loss)
		// print(res, field)
		// fmt.Scanln()
	}

	// fmt.Println(res[maxRowIndex][maxColIndex])

	min := -1
	for _, v := range res[maxRowIndex][maxColIndex] {
		if min == -1 || min > v {
			min = v
		}
	}

	return min
}

func getTurn(dirs string) int {
	last := dirs[len(dirs)-1]

	turn := 0
	for i := len(dirs) - 1; i >= 0; i-- {
		if dirs[i] != last {
			return turn
		}
		turn += 1
	}
	return turn
}

func inside(row, col int) bool {
	return row >= 0 && row <= maxRowIndex && col >= 0 && col <= maxColIndex
}

func fromToDeltas(r rune) []Direction {
	if r == 'R' {
		return []Direction{
			{DR: -1, DC: 0, Dir: 'U'},
			{DR: 0, DC: 1, Dir: 'R'},
			{DR: 1, DC: 0, Dir: 'D'},
		}
	}
	if r == 'L' {
		return []Direction{
			{DR: -1, DC: 0, Dir: 'U'},
			{DR: 0, DC: -1, Dir: 'L'},
			{DR: 1, DC: 0, Dir: 'D'},
		}
	}
	if r == 'U' {
		return []Direction{
			{DR: 0, DC: 1, Dir: 'R'},
			{DR: -1, DC: 0, Dir: 'U'},
			{DR: 0, DC: -1, Dir: 'L'},
		}
	}
	if r == 'D' {
		return []Direction{
			{DR: 0, DC: 1, Dir: 'R'},
			{DR: 1, DC: 0, Dir: 'D'},
			{DR: 0, DC: -1, Dir: 'L'},
		}
	}
	if r == '.' {
		return []Direction{
			{DR: 0, DC: 1, Dir: 'R'},
			{DR: 1, DC: 0, Dir: 'D'},
			{DR: 0, DC: -1, Dir: 'L'},
			{DR: -1, DC: 0, Dir: 'U'},
		}
	}
	panic("fromToDeltas !!!")
}

type Ints [][]int

type Keys map[string]int
type IntsMap [][]Keys

type CheckPoint struct {
	Row, Col int
	Dirs     string
	Heatloss int
}

type Direction struct {
	DR, DC int
	Dir    rune
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
