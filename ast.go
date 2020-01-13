package main

const (
	SEQ_TYPE = 0
	SET_TYPE = 1
	ADD_TYPE = 2
	SUB_TYPE = 3
	MUL_TYPE = 4
	DIV_TYPE = 5
	VAR_TYPE = 6
	NUM_TYPE = 7
	RET_TYPE = 8
)

type Node struct {
	nodeType int
	val      int
	op1      *Node
	op2      *Node
}

func makeNode(nodeType, val int, op1 *Node, op2 *Node) *Node {
	newNode := Node{}

	newNode.nodeType = nodeType
	newNode.val = val
	newNode.op1 = op1
	newNode.op2 = op2
	return &newNode
}
