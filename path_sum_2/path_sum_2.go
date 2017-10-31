/*Difficulty:Medium

Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
For example:
Given the below binary tree and sum = 22,

        5
       / \
      4   8
     /   / \
   11  13  4
  /  \    / \
7    2  5   1

return
[
  [5,4,11,2],
  [5,8,4,5]
]
*/

package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil || (root.Left == nil && root.Right == nil && root.Val != sum) {
		return nil
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return [][]int{{root.Val}}
	}

	result := [][]int{}

	leftRes := pathSum(root.Left, sum - root.Val)
    if leftRes != nil {
    	for _, res := range leftRes {
		    subResult := append([]int{root.Val}, res...)
		    result = append(result, subResult)
	    }
    }

	rightRes := pathSum(root.Right, sum - root.Val)
	if rightRes != nil {
		for _, res := range rightRes {
			subResult := append([]int{root.Val}, res...)
			result = append(result, subResult)
		}
	}

	if len(result) > 0 {
		return result
	}

	return nil
}

func main() {
	fmt.Printf("%v",
		pathSum(
			&TreeNode{
				Val:5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{Val: 7,},
						Right: &TreeNode{Val: 2,},
					},
				},
				Right: &TreeNode{
					Val:8,
					Left: &TreeNode{Val:13,},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{Val:5},
						Right: &TreeNode{Val:1},
					},
				},
			},
			22,
		),
	)
}
