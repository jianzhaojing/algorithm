package main

import (
	"fmt"
	"errors"
)

type Queue struct {
	length int //对列长度
	array [5]int //数组模拟队列
	first int //队列首
	end int //队列尾
}

//添加
func (this *Queue) addQueue(val int) (err error) {
	if this.end == this.length-1 {
		return errors.New("queue full")
	}

	this.end++
	this.array[this.end]=val
	return
}

//展示
func (this Queue) showList(){
	fmt.Println("队列当前的情况是：");

	for i:=this.first+1;i<=this.end;i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i]);
	}
	fmt.Println()
}

func main() {
	queue := &Queue{
		length : 5,
		first : -1,
		end : -1,
	}

	var key string 
	var value int 

	for {
		fmt.Println("1 输入add 表示添加数据到队列")
		fmt.Println("2 输入get 表示添加数据到队列")
		fmt.Println("3 输入show 表示添加数据到队列")
		fmt.Println("1 输入exit 表示添加数据到队列")

		fmt.Scanln(&key)
		switch key {
			case "add" :
				fmt.Println("输入要入队列数值");
				fmt.Scanln(&value)
				err := queue.addQueue(value)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("已加入对列")
				}
			case "show" : {
				queue.showList();
			}
		}
	}
}