package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return constructBst(nums, 0, len(nums)-1)
}

func constructBst(nums []int, startIdx int, endIdx int) *TreeNode {
	if startIdx > endIdx {
		return nil
	}

	midIdx := (startIdx + endIdx) / 2
	bst := &TreeNode{Val: nums[midIdx]}
	bst.Left = constructBst(nums, startIdx, midIdx-1)
	bst.Right = constructBst(nums, midIdx+1, endIdx)

	return bst
}

type nodeInfo struct {
	node  *TreeNode
	level int
}

func printBST(root *TreeNode) [][]int {
	nodes := [][]int{}
	queue := []nodeInfo{{node: root, level: 0}}

	for len(queue) > 0 {
		currentNodeInfo := queue[0]
		queue = queue[1:]
		currentNode, currentLevel := currentNodeInfo.node, currentNodeInfo.level
		if len(nodes) == currentLevel {
			nodes = append(nodes, []int{currentNode.Val})
		} else {
			nodes[currentLevel] = append(nodes[currentLevel], currentNode.Val)
		}

		if currentNode.Left != nil {
			queue = append(queue, nodeInfo{node: currentNode.Left, level: currentLevel + 1})
		}
		if currentNode.Right != nil {
			queue = append(queue, nodeInfo{node: currentNode.Right, level: currentLevel + 1})
		}
	}
	return nodes
}

func bstToArray(bst *TreeNode) []int {
	order := []int{}
	inOrderTravesal(bst, &order)
	return order
}

func inOrderTravesal(node *TreeNode, order *[]int) {
	if node == nil {
		return
	}

	inOrderTravesal(node.Left, order)
	*order = append(*order, node.Val)
	inOrderTravesal(node.Right, order)
}

func main() {
	array := []int{-5, 1, 5, 8, 12, 16, 19, 22, 66}
	bst := sortedArrayToBST(array)
	fmt.Println("bst: ", printBST(bst))
	fmt.Println("sorted array: ", bstToArray(bst))
}

/*
output:
bst:  [[12] [1 19] [-5 5 16 22] [8 66]]
sorted array:  [-5 1 5 8 12 16 19 22 66]
*/
