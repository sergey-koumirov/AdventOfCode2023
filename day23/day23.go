package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// part1()
	part2()
}

var gNodes map[int]Node
var gDD [][]int

func part2() {
	field := loadInput("input-0.txt")

	firstKey := 0*1000 + 1
	lastKey := (len(field)-1)*1000 + (len(field[0]) - 2)

	field[0][1] = 'S'
	field[len(field)-1][len(field[0])-2] = 'F'

	gNodes, gDD = makeNodes(field)
	// printNodes(nodes)
	// printDistances(distances)

	max := 0
	used := map[int]bool{}
	deepCalc(firstKey, lastKey, 0, &max, used)

	fmt.Println("Part-2:")
}

func deepCalc(currKey int, lastKey int, dist int, maxRef *int, used map[int]bool) {
	node := gNodes[currKey]

	if currKey == lastKey {
		if *maxRef < dist {
			fmt.Println(dist)
			*maxRef = dist
		}
	} else {
		for _, k := range node.ChildKeys {
			u, _ := used[k]
			if !u {
				nextNode := gNodes[k]
				used[k] = true
				deepCalc(k, lastKey, dist+gDD[node.Index][nextNode.Index], maxRef, used)
				used[k] = false
			}
		}
	}
}

func printDistances(distances [][]int) {
	for _, dd := range distances {
		for _, d := range dd {
			fmt.Printf("%4d", d)
		}
		fmt.Println()
	}
}

func printNodes(nodes []Node) {
	for _, node := range nodes {
		fmt.Printf("%+v\n", node)
	}
}

type Node struct {
	Row, Col, Key, Index int
	ChildKeys            []int
}

func makeNodes(field Field) (map[int]Node, [][]int) {
	nodes := []Node{}

	nodes = append(
		nodes,
		Node{
			Row: 0,
			Col: 1,
			Key: 1,
		},
	)

	for r := 1; r < len(field)-1; r++ {
		for c := 1; c < len(field[0])-1; c++ {
			nexts := nextDirs(field, r, c, true)
			if field[r][c] != '#' && len(nexts) >= 3 {
				temp := Node{
					Row: r,
					Col: c,
					Key: r*1000 + c,
				}

				nodes = append(nodes, temp)
			}
		}
	}

	nodes = append(
		nodes,
		Node{
			Row:       len(field) - 1,
			Col:       len(field[0]) - 2,
			Key:       1000*(len(field)-1) + (len(field[0]) - 2),
			ChildKeys: []int{},
		},
	)

	distanses := make([][]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		distanses[i] = make([]int, len(nodes))
	}

	for i := 0; i < len(nodes); i++ {
		// fmt.Println(i, "-------------")

		node := nodes[i]
		nexts := nextDirs(field, node.Row, node.Col, true)

		someKeys := map[int]int{}

		for _, next := range nexts {
			newF := duplicate(field)
			newF[node.Row][node.Col] = 'O'

			newF[node.Row+next[0]][node.Col+next[1]] = 'O'

			data := goToCross(newF, node.Row+next[0], node.Col+next[1], true)

			key := 1000*data.Row + data.Col

			_, ex := someKeys[key]
			if ex {
				fmt.Println("key double", key)
			}
			someKeys[key] = 1

			// fmt.Println("key", next, node.Key, "->", key)

			nodes[i].ChildKeys = append(nodes[i].ChildKeys, key)

			for j, n := range nodes {
				if n.Key == 1000*data.Row+data.Col {
					if data.Dist > distanses[i][j] {
						distanses[i][j] = data.Dist
						distanses[j][i] = data.Dist
					}
				}
			}
		}
	}

	res := map[int]Node{}

	for i, node := range nodes {
		nodes[i].Index = i
		res[node.Key] = nodes[i]
	}

	return res, distanses

}

func part1() {
	field := loadInput("input-0.txt")
	field[0][1] = 'S'
	field[1][1] = 'O'
	field[len(field)-1][len(field[0])-2] = 'F'

	distance := 1
	deepWalk(field, 1, 1, 0, &distance, false)

	// print(field)

	fmt.Println("Part-1:", distance)
}

