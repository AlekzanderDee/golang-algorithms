package main

import (
	"errors"
	"fmt"
)

// binarySearch searches the val in the array and returns its position or the error if value was not found
func binarySearch(arr []int, val int) (ind int, err error) {

	var low int = 0
	var high int = len(arr) - 1

	for low <= high {
		var mid int = (low + high) / 2
		var guess int = arr[mid]

		if guess == val {
			return mid, nil
		} else if guess > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, errors.New(fmt.Sprintf("Element %v not found in the array", val))
}


func main() {
	type search_case struct{
		val int
		inpArr []int
		pos int
	}

	cases := []search_case{
		{
			val: 1,
			inpArr: []int{1, 21, 30, 42, 52, 62, 78, 89, 93},
			pos: 0,
		},
		{
			val: 21,
			inpArr: []int{1, 21, 30, 42, 52, 62, 78, 89, 93},
			pos: 1,
		},
		{
			val: 404,
			inpArr: []int{1, 21, 30, 42, 52, 62, 78, 89, 93},
			pos: -1,
		},
		{
			val: 62,
			inpArr: []int{1, 21, 30, 42, 52, 62, 78, 89, 93},
			pos: 5,
		},
		{
			val: 93,
			inpArr: []int{1, 21, 30, 42, 52, 62, 78, 89, 93},
			pos: 8,
		},
		{
			val: -404,
			inpArr: []int{1, 21, 30, 42, 52, 62, 78, 89, 93},
			pos: -1,
		},
	}

	for _, searchCase := range cases {
		pos, err := binarySearch(searchCase.inpArr, searchCase.val)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Element %v found at position %v (expected position %v)\n",
				searchCase.val, pos, searchCase.pos)
		}
	}
}
