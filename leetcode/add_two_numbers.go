/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumHelper(l1, l2, 0)
}

func addTwoNumHelper(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	val := carry
	l1Next := l1
	l2Next := l2
	if l1 != nil {
		val += l1.Val
		l1Next = l1.Next
	}
	if l2 != nil {
		val += l2.Val
		l2Next = l2.Next
	}
	if val > 9 {
		carry = 1
		val %= 10
	} else {
		carry = 0
	}
	result := &ListNode{
		Val:  val,
		Next: nil,
	}
	if l1 != nil || l2 != nil {
		more := addTwoNumHelper(l1Next, l2Next, carry)
		result.Next = more
	}
	return result
}