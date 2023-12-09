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
	part1()
	part2()
}

func part2() {
	seeds, levels := loadInput()

	// fmt.Println("seeds", seeds)
	// fmt.Printf("levels: %+v\n", levels)

	n := len(seeds) / 2

	min := -1

	for i := 0; i < n; i++ {
		deepCheck(levels, 0, seeds[2*i], seeds[2*i]+seeds[2*i+1]-1, &min)
	}

	fmt.Println("Part-2:", min)
}

func deepCheck(levels Levels, deep int, from, last int, min *int) {

	intervals := splitInterval(levels[deep], from, last)
	for _, interval := range intervals {

		temp := interval.First
		for _, levInt := range levels[deep] {
			if levInt.From <= interval.First && interval.First <= levInt.From+levInt.Len-1 {
				temp = levInt.To + (interval.First - levInt.From)
			}
		}

		delta := temp - interval.First

		if deep < len(levels)-1 {
			deepCheck(levels, deep+1, interval.First+delta, interval.Last+delta, min)
		} else {
			if *min == -1 || *min > temp {
				*min = temp
			}
		}
	}

}

func splitInterval(level Level, from, last int) []Period {
	points := map[int]int{}

	if from != last {
		for _, interval := range level {
			if from < interval.From && interval.From < last {
				points[interval.From] = 1
			}

			lastIntIndex := interval.From + interval.Len - 1
			if from < lastIntIndex && lastIntIndex < last {
				points[lastIntIndex+1] = 1
			}
		}
	}

	keys := make([]int, 0, len(points))
	for k := range points {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	// fmt.Println("points", keys)

	if len(points) == 0 {
		return []Period{{First: from, Last: last}}
	}

	res := []Period{}

	for index, point := range keys {
		if index == 0 {
			res = append(res, Period{First: from, Last: point - 1})
		} else {
			res = append(res, Period{First: keys[index-1], Last: point - 1})
		}
		if index == len(keys)-1 {
			res = append(res, Period{First: point, Last: last})
		}
	}

	return res
}

func check(levels Levels, from, len int) int {
	min := -1

	for i := from; i < from+len; i++ {
		next := i

		for _, level := range levels {
			temp := next
			for _, interval := range level {
				if interval.From <= next && next <= interval.From+interval.Len-1 {
					temp = interval.To + (next - interval.From)
				}
			}
			next = temp
		}

		if min == -1 || min > next {
			min = next
		}
	}
	return min
}

func part1() {
	seeds, levels := loadInput()

	min := -1

	for _, seed := range seeds {
		next := seed
		for _, level := range levels {
			temp := next
			for _, interval := range level {
				if interval.From <= next && next <= interval.From+interval.Len-1 {
					temp = interval.To + (next - interval.From)
				}
			}
			next = temp
		}

		if min == -1 || min > next {
			min = next
		}
	}

	fmt.Println("Part-1:", min)
}

type Interval struct {
	From int
	To   int
	Len  int
}

type Level []Interval

type Levels []Level

type Period struct {
	First, Last int
}

func loadInput() ([]int, Levels) {
	file, _ := os.Open("input-0.txt")
	scanner := bufio.NewScanner(file)

	res := make(Levels, 0)

	scanner.Scan()
	seedsStr := scanner.Text()
	parts := strings.Split(seedsStr, ": ")
	seeds := parseInts(parts[1])

	temp := make(Level, 0)
	for scanner.Scan() {
		line := scanner.Text()

		noNums := line == "" || strings.IndexRune(line, ':') > -1

		if noNums && len(temp) > 0 {
			res = append(res, temp)
			temp = make(Level, 0)
		} else if !noNums {
			nums := parseInts(line)
			temp = append(temp, Interval{From: nums[1], To: nums[0], Len: nums[2]})
		}
	}

	if len(temp) > 0 {
		res = append(res, temp)
	}

	return seeds, res
}

func parseInts(line string) []int {
	parts := strings.Split(line, " ")

	res := []int{}
	for _, part := range parts {
		if part != "" {
			n, _ := strconv.Atoi(part)
			res = append(res, n)
		}
	}
	return res
}
