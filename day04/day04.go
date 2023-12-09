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
	data := loadInput()

	points := []int{}
	original := []int{}
	heap := []int{}

	for i, res := range data {
		common := Intersect(res.Set, res.Win)
		data[i].Common = len(common)
		points = append(points, len(common))
		original = append(original, i+1)
		heap = append(heap, i+1)
	}

	copied := 0
	sum := len(original)

	for {
		// fmt.Println("test", points, original, heap)

		temp := []int{}
		for _, t := range heap {
			index := t - 1
			if points[index] > 0 {
				temp = append(temp, original[index+1:index+points[index]+1]...)
			}
		}

		// fmt.Println(len(temp))
		sum += len(temp)

		if len(temp) == 0 {
			break
		}
		copied += 1
		heap = temp
	}

	fmt.Println("Part-2:", sum)
}

func part1() {
	sum := 0

	data := loadInput()

	for _, res := range data {
		common := Intersect(res.Set, res.Win)
		l := len(common)
		if l > 0 {
			points := (1 << (len(common) - 1))
			sum += points
		}
	}

	fmt.Println("Part-1:", sum)
}

type Integers []int

type Result struct {
	Index  int
	Win    Integers
	Set    Integers
	Common int
}

func loadInput() []Result {
	file, _ := os.Open("input-0.txt")
	scanner := bufio.NewScanner(file)

	res := make([]Result, 0)

	index := 1
	for scanner.Scan() {
		line := scanner.Text()

		parts1 := strings.Split(line, ": ")
		parts2 := strings.Split(parts1[1], " | ")

		wins := parseInts(parts2[0])
		sets := parseInts(parts2[1])

		res = append(res, Result{Index: index, Win: wins, Set: sets})
		index += 1
	}

	return res
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

func Intersect[T comparable](arrs ...[]T) []T {
	m := make(map[T]int)

	var (
		tmpArr []T
		count  int
		ok     bool
	)
	for idx1 := range arrs {
		tmpArr = Distinct(arrs[idx1])

		for idx2 := range tmpArr {
			count, ok = m[tmpArr[idx2]]
			if !ok {
				m[tmpArr[idx2]] = 1
			} else {
				m[tmpArr[idx2]] = count + 1
			}
		}
	}

	var (
		ret     []T
		lenArrs int = len(arrs)
	)
	for k, v := range m {
		if v == lenArrs {
			ret = append(ret, k)
		}
	}

	return ret
}

func Distinct[T comparable](arrs ...[]T) []T {
	// put the values of our slice into a map
	// the key's of the map will be the slice's unique values
	m := make(map[T]struct{})
	for idx1 := range arrs {
		for idx2 := range arrs[idx1] {
			m[arrs[idx1][idx2]] = struct{}{}
		}
	}

	// create the output slice and populate it with the map's keys
	res := make([]T, len(m))
	i := 0
	for k := range m {
		res[i] = k
		i++
	}

	return res
}
