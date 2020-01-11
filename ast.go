package main

const (
	SEQ_TYPE = iota
	SET_TYPE
	ADD_TYPE
	SUB_TYPE
	MUL_TYPE
	DIV_TYPE
	VAR_TYPE
	NUM_TYPE
	RET_TYPE
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
