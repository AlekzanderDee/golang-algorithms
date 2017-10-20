//Sort a linked list in O(n log n) time using constant space complexity.

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

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, slowNode, fastNode := head, head, head
	// searching the middle of the linked list
	// as faster pointer is 2 times faster than slower,
	// it means when faster pointer reaches the end our slower pointer should be on the middle
	for fastNode != nil && fastNode.Next != nil {
		prev = slowNode
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}
	// cut linked list into two parts
	prev.Next = nil
	// sort
	sortedL1 := sortList(head)
	sortedL2 := sortList(slowNode)

	// create faked head for connection sorted nodes to
	newHead := &ListNode{}
	curNode := newHead

	// merge
	for true {
		if sortedL1.Val <= sortedL2.Val {
			curNode.Next = sortedL1
			sortedL1 = sortedL1.Next
		} else {
			curNode.Next = sortedL2
			sortedL2 = sortedL2.Next
		}
		curNode = curNode.Next

		if sortedL1 == nil || sortedL2 == nil {
			break
		}
	}

	if sortedL1 == nil {
		curNode.Next = sortedL2
	}

	if sortedL2 == nil {
		curNode.Next = sortedL1
	}

	return newHead.Next
}

func llFromSlice(nums []int) *ListNode {
	head := &ListNode{}
	curNode := head
	for _, item := range nums {
		curNode.Next = &ListNode{
			Val: item,
		}
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
	sl := []int{1,2,3,4,5,6}
	assertEq(
		[]int{1,2,3,4,5,6},
		sliceFromLl(sortList(llFromSlice(sl))),
	)

	sl = []int{}
	assertEq(
		[]int{},
		sliceFromLl(sortList(llFromSlice(sl))),
	)

	sl = []int{6,3,4,2,5,1}
	assertEq(
		[]int{1,2,3,4,5,6},
		sliceFromLl(sortList(llFromSlice(sl))),
	)

	sl = []int{1}
	assertEq(
		[]int{1},
		sliceFromLl(sortList(llFromSlice(sl))),
	)

	fmt.Print("Completed")
}
