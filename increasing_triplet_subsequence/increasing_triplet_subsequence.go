/*
Difficulty:Medium

Given an unsorted array return whether an increasing subsequence of length 3 exists or not in the array.

Formally the function should:

Return true if there exists i, j, k
such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.

Your algorithm should run in O(n) time complexity and O(1) space complexity.

Examples:
Given [1, 2, 3, 4, 5],
return true.

Given [5, 4, 3, 2, 1],
return false
*/

package main

import (
	"math"
	"reflect"
	"fmt"
)

func assertEq(exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		fmt.Printf("Wanted %#v; Got %#v\n", exp, got)
		panic("Assertion error\n")
	}
}

func increasingTriplet(nums []int) bool {
	numsLen := len(nums)
	if numsLen < 3 {
		return false
	}

	var c1 = int(math.MaxInt32)
	var c2 = int(math.MaxInt32)

	for _, element := range nums {
		if element <= c1 {
			c1 = element // c1 is min seen so far (it's a candidate for 1st element)
		} else if element <= c2 { // here when x > c1, i.e. x might be either c2 or c3
			c2 = element // element is better than the current c2, store it
		} else { // here when we have/had c1 < c2 already and x > c2
			return true // the increasing subsequence of 3 elements exists
		}
	}
	return false
}

func main() {
	type TestCase struct {
		Nums []int
		ExpectedOutput bool
	}

	tests := []TestCase{
		{
			Nums: []int{1,2,3,4},
			ExpectedOutput: true,
		},
		{
			Nums: []int{5,4,3,2,1},
			ExpectedOutput: false,
		},
		{
			Nums: []int{1,0,3,2},
			ExpectedOutput: false,
		},
		{
			Nums: []int{1,2},
			ExpectedOutput: false,
		},
	}

	for _, test := range tests {
		res := increasingTriplet(test.Nums)
		assertEq(test.ExpectedOutput, res)
	}

	fmt.Println("completed")
}
