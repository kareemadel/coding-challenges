/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	listsCount := len(lists)
	if listsCount == 1 {
		return lists[0]
	} else if listsCount == 0 {
		return nil
	}
	left := mergeKLists(lists[:listsCount/2])
	right := mergeKLists(lists[listsCount/2 : listsCount])
	return merge2Lists(left, right)
}

func merge2Lists(n1, n2 *ListNode) *ListNode {
	head := ListNode{}
	current := &head
	for n1 != nil && n2 != nil {
		if n1.Val <= n2.Val {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}
	if n1 == nil {
		current.Next = n2
	} else {
		current.Next = n1
	}
	return head.Next
}