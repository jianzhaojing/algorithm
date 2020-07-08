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

//判断前面的结点是否小于后面的结点
func (this *Heap) less(pre, next int) bool {
	if this.Data[pre] < this.Data[next] {
		return true
	}

	return false
}

//交换两个值
func (this *Heap) exch(pre, next int) {
	temp := this.Data[pre]
	this.Data[pre] = this.Data[next]
	this.Data[next] = temp
}

//添加
func (this *Heap) add(data int) {
	this.Num++
	this.Data[this.Num] = data

	//上浮一下，使堆有序
	this.swim(this.Num)
}

//上浮
func (this *Heap) swim(key int) {
	//不断比较当前结点和父节点的值，如果父节点的值小于当前结点，则交换位置
	for {
		if key <= 1 {
			break
		}

		if !this.less(key, key/2) {
			this.exch(key, key/2)
		}

		key = key / 2
	}
}

//下沉
func (this *Heap) down(key int) {
	//如果该结点小于任何一个自结点，选择两个子结点中较大的一个进行交换
	for {
		if this.Data[2*key] == 0 {
			break
		}

		lc := this.Data[2*key]
		rc := this.Data[2*key+1]

		if lc > rc && lc > this.Data[key] {
			this.exch(key, 2*key)
		} else if lc <= rc && rc > this.Data[key] {
			this.exch(key, 2*key+1)
		} else {
			break
		}
	}
}

//删除
func (this *Heap) deleteNode(key int) {
	this.Data[key] = this.Data[this.Num]
	delete(this.Data, this.Num)

	this.down(key)
}

func main() {
	heap := initHeap()

	heap.add(4)
	heap.add(6)
	heap.add(7)
	heap.add(3)
	heap.add(1)
	heap.add(14)
	fmt.Println(heap.Data)
	heap.deleteNode(3)
	fmt.Println(heap.Data)
}
