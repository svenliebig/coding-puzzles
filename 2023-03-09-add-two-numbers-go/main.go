package main

type Node struct {
	Val  int
	Next *Node
}

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
  if l1 == nil && l2 == nil {
    return nil
  }

  sum := 0

  if l1 != nil {
    sum += l1.Val
    l1 = l1.Next
  }

  if l2 != nil {
    sum += l2.Val
    l2 = l2.Next
  }

  if sum > 9 {
    if l1 != nil {
      l1.Val += 1
    } else {
      l1 = &Node{ Val: 1, Next: nil }
    }
  }


  next := addTwoNumbers(l1, l2)
	return &Node{ Val: sum % 10, Next: next }
}
