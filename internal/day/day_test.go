package day

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type testCase struct {
	day    int
	r1, r2 any
}

var testCases []testCase

func TestMain(m *testing.M) {
	testCases = []testCase{
		{
			day: 1,
			r1:  11,
			r2:  31,
		},
		{
			day: 2,
			r1:  2,
			r2:  4,
		},
		{
			day: 3,
			r1:  161,
			r2:  48,
		},
		{
			day: 4,
			r1:  18,
			r2:  9,
		},
	}

	c := m.Run()

	os.Exit(c)
}

func TestAllAOC(t *testing.T) {
	days := DayCollection

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Day %d", tt.day), func(t *testing.T) {
			exec, ok := days[fmt.Sprintf("%d", tt.day)]

			if !ok {
				t.Errorf("no day defined")
			}

			o1, o2 := exec(fmt.Sprintf("../../input/day%d/example_1.txt", tt.day), fmt.Sprintf("../../input/day%d/example_2.txt", tt.day))

			if o1 != tt.r1 {
				t.Errorf("expected result 1 to be %+v, got %+v", tt.r1, o1)
			}

			if o2 != tt.r2 {
				t.Errorf("expected result 2 to be %+v, got %+v", tt.r2, o2)
			}
		})
	}
}

func TestUnitAOC(t *testing.T) {
	now := time.Now()

	if now.Day() > len(testCases) {
		t.Error("no test defined for this day")
	}

	testCase := testCases[now.Day()-1]

	days := DayCollection
	exec, ok := days[fmt.Sprintf("%d", now.Day())]

	if !ok {
		panic("Day not found")
	}

	o1, o2 := exec(fmt.Sprintf("../../input/day%d/example_1.txt", now.Day()), fmt.Sprintf("../../input/day%d/example_2.txt", now.Day()))

	if o1 != testCase.r1 {
		t.Errorf("expected result 1 to be %+v, got %+v", testCase.r1, o1)
	}

	if o2 != testCase.r2 {
		t.Errorf("expected result 2 to be %+v, got %+v", testCase.r2, o2)
	}
}
