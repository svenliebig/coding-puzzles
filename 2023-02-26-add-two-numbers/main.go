package main

type Node struct {
	Val  int
	Next *Node
}

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	head := &Node{}
	r := head

	plus := false
	for node := l1; node != nil; node = node.Next {
		sum := node.Val

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if plus {
			plus = false
			sum++
		}

		if sum > 9 {
			sum -= 10
			plus = true
		}

		r.Val = sum

		if node.Next != nil {
			r.Next = &Node{}
			r = r.Next
		}
	}

	if l2 != nil {
		r.Next = &Node{}
		r = r.Next
	}

	for node := l2; node != nil; node = node.Next {
		sum := node.Val

		if plus {
			plus = false
			sum++
		}

		if sum > 9 {
			sum -= 10
			plus = true
		}

		r.Val = sum
		if node.Next != nil {
			r.Next = &Node{}
			r = r.Next
		}
	}

	if plus {
		r.Next = &Node{Val: 1}
	}

	return head
}
