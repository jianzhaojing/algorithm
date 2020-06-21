package main

import (
	"fmt"
	"errors"
	"os"
)

type Queue struct {
	length int //对列长度
	array [5]int //数组模拟队列,空余一个标识位
	first int //队列首
	end int //队列尾
}

//添加
func (this *Queue) add(value int) (err error){
	//是否队列已满
	if this.isFull() {
		return errors.New("queue is full");
	}
	
	this.array[this.end]=value
	this.end = (this.end +1) % this.length
	return
}

//获取
func (this *Queue) get() (val int, err error){
	if this.isEmpty() {
		return -1, errors.New("queue is empty");
	}

	val = this.array[this.first]
	this.first = (this.first+1)%this.length
	return 
}

//list 
func (this *Queue) list() {
	size := (this.end + this.length - this.first) % this.length
	// if this.end == this.first {
	// 	size = 0
	// } else if this.end > this.first {
	// 	size = this.end - this.first
	// } else {
	// 	size = this.end + this.length - this.first
	// }

	fmt.Println(size);
	if size==0 {
		fmt.Println("queeu is empty");
	}

	tmpFirst := this.first
	for i:=0;i<size;i++ {
		fmt.Printf("array[%d]=%d\t", tmpFirst,this.array[tmpFirst]);
		tmpFirst = (tmpFirst+1) % this.length
	}
	fmt.Println()
}
//判断对列是否已满
func (this *Queue) isFull() bool{
	return (this.end+1) % this.length == this.first;
}

//判断对列是否为空
func (this *Queue) isEmpty() bool{
	return (this.end % this.length) == (this.first % this.length) ;
}

func main() {
	queue := &Queue{
		length : 5,
		first : 0,
		end : 0,
	}

	var key string 
	var value int 

	for {
		fmt.Println("1 输入add 表示添加数据到队列")
		fmt.Println("2 输入get 表示添加数据到队列")
		fmt.Println("3 输入list 表示添加数据到队列")
		fmt.Println("4 输入exit 表示添加数据到队列")

		fmt.Scanln(&key)
		switch key {
			case "add" :
				fmt.Println("输入要入队列数值");
				fmt.Scanln(&value)
				err := queue.add(value)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("已加入对列")
				}
			case "list" : {
				queue.list();
			}
			case "get" : {
				val, err := queue.get();
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(val)
				}
			}
			case "exit" : {
				os.Exit(0)
			}
		}
	}
}