package main

import "fmt"

type Heap struct {
	Data map[int]int
	Num  int
}

//初始化堆
func initHeap() *Heap {
	return &Heap{Num: 0, Data: map[int]int{}}
}

//判断两个值的大小
func less(pre, next int) {

}

//交换两个值
func exch(pre, next int) {

}

//添加
func (this *Heap) add(data int) {
	this.Num++
	this.Data[this.Num] = data

	//上浮一下，使堆有序
	
	// swim(this.Data[this.Num])
}

//上浮
func swim() {

}

func main() {
	heap := initHeap()

	heap.add(4)
	fmt.Println(heap)

}
