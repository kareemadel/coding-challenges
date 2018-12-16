/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	subHead := swapPairs(head.Next.Next)
	secondNode := head.Next
	head.Next = subHead
	secondNode.Next = head
	return secondNode
}