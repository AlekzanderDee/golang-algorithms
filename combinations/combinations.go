/* Difficulty: Medium

Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

For example,
If n = 4 and k = 2, a solution is:

 [
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
 ]*/

package main

import (
	"fmt"
	"reflect"
)

func assertEq(exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		fmt.Printf("Wanted %v; Got %v\n", exp, got)
		panic("Assertion error\n")
	}
}

func combine(n int, k int) [][]int {
	if n < k {
		// invalid inputs
		return nil
	} else if n == k {
		// in this case result would be a single sequence of elements 1..n
		res := []int{}
		for i := 1; i <= n; i++ {
			res = append(res, i)
		}
		return [][]int{res}
	}
	// form initial set of digits as a sequence with size of k
	row := []int{}
	for i := 1; i <= k; i++ {
		row = append(row, i)
	}
	// resulting slice
	res := [][]int{}
	// we start from the most right element
	position := k - 1

	for true {
		if row[position] <= n {
			// adding row that was created on the prev step
			res = append(res, append([]int(nil), row...))
			// if element at the position is less then maximum value, we can increase it
			row[position]++
		} else {
			// if element at the position is equal to the maximum value,
			// then we have to increase element on the previous position (if we have such)
			for position >= 0 {
				// decrease position and check if we can make a sequence
				// that will end with a value less than maximum value
				position--
				if position >= 0 && row[position]+(k-position) <= n {
					newStart := row[position] + 1
					for i := position; i < k; i++ {
						row[i] = newStart
						newStart++
					}
					break
				}
			}
			// in case we were not able to create a new sequence we exit the loop
			if position == -1 {
				break
			}
		}
		// reset position to the end
		position = k - 1
	}
	return res
}

func main() {
	type TestCase struct {
		N              int
		K              int
		ExpectedOutput [][]int
	}

	tests := []TestCase{
		{
			N:              4,
			K:              3,
			ExpectedOutput: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}},
		},
		{
			N:              5,
			K:              3,
			ExpectedOutput: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}},
		},
		{
			N:              4,
			K:              2,
			ExpectedOutput: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
	}
	for _, test := range tests {
		res := combine(test.N, test.K)
		assertEq(test.ExpectedOutput, res)
	}

	fmt.Print("Completed")
}

//Input:4
//3
//Output:[[1,2,3],[1,2,4],[2,3,4]]
//Expected:[[1,2,3],[1,2,4],[1,3,4],[2,3,4]]

//Input:5
//3
//Output:  [[1,2,3],[1,2,4],[1,2,5],[1,3,4],[1,3,5],[2,3,4],[2,3,5],[3,4,5]]
//Expected:[[1,2,3],[1,2,4],[1,2,5],[1,3,4],[1,3,5],[1,4,5],[2,3,4],[2,3,5],[2,4,5],[3,4,5]]

//Your input
//
//4
//2
//
//Your answer [[1,2],[1,3],[1,4]]
//
//Expected answer [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
