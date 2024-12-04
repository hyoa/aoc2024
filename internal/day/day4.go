package day

import (
	"hyoa/aoc2024/internal/utils"
	"strings"
)

func init() {
	DayCollection["4"] = Day4
}

type day4 struct {
	linesInput1 []string
}

func Day4(input1, input2 string) (any, any) {
	d := day4{
		linesInput1: utils.ReadTextFileLinesAsString(input1),
	}

	return d.step1(), d.step2()
}

func (d day4) step1() any {
	type pos struct {
		x, y int
	}

	it := make(map[pos]string)
	x := 0
	for _, v := range d.linesInput1 {
		y := 0

		ss := strings.Split(v, "")
		for _, vs := range ss {
			p := pos{
				x: x,
				y: y,
			}

			it[p] = vs
			y++
		}

		x++
	}

	tt := 0
	for k, v := range it {
		if v == "X" {
			check := make([]string, 0)

			check = append(check, it[k]+it[pos{k.x - 1, k.y}]+it[pos{k.x - 2, k.y}]+it[pos{k.x - 3, k.y}])
			check = append(check, it[k]+it[pos{k.x + 1, k.y}]+it[pos{k.x + 2, k.y}]+it[pos{k.x + 3, k.y}])
			check = append(check, it[k]+it[pos{k.x, k.y - 1}]+it[pos{k.x, k.y - 2}]+it[pos{k.x, k.y - 3}])
			check = append(check, it[k]+it[pos{k.x, k.y + 1}]+it[pos{k.x, k.y + 2}]+it[pos{k.x, k.y + 3}])

			check = append(check, it[k]+it[pos{k.x - 1, k.y - 1}]+it[pos{k.x - 2, k.y - 2}]+it[pos{k.x - 3, k.y - 3}])
			check = append(check, it[k]+it[pos{k.x - 1, k.y + 1}]+it[pos{k.x - 2, k.y + 2}]+it[pos{k.x - 3, k.y + 3}])
			check = append(check, it[k]+it[pos{k.x + 1, k.y - 1}]+it[pos{k.x + 2, k.y - 2}]+it[pos{k.x + 3, k.y - 3}])
			check = append(check, it[k]+it[pos{k.x + 1, k.y + 1}]+it[pos{k.x + 2, k.y + 2}]+it[pos{k.x + 3, k.y + 3}])

			for _, c := range check {
				if c == "XMAS" {
					tt++
				}
			}

		}
	}

	return tt
}

func (d day4) step2() any {
	type pos struct {
		x, y int
	}

	it := make(map[pos]string)
	x := 0
	for _, v := range d.linesInput1 {
		y := 0

		ss := strings.Split(v, "")
		for _, vs := range ss {
			p := pos{
				x: x,
				y: y,
			}

			it[p] = vs
			y++
		}

		x++
	}

	tt := 0
	for k, v := range it {
		if v == "A" {
			l1 := it[pos{k.x - 1, k.y - 1}] + it[k] + it[pos{k.x + 1, k.y + 1}]
			l2 := it[pos{k.x - 1, k.y + 1}] + it[k] + it[pos{k.x + 1, k.y - 1}]

			if (l1 == "MAS" || l1 == "SAM") && (l2 == "MAS" || l2 == "SAM") {
				tt++
			}

		}
	}

	return tt
}
