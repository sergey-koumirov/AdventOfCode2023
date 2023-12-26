package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part2() {
	modules := loadInput("input-0.txt")

	modules["rx"] = &Module{
		Name:         "rx",
		Type:         "~",
		Destinations: []string{},
		Memory:       map[string]bool{},
		OnOff:        true,
	}

	cnt := 1
	for modules["rx"].OnOff {
		pushButton(modules, cnt)

		// printDebug(modules)

		// fmt.Scanln()

		cnt++
		if cnt%1_000_000 == 0 {
			fmt.Println(cnt)
		}

		if cnt > 20_000 {
			break
		}
	}

	fmt.Println("Part-2:", cnt)
}

func printDebug(modules map[string]*Module) {
	keys := []string{}
	for k, ref := range modules {
		if ref.Type == "%" || ref.Type == "&" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	for _, key := range keys {
		if modules[key].Type == "%" {
			fmt.Println(key, sbool(modules[key].OnOff))
		} else {
			fmt.Println(key, sboolArr(modules[key].Memory))
		}

	}
}

func sbool(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func sboolArr(bb map[string]bool) string {
	keys := []string{}
	for k := range bb {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	res := ""

	for _, key := range keys {
		res = res + sbool(bb[key])
	}

	return res
}

func part1() {
	modules := loadInput("input-0.txt")

	// for _, m := range modules {
	// 	fmt.Println(m)
	// }
	sumLow := 0
	sumHigh := 0

	for i := 0; i < 1000; i++ {
		cntHigh, cntLow := pushButton(modules, i)
		sumLow += cntLow
		sumHigh += cntHigh
	}

	fmt.Println("Part-1:", sumHigh*sumLow)
}

func pushButton(modules map[string]*Module, cnt int) (int, int) {
	order := []Pulse{{Source: "button", DestName: "broadcaster", Level: false}}

	cntLow := 0
	cntHigh := 0

	for len(order) > 0 {
		first := order[0]
		order = order[1:]

		if first.Level {
			cntHigh++
		} else {
			cntLow++
		}

		module, ex := modules[first.DestName]

		if first.DestName == "cl" && first.Level {
			fmt.Printf("%s -> %t -> %s  [%d]\n", first.Source, first.Level, first.DestName, cnt)
		}

		if !ex {
			fmt.Println(first.Source, " ", pname(first.Level), " ", first.DestName)
		} else {
			if module.Type == "broadcaster" {
				// fmt.Print(first.Source, " ", pname(first.Level), " broadcaster [")
				for _, name := range module.Destinations {
					order = append(order, Pulse{Source: module.Name, DestName: name, Level: first.Level})
					// fmt.Print(" ", name)
				}
				// fmt.Println("]")
			}

			if module.Type == "%" {
				// fmt.Print(first.Source, " ", pname(first.Level), " %", module.Name, " [")
				if !first.Level {
					module.OnOff = !module.OnOff
					for _, name := range module.Destinations {
						order = append(order, Pulse{Source: module.Name, DestName: name, Level: module.OnOff})
						// fmt.Print(" ", name)
					}
				}
				// fmt.Println("]")
			}

			if module.Type == "&" {
				// fmt.Print(first.Source, " ", pname(first.Level), " &", module.Name, " [")

				module.Memory[first.Source] = first.Level
				allHigh := true
				for _, v := range module.Memory {
					if !v {
						allHigh = false
						break
					}
				}

				for _, name := range module.Destinations {
					order = append(order, Pulse{Source: module.Name, DestName: name, Level: !allHigh})
					// fmt.Print(" ", name)
				}

				// fmt.Println("]")
			}

			if module.Type == "~" {
				if !first.Level {
					fmt.Println("~~~~~~~~~~~")
				}

				if module.OnOff && !first.Level {
					module.OnOff = false
				}
			}
		}

	}
	return cntHigh, cntLow
}

func pname(b bool) string {
	if b {
		return "-high->"
	}
	return "-low->"
}

type Module struct {
	Name         string
	Type         string // % & B
	Destinations []string
	Memory       map[string]bool
	OnOff        bool
}

type Pulse struct {
	Source   string
	DestName string
	Level    bool
}

func loadInput(fn string) map[string]*Module {
	file, _ := os.Open(fn)
	scanner := bufio.NewScanner(file)

	res := map[string]*Module{}

	for scanner.Scan() {
		line := scanner.Text()
		parts1 := strings.Split(line, " -> ")
		parts2 := strings.Split(parts1[1], ", ")

		var (
			n string
			t string
		)

		if parts1[0] == "broadcaster" {
			n = "broadcaster"
			t = "broadcaster"
		} else {
			n = parts1[0][1:]
			t = string(parts1[0][0])
		}

		temp := Module{
			Name:         n,
			Type:         t,
			Destinations: parts2,
			Memory:       map[string]bool{},
			OnOff:        false,
		}

		res[n] = &temp
	}

	// init memory
	for _, ref := range res {
		for _, destName := range ref.Destinations {
			if res[destName] != nil && res[destName].Type == "&" {
				res[destName].Memory[ref.Name] = false
			}
		}
	}

	return res
}

// broadcaster -> ct, hr, ft, qm
// %bh -> lj
// &zz -> th, hr, jk, bh, js
