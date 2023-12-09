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

func part2() {
	nums, syms := loadInput()
	sum := 0

	indexes := []int{0, 0, 0, 0}

	for _, sym := range syms {
		cnt := 0
		for indexNum, num := range nums {
			xInside := num.X-1 <= sym.X && sym.X <= num.X+num.Len
			yInside := num.Y-1 <= sym.Y && sym.Y <= num.Y+1
			if xInside && yInside {
				indexes[cnt] = indexNum
				cnt += 1
			}
		}
		if cnt == 2 {
			power := nums[indexes[0]].N * nums[indexes[1]].N
			sum += power
		}
	}

	fmt.Println("Part-2:", sum)
}

func part1() {
	nums, syms := loadInput()

	sum := 0

	for _, num := range nums {
		add := false

		for _, sym := range syms {
			xInside := num.X-1 <= sym.X && sym.X <= num.X+num.Len
			yInside := num.Y-1 <= sym.Y && sym.Y <= num.Y+1

			// fmt.Println(xInside, num.X-1, sym.X, num.X+num.Len+1)

			if xInside && yInside {
				add = true
			}

		}
		// fmt.Println(add, num)
		if add {
			sum += num.N
		}
	}

	fmt.Println("Part-1:", sum)
}

type Number struct {
	N, X, Y, Len int
}

type SXY struct {
	S    string
	X, Y int
}

type SIndex struct {
	S     string
	IsNum bool
	Index int
}

func loadInput() ([]Number, []SXY) {
	file, _ := os.Open("input-0.txt")
	scanner := bufio.NewScanner(file)

	nums := make([]Number, 0)
	syms := make([]SXY, 0)

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := split(line)

		// fmt.Println(line)
		// fmt.Printf("%+v\n\n", parts)

		for _, part := range parts {
			if part.IsNum {
				n, _ := strconv.Atoi(part.S)

				nums = append(nums, Number{
					N:   n,
					X:   part.Index,
					Y:   index,
					Len: len(part.S),
				})
			} else {
				syms = append(syms, SXY{
					S: part.S,
					X: part.Index,
					Y: index,
				})
			}
		}
		index += 1
	}

	return nums, syms
}

func split(line string) []SIndex {
	res := []SIndex{}

	startIndex := 0
	for index, char := range line {
		if strings.IndexRune("0123456789", char) > -1 {
			if index == len(line)-1 || strings.IndexByte("0123456789", line[index+1]) == -1 {
				res = append(res, SIndex{S: line[startIndex : index+1], IsNum: true, Index: startIndex})
				startIndex = index + 1
			}
		} else if char != '.' {
			res = append(res, SIndex{S: line[index : index+1], IsNum: false, Index: index})
			startIndex = index + 1
		} else {
			startIndex = index + 1
		}
	}
	return res
}
