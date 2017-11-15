/*
Difficulty:Hard

 Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

You may not alter the values in the nodes, only nodes itself may be changed.

Only constant memory is allowed.

For example,
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	fakeHead := &ListNode{}
	// pointer that points to the last processed node in initial list (last reversed)
	newEnd := fakeHead
	fakeHead.Next = head

	// first pointer will traverse thru the list
	firstPointer := fakeHead.Next
	// secondPointer will show where the group ends
	secondPointer := fakeHead.Next

	for secondPointer != nil {
		var i int
		// moving secondPointer to K positions
		for i = 0; i < k && secondPointer != nil; i++ {
			secondPointer = secondPointer.Next
		}
		// if we reached the end of the list before we have moved to the K positions, then break
		if secondPointer == nil && i < k {
			break
		}
		// create new head for a sublist which will be reverted
		revertedListHead := &ListNode{}
		// a node that represents a node in the reverted list
		var node *ListNode

		for firstPointer != secondPointer {
			// we will add new nodes to the beginning of the list,
			// so we have to save the remaining part first
			oldNext := revertedListHead.Next
			// the next node that will be added to the beginning of the reverted list
			// is the node at firstPointer position
			node = firstPointer
			// moving firstPointer before changing the node in order to avoid altering it
			firstPointer = firstPointer.Next
			// connect remaining part of the reverted list to the new node (head of reverted list)
			node.Next = oldNext
			// adding a new node to the beginning of the reverted list
			revertedListHead.Next = node
		}
		// moving node pointer to the end of the reverted list
		// in order to attach the remaining part of initial list later
		for node.Next != nil {
			node = node.Next
		}
		// attaching the remaining part of the initial list to which secondPointer is pointed
		node.Next = secondPointer

		// connecting new reversed list and prior part of already processed list
		newEnd.Next = revertedListHead.Next
		// update newEnd with the last processed node
		newEnd = node
	}

	return fakeHead.Next
}

func main() {
	type TestCase struct {
		InputList []int
		K int
		OutputList []int
	}

	tests := []TestCase{
		{
			InputList: []int{1,2,3,4,5},
			K: 2,
			OutputList: []int{2,1,4,3,5},
		},
		{
			InputList: []int{1,2,3,4,5},
			K: 3,
			OutputList: []int{3,2,1,4,5},
		},
		{
			InputList: []int{1,2,3,4,5},
			K: 1,
			OutputList: []int{1,2,3,4,5},
		},
		{
			InputList: []int{1,2,3,4,5},
			K: 4,
			OutputList: []int{4, 3, 2, 1, 5},
		},
		{
			InputList: []int{1,2,3,4,5},
			K: 5,
			OutputList: []int{5,4,3,2,1},
		},
	}

	for _, test := range tests {
		inputList := llFromSlice(test.InputList)
		res := reverseKGroup(inputList, test.K)

		outputList := sliceFromLl(res)
		assertEq(test.OutputList, outputList)
	}

	fmt.Println("Completed")
}
