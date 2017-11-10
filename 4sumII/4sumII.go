/*
Difficulty:Medium

Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l)
there are such that A[i] + B[j] + C[k] + D[l] is zero.

To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500.
All integers are in the range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.

Example:

Input:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

Output:
2

Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/

package main

import (
	"reflect"
	"fmt"
)

func assertEq(exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		fmt.Printf("Wanted %v; Got %v\n", exp, got)
		panic("Assertion error\n")
	}
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0

	// map that shows how many A and B value combinations correspond to certain sum value
	sumAB := map[int]int{}
	// count sum for all A and B combinations
	for _, a := range A {
		for _, b := range B {
			sumAB[a+b]++
		}
	}

	// iterate thru all possible C and D value combinations
	for _, c := range C {
		for _, d := range D {
			// get the count of A and B sum values computed on the previous step with reversed sign (final sum should be equal 0)
			// that correspond to the sum of C and D
			l := sumAB[-1 * (c+d)]
			res += l
		}
	}

	return res
}

func main() {
	type TestCase struct {
		A         []int
		B         []int
		C         []int
		D         []int
		TuplesCnt int
	}

	tests := []TestCase{
		{
			A: []int{1, 2},
			B: []int{-2,-1},
			C: []int{-1, 2},
			D: []int{ 0, 2},
			TuplesCnt: 2,
		},
		{
			A: []int{-1, -1},
			B: []int{-1, 1},
			C: []int{-1, 1},
			D: []int{ 1, -1},
			TuplesCnt: 6,
		},
	}

	for _, test := range tests {
		assertEq(test.TuplesCnt, fourSumCount(test.A, test.B, test.C, test.D))
	}

	fmt.Println("Completed")
}
