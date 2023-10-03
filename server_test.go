package main

import (
	"fmt"
	"testing"
)

type TestTable[T int | float64] struct {
	a, b, result T
}

func TestSum(t *testing.T) {
	intTests := []TestTable[int]{
		{1, 2, 3},
		{4, 5, 9},
	}

	floatTests := []TestTable[float64]{
		{5.5, 3.5, 9.0},
		{4.0, 6.0, 10},
	}
	RunSumTests(t, intTests)
	RunSumTests(t, floatTests)

}

func RunSumTests[T int | float64](t *testing.T, tests []TestTable[T]) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v+%v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Sum(tt.a, tt.b)
			if ans != tt.result {
				t.Errorf("Wanted %v Got %v", tt.result, ans)
			}
		})
	}
}
