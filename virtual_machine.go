package main

import "fmt"

var curByte = 0

func nextByte() int {
	curByte += 1
	return curByte
}

func run() {
	var (
		arg1 int
		arg2 int
		i    int
	)

next_op:
	switch nextByte() {
	case PUSH:
		push(nextByte())
		goto next_op
	case READ:
		push(getSym(nextByte()).val)
		goto next_op
	case WRITE:
		setSym(nextByte(), pop())
		goto next_op
	case ADD:
		arg1 = pop()
		arg2 = pop()
		push(arg1 + arg2)
		goto next_op
	case SUB:
		arg1 = pop()
		arg2 = pop()
		push(arg1 - arg2)
		goto next_op
	case MUL:
		arg1 = pop()
		arg2 = pop()
		push(arg1 * arg2)
		goto next_op
	case DIV:
		arg1 = pop()
		arg2 = pop()
		push(arg1 / arg2)
		goto next_op
	case RET:
		for i = 0; i < table_size; i++ {
			fmt.Printf("%s = %i\n", getSym(i).name, getSym(i).val)
		}

		return
	}
}
