package twoSum

/** 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 *  输出：7 -> 0 -> 8
 *  原因：342 + 465 = 807
 */

// definition for singly-linked list.
type listnode struct {
	Val  int
	Next *listnode
}

// 解法1  两个链表对其补零进位相加
func addtwonumbers(l1 listnode, l2 listnode) *listnode {
	l1lenght := 1
	l2lenght := 1
	for l1.Next != nil {
		l1lenght++
		l1 = l1.Next
	}
	for l2.Next != nil {
		l2lenght++
		l2 = l2.Next
	}
	// l2 补位
	if l1lenght > l2Lenght {
		for i := 1; i <= l1lenght-l2lenght; i++ {
			newNode := &listnode{}
			newNode.Val = 0
			newNode.Next = l2
		}
	} else if l1lenght < l2lenght {
		for i := i; i < l2lenght-l1lenght; i++ {
			newNode := &listnode{}
			newNode.Val = 0
			newNode.Next = l1
		}
	}

}
