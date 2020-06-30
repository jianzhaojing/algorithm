package main

import "fmt"

type Hero struct {
	no         int
	name       string
	leftChild  *Hero
	rightChild *Hero
}

//list前序
func preList(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d,name=%s\n", node.no, node.name)
		preList(node.leftChild)
		preList(node.rightChild)
	}
}

//list中序
func middleList(node *Hero) {
	if node != nil {
		preList(node.leftChild)
		fmt.Printf("no=%d,name=%s\n", node.no, node.name)
		preList(node.rightChild)
	}
}

//后序
func nextList(node *Hero) {
	if node != nil {
		preList(node.leftChild)
		preList(node.rightChild)
		fmt.Printf("no=%d,name=%s\n", node.no, node.name)

	}
}

func main() {
	var root = &Hero{
		no:   1,
		name: "宋江",
	}

	var oneLeftTree = &Hero{
		no:   2,
		name: "卢俊义",
	}

	var oneRightTree = &Hero{
		no:   3,
		name: "吴用",
	}

	root.leftChild = oneLeftTree
	root.rightChild = oneRightTree

	var twoLeftTree = &Hero{
		no:   4,
		name: "武松",
	}

	var twoRightTree = &Hero{
		no:   5,
		name: "鲁智深",
	}

	oneLeftTree.leftChild = twoLeftTree
	oneLeftTree.rightChild = twoRightTree

	preList(root)
	fmt.Println()
	middleList(root)
	fmt.Println()
	nextList(root)
}
