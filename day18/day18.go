package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part2() {
	f := loadInput2("input-0.txt")
	fmt.Printf("(%d,%d) - (%d,%d)\n", f.MinX, f.MinY, f.MaxX, f.MaxY)
	// fmt.Println("-----------")
	// for _, line := range f.Lines {
	// 	fmt.Printf("(%d,%d) - (%d,%d)\n", line.X1, line.Y1, line.X2, line.Y2)
	// }
	// fmt.Println()

	// for i := 1; i < len(f.Lines); i++ {
	// 	line1 := f.Lines[i-1]
	// 	line2 := f.Lines[i]
	// 	if line1.X1 == line1.X2 && line2.X1 == line2.X2 || line1.Y1 == line1.Y2 && line2.Y1 == line2.Y2 {
	// 		fmt.Println("double !!!")
	// 	}
	// }

	sum := 0
	for y := f.MaxY; y >= f.MinY; y-- {

		if y%1000 == 0 {
			fmt.Println(y)
		}

		borders := []int{f.MinX - 1}

		for _, line := range f.Lines {
			if line.X1 == line.X2 && (line.Y1*2 <= y*2+1 && y*2+1 <= line.Y2*2 || line.Y2*2 <= y*2+1 && y*2+1 <= line.Y1*2) {
				borders = append(borders, line.X1)
			}
		}

		borders = append(borders, f.MaxX+1)

		sort.Ints(borders)

		// fmt.Println(y, "borders", borders)

		for i := 1; i < len(borders); i++ {
			if i%2 == 0 {
				left := borders[i-1]
				right := borders[i]
				test := right - left - 1

				for _, line := range f.Lines {
					if line.Y1 == y && line.Y1 == line.Y2 && left <= line.X1 && line.X1 <= right && left <= line.X2 && line.X2 <= right {
						// fmt.Println("    L", left, line.X1, line.X2, right)
						if left < line.X1 && line.X1 < right && left < line.X2 && line.X2 < right {
							test = test - abs(line.X1-line.X2) - 1
							// fmt.Println("    A", abs(line.X1-line.X2)+1)
						} else if (line.X1 == left || line.X1 == right) && left < line.X2 && line.X2 < right {
							test = test - abs(line.X1-line.X2)
							// fmt.Println("    B", abs(line.X1-line.X2))
						} else if (line.X2 == left || line.X2 == right) && left < line.X1 && line.X1 < right {
							test = test - abs(line.X1-line.X2)
							// fmt.Println("    C", abs(line.X1-line.X2))
						} else if line.X1 == left && line.X2 == right || line.X2 == left && line.X1 == right {
							test = test - abs(line.X1-line.X2) + 1
							// fmt.Println("    D", abs(line.X1-line.X2)-1)
						}
					}
				}

				// fmt.Println("    add", test)

				sum += test
			}
		}
	}

	for _, line := range f.Lines {
		sum += abs(line.X1-line.X2) + abs(line.Y1-line.Y2)
	}

	fmt.Println("Part-2:", sum)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func loadInput2(fn string) Field {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := Field{MaxX: -1, MaxY: -1, MinX: -1, MinY: -1}
	res.Lines = []Line{}

	x := 0
	y := 0

	applyMaxMin(&res, x, y)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		hexParts := parts[2][1 : len(parts[2])-1]

		l, err := strconv.ParseInt(hexParts[1:6], 16, 64)

		if err != nil {
			fmt.Println(err)
		}

		dx, dy := DirToDeltas(string(hexParts[6]))

		temp := Line{
			S:   line,
			Len: int(l),
			Hex: hexParts,
			X1:  x,
			Y1:  y,
			X2:  x + dx*int(l),
			Y2:  y + dy*int(l),
		}

		x += dx * int(l)
		y += dy * int(l)

		applyMaxMin(&res, x, y)

		res.Lines = append(res.Lines, temp)
	}

	return res
}

func part1() {
	f := loadInput("input-0.txt")
	print(f)
	fmt.Println(f.MinX, f.MinY, f.MaxX, f.MaxY)
	for _, line := range f.Lines {
		fmt.Printf("(%d,%d) - (%d,%d)\n", line.X1, line.Y1, line.X2, line.Y2)
	}
	fmt.Println()

	sum := 0
	for y := f.MaxY; y >= f.MinY; y-- {
		for x := f.MinX; x <= f.MaxX; x++ {
			in, _ := inline(&f, x, y)

			increased := false

			if in {
				sum += 1
				increased = true
			} else {
				lessThan := 0

				for _, line := range f.Lines {
					if onSameLevelVer(line.X1, line.Y1, line.X2, line.Y2, x, y) {
						lessThan += 1
					}
				}
				if lessThan%2 == 1 {
					sum += 1
					increased = true
				}
			}

			if increased {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}

	fmt.Println("Part-1:", sum)
}

func onSameLevelVer(x1, y1, x2, y2, x, y int) bool {
	return x1 == x2 && x1 < x && (y1*2 <= y*2+1 && y*2+1 <= y2*2 || y2*2 <= y*2+1 && y*2+1 <= y1*2)
}

type Line struct {
	S              string
	Code           string
	Len            int
	Hex            string
	X1, Y1, X2, Y2 int
}

type Field struct {
	Lines                  []Line
	MaxX, MaxY, MinX, MinY int
}

func print(f Field) {
	for y := f.MaxY; y >= f.MinY; y-- {
		for x := f.MinX; x <= f.MaxX; x++ {
			in, _ := inline(&f, x, y)
			if in {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func inline(f *Field, x, y int) (bool, string) {
	for _, line := range f.Lines {
		if line.X1 == x && line.X2 == x && (line.Y1 <= y && y <= line.Y2 || line.Y2 <= y && y <= line.Y1) {
			return true, line.Code
		}
		if line.Y1 == y && line.Y2 == y && (line.X1 <= x && x <= line.X2 || line.X2 <= x && x <= line.X1) {
			return true, line.Code
		}
	}
	return false, "."
}

func loadInput(fn string) Field {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := Field{MaxX: -1, MaxY: -1, MinX: -1, MinY: -1}
	res.Lines = []Line{}

	x := 0
	y := 0

	applyMaxMin(&res, x, y)

	code := "#"
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		hexParts := parts[2][1 : len(parts[2])-2]

		l, _ := strconv.Atoi(parts[1])
		dx, dy := DirToDeltas(parts[0])

		temp := Line{
			S:    line,
			Code: code,
			Len:  l,
			Hex:  hexParts,
			X1:   x,
			Y1:   y,
			X2:   x + dx*l,
			Y2:   y + dy*l,
		}

		if code == "#" {
			code = "@"
		} else {
			code = "#"
		}

		x += dx * l
		y += dy * l

		applyMaxMin(&res, x, y)

		res.Lines = append(res.Lines, temp)
	}

	// U 2 (#7a21e3)

	return res
}

func applyMaxMin(f *Field, x, y int) {
	if f.MinX == -1 || f.MinX > x {
		f.MinX = x
	}
	if f.MinY == -1 || f.MinY > y {
		f.MinY = y
	}

	if f.MaxX == -1 || f.MaxX < x {
		f.MaxX = x
	}
	if f.MaxY == -1 || f.MaxY < y {
		f.MaxY = y
	}
}

func DirToDeltas(dir string) (int, int) {
	if dir == "R" || dir == "0" {
		return 1, 0
	}
	if dir == "L" || dir == "2" {
		return -1, 0
	}
	if dir == "U" || dir == "3" {
		return 0, 1
	}
	if dir == "D" || dir == "1" {
		return 0, -1
	}
	return 0, 0
}
