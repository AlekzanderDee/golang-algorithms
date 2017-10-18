//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order
// and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8

package main

import "fmt"

//Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func LinkedListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return &ListNode{}
	}
	newHeadNode := &ListNode{}
	curNode := newHeadNode
	for _, elmn := range nums {
		curNode.Next = &ListNode{
			Val: elmn,
		}
		curNode = curNode.Next
	}
	return newHeadNode.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	var reminder int
	newHeadNode := &ListNode{}
	sumNode := newHeadNode
	var v1, v2 int
	v1 = l1.Val
	v2 = l2.Val
	for {
		digitsSum := v1 + v2 + reminder
		if digitsSum >= 10 {
			reminder = 1
			digitsSum = digitsSum - 10
		} else {
			reminder = 0
		}
		sumNode.Next = &ListNode{
			Val: digitsSum,
		}
		sumNode = sumNode.Next

		if l1.Next == nil && l2.Next == nil {
			break
		}

		if l1.Next != nil {
			l1 = l1.Next
			v1 = l1.Val
		} else {
			v1 = 0
		}

		if l2.Next != nil {
			l2 = l2.Next
			v2 = l2.Val
		} else {
			v2 = 0
		}
	}

	if reminder != 0 {
		sumNode.Next = &ListNode{
			Val: 1,
		}
	}

	return newHeadNode.Next
}

func main() {
	type TestCase struct {
		L1      *ListNode
		L2      *ListNode
		SumNode *ListNode
	}

	tests := []TestCase{
		{
			L1:      LinkedListFromSlice([]int{2, 4, 3}),
			L2:      LinkedListFromSlice([]int{5, 6, 4}),
			SumNode: LinkedListFromSlice([]int{7, 0, 8}),
		},
		{
			L1:      LinkedListFromSlice([]int{0}),
			L2:      LinkedListFromSlice([]int{5, 6, 4}),
			SumNode: LinkedListFromSlice([]int{5, 6, 4}),
		},
		{
			L1:      LinkedListFromSlice([]int{}),
			L2:      LinkedListFromSlice([]int{5, 6, 4}),
			SumNode: LinkedListFromSlice([]int{5, 6, 4}),
		},
		{
			L1:      LinkedListFromSlice([]int{2, 4, 7}),
			L2:      LinkedListFromSlice([]int{5, 6, 4}),
			SumNode: LinkedListFromSlice([]int{7, 0, 2, 1}),
		},
		{
			L1:      LinkedListFromSlice([]int{9, 8}),
			L2:      LinkedListFromSlice([]int{1}),
			SumNode: LinkedListFromSlice([]int{0, 9}),
		},
		{
			L1:      LinkedListFromSlice([]int{9, 9}),
			L2:      LinkedListFromSlice([]int{1}),
			SumNode: LinkedListFromSlice([]int{0, 0, 1}),
		},
		{
			L1:      LinkedListFromSlice([]int{8, 9, 9}),
			L2:      LinkedListFromSlice([]int{2}),
			SumNode: LinkedListFromSlice([]int{0, 0, 0, 1}),
		},
		{
			L1:      LinkedListFromSlice([]int{0, 8, 6, 5, 6, 8, 3, 5, 7}),
			L2:      LinkedListFromSlice([]int{6, 7, 8, 0, 8, 5, 8, 9, 7}),
			SumNode: LinkedListFromSlice([]int{6, 5, 5, 6, 4, 4, 2, 5, 5, 1}),
		},
	}

	for _, test := range tests {
		res := addTwoNumbers(test.L1, test.L2)
		node := test.SumNode
		for {
			if node.Val != res.Val {
				fmt.Printf("Error. Expected Val=%v; fact Val=%v\n", node.Val, res.Val)
			}

			if (res.Next == nil && node.Next != nil) || (res.Next != nil && node.Next == nil) {
				fmt.Print("Error. Different result lengths \n")
				return
			}

			if res.Next == nil || node.Next == nil {
				break
			}
			res = res.Next
			node = node.Next
		}
	}

	fmt.Print("Done")
}
