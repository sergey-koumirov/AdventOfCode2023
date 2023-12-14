package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 0 - 7843 + 525152
	datas := loadInput("input-0.txt", true)

	sum := 0
	for i, dl := range datas {
		fmt.Println(dl.Pattern, dl.Groups)
		fmt.Println()

		cnt := deep(dl)

		fmt.Println(">>", i, cnt)
		fmt.Println()

		sum += cnt
	}

	fmt.Println("Part-2.2:", sum)
}

func deep(dl DataLine) int {
	res := 0

	if len(dl.Groups) > 2 {
		centerGroup := len(dl.Groups) / 2
		cg := dl.Groups[centerGroup]
		for _, start := range cg.Starts {
			leftStr := dl.Pattern[0 : start-1]
			leftGroup := dl.Groups[0:centerGroup]
			leftVals := make([]int, len(leftGroup))
			for i, g := range leftGroup {
				leftVals[i] = g.Val
			}
			leftDL := makeDataLine(leftStr, leftVals)
			leftCnt := deep(leftDL)

			rightStr := dl.Pattern[start+cg.Val+1:]
			rightGroup := dl.Groups[centerGroup+1:]
			rightVals := make([]int, len(rightGroup))
			for i, g := range rightGroup {
				rightVals[i] = g.Val
			}
			rightDL := makeDataLine(rightStr, rightVals)
			rightCnt := deep(rightDL)

			// fmt.Println(leftStr, cg.Val, rightStr, start)
			// fmt.Println(leftVals, cg.Val, rightVals)
			// fmt.Println("L", leftCnt, "R", rightCnt)
			// fmt.Println()

			res += leftCnt * rightCnt
		}
	} else {
		cnt := 0
		deepEnum(dl, []int{}, 0, dl.Pattern, &cnt)

		// fmt.Println("E", cnt, dl.Pattern, dl.Groups)
		res += cnt
	}

	return res
}

func deepEnum(dl DataLine, positions []int, index int, pattern string, cntRef *int) {
	val := dl.Groups[index].Val

	for _, start := range dl.Groups[index].Starts {
		if (index == 0 || index > 0 && start > positions[len(positions)-1]) && possible(pattern, val, start) && strings.IndexRune(pattern[0:start], '#') == -1 {

			next := append(positions, start)
			newPattern := pattern[0:start] + strings.Repeat("$", val) + pattern[start+val:]

			if index == len(dl.Groups)-1 {
				if strings.IndexRune(newPattern, '#') == -1 {
					// fmt.Println(newPattern)
					*cntRef++
				}
			} else {
				deepEnum(dl, next, index+1, newPattern, cntRef)
			}
		}
	}
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

type GroupData struct {
	Val    int
	Starts []int
}

type DataLine struct {
	Pattern string
	Groups  []GroupData
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

		vals := parseInts(parts[1])
		datas = append(datas, makeDataLine(parts[0], vals))
	}

	return datas
}

func makeDataLine(pattern string, vals []int) DataLine {
	groups := make([]GroupData, len(vals))

	for i, val := range vals {
		groups[i].Val = val
		groups[i].Starts = calcPositions(pattern, vals[0:i], val, vals[i+1:])
	}

	temp := DataLine{
		Pattern: pattern,
		Groups:  groups,
	}

	return temp
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
		if !noFitCondition(pattern, i, g) && strings.IndexRune(pattern[i:i+g], '.') == -1 {
			res = append(res, i)
		}
	}

	// fmt.Println(len(res), res, start, end)
	// fmt.Print(len(res), " * ")

	return res
}

func noFitCondition(pattern string, start, val int) bool {
	body := pattern[start : start+(val-1)]
	return start > 0 && pattern[start-1] == '#' || strings.IndexRune(body, '.') > -1 || start+(val-1) < len(pattern)-1 && pattern[start+val] == '#'
}
