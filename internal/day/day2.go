package day

import (
	"hyoa/aoc2024/internal/utils"
	"math"
	"strconv"
	"strings"
)

func init() {
	DayCollection["2"] = Day2
}

type day2 struct {
	reports [][]int
}

func Day2(input1, input2 string) (any, any) {
	reports := make([][]int, 0)

	lines := utils.ReadTextFileLinesAsString(input1)
	// INIT
	for _, v := range lines {
		report := make([]int, 0)
		splitted := strings.Split(v, " ")

		for _, vS := range splitted {
			n, _ := strconv.Atoi(vS)

			report = append(report, n)
		}

		reports = append(reports, report)
	}

	d := day2{
		reports: reports,
	}

	return d.step1(), d.step2()
}

func (d day2) step1() any {
	safeReport := 0

mainLoop:
	for _, r := range d.reports {
		increase := false
		for i := 0; i < len(r)-1; i++ {
			diff := int(math.Abs(float64(r[i]) - float64(r[i+1])))

			if diff > 3 || diff < 1 {
				continue mainLoop
			}

			if i == 0 {
				if r[i] < r[i+1] {
					increase = true
				}

				continue
			}

			if (increase && r[i] > r[i+1]) || (!increase && r[i] < r[i+1]) {
				continue mainLoop
			}
		}

		safeReport++
	}

	return safeReport
}

func (d day2) step2() any {
	safeReport := 0

	for _, r := range d.reports {
		safe, c := checkReport(r)
		for _, v := range c {
			safe, _ = checkReport(v)

			if safe {
				break
			}
		}

		if safe {
			safeReport++
		}
	}

	return safeReport
}

func checkReport(r []int) (bool, [][]int) {
	success := true
	increase := false
	cleanedReport := make([][]int, 0)

	faultyKeys := make([]int, 0)

	for i := 0; i < len(r)-1; i++ {
		diff := int(math.Abs(float64(r[i]) - float64(r[i+1])))

		if diff > 3 || diff < 1 {
			success = false
			faultyKeys = append(faultyKeys, i)
			faultyKeys = append(faultyKeys, i+1)

			if i > 0 {
				faultyKeys = append(faultyKeys, i-1)
			}

			break
		}

		if i == 0 {
			if r[i] < r[i+1] {
				increase = true
			}

			continue
		}

		if (increase && r[i] > r[i+1]) || (!increase && r[i] < r[i+1]) {
			success = false
			faultyKeys = append(faultyKeys, i)
			faultyKeys = append(faultyKeys, i+1)

			if i > 0 {
				faultyKeys = append(faultyKeys, i-1)
			}

			break
		}
	}

	if !success {
		for _, v := range faultyKeys {
			c := make([]int, 0)
			for k := range r {
				if k == v {
					continue
				}

				c = append(c, r[k])
			}

			cleanedReport = append(cleanedReport, c)
		}
	}

	return success, cleanedReport
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
