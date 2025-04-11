package linked_list

// ReverseLinkedList переворачивает связанный список
// Пример: (1 → 2 → 3 → 4) → (4 → 3 → 2 → 1)
func ReverseLinkedList(head *Node) *Node {

	var prev *Node = nil
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
