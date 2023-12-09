package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := loadInput()

	sum := 0

	for _, line := range lines {
		firstIndex := strings.IndexAny(line, "123456789")
		lastIndex := strings.LastIndexAny(line, "123456789")
		num := int(line[firstIndex]-48)*10 + int(line[lastIndex]-48)
		sum += num
	}

	fmt.Println("PART-1", sum)

}

func part2() {
	subs := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	strToInt := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	lines := loadInput()

	sum := 0

	for _, line := range lines {
		_, firstEl := IndexAnyArray(line, subs)
		_, lastEl := LastIndexAnyArray(line, subs)

		num1 := strToInt[subs[firstEl]]
		num2 := strToInt[subs[lastEl]]

		num := num1*10 + num2
		sum += num
	}

	fmt.Println("PART-2", sum)
}

func loadInput() []string {
	file, _ := os.Open("input-0.txt")
	scanner := bufio.NewScanner(file)

	res := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)

	}

	return res
}

func IndexAnyArray(s string, subs []string) (int, int) {
	res := -1
	elPos := -1

	for index, sub := range subs {
		temp := strings.Index(s, sub)
		if temp > -1 && (res == -1 || temp < res) {
			res = temp
			elPos = index
		}
	}

	return res, elPos
}

func LastIndexAnyArray(s string, subs []string) (int, int) {
	res := -1
	elPos := -1

	for index, sub := range subs {
		temp := strings.LastIndex(s, sub)
		if temp > res {
			res = temp
			elPos = index
		}
	}

	return res, elPos
}
