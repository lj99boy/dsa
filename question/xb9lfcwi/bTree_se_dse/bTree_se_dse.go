package bTree_se_dse

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }


type Codec struct {
	preTreeArr []string
}

func Constructor() Codec {
	return Codec{preTreeArr: make([]string,0)}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	//前序start
	str := preReverseSe(root)
	str = strings.TrimRight(str,",")
	return str
	//前序end
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	//前序start

	dataArr := strings.Split(data,",")
	this.preTreeArr = dataArr
	return this.preReverseDe()
	//前序end
}

func preReverseSe(root *TreeNode) string {
	if root == nil {
		return "#,"
	}

	res := strconv.Itoa(root.Val)+","

	res += preReverseSe(root.Left)
	res += preReverseSe(root.Right)


	return res
}

func (this *Codec)preReverseDe() *TreeNode {
	if len(this.preTreeArr) <1{
		return nil
	}
	val := this.preTreeArr[0]
	this.preTreeArr = this.preTreeArr[1:]
	if val == "#" {
		return nil
	}
	strVal,_ := strconv.Atoi(val)
	root := &TreeNode{Val:strVal}
	root.Left = this.preReverseDe()
	root.Right = this.preReverseDe()

	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */