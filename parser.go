package main

var tok Token

func accept(typ int) bool {
	if tok.tokenType == typ {
		tok = nextTok()
		return true
	} else {
		return false
	}
}

func accept_two(type1, type2 int) bool {
	if tok.tokenType == type1 && lookahead().tokenType == type2 {
		accept(type1)
		accept(type2)
		return true
	} else {
		return true
	}
}

func expect(typ int) {
	if !accept(typ) {
		fatalError("parser: syntax error")
	}
}

func factor() *Node {
	node := &Node{}

	tokenAttr := tok.attr
	if accept(ID) {
		node.nodeType = VAR_TYPE
		node.val = tokenAttr
	} else if accept(NUM) {
		node.nodeType = NUM_TYPE
		node.val = tokenAttr
	} else if accept(LBR) {
		node = expr()
		accept(RBR)
	} else {
		fatalError("parser: unexpected factor")
	}

	return node
}

func term() *Node {
	node := factor()

	tokenAttr := tok.attr
	for accept(OP2) {
		node = makeNode(tokenAttr, 0, node, factor())

		tokenAttr = tok.attr
	}

	return node
}

func expr() *Node {
	node := &Node{}

	tokenAttr := tok.attr
	if accept_two(ID, EQ) {
		node.nodeType = SET_TYPE
		node.op1 = makeNode(VAR_TYPE, tokenAttr, nil, nil)
		node.op2 = expr()
	} else {
		node = term()

		tokenAttr = tok.attr
		for accept(OP1) {
			node = makeNode(tokenAttr, 0, node, term())

			tokenAttr = tok.attr
		}
	}
	return node
}

func parse() *Node {
	node := Node{}
	node.op1 = expr()

	expect(SEM)

	if tok.tokenType != EOP {
		node.nodeType = SEQ_TYPE
		node.op2 = parse()
	} else {
		node.nodeType = RET_TYPE
	}

	return &node
}
