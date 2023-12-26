package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part2() {
	points := loadInput("input-24-0.txt")
	// use SageMath
	// var('x y z vx vy vz t1 t2 t3')
	// eq1 = x+vx*t1==252878112005945+18*t1
	// eq2 = y+vy*t1==268091113812521-28*t1
	// eq3 = z+vz*t1==383680590897028-115*t1

	// eq4 = x+vx*t2==367457360760990-128*t2
	// eq5 = y+vy*t2==146734765528506+54*t2
	// eq6 = z+vz*t2==247459218958663+82*t2

	// eq7 = x+vx*t3==224753670596084+54*t3
	// eq8 = y+vy*t3==211427773389275+41*t3
	// eq9 = z+vz*t3==320164559684362-21*t3

	// solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9],x,y,z,vx,vy,vz,t1,t2,t3)

	// x == 309991770591665, y == 460585296453281, z == 234197928919588
	// vx == -63, vy == -301, vz == 97
	// t1 == 705106896120, t2 == 884086002605, t3 == 728530769193

	x0 := int64(309991770591665)
	vx0 := int64(-63)

	for i := 0; i < len(points); i++ {
		// x + vx*t = x0 + vx0*t
		// t * (vx - vx0) = x0 - x

		x := points[i].X
		vx := points[i].Vx

		if vx-vx0 == 0 {
		} else {
			t := (x0 - x) / (vx - vx0)
			fmt.Println(t)
		}
	}

	fmt.Println("Part-2:", 309991770591665+460585296453281+234197928919588)
}

func part1() {
	// points := loadInput("input-24-1.txt")
	// min := big.NewFloat(7)
	// max := big.NewFloat(27)

	points := loadInput("input-24-0.txt")
	min := big.NewFloat(200000000000000)
	max := big.NewFloat(400000000000000)

	result := map[string]int{}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			x, y, ex, past := intersectXY(p1.X, p1.Y, p1.Vx, p1.Vy, p2.X, p2.Y, p2.Vx, p2.Vy)
			fmt.Printf("%+v\n", points[i])
			fmt.Printf("%+v\n", points[j])

			fmt.Println(min.Cmp(x), max.Cmp(x))
			fmt.Println(min.Cmp(y), max.Cmp(y))

			if ex && !past && min.Cmp(x) < 1 && max.Cmp(x) > -1 && min.Cmp(y) < 1 && max.Cmp(y) > -1 {
				fmt.Print("In ")
				result["In"] += 1
			} else if ex && !past {
				fmt.Print("Out ")
				result["Out"] += 1
			} else if ex && past {
				fmt.Print("Past ")
				result["Past"] += 1
			} else {
				fmt.Print("? ")
				result["?"] += 1
			}
			fmt.Println(x, y, ex, past)
		}
	}

	fmt.Println("Part-1:", result)
}

func intersectXY(p1X_, p1Y_, p1Vx_, p1Vy_, p2X_, p2Y_, p2Vx_, p2Vy_ int64) (*big.Float, *big.Float, bool, bool) {
	if p1Vy_ == 0 && p2Vy_ == 0 || p1Vx_ == 0 && p2Vx_ == 0 {
		return big.NewFloat(0), big.NewFloat(0), false, false
	}

	p1Y := big.NewFloat(float64(p1Y_))
	p2Y := big.NewFloat(float64(p2Y_))

	p1X := big.NewFloat(float64(p1X_))
	p2X := big.NewFloat(float64(p2X_))

	p1Vy := big.NewFloat(float64(p1Vy_))
	p2Vy := big.NewFloat(float64(p2Vy_))

	p1Vx := big.NewFloat(float64(p1Vx_))
	p2Vx := big.NewFloat(float64(p2Vx_))

	temp1 := big.NewFloat(0)
	temp1.Mul(p1Vy, p2Vx)

	temp2 := big.NewFloat(0)
	temp2.Mul(p2Vy, p1Vx)

	A := big.NewFloat(0)
	A.Sub(temp1, temp2)
	// A = p1.Vy*p2.Vx - p2.Vy*p1.Vx

	if A.Sign() == 0 {
		return big.NewFloat(0), big.NewFloat(0), false, false
	}

	B := big.NewFloat(0)
	B.Mul(p1Vy, p2Vx)
	B.Mul(B, p2Y)
	// B := p1.Vy * p2.Vx * p2.Y

	C := big.NewFloat(0)
	C.Mul(p2Vy, p1Vx)
	C.Mul(C, p1Y)
	// C := p2.Vy * p1.Vx * p1.Y

	D := big.NewFloat(0)
	D.Mul(p2Vy, p1Vy)
	D.Mul(D, p1X)
	// D := p2.Vy * p1.Vy * p1.X

	E := big.NewFloat(0)
	E.Mul(p2Vy, p1Vy)
	E.Mul(E, p2X)
	// E := p2.Vy * p1.Vy * p2.X

	y := big.NewFloat(0)
	y.Sub(B, C)
	y.Add(y, D)
	y.Sub(y, E)
	y.Quo(y, A)
	// y := float64(B-C+D-E) / float64(A)

	x := big.NewFloat(0)
	if p1Vx_ == 0 {
		x.Sub(y, p2Y)
		x.Mul(x, p2Vx)
		x.Quo(x, p2Vy)
		x.Add(x, p2X)
		// x = (y-p2.Y)*p2.Vx/p2.Vy + p2.X
	} else {
		x.Sub(y, p1Y)
		x.Mul(x, p1Vx)
		x.Quo(x, p1Vy)
		x.Add(x, p1X)
		// x = (y-p1.Y)*p1.Vx/p1.Vy + p1.X
	}

	t1 := big.NewFloat(0)
	if p1Vx_ != 0 {
		t1.Sub(x, p1X)
		t1.Quo(t1, p1Vx)
		// t1 = (x - p1.X) / p1.Vx
	} else {
		t1.Sub(y, p1Y)
		t1.Quo(t1, p1Vy)
		// t1 = (y - p1.Y) / p1.Vy
	}

	t2 := big.NewFloat(0)
	if p2Vx_ != 0 {
		t2.Sub(x, p2X)
		t2.Quo(t2, p2Vx)
		// t2 = (x - p2.X) / p2.Vx
	} else {
		t2.Sub(y, p2Y)
		t2.Quo(t2, p2Vy)
		// t2 = (y - p2.Y) / p2.Vy
	}

	// fmt.Println("time", t1, t1.Sign(), t2, t2.Sign())

	return x, y, true, t1.Sign() == -1 || t2.Sign() == -1
}

type Point struct {
	X, Y, Z    int64
	Vx, Vy, Vz int64
}

func loadInput(fn string) []Point {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := make([]Point, 0)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " @ ")
		partsA := parseInts(parts[0])
		partsB := parseInts(parts[1])

		temp := Point{}

		temp.X = partsA[0]
		temp.Y = partsA[1]
		temp.Z = partsA[2]

		temp.Vx = partsB[0]
		temp.Vy = partsB[1]
		temp.Vz = partsB[2]

		res = append(res, temp)
	}

	return res
}

func parseInts(line string) []int64 {
	parts := strings.Split(line, ", ")

	res := []int64{}
	for _, part := range parts {
		if part != "" {
			n, _ := strconv.Atoi(strings.TrimSpace(part))
			res = append(res, int64(n))
		}
	}
	return res
}
