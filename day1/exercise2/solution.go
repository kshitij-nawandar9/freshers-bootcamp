//package exercise2

package main

import "fmt"

type tree struct {
	val string
	left *tree
	right *tree
}
func constructNode(val string) *tree{
	node:=tree{val,nil,nil}
	return &node
}

func preOrder(node *tree) {
	fmt.Printf(node.val)
	if node.left != nil{
		preOrder(node.left)
	}
	if node.right!=nil{
		preOrder(node.right)
	}
}

func postOrder(node *tree) {
	if node.left != nil{
		postOrder(node.left)
	}
	if node.right!=nil{
		postOrder(node.right)
	}
	fmt.Printf(node.val)
}

func main(){
	root:=constructNode("+")
	root.left=constructNode("a")
	root.right=constructNode("-")
	root.right.left=constructNode("b")
	root.right.right=constructNode("c")

	fmt.Println("Printing preorder traversal:")
	preOrder(root)
	fmt.Println()
	fmt.Println("Printing preorder traversal:")
	postOrder(root)
	fmt.Println()


}

