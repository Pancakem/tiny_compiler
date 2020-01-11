package main

const MAXPROGRAMSIZE = 300

var (
	obj = make([]int, MAXPROGRAMSIZE)
	cur = 0
)

func gen(code int) {
	cur += 1
	obj[cur] = code
}

func compile(node *Node) {
	if node == nil {
		return
	}

	switch node.nodeType {
	case SEQ_TYPE:
		compile(node.op1)
		compile(node.op2)
	case SET_TYPE:
		compile(node.op2)
		gen(WRITE)
		gen(node.op1.val)
	case VAR_TYPE:
		gen(READ)
		gen(node.val)
	case NUM_TYPE:
		gen(PUSH)
		gen(node.val)
	case ADD_TYPE:
		compile(node.op1)
		compile(node.op2)
		gen(ADD)
	case SUB_TYPE:
		compile(node.op1)
		compile(node.op2)
		gen(SUB)
	case MUL_TYPE:
		compile(node.op1)
		compile(node.op2)
		gen(MUL)
	case DIV_TYPE:
		compile(node.op1)
		compile(node.op2)
		gen(DIV)
	case RET_TYPE:
		compile(node.op1)
		compile(node.op2)
		gen(RET)
	}
}
