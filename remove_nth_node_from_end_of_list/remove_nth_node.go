/*
Difficulty:Medium

Given a linked list, remove the nth node from the end of list and return its head.

For example,

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:
Given n will always be valid.
Try to do this in one pass.
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

type ListNode struct {
	Val  int
	Next *ListNode
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
	res := []int{ll.Val}
	for ll.Next != nil {
		ll = ll.Next
		res = append(res, ll.Val)
	}
	return res
}

// removeNthFromEnd removes the Nth node counting from the end of linked list.
// Algorithm creates 2 pointers with the gap size of N and moves traverses list
// until the second pointer reaches the end of the list.
// At this moment the first pointer will be pointing to the node before the Nth,
// so we just have to re-link the nodes
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := &ListNode{}
	fakeHead.Next = head

	firstPointer := fakeHead
	secondPointer := fakeHead
	for i := 0; i < n ; i++ {
		secondPointer = secondPointer.Next
	}

	for secondPointer.Next != nil {
		firstPointer = firstPointer.Next
		secondPointer = secondPointer.Next
	}

	firstPointer.Next = firstPointer.Next.Next

	return fakeHead.Next
}

func main() {
	type TestCase struct {
		InputList []int
		N int
		OutputList []int
	}

	tests := []TestCase{
		{
			InputList: []int{1,2,3,4,5},
			N: 1,
			OutputList: []int{1,2,3,4},
		},
		{
			InputList: []int{1,2,3,4,5},
			N: 2,
			OutputList: []int{1,2,3,5},
		},
		{
			InputList: []int{1,2,3},
			N: 3,
			OutputList: []int{2,3},
		},
		{
			InputList: []int{1,},
			N: 1,
			OutputList: []int{},
		},
	}

	for _, test := range tests {
		inputList := llFromSlice(test.InputList)
		res := removeNthFromEnd(inputList, test.N)

		outputList := sliceFromLl(res)
		assertEq(test.OutputList, outputList)
	}

	fmt.Println("Completed")
}
