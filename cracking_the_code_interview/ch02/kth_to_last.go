package main

func main() {

}

func kthToLast(l *node, k int) *node {
	target, ahead := l, l
	for i := 0; i < k; i++ {
		if ahead == nil {
			return nil
		}
		ahead = ahead.Next
	}
	for ahead != nil {
		ahead = ahead.Next
		target = target.Next
	}
	return target
}

func kthToLastRecursive(l *node, k int) (*node, int) {
	if l == nil {
		return nil, 0
	}
	_, counter := kthToLastRecursive(l.Next, k)
	counter++
	var result *node
	if counter == k {
		result = l
	}
	return result, counter
}

type node struct {
	Val  int
	Next *node
}
