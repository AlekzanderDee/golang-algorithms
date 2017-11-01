/*
Difficulty:Medium

Given a singly linked list L: L0 → L1 → … → Ln-1 → Ln,
reorder it to: L0 → Ln → L1 → Ln-1 → L2→ Ln-2 → …

You must do this in-place without altering the nodes' values.

For example,
Given {1,2,3,4}, reorder it to {1,4,2,3}.
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var llLen int
	// create list of all nodes for later easy access by the position index
	nodesList := []*ListNode{}
	node := head
	for node != nil {
		llLen++
		nodesList = append(nodesList, node)
		node = node.Next
	}

	newHead := head
	frontPosition := 1
	tailPosition := llLen - 1
	for i := 0; i < llLen; i++ {
		if i % 2 == 0 {
			newHead.Next = nodesList[tailPosition]
			tailPosition--
		} else {
			newHead.Next = nodesList[frontPosition]
			frontPosition++
		}

		newHead = newHead.Next
	}
	newHead.Next = nil
}

func llFromSlice(nums []int) *ListNode {
	head := &ListNode{}
	curNode := head
	for _, item := range nums {
		newNode := &ListNode{
			Val: item,
		}
		curNode.Next = newNode
		curNode = curNode.Next
	}
	return head.Next
}

func sliceFromLl(ll *ListNode) []int {
	if ll == nil {
		return []int{}
	}
	res := []int{ll.Val,}
	for ll.Next != nil {
		ll = ll.Next
		res = append(res, ll.Val)
	}
	return res
}

func main() {
	type TestCase struct {
		InputNums []int
		ExpectedOutputNums []int
	}

	tests := []TestCase{
		{
			InputNums: []int{},
			ExpectedOutputNums: []int{},
		},
		{
			InputNums: []int{1,},
			ExpectedOutputNums: []int{1,},
		},
		{
			InputNums: []int{1,2},
			ExpectedOutputNums: []int{1,2},
		},
		{
			InputNums: []int{1,2,3},
			ExpectedOutputNums: []int{1,3,2},
		},
		{
			InputNums: []int{1,2,3,4},
			ExpectedOutputNums: []int{1,4,2,3},
		},
		{
			InputNums: []int{1,2,3,4,5},
			ExpectedOutputNums: []int{1,5,2,4,3},
		},
		{
			InputNums: []int{1,2,3,4,5,6},
			ExpectedOutputNums: []int{1,6,2,5,3,4},
		},
		{
			InputNums: []int{1,2,3,4,5,6,7},
			ExpectedOutputNums: []int{1,7,2,6,3,5,4},
		},
	}
	for _, test := range tests {
		head := llFromSlice(test.InputNums)
		reorderList(head)
		assertEq(test.ExpectedOutputNums, sliceFromLl(head))
	}

	fmt.Print("Completed")
}
