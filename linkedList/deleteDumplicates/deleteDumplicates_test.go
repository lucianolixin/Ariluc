package deleteDumplicates

import "testing"

func TestDeleteDumpulicates(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
	}
	list2 := &ListNode{
		Val: 1,
	}
	list3 := &ListNode{
		Val: 2,
	}
	list1.Next = list2
	list2.Next = list3

	res := delDump(list1)
	curr := res
	t.Log(curr.Val)
	for curr.Next != nil {
		t.Log(curr.Val)
		curr = res.Next
	}
}
