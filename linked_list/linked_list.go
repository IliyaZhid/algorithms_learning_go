package linked_list

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (ll *LinkedList) AddNewHead(data int) {
	newNode := &Node{data: data, next: ll.head}

	if ll.head == nil {
		ll.tail = newNode
	} else {
		newNode.next = ll.head
	}

	ll.head = newNode
	ll.size++
}

func (ll *LinkedList) AddNewTail(data int) {
	newNode := &Node{data: data, next: ll.head}

	if ll.tail == nil {
		ll.head = newNode
	} else {
		ll.tail.next = newNode
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
		search = search.next
	}

	if search == nil {
		return
	}

	newNode := &Node{data: data, next: nil}

	if search == ll.tail {
		ll.tail = newNode
	}

	newNode.next = search.next
	search.next = newNode

	ll.size++
}

func (ll *LinkedList) Search(data int) *Node {
	search := ll.head

	for search != nil {
		if search.data == data {
			return search
		}
		search = search.next
	}

	return nil
}

func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}
