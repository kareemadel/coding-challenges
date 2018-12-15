package main

func main() {

}

// Floyd's Hare and tortoise algorithm
func findLoopBegining(n *node) *node {
	slow, fast := n, n
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = n
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

type node struct {
	Val  int
	Next *node
}
