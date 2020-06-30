package main

import "fmt"

type CatNode struct {
	no   int
	next *CatNode
}

//创建链表
func createList(n int) *CatNode {
	var head = &CatNode{}
	var temp = &CatNode{}

	for i := 1; i <= n; i++ {
		node := &CatNode{
			no: i,
		}

		if i == 1 {
			head = node
			temp = node
			temp.next = head
		} else {
			temp.next = node
			temp = node
			temp.next = head
		}
	}

	return head
}

//打印
func list(head *CatNode) {
	temp := head

	if head.next == nil {
		fmt.Println("这是一个空链表")
	}

	for {
		fmt.Printf("cat的no=%d\n", temp.no)

		if temp.next == head {
			return
		}

		temp = temp.next
	}
}

//从m开始数数，n个一循环
func begin(head *CatNode, startNo, num int) {
	temp := head
	var count int
	var tail *CatNode

	for {
		count++

		if temp.next == head {
			break
		}
		temp = temp.next
	}

	//找到尾部
	tail = temp

	if startNo < 0 || startNo > count {
		fmt.Println("m不在链表范围中")
		return
	}

	//数startNo-1个数
	for i := 1; i < startNo; i++ {
		head = head.next
		tail = tail.next
	}
	// fmt.Println(head)
	// fmt.Println(tail)

	for {
		//开始循环n次
		for j := 1; j < num; j++ {
			head = head.next
			tail = tail.next
		}

		//提出读n的结点
		fmt.Println("小猫no=%d的出列", head.no)

		head = head.next
		tail.next = head

		if tail == head {
			break
		}
	}
	fmt.Println("小猫no=%d的出列", head.no)
}

func main() {
	//添加节点，n个
	n := 10
	head := createList(n)
	list(head)
	begin(head, 4, 2)
}
