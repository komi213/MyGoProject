package main

import "fmt"
//结构体

//1.二叉树节点的数据结构定义如下：
type TreeNode struct{
	Value int
	Left,Right *TreeNode
}

func createTreeNode(value int) *TreeNode{
	return &TreeNode{Value:value}
}


func (node TreeNode) Print(){
	fmt.Println(node.Value, " ")
}

func (node *TreeNode)traverse(){
	if node != nil{
		//递归遍历左子树
		node.Left.traverse()
		node.Print()
		//递归遍历右子树
		node.Right.traverse()
	}
}

//4、结构体指针
//如果结构体的方法中需要对结构体成员的值进行修改，必须使用结构体指针作为方法的接收者
func (node *TreeNode)setValue(value int){
	if node!= nil{
		node.Value = value
	}else {
		fmt.Println("The node is nil.")
	}
}


func main() {
	var root TreeNode
	fmt.Println(root)

	root.Value = 0
	root.Left = &TreeNode{Value:1}
	root.Right = &TreeNode{2,nil,nil}
	root.Left.Left = &TreeNode{Value:3}
	root.traverse()

	root.setValue(100)
	root.traverse()

}