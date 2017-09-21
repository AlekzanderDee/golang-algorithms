package main

import (
	"fmt"
	"sort"
)

func quickSort(inpArr []int) []int {
	if len(inpArr) < 2 {
		return inpArr
	}

	// swapping the last and the mid elements in the slice
	// in order to prevent worse case of presorted input array
	inpArr[len(inpArr) - 1], inpArr[len(inpArr) / 2] = inpArr[len(inpArr) / 2], inpArr[len(inpArr) - 1]
	// let pivot element to be a last element of the slice
	pivot := inpArr[len(inpArr) - 1]

	var lessArr []int
	var greaterArr []int

	// don't process the last element as it is our pivot
	for i:=0; i < len(inpArr) - 1; i++ {
		if inpArr[i] <= pivot {
			lessArr = append(lessArr, inpArr[i])
		} else {
			greaterArr = append(greaterArr, inpArr[i])
		}
	}


	lessSorted:= quickSort(lessArr)
	greaterSorted := quickSort(greaterArr)

	// copying sorted slices and pivot element into the resulting slice
	// this code works faster then just appending slices into ech other
	result := make([]int, len(lessSorted) + len(greaterSorted) + 1)
	copy(result, lessSorted)
	result[len(lessSorted)] = pivot
	copy(result[len(lessSorted)+1:],greaterSorted )
	return result
}


func main() {
	sortCases := [][]int{
		{123, 434, -2, 232, 3, 55, 65, 42, 42, 4, 42, -100, 0, 33, 22, 989},
		{0},
		{},
	}

	for _, item := range sortCases {
		sortedArr := quickSort(item)
		fmt.Printf("Result: %v; Slice is sorted: %v \n", sortedArr, sort.IntsAreSorted(sortedArr))
	}
}
