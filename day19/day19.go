package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part2() {
	ii, _ := loadInput("input-0.txt")
	sum := 0

	ranges := Ranges{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}

	deep(ii, ranges, "in", 0, &sum)

	fmt.Println(ranges, ii["sg"])

	fmt.Println("Part-2:", sum)
}

func deep(ii Instructions, ranges Ranges, next string, index int, sumRef *int) {
	if next == "R" {
		return
	}
	if next == "A" {
		applySum(ranges, sumRef)
		return
	}

	ops := ii[next]
	op := ops[index]

	if op.Action == "goto" && op.Dest != "A" && op.Dest != "R" {
		newRanges := copyRanges(ranges)
		deep(ii, newRanges, op.Dest, 0, sumRef)
	} else if op.Action == "goto" && op.Dest == "A" {
		applySum(ranges, sumRef)
	} else if op.Action == "goto" && op.Dest == "R" {
		return
	} else if op.Action == "<" {
		r10 := -1
		r11 := -1
		r20 := -1
		r21 := -1

		if ranges[op.ParamName][1] < op.Num {
			r10 = ranges[op.ParamName][0]
			r11 = ranges[op.ParamName][1]
		}
		if op.Num <= ranges[op.ParamName][0] {
			r20 = ranges[op.ParamName][0]
			r21 = ranges[op.ParamName][1]
		}
		if ranges[op.ParamName][0] < op.Num && op.Num <= ranges[op.ParamName][1] {
			r10 = ranges[op.ParamName][0]
			r11 = op.Num - 1
			r20 = op.Num
			r21 = ranges[op.ParamName][1]
		}

		if r10 != -1 && r11 != -1 {
			newRanges := copyRanges(ranges)
			newRanges[op.ParamName][0] = r10
			newRanges[op.ParamName][1] = r11
			deep(ii, newRanges, op.Dest, 0, sumRef)
		}

		if r20 != -1 && r21 != -1 {
			newRanges := copyRanges(ranges)
			newRanges[op.ParamName][0] = r20
			newRanges[op.ParamName][1] = r21
			deep(ii, newRanges, next, index+1, sumRef)
		}
	} else if op.Action == ">" {
		r10 := -1
		r11 := -1
		r20 := -1
		r21 := -1

		if ranges[op.ParamName][1] <= op.Num {
			r10 = ranges[op.ParamName][0]
			r11 = ranges[op.ParamName][1]
		}
		if op.Num < ranges[op.ParamName][0] {
			r20 = ranges[op.ParamName][0]
			r21 = ranges[op.ParamName][1]
		}
		if ranges[op.ParamName][0] <= op.Num && op.Num < ranges[op.ParamName][1] {
			r10 = ranges[op.ParamName][0]
			r11 = op.Num
			r20 = op.Num + 1
			r21 = ranges[op.ParamName][1]
		}
		if r10 != -1 && r11 != -1 {
			newRanges := copyRanges(ranges)
			newRanges[op.ParamName][0] = r10
			newRanges[op.ParamName][1] = r11
			deep(ii, newRanges, next, index+1, sumRef)
		}
		if r20 != -1 && r21 != -1 {
			newRanges := copyRanges(ranges)
			newRanges[op.ParamName][0] = r20
			newRanges[op.ParamName][1] = r21
			deep(ii, newRanges, op.Dest, 0, sumRef)
		}
	}
}

func applySum(rr Ranges, sumRef *int) {
	*sumRef += (rr["x"][1] - rr["x"][0] + 1) * (rr["m"][1] - rr["m"][0] + 1) * (rr["a"][1] - rr["a"][0] + 1) * (rr["s"][1] - rr["s"][0] + 1)
}

func copyRanges(source Ranges) Ranges {
	return Ranges{
		"x": {source["x"][0], source["x"][1]},
		"m": {source["m"][0], source["m"][1]},
		"a": {source["a"][0], source["a"][1]},
		"s": {source["s"][0], source["s"][1]},
	}
}

type Ranges map[string][]int

func part1() {
	ii, pp := loadInput("input-0.txt")

	sum := 0

	for _, p := range pp {
		ops := ii["in"]
		next := run(ops, p)

		fmt.Print(p, "  in -> ", next)

		for next != "R" && next != "A" {
			ops = ii[next]
			next = run(ops, p)
			fmt.Print(" -> ", next)
		}

		if next == "A" {
			sum += p["x"] + p["m"] + p["a"] + p["s"]
		}
		fmt.Println()
	}

	fmt.Println("Part-1:", sum)
}

func run(ops []Operation, p Param) string {
	for _, op := range ops {
		if op.Action == "goto" {
			return op.Dest
		}

		if op.Action == "<" && p[op.ParamName] < op.Num {
			return op.Dest
		}

		if op.Action == ">" && p[op.ParamName] > op.Num {
			return op.Dest
		}
	}
	panic("run")
}

type Instructions map[string][]Operation

type Operation struct {
	ParamName string
	Action    string
	Num       int
	Dest      string
}

type Param map[string]int

type Params []Param

func loadInput(fn string) (Instructions, Params) {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	instructions := Instructions{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "{")

		instructionName := parts[0]
		actionsStr := parts[1][:len(parts[1])-1]
		actions := strings.Split(actionsStr, ",")

		instructions[instructionName] = []Operation{}

		for index, action := range actions {
			if index == len(actions)-1 {
				temp := Operation{
					ParamName: "-",
					Action:    "goto",
					Num:       0,
					Dest:      action,
				}
				instructions[instructionName] = append(instructions[instructionName], temp)
			} else {
				partsA := strings.Split(action, ":")

				num, _ := strconv.Atoi(partsA[0][2:])

				temp := Operation{
					ParamName: partsA[0][0:1],
					Action:    partsA[0][1:2],
					Num:       num,
					Dest:      partsA[1],
				}
				instructions[instructionName] = append(instructions[instructionName], temp)
			}
		}
	}

	params := Params{}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line[1:len(line)-1], ",")

		temp := Param{}
		for _, part := range parts {
			subParts := strings.Split(part, "=")
			temp[subParts[0]], _ = strconv.Atoi(subParts[1])
		}
		params = append(params, temp)
	}

	return instructions, params
}
