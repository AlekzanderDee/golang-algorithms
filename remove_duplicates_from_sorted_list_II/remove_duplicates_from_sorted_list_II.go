/*
Difficulty:Medium

Given a sorted linked list, delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list.

For example,
Given 1->2->3->3->4->4->5, return 1->2->5.
Given 1->1->1->2->3, return 2->3.
*/

package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func assertEq(exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		fmt.Printf("Wanted %#v; Got %#v\n", exp, got)
		panic("Assertion error\n")
	}
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

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{}

	prevNode := newHead
	prevNode.Next = head
	curNode := head

	for curNode != nil {
		// move cur pointer to the next node while next element has the same value
		// at the end of the loop the curNode pointer points to the last duplicate
		for curNode.Next != nil && curNode.Val == curNode.Next.Val {
			curNode = curNode.Next
		}
		// if next element after the prevNode is curNode it means
		// that no duplicates were found (we had 0 loop iterations)
		// on prev step and we can just move our pointer one position right
		if prevNode.Next == curNode {
			prevNode = prevNode.Next
		} else {
			// if they don't match it means that duplicates were found during the search
			// and the curNode points to the last duplicate,
			// so in order to create next element we have to take its next node
			prevNode.Next = curNode.Next
		}
		// moving the curNode pointer
		curNode = curNode.Next
	}

	return newHead.Next
}

func main() {
	type TestCase struct {
		InpArray []int
		ExpectedResult []int
	}
	tests := []TestCase{
		{
			InpArray: []int{},
			ExpectedResult: []int{},
		},
		{
			InpArray: []int{1,2,3,4},
			ExpectedResult: []int{1,2,3,4},
		},
		{
			InpArray: []int{1,1,1},
			ExpectedResult: []int{},
		},
		{
			InpArray: []int{1,2,3,3,4,5,5,6},
			ExpectedResult: []int{1,2,4,6},
		},
		{
			InpArray: []int{1,1,1,2},
			ExpectedResult: []int{2},
		},
		{
			InpArray: []int{1,2,2,2},
			ExpectedResult: []int{1},
		},
	}

	for _, test := range tests {
		ll := llFromSlice(test.InpArray)
		res := deleteDuplicates(ll)
		assertEq(test.ExpectedResult, sliceFromLl(res))
	}
	fmt.Println("Completed")
}
