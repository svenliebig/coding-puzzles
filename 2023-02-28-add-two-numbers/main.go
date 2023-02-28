package main

type Node struct {
	Val  int
	Next *Node
}

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	if l1 == nil && l2 == nil {
		return nil
	}

	var n1 *Node = nil
	var n2 *Node = nil

	sum := 0
	if l1 != nil {
		sum += l1.Val
		n1 = l1.Next
	}

	if l2 != nil {
		sum += l2.Val
		n2 = l2.Next
	}

	if sum > 9 {
		if l1 == nil {
			l1 = &Node{}
		}

		if l1.Next != nil {
			l1.Next.Val++
		} else {
			l1.Next = &Node{Val: 1}
		}

		n1 = l1.Next
	}

	next := addTwoNumbers(n1, n2)

	return &Node{Val: sum % 10, Next: next}
}
