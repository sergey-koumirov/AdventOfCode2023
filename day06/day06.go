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
	t, d := loadInput2()
	fmt.Println(t, d)

	index1 := 0
	for {
		test := index1 * (t - index1)
		if test > d {
			break
		}
		index1++
	}

	index2 := t
	for {
		test := index2 * (t - index2)
		if test > d {
			break
		}
		index2--
	}

	fmt.Println("Part-2:", index1, index2, index2-index1+1)
}

func part1() {
	times, distances := loadInput()

	// fmt.Println(times, distances)

	mult := 1
	for i, t := range times {

		cnt := 0
		for n := 0; n <= t; n++ {
			test := n * (t - n)

			// fmt.Println(distances[i], test, test > distances[i])
			if test > distances[i] {
				cnt += 1
			}
		}

		mult *= cnt

	}

	fmt.Println("Part-1:", mult)
}

func loadInput() ([]int, []int) {
	file, _ := os.Open("input-0.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	str1 := scanner.Text()
	parts1 := strings.Split(str1, ":")
	nums1 := parseInts(parts1[1])

	scanner.Scan()
	str2 := scanner.Text()
	parts2 := strings.Split(str2, ":")
	nums2 := parseInts(parts2[1])

	return nums1, nums2
}

func loadInput2() (int, int) {
	file, _ := os.Open("input-0.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	str1 := scanner.Text()
	num1, _ := strconv.Atoi(copyOnly(str1, "0123456789"))

	scanner.Scan()
	str2 := scanner.Text()
	num2, _ := strconv.Atoi(copyOnly(str2, "0123456789"))

	return num1, num2
}

func copyOnly(line, abc string) string {
	return strings.Map(
		func(r rune) rune {
			if strings.IndexRune(abc, r) >= 0 {
				return r
			}
			return -1
		},
		line,
	)
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
