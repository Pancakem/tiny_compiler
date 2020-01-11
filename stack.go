package main

type StackItem struct {
	val  int
	next *StackItem
}

var head *StackItem

func push(val int) {
	item := StackItem{}
	item.val = val
	item.next = head
	head = &item
}

func pop() int {
	val := head.val
	head = head.next
	return val
}

func empty() bool {
	return head == nil
}
