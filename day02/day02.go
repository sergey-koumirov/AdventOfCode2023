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

func part1() {
	batches := loadInput()

	sum := 0

	for _, record := range batches {
		possible := true

		for _, game := range record.Games {
			if !isPossible(game) {
				possible = false
			}
		}

		if possible {
			sum += record.Id
		}
	}

	fmt.Println("PART-1", sum)
}

func part2() {
	batches := loadInput()

	sum := 0

	for _, record := range batches {
		min := Game{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, game := range record.Games {
			for color, val := range game {
				if min[color] < val {
					min[color] = val
				}
			}
		}

		power := min["red"] * min["green"] * min["blue"]
		sum += power
	}

	fmt.Println("PART-2", sum)
}

func isPossible(game Game) bool {
	return game["red"] <= 12 && game["green"] <= 13 && game["blue"] <= 14
}

type Game map[string]int

type GamesArr []Game

type Record struct {
	Id    int
	Games GamesArr
}

func loadInput() []Record {
	file, _ := os.Open("input-0.txt")
	scanner := bufio.NewScanner(file)

	res := make([]Record, 0)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")

		leftParts := strings.Split(parts[0], " ")
		id, err := strconv.Atoi(leftParts[1])
		if err != nil {
			fmt.Println(line, err)
		}

		rightParts := strings.Split(parts[1], "; ")

		games := GamesArr{}

		for _, part := range rightParts {
			gameParts := strings.Split(part, ", ")

			game := Game{}
			for _, gamePart := range gameParts {
				values := strings.Split(gamePart, " ")
				num, err1 := strconv.Atoi(values[0])
				if err1 != nil {
					fmt.Println(values, err1)
				}
				game[values[1]] = num
			}
			games = append(games, game)
		}

		temp := Record{
			Id:    id,
			Games: games,
		}

		res = append(res, temp)
		// Game 16: 9 blue, 11 green; 8 green, 2 blue; 1 red, 6 green, 4 blue
	}

	return res
}
