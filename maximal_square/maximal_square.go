/*
Difficulty:Medium

Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

For example, given the following matrix:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

Return 4.


Dynamic programming approach
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

const ONEByteCode = 49

func minIntSlice(nums []int) int {
	sliceLen := len(nums)
	if sliceLen == 0 {
		panic("empty int slice")
	}
	min := nums[0]
	for i := 1; i < sliceLen; i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	return min
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var maxBorderLen = 0
	rowCnt := len(matrix)
	colCnt := len(matrix[0])
	// form another matrix with the same dimensions for dynamic programming
	dp := make([][]int, rowCnt, rowCnt)
	for row := range dp {
		dp[row] = make([]int, colCnt, colCnt)
	}

	for rowInd := 0; rowInd < rowCnt; rowInd++ {
		for colInd := 0; colInd < colCnt; colInd++ {
			// elements in the first row and in the first column can form only square with border size = 1
			// so we just fill dp matrix with ONEs if we meet ONE on the position in the original matrix
			if rowInd == 0 || colInd == 0 {
				if matrix[rowInd][colInd] == ONEByteCode {
					dp[rowInd][colInd] = 1
					// if we meet ONE and the maxBorderLen hasn't been changed from the initial one so far,
					// then we say that we can build a square with border size = 1
					if maxBorderLen == 0 {
						maxBorderLen = 1
					}
				} else {
					dp[rowInd][colInd] = 0
				}

			} else {
				if matrix[rowInd][colInd] == ONEByteCode {
					// DP formula that checks that current cell
					// can be a right bottom corner of a larger square
					maxBorderFoFar := minIntSlice([]int{
						dp[rowInd-1][colInd],
						dp[rowInd-1][colInd-1],
						dp[rowInd][colInd-1],
					}) + 1

					dp[rowInd][colInd] = maxBorderFoFar
					// keeping maxBorderLen up to date
					if maxBorderFoFar > maxBorderLen {
						maxBorderLen = maxBorderFoFar
					}
				}
			}
		}
	}

	return maxBorderLen * maxBorderLen
}

func main() {

	type TestCase struct {
		Matrix        [][]byte
		MaximumSquare int
	}
	tests := []TestCase{
		{
			Matrix: [][]byte{
				[]byte("10100"),
				[]byte("10111"),
				[]byte("11111"),
				[]byte("10010"),
			},
			MaximumSquare: 4,
		},
		{
			Matrix: [][]byte{
				[]byte("11111111"),
				[]byte("11111110"),
				[]byte("11111110"),
				[]byte("11111000"),
				[]byte("01111000"),
			},
			MaximumSquare: 16,
		},
		{
			Matrix: [][]byte{
				[]byte("10100"),
				[]byte("10111"),
				[]byte("11111"),
				[]byte("10010"),
			},
			MaximumSquare: 4,
		},
		{
			Matrix: [][]byte{
				[]byte("10100"),
				[]byte("10111"),
				[]byte("11111"),
				[]byte("10111"),
			},
			MaximumSquare: 9,
		},
		{
			Matrix: [][]byte{
				[]byte("10"),
				[]byte("10"),
				[]byte("11"),
				[]byte("10"),
			},
			MaximumSquare: 1,
		},
		{
			Matrix:        [][]byte{},
			MaximumSquare: 0,
		},
		{
			Matrix: [][]byte{
				[]byte("10100"),
				[]byte("11111"),
				[]byte("11111"),
				[]byte("10110"),
			},
			MaximumSquare: 4,
		},
		{
			Matrix: [][]byte{
				[]byte("1"),
			},
			MaximumSquare: 1,
		},
		{
			Matrix: [][]byte{
				[]byte("0"),
			},
			MaximumSquare: 0,
		},
	}

	for _, test := range tests {
		sq := maximalSquare(test.Matrix)
		assertEq(test.MaximumSquare, sq)
	}
	fmt.Print("Completed")
}
