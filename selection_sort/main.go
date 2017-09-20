package main

import (
	"fmt"
	"sort"
)

// Implementation of the SelectionSort algorithm.
// Complexity O = n^2

// pop returns an element at the given position of the slice of ints and a new slice of ints without returned element
// order of the elements in slice is not preserved
func pop(inpArr []int, pos int) (int, []int, error) {
	// swapping last element and element at the given position in the slice
	inpArr[pos], inpArr[len(inpArr)-1] = inpArr[len(inpArr)-1], inpArr[pos]
	return inpArr[len(inpArr)-1], inpArr[:len(inpArr)-1], nil
}

// findMinElementIndex returns the position of the minimum element in the slice of ints
func findMinElementIndex(inpArr []int) (int, error) {
	minElement := inpArr[0]
	minxIndex := 0
	for ind, element := range inpArr {
		if element < minElement {
			minElement = element
			minxIndex = ind
		}
	}
	return minxIndex, nil
}

// selectionSort creates and returns a new slice of ints sorted in the ascending order from the input slice
func selectionSort(inpArr []int) (sortedArr []int, err error) {
	var nextElement, elementIndex int
	for range inpArr {
		elementIndex, err = findMinElementIndex(inpArr)
		if err != nil {
			return nil, err
		}

		nextElement, inpArr, err = pop(inpArr, elementIndex)
		if err != nil {
			return nil, err
		}

		sortedArr = append(sortedArr, nextElement)
	}

	return
}

func main() {
	sortCases := [][]int{
		{6, 2, 3, 8, 0, 1, 2, 4, 1},
		{6, 2, 3, 8, 0, 1, 2, 4, 1},
		{0, -2, 3, 12, 0, 1, 2, 4, 1},
		{6, 22, 3, -238, 0, 1, 2, 4, 1},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1},
		{-1, -1, -1, -1},
		{0},
		{},
	}
	for _, item := range sortCases {
		sortedArr, _ := selectionSort(item)
		fmt.Printf("Result: %v; Slice is sorted: %v \n", sortedArr, sort.IntsAreSorted(sortedArr))
	}
}
