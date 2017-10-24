// Difficulty: Medium
//
//Given preorder and inorder traversal of a tree, construct the binary tree.
//
//Note:
//You may assume that duplicates do not exist in the tree.
//
// Tree example:
//                                  a
//                              /      \
//                           b           i
//                        /    \       /   \
//                     c        d     j      k
//                   /   \    /  \
//                 e      f  g    h
//
//                                 ******************
// Traversals example:            \/                *
// index     0  1  2  3  4  5  6  7  8  9  10       *
// preorder: a  b  c  e  f  d  g  h  i  j  k        *
//           ^  |_________________|  |_____|        *
//           |         |                |           *
//         root    left branch     right branch     *
//                  preorder         preorder       *
//                                                  *
//                                 ******************
//                                \/
// index     0  1  2  3  4  5  6  7  8  9  10
//  inorder: e  c  f  b  g  d  h  a  j  i  k
//           |_________________|  ^  |_____|
//                     |          |     |
//                left branch   root  right branch
//                 inorder              inorder
//
// _Left_ branch preorder traversal ends at position = [index of root element in preorder] = 7

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

// Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func PreorderTaversal(node *TreeNode) []int {
	res := []int{}
	if node != nil {
		res = append(res, node.Val)
		res = append(res, PreorderTaversal(node.Left)...)
		res = append(res, PreorderTaversal(node.Right)...)
	}
	return res
}

func InorderTaversal(node *TreeNode) []int{
	res := []int{}
	if node != nil {
		res = append(res, InorderTaversal(node.Left)...)
		res = append(res, node.Val)
		res = append(res, InorderTaversal(node.Right)...)
	}
	return res
}

func constructTree() *TreeNode{
	aNode := TreeNode{Val:1,}
	bNode := TreeNode{Val:2,}
	cNode := TreeNode{Val:3,}
	dNode := TreeNode{Val:4,}
	eNode := TreeNode{Val:5,}
	fNode := TreeNode{Val:6,}
	gNode := TreeNode{Val:7,}
	hNode := TreeNode{Val:8,}
	iNode := TreeNode{Val:9,}
	jNode := TreeNode{Val:10,}
	kNode := TreeNode{Val:11,}
	lNode := TreeNode{Val:12,}
	mNode := TreeNode{Val:13,}
	nNode := TreeNode{Val:14,}
	oNode := TreeNode{Val:15,}
	pNode := TreeNode{Val:16,}
	qNode := TreeNode{Val:17,}
	rNode := TreeNode{Val:18,}
	sNode := TreeNode{Val:19,}
	tNode := TreeNode{Val:20,}
	uNode := TreeNode{Val:21,}
	vNode := TreeNode{Val:22,}
	wNode := TreeNode{Val:23,}
	xNode := TreeNode{Val:24,}
	yNode := TreeNode{Val:25,}

	aNode.Left = &bNode
	aNode.Right = &oNode

	bNode.Left = &cNode
	bNode.Right = &dNode

	cNode.Left = &eNode
	cNode.Right = &fNode

	dNode.Left = &gNode
	dNode.Right = &hNode

	gNode.Left = &iNode
	gNode.Right = &jNode

	iNode.Left = &mNode
	iNode.Right = &nNode

	jNode.Left = &vNode
	jNode.Right = &wNode

	vNode.Left = &xNode
	vNode.Right = & yNode

	hNode.Left = &kNode
	hNode.Right = &lNode

	oNode.Left = &pNode
	oNode.Right = &qNode

	pNode.Left = &rNode
	pNode.Right = &sNode

	qNode.Left = &tNode
	qNode.Right = &uNode

	return &aNode
}

func inorderRootIndex(inorder []int, value int) int {
	for ind, item := range inorder {
		if item == value {
			return ind
		}
	}
	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// first element of the preorder traversal is the Root element of the tree
	rootElement := preorder[0]
	// getting root's element index in the inorder traversal
	// inorder traversal is split by root element into left and right branches
	rootInorderIndex := inorderRootIndex(inorder, rootElement)
	if rootInorderIndex == -1 {
		return nil
	}
	// construct the tree node and construct its left and right branches recursively
	// (see the traversals example above)
	preorderLeftBranch := preorder[1:rootInorderIndex+1]
	inorderLeftBranch := inorder[:rootInorderIndex]
	preorderRightBranch := preorder[rootInorderIndex+1:]
	inorderRightBranch := inorder[rootInorderIndex+1:]
	node := TreeNode{
		// value is the root element
		Val:rootElement,
		Left: buildTree(preorderLeftBranch, inorderLeftBranch),
		Right: buildTree(preorderRightBranch, inorderRightBranch),
		}

	return &node
}

func main(){
	// manually construct the tree
	tree := constructTree()
	// generate tree traversals
	preorder := PreorderTaversal(tree)
	inorder := InorderTaversal(tree)
	fmt.Print("Preorder traversal: ", preorder)
	fmt.Print("\n\n Inorder traversal: ", inorder)

	// generate the tree from preorder and inorder traversals
	tree2 := buildTree(preorder, inorder)
	preorder2 := PreorderTaversal(tree2)
	inorder2 := InorderTaversal(tree2)
	fmt.Print("\n\nPreorder traversal 2 : ", preorder2)
	fmt.Print("\n\n Inorder traversal 2 : ", inorder2)
	// traversals should be the same
	assertEq(preorder, preorder2)
	assertEq(inorder, inorder2)
}
