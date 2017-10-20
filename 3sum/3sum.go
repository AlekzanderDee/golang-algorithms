//Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
//
//Note: The solution set must not contain duplicate triplets.
//
//For example, given array S = [-1, 0, 1, 2, -1, -4],
//
//A solution set is:
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

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

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	result := [][]int{}

	for aInd := 0; aInd < len(nums)-2; aInd++ {
		// if prev item's value is the same, and it was possible to make a sum with it then:
		//   1. it means that it was already done on the prev step
		//   2. if we use the same value again we will produce the same triplet at the end
		// so we skip it to avoid duplicates:
		if aInd > 0 && nums[aInd] == nums[aInd-1] {
			continue
		}
		bInd := aInd + 1
		cInd := len(nums) - 1
		for bInd < cInd {
			curSum := nums[aInd] + nums[bInd] + nums[cInd]
			if curSum == 0 {
				result = append(result, []int{nums[aInd], nums[bInd], nums[cInd]})
				// skip the same values as they will produce duplicated triplets:
				for bInd < cInd && nums[bInd] == nums[bInd+1] {
					bInd++
				}
				for bInd < cInd && nums[cInd] == nums[cInd-1] {
					cInd--
				}
				bInd++
				cInd--
			} else if curSum > 0 {
				cInd--
			} else {
				bInd++
			}
		}
	}
	return result
}

func main() {
	assertEq(
		[][]int{{-2, -2, 4}, {-2, -1, 3}, {-2, 0, 2}, {-2, 1, 1}, {-1, -1, 2}, {-1, 0, 1}},
		threeSum([]int{-2, -2, -1, -1, 0, 1, 1, 2, 3, 4}))

	assertEq(
		[][]int{{-2, -1, 3}, {-2, 0, 2}, {-2, 1, 1}, {-1, -1, 2}, {-1, 0, 1}},
		threeSum([]int{-2, -2, -1, -1, 0, 1, 1, 2, 3}))
	assertEq(
		[][]int{{0, 0, 0}},
		threeSum([]int{0, 0, 0, 0}))

	assertEq(
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
		threeSum([]int{-1, 0, 1, 2, -1, -4}))

	assertEq(
		[][]int{{-15, 1, 14}, {-15, 2, 13}, {-15, 3, 12}, {-15, 4, 11}, {-15, 5, 10}, {-15, 6, 9}, {-15, 7, 8}, {-14, 0, 14}, {-14, 1, 13}, {-14, 2, 12}, {-14, 3, 11}, {-14, 4, 10}, {-14, 5, 9}, {-14, 6, 8}, {-14, 7, 7}, {-13, -1, 14}, {-13, 0, 13}, {-13, 1, 12}, {-13, 2, 11}, {-13, 3, 10}, {-13, 4, 9}, {-13, 5, 8}, {-13, 6, 7}, {-12, -2, 14}, {-12, -1, 13}, {-12, 0, 12}, {-12, 1, 11}, {-12, 2, 10}, {-12, 3, 9}, {-12, 4, 8}, {-12, 5, 7}, {-12, 6, 6}, {-11, -3, 14}, {-11, -2, 13}, {-11, -1, 12}, {-11, 0, 11}, {-11, 1, 10}, {-11, 2, 9}, {-11, 3, 8}, {-11, 4, 7}, {-11, 5, 6}, {-10, -4, 14}, {-10, -3, 13}, {-10, -2, 12}, {-10, -1, 11}, {-10, 0, 10}, {-10, 1, 9}, {-10, 2, 8}, {-10, 3, 7}, {-10, 4, 6}, {-10, 5, 5}, {-9, -5, 14}, {-9, -4, 13}, {-9, -3, 12}, {-9, -2, 11}, {-9, -1, 10}, {-9, 0, 9}, {-9, 1, 8}, {-9, 2, 7}, {-9, 3, 6}, {-9, 4, 5}, {-8, -6, 14}, {-8, -5, 13}, {-8, -4, 12}, {-8, -3, 11}, {-8, -2, 10}, {-8, -1, 9}, {-8, 0, 8}, {-8, 1, 7}, {-8, 2, 6}, {-8, 3, 5}, {-8, 4, 4}, {-7, -7, 14}, {-7, -6, 13}, {-7, -5, 12}, {-7, -4, 11}, {-7, -3, 10}, {-7, -2, 9}, {-7, -1, 8}, {-7, 0, 7}, {-7, 1, 6}, {-7, 2, 5}, {-7, 3, 4}, {-6, -5, 11}, {-6, -4, 10}, {-6, -3, 9}, {-6, -2, 8}, {-6, -1, 7}, {-6, 0, 6}, {-6, 1, 5}, {-6, 2, 4}, {-6, 3, 3}, {-5, -5, 10}, {-5, -4, 9}, {-5, -3, 8}, {-5, -2, 7}, {-5, -1, 6}, {-5, 0, 5}, {-5, 1, 4}, {-5, 2, 3}, {-4, -4, 8}, {-4, -3, 7}, {-4, -2, 6}, {-4, -1, 5}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-3, -3, 6}, {-3, -2, 5}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -2, 4}, {-2, -1, 3}, {-2, 0, 2}, {-2, 1, 1}, {-1, -1, 2}, {-1, 0, 1}, {0, 0, 0}},
		threeSum([]int{0, 8, 2, -9, -14, 5, 2, -5, -5, -9, -1, 3, 1, -8, 0, -3, -12, 2, 11, 9, 13, -14, 2, -15, 4, 10, 9, 7, 14, -8, -2, -1, -15, -15, -2, 8, -3, 7, -12, 8, 6, 2, -12, -8, 1, -4, -3, 5, 13, -7, -1, 11, -13, 8, 4, 6, 3, -2, -2, 3, -2, 3, 9, -10, -4, -8, 14, 8, 7, 9, 1, -2, -3, 5, 5, 5, 8, 9, -5, 6, -12, 1, -5, 12, -6, 14, 3, 5, -11, 8, -7, 2, -12, 9, 8, -1, 9, -1, -7, 1, -7, 1, 14, -3, 13, -4, -12, 6, -9, -10, -10, -14, 7, 0, 13, 8, -9, 1, -2, -5, -14}))

	fmt.Print("Completed")
}
