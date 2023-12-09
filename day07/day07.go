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
	hands := loadInput2()

	sort.Sort(hands)

	// fmt.Println(hands)

	sum := 0
	for i, hand := range hands {
		// fmt.Println(hand.Original, hand.Power)
		sum += hand.Bid * (i + 1)
	}
	fmt.Println("Part-2:", sum)
}

func part1() {
	hands := loadInput()

	sort.Sort(hands)

	// fmt.Println(hands)

	sum := 0
	for i, hand := range hands {
		// fmt.Println(hand.Bid, i+1, hand)
		sum += hand.Bid * (i + 1)
	}

	fmt.Println("Part-1:", sum)
}

type Hand struct {
	Original, Power, Code string
	Bid                   int
}

type Hands []Hand

func (a Hands) Len() int      { return len(a) }
func (a Hands) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Hands) Less(i, j int) bool {
	return a[i].Code < a[j].Code
}

func handPowerJoker(h string) string {

	test := map[rune]int{}
	for _, r := range h {
		test[r] += 1
	}

	counts := [5]int{}

	cntJoker := test['J']

	for r, cnt := range test {
		if r != 'J' {
			counts[cnt-1]++
		}
	}

	if cntJoker == 5 {
		counts[4]++
	} else if cntJoker == 4 && counts[0] == 1 {
		counts[4]++
		counts[0]--
	} else if cntJoker == 3 && counts[1] == 1 {
		counts[4]++
		counts[1]--
	} else if cntJoker == 3 && counts[0] == 2 {
		counts[3]++
		counts[0]--
	} else if cntJoker == 2 && counts[2] == 1 {
		counts[4]++
		counts[2]--
	} else if cntJoker == 2 && counts[1] == 1 && counts[0] == 1 {
		counts[3]++
		counts[1]--
	} else if cntJoker == 2 && counts[0] == 3 {
		counts[2]++
		counts[0]--
	} else if cntJoker == 1 && counts[3] == 1 {
		counts[4]++
		counts[3]--
	} else if cntJoker == 1 && counts[2] == 1 {
		counts[3]++
		counts[2]--
	} else if cntJoker == 1 && counts[1] > 0 {
		counts[2]++
		counts[1]--
	} else if cntJoker == 1 && counts[0] == 4 {
		counts[1]++
		counts[0]--
	} else if cntJoker > 0 {
		fmt.Println("JOKER NOT USED", h, counts)
	}

	if counts[4] == 1 {
		return "7"
	}
	if counts[3] == 1 {
		return "6"
	}
	if counts[2] == 1 && counts[1] == 1 {
		return "5"
	}
	if counts[2] == 1 && counts[0] == 2 {
		return "4"
	}
	if counts[1] == 2 {
		return "3"
	}
	if counts[1] == 1 {
		return "2"
	}
	if counts[0] == 5 {
		return "1"
	}
	fmt.Println(h, "ERROR 0 !!!", counts)
	return "0"
}

func handPower(h string) string {

	test := map[rune]int{}
	for _, r := range h {
		test[r] += 1
	}

	counts := [5]int{}

	for _, cnt := range test {
		counts[cnt-1]++
	}

	if counts[4] == 1 {
		return "7"
	}
	if counts[3] == 1 {
		return "6"
	}
	if counts[2] == 1 && counts[1] == 1 {
		return "5"
	}
	if counts[2] == 1 && counts[0] == 2 {
		return "4"
	}
	if counts[1] == 2 {
		return "3"
	}
	if counts[1] == 1 {
		return "2"
	}
	if counts[0] == 5 {
		return "1"
	}
	fmt.Println(h, "ERROR 0 !!!", counts)
	return "0"
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2

var CardPowers = map[rune]string{
	'A': "O",
	'K': "N",
	'Q': "M",
	'J': "L",
	'T': "K",
	'9': "H",
	'8': "G",
	'7': "F",
	'6': "E",
	'5': "D",
	'4': "C",
	'3': "B",
	'2': "A",
}

var CardPowersJoker = map[rune]string{
	'A': "O",
	'K': "N",
	'Q': "M",
	'T': "L",
	'9': "K",
	'8': "H",
	'7': "G",
	'6': "F",
	'5': "E",
	'4': "D",
	'3': "C",
	'2': "B",
	'J': "A",
}

func handToCode(h string, abc map[rune]string) string {
	res := ""
	for _, r := range h {
		res += abc[r]
	}

	return res
}

func loadInput2() Hands {
	file, _ := os.Open("input-0.txt")
	scanner := bufio.NewScanner(file)

	res := make(Hands, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		bid, _ := strconv.Atoi(parts[1])
		power := handPowerJoker(parts[0])
		cards := handToCode(parts[0], CardPowersJoker)

		res = append(
			res,
			Hand{
				Code:     power + cards,
				Power:    power,
				Original: parts[0],
				Bid:      bid,
			},
		)
	}
	scanner.Scan()

	return res
}

func loadInput() Hands {
	file, _ := os.Open("input-0.txt")
	scanner := bufio.NewScanner(file)

	res := make(Hands, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		bid, _ := strconv.Atoi(parts[1])
		power := handPower(parts[0])
		cards := handToCode(parts[0], CardPowers)

		res = append(
			res,
			Hand{
				Code:     power + cards,
				Power:    power,
				Original: parts[0],
				Bid:      bid,
			},
		)
	}
	scanner.Scan()

	return res
}
