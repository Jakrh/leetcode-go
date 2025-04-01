package solution

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(preorder **[]int, inorderMap map[int]int, left, right int) *TreeNode {
	if len(**preorder) == 0 {
		return nil
	}

	val := (**preorder)[0]
	tmp := (**preorder)[1:]
	*preorder = &tmp

	idx, ok := inorderMap[val]
	if !ok {
		panic(fmt.Errorf("no such value by key %d", val))
	}

	root := &TreeNode{
		Val: val,
	}

	if idx != left {
		root.Left = build(preorder, inorderMap, left, idx-1)
	}
	if idx != right {
		root.Right = build(preorder, inorderMap, idx+1, right)
	}

	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := map[int]int{}
	for i, val := range inorder {
		inorderMap[val] = i
	}

	ptrPreorder := &preorder
	return build(&ptrPreorder, inorderMap, 0, len(inorder)-1)
}
