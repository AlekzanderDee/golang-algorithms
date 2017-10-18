//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

package main

import "fmt"

func twoSum(nums []int, target int) []int {

	elementsMap := map[int]int{}
	for ind, elm := range nums {
		complement := target - elm
		if val, ok := elementsMap[complement]; ok == true && val != ind {
			return []int{val, ind}
		}
		elementsMap[elm] = ind
	}
	return nil
}

func main() {
	type TestCase struct {
		InpNums        []int
		Target         int
		ExpectedResult []int
	}

	tests := []TestCase{
		{
			InpNums: []int{2, 7, 11, 15},
			Target: 9,
			ExpectedResult: []int{0, 1},
		},
		{
			InpNums: []int{5, 5},
			Target: 10,
			ExpectedResult: []int{0, 1},
		},
		{
			InpNums: []int{1, 5, 6, 12, 11},
			Target: 7,
			ExpectedResult: []int{0, 2},
		},
	}

	for _, test := range tests {
		res := twoSum(test.InpNums, test.Target)
		pass := true
		if len(res) != len(test.ExpectedResult) {
			pass = false
		}
		for ind := range res {
			if res[ind] != test.ExpectedResult[ind] {
				pass = false
			}
		}
		if pass == false {
			fmt.Printf("nums=%v; Tagret=%v; Expected result=%v; Fact result=%v\n",
				test.InpNums, test.Target, test.ExpectedResult, res)
		}

	}
	fmt.Print("finished")
}
