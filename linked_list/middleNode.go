package linked_list

// MiddleNode возвращает середину списка
func MiddleNode(head *Node) *Node {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
