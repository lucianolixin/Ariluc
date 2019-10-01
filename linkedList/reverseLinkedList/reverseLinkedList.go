package reverseLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
* 1.求链表的长度n
* 2.k = k%n ,k取余就是要右移的距离。
* 3.找到倒数第k个位置。可以用双指针法
* 4.记录慢指针的next节点，这就是最后要返回的节点。
 */

func reverseLinkedList(head *ListNode, k int) *ListNode {
	//求链表的长度
	currNode := head
	node := head
	n := 1
	for currNode.Next != nil {
		currNode = currNode.Next
		n++
	}
	//求右移的距离
	k := k % n
	if k == 0 {
		return currNode
	}
	//找到倒数第k个位置。

	currNode = head
	for k > 0 {
		k--
		currNode = currNode.Next
	}
	for currNode.Next != nil {
		head = head.Next
		currNode = currNode.Next
	}
	//res 为最后返回的头节点
	res = head.Next
	head.Next = nil
	currNode.Next = node
	return res
}
