package main

func main() {
	root := new(TreeNode)
	rightSideView(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int
func rightSideView(root *TreeNode) []int {
	dfs(root,1)
	return res
}

func dfs(node *TreeNode, level int) {
	if node == nil{
		return
	}

	if len(res) < level{
		res = append(res,node.Val)
	}
	dfs(node.Right,level+1)
	dfs(node.Left,level+1)
}
