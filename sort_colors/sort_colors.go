/*
Difficulty:Medium

Given an array with n objects colored red, white or blue, sort them so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note:
You are not suppose to use the library's sort function for this problem.

click to show follow up.

Follow up:
A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.

Could you come up with an one-pass algorithm using only constant space?
*/

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

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	redMarker := 0
	whiteMarker := 0
	blueMarker := len(nums) - 1

	for whiteMarker <= blueMarker {
		if nums[whiteMarker] == 0 {
			nums[redMarker], nums[whiteMarker] = nums[whiteMarker], nums[redMarker]
			redMarker++
			whiteMarker++
		} else if nums[whiteMarker] == 2 {
			nums[blueMarker], nums[whiteMarker] = nums[whiteMarker], nums[blueMarker]
			blueMarker--
		} else {
			whiteMarker++
		}
	}
}

func main() {
	type TestCase struct {
		InputNums      []int
		ExpectedOutput []int
	}
	tests := []TestCase{
		{
			InputNums:      []int{0, 0},
			ExpectedOutput: []int{0, 0},
		},
		{
			InputNums:      []int{1, 0},
			ExpectedOutput: []int{0, 1},
		},
		{
			InputNums:      []int{1, 1, 1, 0},
			ExpectedOutput: []int{0, 1, 1, 1},
		},
		{
			InputNums:      []int{2, 0},
			ExpectedOutput: []int{0, 2},
		},
		{
			InputNums:      []int{2, 2},
			ExpectedOutput: []int{2, 2},
		},
		{
			InputNums:      []int{2,},
			ExpectedOutput: []int{2,},
		},
		{
			InputNums:      []int{1, 2, 2, 2, 0, 0, 1, 0, 0, 2, 1},
			ExpectedOutput: []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 2},
		},
		{
			InputNums:      []int{},
			ExpectedOutput: []int{},
		},
	}

	for _, test := range tests {
		nums := test.InputNums
		sortColors(nums)
		assertEq(test.ExpectedOutput, nums)
	}

	fmt.Print("Completed")
}
