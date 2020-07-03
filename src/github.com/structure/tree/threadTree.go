package main

import "fmt"

type treeNode struct {
	l     *treeNode
	r     *treeNode
	p     *treeNode
	n     *treeNode
	value int
}

//根结点
type tree struct {
	root *treeNode
}

//创建树
func newTree() *tree {
	return &tree{root: nil}
}

//创建树结点
func newTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}

//添加树结点
func (t *tree) append(value int) {
	if t.root == nil {
		t.root = &treeNode{value: value}
		return
	}

	p := t.root
	for p != nil {
		if value == p.value {
			return
		}

		if value < p.value {
			if p.l == nil {
				p.l = newTreeNode(value)
				return
			}
			p = p.l
		} else {
			if p.r == nil {
				p.r = newTreeNode(value)
				return
			}
			p = p.r
		}
	}
}

//查找结点
func (t *tree) search(value int) (n, p *treeNode) {
	if t.root == nil {
		return nil, nil
	}

	n, p = t.root, nil
	for n != nil && n.value != value {
		p = n
		if value < n.value {
			n = n.l
		} else {
			n = n.r
		}
	}

	return n, p
}

//删除结点
func (t *tree) delete(value int) {
	c, p := t.search(value)
	if c == nil {
		return
	}

	if c.l == nil {
		if p == nil {
			t.root = c.r
		} else {
			t.swapChildNode(c.r, c, p)
		}
		return
	}
	if c.r == nil {
		if p == nil {
			t.root = c.l
		} else {
			t.swapChildNode(c.l, c, p)
		}
		return
	}

	s := t.dangleLeftMostNode(c.r, c)
	s.l, s.r = c.l, c.r
	if p == nil {
		t.root = s
	} else {
		t.swapChildNode(s, c, p)
	}
}

func (t *tree) dangleLeftMostNode(n, p *treeNode) *treeNode {
	for n != nil && n.l != nil {
		n, p = n.l, n
	}
	if n == p.l {
		p.l = n.r
		n.r = nil
	} else {
		p.r = n.r
		n.r = nil
	}
	return n
}

func (t *tree) printRecursiveLeftToRight() {
	t.printRecursiveLeftToRightImplement(t.root)
}

func (t *tree) printRecursiveLeftToRightImplement(n *treeNode) {
	if n == nil {
		return
	}
	t.printRecursiveLeftToRightImplement(n.l)
	fmt.Print(n.value, "\t")
	t.printRecursiveLeftToRightImplement(n.r)
}

func (t *tree) swapChildNode(n, c, p *treeNode) {
	if c == p.l {
		p.l = n
	} else {
		p.r = n
	}
}

var thread *treeNode

func (t *tree) threaded() {
	thread = nil
	t.threadImplement(t.root)
}

func (t *tree) printThreadedTree() {
	fmt.Print("threaded")
	h := t.root
	for h.l != nil {
		h = h.l
	}
	for h != nil {
		fmt.Print(" -> ", h.value)
		h = h.n
	}
}

func (t *tree) threadImplement(n *treeNode) {
	if n == nil {
		return
	}
	t.threadImplement(n.l)
	if n.p == nil {
		n.p = thread
	}
	if thread != nil && thread.n == nil {
		thread.n = n
	}
	thread = n
	t.threadImplement(n.r)
}

func main() {
	t := newTree()
	fmt.Println(t)
	t.append(10)
	t.printThreadedTree()
	fmt.Println()
	t.append(0)
	t.append(20)
	t.append(-10)
	t.append(5)
	t.append(25)
	t.append(30)
	t.append(15)
	t.delete(15)
	t.printRecursiveLeftToRight()
	t.threaded()
	t.printThreadedTree()
}
