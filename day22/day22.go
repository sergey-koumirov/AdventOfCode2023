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
	bricks := loadInput("input-0.txt")
	fallen := dropIt(bricks)
	sum := countDisDeep(fallen)
	fmt.Println("Part-2:", sum)
}

func countDisDeep(bricks Bricks) int {
	// fmt.Println(bricks)

	sum := 0
	for i := range bricks {
		test := deepCnt(bricks, i)
		fmt.Println(test)
		sum += test
	}

	return sum
}

func deepCnt(bricks Bricks, index int) int {

	touched := map[int]int{bricks[index].Code: 1}

	for {
		cntNew := 0
		for _, b := range bricks {

			_, brickTouched := touched[b.Code]

			if !brickTouched {
				cnt := 0
				for _, refBelow := range b.Below {
					_, ex := touched[refBelow.Code]
					if ex {
						cnt++
					}
				}

				if cnt == len(b.Below) && len(b.Below) > 0 {
					touched[b.Code] = 1
					cntNew++
				}
			}

		}

		if cntNew == 0 {
			break
		}
	}

	return len(touched) - 1
}

func part1() {
	bricks := loadInput("input-0.txt")
	fallen := dropIt(bricks)
	cnt := countDis(fallen)
	fmt.Println("Part-1:", cnt)
}

func countDis(bricks Bricks) int {
	cnt := 0
	for i := range bricks {
		canFall := true
		for _, ab := range bricks[i].Abowe {
			if len(ab.Below) == 1 {
				canFall = false
			}
		}

		if canFall {
			cnt++
		}
	}

	return cnt
}

func dropIt(bricks Bricks) Bricks {
	sort.Sort(bricks)

	fallen := Bricks{}

	for i1 := range bricks {
		if bricks[i1].Z1 > 1 {
			upperZ := 0

			for i2 := range fallen {
				if intersectXY(bricks[i1], fallen[i2]) {
					if upperZ < fallen[i2].Z2 {
						upperZ = fallen[i2].Z2
					}
				}
			}

			for i2 := 0; i2 < len(fallen); i2++ {
				if intersectXY(bricks[i1], fallen[i2]) && fallen[i2].Z2 == upperZ {
					bricks[i1].Below = append(bricks[i1].Below, &fallen[i2])
					fallen[i2].Abowe = append(fallen[i2].Abowe, &bricks[i1])
				}
			}

			if bricks[i1].Z1 == bricks[i1].Z2 {
				bricks[i1].Z1 = upperZ + 1
				bricks[i1].Z2 = upperZ + 1
			} else {
				delta := bricks[i1].Z2 - bricks[i1].Z1
				bricks[i1].Z1 = upperZ + 1
				bricks[i1].Z2 = upperZ + 1 + delta
			}

			fallen = append(fallen, bricks[i1])
		} else {
			fallen = append(fallen, bricks[i1])
		}
	}

	return fallen
}

func intersect(b1, b2 Brick) bool {
	ix := b1.X1 <= b2.X2 && b2.X1 <= b1.X2
	iy := b1.Y1 <= b2.Y2 && b2.Y1 <= b1.Y2
	iz := b1.Z1 <= b2.Z2 && b2.Z1 <= b1.Z2

	return ix && iy && iz
}

func intersectXY(b1, b2 Brick) bool {
	ix := b1.X1 <= b2.X2 && b2.X1 <= b1.X2
	iy := b1.Y1 <= b2.Y2 && b2.Y1 <= b1.Y2

	return ix && iy
}

type Brick struct {
	Code       int
	X1, Y1, Z1 int
	X2, Y2, Z2 int
	Abowe      []*Brick
	Below      []*Brick
}

type Bricks []Brick

func (a Bricks) Len() int           { return len(a) }
func (a Bricks) Less(i, j int) bool { return a[i].Z1 < a[j].Z1 }
func (a Bricks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func loadInput(fn string) Bricks {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := make(Bricks, 0)
	code := 50000
	for scanner.Scan() {
		line := scanner.Text()
		code++

		strs := strings.Split(line, "~")
		parts1 := strings.Split(strs[0], ",")
		parts2 := strings.Split(strs[1], ",")

		x1, _ := strconv.Atoi(parts1[0])
		y1, _ := strconv.Atoi(parts1[1])
		z1, _ := strconv.Atoi(parts1[2])

		x2, _ := strconv.Atoi(parts2[0])
		y2, _ := strconv.Atoi(parts2[1])
		z2, _ := strconv.Atoi(parts2[2])

		res = append(
			res,
			Brick{
				Code: code,
				X1:   x1, Y1: y1, Z1: z1,
				X2: x2, Y2: y2, Z2: z2,
				Abowe: []*Brick{},
				Below: []*Brick{},
			},
		)
	}

	return res
}
