package main

import "fmt"

type Node struct {
	no       int
	name     string
	nickname string
	next     *Node
	pre      *Node
}

//尾部添加
func eddAdd(head, node *Node) {
	temp := head

	for {
		if temp.next == nil {
			break
		}

		temp = temp.next
	}

	temp.next = node
	node.pre = temp
	fmt.Println("添加成功")
}

//顺序添加
func orderAdd(head, node *Node) {
	temp := head

	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > node.no {
			flag = true
			break
		}

		temp = temp.next
	}

	if flag { //找到了
		tempNode := temp.next
		temp.next = node
		node.pre = temp
		node.next = tempNode
		tempNode.pre = node
	} else { //没找到
		temp.next = node
		node.pre = temp
	}

	fmt.Println()
}

//list展示（正向）
func list(head *Node) {
	temp := head

	if temp.next == nil {
		fmt.Println("这是一个空链表")
		return
	}

	for {
		fmt.Printf("no=%d,name=%s,nickname=%s=>\t", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next

		if temp.next == nil {
			break
		}
	}

	fmt.Println()
}

//list展示（反向）
func reverseList(head *Node) {
	temp := head

	if temp.next == nil {
		fmt.Println("这是一个空链表")
		return
	}

	for {
		if temp.next == nil {
			break
		}

		temp = temp.next
	}

	// fmt.Println(temp)
	for {
		fmt.Printf("no=%d,name=%s,nickname=%s=>\t", temp.no, temp.name, temp.nickname)
		temp = temp.pre

		if temp.pre == nil {
			break
		}
	}

	fmt.Println()
}

//修改
func update(head, node *Node) {
	temp := head

	for {
		if temp.next == nil {
			fmt.Println("未找到修改项")
			return
		}

		if temp.next.no == node.no {
			temp.next = node
			node.pre = temp
			if temp.next.next != nil {
				tempNode := temp.next.next
				node.next = tempNode
				tempNode.pre = node
			}

			break
		}

		temp = temp.next
	}

	fmt.Println("修改完成")
}

//删除
func delete(head *Node, no int) {
	temp := head

	for {
		if temp.next == nil {
			fmt.Println("未找到删除项")
			return
		}

		if temp.next.no == no {
			if temp.next.next != nil {
				tempNode := temp.next.next
				temp.next = tempNode
				tempNode.pre = temp
			} else {
				temp.next = nil
			}

			break
		}

		temp = temp.next
	}

	fmt.Println("删除完成")
}

func main() {
	songjiang := Node{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	linchong := Node{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	wusong := Node{
		no:       3,
		name:     "武松",
		nickname: "天伤星",
	}

	lujunyi := Node{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	//初始化
	var head = Node{}
	// list(&head)
	orderAdd(&head, &songjiang)
	// list(&head)
	orderAdd(&head, &linchong)
	// list(&head)
	orderAdd(&head, &lujunyi)
	list(&head)
	// reverseList(&head)
	update(&head, &wusong)
	list(&head)
	delete(&head, 2)
	list(&head)
}
