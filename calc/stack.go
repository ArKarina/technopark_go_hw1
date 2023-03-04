package main

type node struct {
	data interface{}
	prev *node
}

type Stack struct {
	top  *node
	size int
}

func (stack *Stack) Push(elem interface{}) {
	n := &node{elem, stack.top}
	stack.top = n
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	if stack.size == 0 {
		return nil
	}

	n := stack.top
	stack.top = n.prev
	stack.size--
	return n.data
}

func (stack *Stack) Top() interface{} {
	if stack.size == 0 {
		return nil
	}

	return stack.top.data
}

func (stack *Stack) Size() (size int) {
	return stack.size
}

func (stack *Stack) Empty() bool {
	return stack.size == 0
}
