package main

import (
	"fmt"
)


type Node struct {
	no int
	name string
	nickname string
	next *Node
}

//尾部添加
func endAddNode(head, newNode *Node) {
	temp := head
	for {
		if temp.next == nil {
			break
		}

		temp = temp.next
	}

	temp.next = newNode
}

//按顺序添加
func addByOrder(head, node *Node) {
	temp := head

	for {
		if temp.next == nil {
			temp.next = node
			fmt.Println("添加完毕")
			break
		} else if temp.next.no > node.no {
			nextNode := temp.next
			temp.next = node
			node.next = nextNode
			fmt.Println("添加完毕")
			break
		} else if temp.next.no == node.no {
			fmt.Println("该节点已经存在，请检查")
			break
		} else {
			temp = temp.next
		}
	}
	
	fmt.Println()
}

//展示list
func showList(head *Node) {
	temp := head

	if temp.next == nil {
		fmt.Println("这是一个空链表");
	}

	for {
		fmt.Printf("%d %s %s==>\t", temp.next.no, temp.next.name, temp.next.nickname);
		temp = temp.next

		if temp.next == nil { 
			break
		}		
	}
	
	fmt.Println()
}

func deleteNode(head *Node, no int) {
	temp := head

	for {
		if temp.next == nil {
			fmt.Println("未找到要删除项");
			break
		} else if temp.next.no == no {
			temp.next = temp.next.next
			break
		} else {
			temp = temp.next
		}
	}

	fmt.Println("删除成功");
}

func main() {
	//初始化
	head := Node{}

	//添加
	songjiang := Node {
		no : 1,
		name : "宋江",
		nickname : "及时雨",
	}

	linchong := Node {
		no : 3,
		name : "林冲",
		nickname : "豹子头",
	}

	lujunyi := Node {
		no : 2,
		name : "卢俊义",
		nickname : "玉麒麟",
	}
	addByOrder(&head, &songjiang)
	showList(&head)
	addByOrder(&head, &linchong)
	showList(&head)
	addByOrder(&head, &lujunyi)
	showList(&head)
	// deleteNode(&head, 2)
	// showList(&head)
}