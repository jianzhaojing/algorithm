package main

import "fmt"

type Node struct {
	no   int
	name string
	next *Node
}

//添加
func insert(head, newNode *Node) {
	//判断是否为空
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head
		return
	}

	//非空
	temp := head
	//查找尾巴
	for {
		if temp.next == head {
			break
		}

		temp = temp.next
	}

	//在尾巴处添加
	temp.next = newNode
	newNode.next = head
}

//展示
func list(head *Node) {
	temp := head
	if temp.next == nil {
		fmt.Println("这是一个空循环链表")
	}

	for {
		fmt.Printf("id=%d,no=%s==>\n", temp.no, temp.name)

		if temp.next == head {
			break
		}

		temp = temp.next
	}
}

func main() {
	var head = &Node{}
	node1 := &Node{
		no:   1,
		name: "tom",
	}

	node2 := &Node{
		no:   2,
		name: "jack",
	}

	node3 := &Node{
		no:   3,
		name: "nini",
	}

	insert(head, node1)
	insert(head, node2)
	insert(head, node3)
	list(head)
}
