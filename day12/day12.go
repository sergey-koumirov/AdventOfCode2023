package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1() // 0 - 7843    / 12615
	// part2() // 1 - 525_152 / 686_495 / 656_932 / 654_576
	part2v2() // 1 - 525_152 / 686_495 / 656_932 / 654_576

	// 2 -  1_960_400 / 1_649_649_056 / 868_879_266 / 427_000_309 / 5_309_179
}

func part2v2() {
	datas := loadInputV2("input-3.txt", false)
	// fmt.Println()
	sum := 0
	deepCalls := 0

	for i, d := range datas {
		cnt := 0
		fmt.Println(d.Pattern, d.Org)

		deepV2(d, []int{}, 0, d.Pattern, &cnt, &deepCalls)

		fmt.Println(">>", i, cnt)
		fmt.Println()

		sum += cnt
	}
	fmt.Println(deepCalls)
	fmt.Println("Part-2.1:", sum)
}

func deepV2(dl DataLineV2, positions []int, index int, pattern string, cntRef *int, deepCalls *int) {
	*deepCalls++

	val := dl.Groups[index].Val

	for _, start := range dl.Groups[index].Starts {
		if (index == 0 || index > 0 && start > positions[len(positions)-1]) && possible(pattern, val, start) && strings.IndexRune(pattern[0:start], '#') == -1 {

			next := append(positions, start)
			newPattern := pattern[0:start] + strings.Repeat("$", val) + pattern[start+val:]

			if index == len(dl.Groups)-1 {
				if strings.IndexRune(newPattern, '#') == -1 {
					fmt.Println(newPattern)
					*cntRef++
				}
			} else {
				deepV2(dl, next, index+1, newPattern, cntRef, deepCalls)
			}
		}
	}
}

func part2() {
	fmt.Println(canFit("###", []int{1, 1}))

	datas := loadInput("input-2.txt", true)

	// fmt.Println(datas[0].Starts)
	// fmt.Println(datas[0].Ends)

	cnt := 0
	deepCalls := 0
	for i, d := range datas {
		fmt.Println(i, d.Pattern, d.Groups)
		// fmt.Println()
		deep(d, 0, d.Pattern, []int{}, &cnt, &deepCalls)
		// fmt.Println()
		// fmt.Println()
	}

	fmt.Println(deepCalls)
	fmt.Println("Part-2:", cnt)
}

func part1() {
	datas := loadInput("input-0.txt", false)
	cnt := 0
	deepCalls := 0
	for _, d := range datas {
		deep(d, 0, d.Pattern, []int{}, &cnt, &deepCalls)
	}
	fmt.Println(deepCalls)
	fmt.Println("Part-1:", cnt)
}

func deep(dl DataLine, index int, pattern string, positions []int, cntRef *int, deepCalls *int) {
	*deepCalls++
	from := dl.Starts[index]
	g := dl.Groups[index]

	if len(positions) > 0 && from <= positions[len(positions)-1] {
		from = positions[len(positions)-1] + dl.Groups[index-1] + 1
	}

	for pos := from; pos <= dl.Ends[index]; pos++ {
		if possible(pattern, g, pos) {

			next := append(positions, pos)
			newPattern := pattern[0:pos] + strings.Repeat("$", g) + pattern[pos+g:]

			if index == len(dl.Groups)-1 {
				if strings.IndexRune(newPattern, '#') == -1 {
					*cntRef++
				}
			} else {
				if canFit(pattern[pos+g:], dl.Groups[index+1:]) {
					deep(dl, index+1, newPattern, next, cntRef, deepCalls)
				}

			}
		}
	}
}

func canFit(pattern string, gg []int) bool {

	index := 1
	lp := len(pattern)

	if len(gg) == 0 {
		return true
	}
	// fmt.Println(pattern, gg)

	for _, g := range gg {
		// fmt.Println(index, g, pattern[index:index+g])
		// fmt.Println(index+g-1, lp-1, "C")

		if index+g-1 >= lp {
			return false
		}

		for strings.IndexRune(pattern[index:index+g], '.') > -1 {
			index++
			if index+g-1 >= lp {
				return false
			}
			// fmt.Println(index, "L")
		}

		index += (g - 1) + 1 + 1
		// fmt.Println(index, "E")
	}

	return true
}

func printCurrent(groups []int, pp []int, l int) {
	current := 0
	for i, p := range pp {
		for j := current; j < p; j++ {
			fmt.Print(".")
			current++
		}

		for j := 0; j < groups[i]; j++ {
			fmt.Print("#")
			current++
		}
	}
	for j := current; j < l; j++ {
		fmt.Print(".")
	}
	fmt.Println()
}

