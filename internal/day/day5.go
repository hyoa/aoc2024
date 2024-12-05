package day

import (
	"fmt"
	"hyoa/aoc2024/internal/utils"
	"slices"
	"strconv"
	"strings"
)

func init() {
	DayCollection["5"] = Day5
}

type day5 struct {
	rules   map[int][]int
	updates [][]int
}

func Day5(input1, input2 string) (any, any) {
	lines := utils.ReadTextFileLinesAsString(input1)

	idxUpdate := 0
	rules := make(map[int][]int)
	for k, v := range lines {
		if v == "" {
			idxUpdate = k
			break
		}

		var n, b int
		fmt.Sscanf(v, "%d|%d", &n, &b)

		if _, ok := rules[n]; ok {
			rules[n] = append(rules[n], b)
		} else {
			rules[n] = make([]int, 0)
			rules[n] = append(rules[n], b)
		}
	}

	updates := make([][]int, 0)
	for _, v := range lines[idxUpdate+1:] {
		s := strings.Split(v, ",")
		u := make([]int, 0)
		for _, vs := range s {
			n, _ := strconv.Atoi(vs)
			u = append(u, n)
		}

		updates = append(updates, u)
	}

	d := day5{
		updates: updates,
		rules:   rules,
	}

	return d.step1(), d.step2()
}

// ITERATION 1
// func (d day5) step1() any {
// 	valid := make([][]int, 0)
// main:
// 	for _, u := range d.updates {
// 		for k, n := range u {
// 			_, ok := d.rules[n]
// 			if ok && len(u[:k]) > 0 {
// 				for _, r1 := range d.rules[n] {
// 					for _, c := range u[:k] {
// 						if r1 == c {
// 							continue main
// 						}
// 					}
// 				}
// 			}
// 		}

// 		valid = append(valid, u)
// 	}

// 	tt := 0
// 	for _, u := range valid {
// 		m := (len(u) / 2)

// 		tt += u[m]
// 	}

// 	return tt
// }

// func (d day5) step2() any {
// 	invalid := make([][]int, 0)
// main:
// 	for _, u := range d.updates {
// 		for k, n := range u {
// 			_, ok := d.rules[n]
// 			if ok && len(u[:k]) > 0 {

// 				for _, r1 := range d.rules[n] {
// 					for _, c := range u[:k] {
// 						if r1 == c {
// 							invalid = append(invalid, u)
// 							continue main
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	tt := 0
// 	for _, u := range invalid {
// 		slices.SortFunc(u, func(a int, b int) int {
// 			v1, _ := d.rules[a]

// 			if slices.Contains(v1, b) {
// 				return -1
// 			}

// 			v2, _ := d.rules[b]
// 			if slices.Contains(v2, a) {
// 				return 1
// 			}

// 			return 0
// 		})

// 		m := (len(u) / 2)
// 		tt += u[m]
// 	}

// 	return tt
// }

// ITERATION 2 (Found after using sortFunc for step2 in iteration 1)
var sortUpdateDay5 = func(rules map[int][]int) func(a int, b int) int {
	return func(a, b int) int {
		v1, _ := rules[a]

		if slices.Contains(v1, b) {
			return -1
		}

		v2, _ := rules[b]
		if slices.Contains(v2, a) {
			return 1
		}

		return 0
	}
}

func (d day5) step1() any {
	tt := 0
	for _, u := range d.updates {
		ok := slices.IsSortedFunc(u, sortUpdateDay5(d.rules))

		if ok {
			m := (len(u) / 2)
			tt += u[m]
		}
	}

	return tt
}

func (d day5) step2() any {
	tt := 0
	for _, u := range d.updates {
		ok := slices.IsSortedFunc(u, sortUpdateDay5(d.rules))

		if !ok {
			slices.SortFunc(u, sortUpdateDay5(d.rules))

			m := (len(u) / 2)
			tt += u[m]
		}
	}

	return tt
}
