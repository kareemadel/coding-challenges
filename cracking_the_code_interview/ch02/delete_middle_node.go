package main

func main() {

}

func deleteMiddleNode(l *node) bool {
	if l == nil || l.Next == nil {
		return false
	}
	next := l.Next
	l.Val = next.Val
	l.Next = next.Next
	return true
}

type node struct {
	Val  int
	Next *node
}