func possible(pattern string, g int, pos int) bool {
	if pos > 0 && (pattern[pos-1] == '#' || pattern[pos-1] == '$') {
		return false
	}

	to := pos + g - 1
	if to < len(pattern)-1 && (pattern[to+1] == '#' || pattern[to+1] == '$') {
		return false
	}

	for _, r := range pattern[pos : to+1] {
		if r == '.' || r == '$' {
			return false
		}
	}
	return true
}

type DataLine struct {
	Pattern string
	Groups  []int
	Starts  []int
	Ends    []int
}

func loadInput(filename string, five bool) []DataLine {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	datas := make([]DataLine, 0)

	for scanner.Scan() {
		line := scanner.Text()

		rawParts := strings.Split(line, " ")

		parts := []string{}

		if five {
			parts = []string{
				rawParts[0] + "?" + rawParts[0] + "?" + rawParts[0] + "?" + rawParts[0] + "?" + rawParts[0],
				rawParts[1] + "," + rawParts[1] + "," + rawParts[1] + "," + rawParts[1] + "," + rawParts[1],
			}
		} else {
			parts = rawParts
		}

		groups := parseInts(parts[1])

		starts := []int{}
		index := 0
		for i, g := range groups {
			if i == 0 {
				starts = append(starts, 0)
			} else {
				starts = append(starts, index)
			}
			index += g + 1
		}

		ends := []int{}
		sum := 0
		for i := len(groups) - 1; i >= 0; i-- {
			g := groups[i]
			if i == len(groups)-1 {
				ends = append(ends, len(parts[0])-1-(g-1))
			} else {
				// fmt.Println(parts[0], len(parts[0]), sum, g)
				ends = append([]int{len(parts[0]) - (sum + g)}, ends...)
			}
			sum += g + 1
		}

		// starts = append(starts, len(parts[0])-1-(g-1))

		temp := DataLine{
			Pattern: parts[0],
			Groups:  groups,
			Starts:  starts,
			Ends:    ends,
		}
		datas = append(datas, temp)
	}

	return datas
}

func parseInts(line string) []int {
	parts := strings.Split(line, ",")

	res := []int{}
	for _, part := range parts {
		if part != "" {
			n, _ := strconv.Atoi(part)
			res = append(res, n)
		}
	}
	return res
}

type GroupData struct {
	Val    int
	Starts []int
}

type DataLineV2 struct {
	Org     string
	Pattern string
	Groups  []GroupData
}

func loadInputV2(filename string, five bool) []DataLineV2 {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	datas := make([]DataLineV2, 0)

	for scanner.Scan() {
		line := scanner.Text()

		rawParts := strings.Split(line, " ")

		parts := []string{}

		if five {
			parts = []string{
				rawParts[0] + "?" + rawParts[0] + "?" + rawParts[0] + "?" + rawParts[0] + "?" + rawParts[0],
				rawParts[1] + "," + rawParts[1] + "," + rawParts[1] + "," + rawParts[1] + "," + rawParts[1],
			}
		} else {
			parts = rawParts
		}

		vals := parseInts(parts[1])
		groups := make([]GroupData, len(vals))

		for i, val := range vals {
			groups[i].Val = val
			groups[i].Starts = calcPositions(parts[0], vals[0:i], val, vals[i+1:])
		}

		temp := DataLineV2{
			Pattern: parts[0],
			Org:     parts[1],
			Groups:  groups,
		}
		datas = append(datas, temp)
	}

	return datas
}

func calcPositions(pattern string, before []int, g int, after []int) []int {
	// fmt.Println(pattern, before, g, after)

	start := 0
	for _, bf := range before {
		for noFitCondition(pattern, start, bf) {
			start += 1
		}
		start += (bf - 1) + 2
	}

	end := len(pattern) - 1 - (g - 1)

	if len(after) > 0 {
		end = len(pattern) - 1 - (after[len(after)-1] - 1)
	}

	for i := len(after) - 1; i >= 0; i-- {
		af := after[i]

		for noFitCondition(pattern, end, af) {
			end -= 1
		}

		if i != 0 {
			end -= (after[i-1] - 1) + 2
		} else {
			end -= (g - 1) + 2
		}
	}

	res := []int{}
	for i := start; i <= end; i++ {
		if strings.IndexRune(pattern[i:i+g], '.') == -1 {
			res = append(res, i)
		}
	}

	// fmt.Println(len(res), res, start, end)
	fmt.Print(len(res), " * ")

	return res
}

func noFitCondition(pattern string, start, val int) bool {
	body := pattern[start : start+(val-1)]
	return start > 0 && pattern[start-1] == '#' || strings.IndexRune(body, '.') > -1 || start+(val-1) < len(pattern)-1 && pattern[start+val] == '#'
}
