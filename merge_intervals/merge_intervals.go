//Difficulty:Medium
//
//Given a collection of intervals, merge all overlapping intervals.
//
//For example,
//Given [1,3],[2,6],[8,10],[15,18],
//return [1,6],[8,10],[15,18].

package main

import (
	"fmt"
	"reflect"
	"sort"
)

func assertEq(exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		fmt.Printf("Wanted %v; Got %v\n", exp, got)
		panic("Assertion error\n")
	}
}

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

type ByStart []Interval
func (s ByStart) Len() int {
	return len(s)
}
func (s ByStart) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByStart) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Sort(ByStart(intervals))
	result := []Interval{}
	curInterval := intervals[0]
	intervalsCnt := len(intervals)
	for ind := 1; ind < intervalsCnt; ind++ {
		if intervals[ind].Start <= curInterval.End {
			if curInterval.End < intervals[ind].End {
				curInterval.End = intervals[ind].End
			}

		} else {
			result = append(result, curInterval)
			curInterval = intervals[ind]
		}
	}
	result = append(result, curInterval)

	return result
}

func main() {
	type TestCase struct {
		Intervals      []Interval
		ExpectedResult []Interval
	}

	tests := []TestCase{
		{
			Intervals:      []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			ExpectedResult: []Interval{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			Intervals:      []Interval{{1, 2}, {2, 6}, {8, 10}, {15, 18}},
			ExpectedResult: []Interval{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			Intervals:      []Interval{{1, 3}, {5, 7}, {5, 10}, {15, 18}},
			ExpectedResult: []Interval{{1, 3}, {5, 10}, {15, 18}},
		},
		{
			Intervals:      []Interval{{1, 3},},
			ExpectedResult: []Interval{{1, 3},},
		},
		{
			Intervals:      []Interval{{1, 3}, {2, 6}, {8, 10}, {7, 18}},
			ExpectedResult: []Interval{{1, 6}, {7, 18}, },
		},
		{
			Intervals:      []Interval{{1, 4}, {0, 4},},
			ExpectedResult: []Interval{{0, 4}, },
		},
		{
			Intervals:      []Interval{{1, 3}, {2, 6}, {-1, 100}, {7, 18}},
			ExpectedResult: []Interval{{-1, 100}, },
		},
	}

	for _, test:= range tests {
		res := merge(test.Intervals)
		assertEq(test.ExpectedResult, res)
	}

}
