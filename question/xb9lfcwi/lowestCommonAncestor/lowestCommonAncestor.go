package lowestCommonAncestor

type TreeNode struct {
	val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left,p,q)
	r:=lowestCommonAncestor(root.Right,p,q)

	if l !=nil && r!=nil {
		return root
	}

	if l == nil && r==nil {
		return nil
	}

	if l==nil{
		return r
	}else{
		return l
	}

}
