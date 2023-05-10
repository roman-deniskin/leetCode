package root_equals_sum_of_children

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	return root.Right.Val+root.Left.Val == root.Val
}
