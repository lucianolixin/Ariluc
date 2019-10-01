package deleteDumplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func delDump(head *ListNode) *ListNode {
	currNode := head
	for currNode != nil && currNode.Next != nil {
		if currNode.Val == currNode.Next.Val {
			currNode.Next = currNode.Next.Next
		} else {
			currNode = currNode.Next
		}
	}

	return head
}
