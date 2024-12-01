package day

import (
	"hyoa/aoc2024/internal/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

func init() {
	DayCollection["1"] = Day1
}

type day1 struct {
	list1, list2 []int
}

func Day1(input1, input2 string) (any, any) {
	lines := utils.ReadTextFileLinesAsString(input1)

	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for _, v := range lines {
		if v == "" {
			continue
		}

		s := strings.Split(v, "   ")

		v1, _ := strconv.Atoi(s[0])
		v2, _ := strconv.Atoi(s[1])

		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}

	d := day1{
		list1: l1,
		list2: l2,
	}

	return d.step1(), d.step2()
}

func (d day1) step1() any {
	sort.Slice(d.list1, func(i, j int) bool {
		return d.list1[i] < d.list1[j]
	})

	sort.Slice(d.list2, func(i, j int) bool {
		return d.list2[i] < d.list2[j]
	})

	tt := 0

	for i := 0; i < len(d.list1); i++ {
		v := math.Abs(float64(d.list1[i]) - float64(d.list2[i]))

		tt += int(v)
	}

	return tt
}

func (d day1) step2() any {
	lMapped := make(map[int]int)
	for _, v := range d.list2 {
		lMapped[v]++
	}

	tt := 0
	for _, v := range d.list1 {
		n, ok := lMapped[v]

		if ok {
			tt += v * n
		}
	}

	return tt
}
