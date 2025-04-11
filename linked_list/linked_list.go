package linked_list

import "fmt"

type Node struct {
	data int
	Next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (ll *LinkedList) AddNewHead(data int) {
	newNode := &Node{data: data, Next: ll.head}

	if ll.head == nil {
		ll.tail = newNode
	} else {
		newNode.Next = ll.head
	}

	ll.head = newNode
	ll.size++
}

func (ll *LinkedList) AddNewTail(data int) {
	newNode := &Node{data: data, Next: ll.head}

	if ll.tail == nil {
		ll.head = newNode
	} else {
		ll.tail.Next = newNode
	}

	ll.tail = newNode
	ll.size++
}

func (ll *LinkedList) Insert(after int, data int) {
	search := ll.head
	for search != nil {
		if search.data == after {
			break
		}
		search = search.Next
	}

	if search == nil {
		return
	}

	newNode := &Node{data: data, Next: nil}

	if search == ll.tail {
		ll.tail = newNode
	}

	newNode.Next = search.Next
	search.Next = newNode

	ll.size++
}

func (ll *LinkedList) Search(data int) *Node {
	search := ll.head

	for search != nil {
		if search.data == data {
			return search
		}
		search = search.Next
	}

	return nil
}

func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}
