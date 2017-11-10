/*
Difficulty:Easy
Suppose Andy and Doris want to choose a restaurant for dinner, and they both have a list of favorite restaurants represented by strings.

You need to help them find out their common interest with the least list index sum. If there is a choice tie between answers, output all of them with no order requirement. You could assume there always exists an answer.

Example 1:

Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
Output: ["Shogun"]
Explanation: The only restaurant they both like is "Shogun".

Example 2:

Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["KFC", "Shogun", "Burger King"]
Output: ["Shogun"]
Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).

Note:

The length of both lists will be in the range of [1, 1000].
The length of strings in both lists will be in the range of [1, 30].
The index is starting from 0 to the list length minus 1.
No duplicates in both lists.
*/


package main

import (
	"reflect"
	"fmt"
)

func assertEq(exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		fmt.Printf("Wanted %#v; Got %#v\n", exp, got)
		panic("Assertion error\n")
	}
}

func findRestaurant(list1 []string, list2 []string) []string {
	res := []string{}
	minSumIndex := 2000

	for aInd, aVal := range list1 {
		if aInd > minSumIndex {
			return res
		}

		for bInd, bVal := range list2 {
			if aVal == bVal {
				if aInd + bInd < minSumIndex {
					res = []string{aVal}
					minSumIndex = aInd + bInd
				} else if aInd + bInd == minSumIndex {
					res = append(res, aVal)
				}
			}
		}
	}

	return res
}

func main() {
	type TestCase struct {
		List1 []string
		List2 []string
		Result []string
	}

	tests := []TestCase{
		{
			List1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			List2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			Result: []string{"Shogun"},
		},
		{
			List1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			List2: []string{"KFC", "Shogun", "Burger King"},
			Result: []string{"Shogun"},
		},
		{
			List1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			List2: []string{"Mac", "TacoBell"},
			Result: []string{},
		},
		{
			List1: []string{"Shogun","Tapioca Express","Burger King","KFC"},
			List2: []string{"KFC","Burger King","Tapioca Express","Shogun"},
			Result: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		},
	}

	for _, test := range tests {
		assertEq(test.Result, findRestaurant(test.List1, test.List2))
	}

	fmt.Println("Completed")
}
