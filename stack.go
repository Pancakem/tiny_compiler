package main

type Stack struct {
	stack []int
	size  int
}

var stack *Stack

// will init stack
func init() {
	initStack()
}

func initStack() {
	internal := make([]int, 0)
	stack = &Stack{stack: internal, size: 0}
}

func push(val int) {
	stack.size += 1
	stack.stack = append(stack.stack, val)
}

func pop() int {
	stack.size -= 1
	val := stack.stack[stack.size]
	replacementSlice := make([]int, stack.size)
	copy(replacementSlice, stack.stack)
	stack.stack = replacementSlice
	return val
}

func empty() bool {
	return stack.size == 0
}
