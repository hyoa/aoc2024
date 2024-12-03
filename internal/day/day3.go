package day

import (
	"fmt"
	"hyoa/aoc2024/internal/utils"
	"regexp"
)

func init() {
	DayCollection["3"] = Day3
}

type day3 struct {
	linesInput1 string
	linesInput2 string
}

func Day3(input1, input2 string) (any, any) {
	d := day3{
		linesInput1: utils.ReadTextFile(input1),
		linesInput2: utils.ReadTextFile(input2),
	}

	return d.step1(), d.step2()
}

func (d day3) step1() any {
	re := regexp.MustCompile(`(?m)(mul\(\d{1,3},\d{1,3}\))`)

	tt := 0
	for _, match := range re.FindAllString(d.linesInput1, -1) {
		var x, y int

		fmt.Sscanf(match, "mul(%d,%d)", &x, &y)
		tt += (x * y)
	}

	return tt
}

func (d day3) step2() any {
	re := regexp.MustCompile(`(?m)(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don\'t\(\))`)

	tt := 0
	valid := true
	for _, match := range re.FindAllString(d.linesInput2, -1) {
		switch match {
		case "don't()":
			valid = false
		case "do()":
			valid = true
		default:
			if valid {
				var x, y int

				fmt.Sscanf(match, "mul(%d,%d)", &x, &y)
				tt += (x * y)
			}
		}
	}

	return tt
}