type Cache struct {
	Row, Col int
	Dist     int
	Cells    [][]int
	IsFinish bool
}

var cache = map[int]map[int]Cache{}

func deepWalk(f Field, row, col int, currDist int, distRef *int, useSlops bool) {
	data := goToCross(f, row, col, useSlops)

	// print(f)
	// fmt.Println(*distRef, currDist+dist)
	// fmt.Scanln()

	if data.IsFinish {
		if *distRef < currDist+data.Dist {
			*distRef = currDist + data.Dist
			fmt.Println(*distRef)
		}
	} else {
		for _, cell := range data.Cells {
			newF := duplicate(f)
			deepWalk(newF, cell[0], cell[1], currDist+data.Dist, distRef, useSlops)
		}
	}
}

var deltas = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func goToCross(f Field, row, col int, useSlops bool) Cache {
	res := Cache{
		Dist: 0,
	}

	// fmt.Println("goToCross: from", row, col)

	for {
		next := nextDirs(f, row, col, useSlops)
		res.Dist++

		if len(next) == 1 {
			row += next[0][0]
			col += next[0][1]

			// fmt.Println("goToCross:", row, col, string(f[row][col]))

			if f[row][col] == 'F' || f[row][col] == 'S' {
				res.Row = row
				res.Col = col
				res.IsFinish = true
				res.Cells = [][]int{{row, col}}
				res.Dist++
				return res
			} else {
				f[row][col] = 'O'
			}
		} else {
			// fmt.Println("goToCross *:", next)

			nextCells := make([][]int, len(next))

			for i := 0; i < len(next); i++ {
				nextCells[i] = make([]int, 2)
				nextCells[i][0] = row + next[i][0]
				nextCells[i][1] = col + next[i][1]

				if useSlops {
					f[nextCells[i][0]][nextCells[i][1]] = 'O'
				}
			}

			res.Row = row
			res.Col = col
			res.IsFinish = false
			res.Cells = nextCells

			return res
		}
	}
}

func nextDirs(f Field, row, col int, useSlops bool) [][]int {
	res := [][]int{}

	for _, dd := range deltas {
		aRow := row + dd[0]
		aCol := col + dd[1]

		// if row == 5 && col == 3 {
		// 	print(f)
		// 	fmt.Println("5:3", aRow >= 0 && aRow < len(f) && aCol >= 0 && aCol < len(f[0]))
		// }

		if aRow >= 0 && aRow < len(f) && aCol >= 0 && aCol < len(f[0]) {
			if f[aRow][aCol] == '.' || f[aRow][aCol] == 'F' {
				res = append(res, dd)
			}
			right := f[aRow][aCol] == '>' && (dd[0] == 0 && dd[1] == 1 || useSlops)
			left := f[aRow][aCol] == '<' && (dd[0] == 0 && dd[1] == -1 || useSlops)
			up := f[aRow][aCol] == '^' && (dd[0] == -1 && dd[1] == 0 || useSlops)
			down := f[aRow][aCol] == 'v' && (dd[0] == 1 && dd[1] == 0 || useSlops)

			if right || left || up || down {
				res = append(res, dd)
			}
		}
	}

	return res
}

func duplicate(f Field) Field {
	duplicate := make(Field, len(f))
	for i := range f {
		duplicate[i] = make([]rune, len(f[i]))
		copy(duplicate[i], f[i])
	}
	return duplicate
}

func print(f Field) {
	for row := 0; row < len(f); row++ {
		for col := 0; col < len(f[0]); col++ {
			fmt.Print(string(f[row][col]))
		}
		fmt.Println()
	}
	fmt.Println()
}

type Field [][]rune

func loadInput(fn string) Field {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := make(Field, 0)
	for scanner.Scan() {
		line := scanner.Text()

		temp := []rune{}

		for _, b := range line {
			temp = append(temp, rune(b))
		}

		res = append(res, temp)
	}

	return res
}
