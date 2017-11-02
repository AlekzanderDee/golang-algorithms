/*
Difficulty:Hard

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
*/

package main

import (
	"fmt"
	"container/heap"
	"reflect"
)

func assertEq(exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		fmt.Printf("Wanted %v; Got %v\n", exp, got)
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

func mergeKLists_OneByOne(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	newHead := &ListNode{}
	node := newHead
	for {
		movingInd := -1
		for ind, list := range lists {
			if list == nil {
				continue
			}

	        if movingInd == -1 {
			    movingInd = ind
		    } else if lists[movingInd].Val > list.Val {
			    movingInd = ind
		    }
	    }
	    if movingInd == -1 {
	    	break
	    }
	    node.Next = lists[movingInd]
		node = node.Next
		lists[movingInd] = lists[movingInd].Next
    }

	return newHead.Next
}

func mergeKListsRecur(lists []*ListNode) *ListNode {
	len := len(lists)
	if len == 0 {
		return nil
	} else if len == 1 {
		return lists[0]
	} else if len == 2 {
		newHead := &ListNode{}
		node := newHead
		for lists[0] != nil && lists[1] != nil {
			if lists[0].Val < lists[1].Val {
				node.Next = &ListNode{Val:lists[0].Val}
				node = node.Next
				lists[0] = lists[0].Next
			} else {
				node.Next = &ListNode{Val:lists[1].Val}
				node = node.Next
				lists[1] = lists[1].Next
			}
		}

		for lists[0] != nil {
			node.Next = &ListNode{Val:lists[0].Val}
			node = node.Next
			lists[0] = lists[0].Next
		}

		for lists[1] != nil {
			node.Next = &ListNode{Val:lists[1].Val}
			node = node.Next
			lists[1] = lists[1].Next
		}

		return newHead.Next
	}

	res := mergeKLists([]*ListNode{
		lists[0],
		mergeKLists(lists[1:]),
	})

	return res

}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Using built-in Heap
func mergeKLists(lists []*ListNode) *ListNode {
	newHead := &ListNode{}
	curNode := newHead

	h := &IntHeap{}
	heap.Init(h)

	for _, list := range lists {
		node := list
		for node != nil {
			heap.Push(h, node.Val)
			node = node.Next
		}
	}
	for h.Len() > 0 {
		curNode.Next = &ListNode{Val: heap.Pop(h).(int)}
		curNode = curNode.Next
	}
	return newHead.Next
}


func main() {
	type TestCase struct {
		Lists [][]int
		SortedList []int
	}

	tests := []TestCase{
		{
			Lists: [][]int{{1,2,5,8}, {0,},},
			SortedList: []int{0,1,2,5,8},
		},
		{
			Lists: [][]int{{1,2,2}, {1,1,2},},
			SortedList: []int{1,1,1,2,2,2},
		},
	}

	for _, test := range tests {
		lists := []*ListNode{}
		for _, list := range test.Lists {
			lists = append(lists, llFromSlice(list))
		}
		resLL := mergeKLists(lists)
		assertEq(test.SortedList, sliceFromLl(resLL))
	}

	fmt.Print("Completed")
}
