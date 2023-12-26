package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
}

func part1() {
	names, nodes := loadInput("input-25-0.txt")

	fmt.Println(len(names))
	fmt.Println(len(nodes))

	nlen := len(names)
	half := nlen / 2

	magic := 3

	left := map[string]bool{}
	right := map[string]bool{}

	for i := range names {
		left[names[i]] = i >= half
		right[names[i]] = i < half
	}

	// fmt.Println(nodes)
	// fmt.Println(left)
	// fmt.Println(right)

	for {
		leftCnt := countLinks(left, nodes)

		actions := false

		// TODO do it once, recalc for changed
		hasExternal := map[string]bool{}
		for name, val := range left {
			if val {
				for _, n := range nodes[name] {
					if right[n] {
						hasExternal[name] = true
						break
					}
				}
			}
		}
		for name, val := range right {
			if val {
				for _, n := range nodes[name] {
					if left[n] {
						hasExternal[name] = true
						break
					}
				}
			}
		}

		// TDOD move all, do not break
		for name, val := range left {
			if val && hasExternal[name] {
				left[name] = false
				right[name] = true

				test := countLinks(left, nodes)
				if test >= magic && test < leftCnt {
					actions = true
					break
				}

				left[name] = true
				right[name] = false
			}
		}

		// TDOD move all, do not break
		for name, val := range right {
			if val && hasExternal[name] {
				left[name] = true
				right[name] = false

				test := countLinks(left, nodes)
				if test >= magic && test < leftCnt {
					actions = true
					break
				}

				left[name] = false
				right[name] = true
			}
		}

		for leftName, leftVal := range left {
			if leftVal && hasExternal[leftName] {
				for rightName, rightVal := range right {
					if rightVal && hasExternal[rightName] {
						left[leftName] = false
						right[leftName] = true
						left[rightName] = true
						right[rightName] = false

						test := countLinks(left, nodes)
						if test >= magic && test < leftCnt {
							actions = true
							break
						}

						left[leftName] = true
						right[leftName] = false
						left[rightName] = false
						right[rightName] = true
					}
				}
			}
		}

		fmt.Println("leftCnt", leftCnt)

		if !actions {
			break
		}
	}

	leftRes := 0
	rightRes := 0
	for i := range names {
		if left[names[i]] {
			leftRes++
		}
		if right[names[i]] {
			rightRes++
		}

	}

	fmt.Println("Part1: ", nlen, leftRes, rightRes, leftRes*rightRes)
}

func countLinks(half map[string]bool, nodes map[string][]string) int {
	res := 0

	for name, val := range half {
		if val {
			for _, chName := range nodes[name] {
				if !half[chName] {
					res++
				}
			}
		}
	}

	return res

}

func loadInput(fn string) ([]string, map[string][]string) {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	names := map[string]int{}
	links := map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")
		children := strings.Split(parts[1], " ")

		names[parts[0]] = 1

		for i := range children {

			v := children[i]

			names[v] = 1

			links[parts[0]+"-"+v] = 1
			links[v+"-"+parts[0]] = 1
		}
	}

	namesArr := []string{}

	for k := range names {
		namesArr = append(namesArr, k)
	}

	nodes := map[string][]string{}

	for k := range links {
		parts := strings.Split(k, "-")

		_, ex := nodes[parts[0]]
		if !ex {
			nodes[parts[0]] = []string{}
		}

		nodes[parts[0]] = append(nodes[parts[0]], parts[1])

	}

	return namesArr, nodes
}
