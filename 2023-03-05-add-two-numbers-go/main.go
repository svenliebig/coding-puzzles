package main

type Node struct {
	Val  int
	Next *Node
}

// Complexity
// Time: O(n+1) where n is the longest tree and +1 when the last sum is > 9
// Space: O(n+1) we create one new node for each node in the longest tree +1 when the last sum is > 9
func addTwoNumbers(l1 *Node, l2 *Node) *Node {
  if l1 == nil && l2 == nil {
    return nil
  }

  sum := 0
  var l1n, l2n *Node
  if l1 != nil {
    sum += l1.Val
    l1n = l1.Next
  }

  if l2 != nil {
    sum += l2.Val
    l2n = l2.Next
  }

  if sum > 9 {
    if l1n != nil {
      l1n.Val++
    } else {
      l1n = &Node{Val:1}
    }
  }

  next := addTwoNumbers(l1n, l2n)
  return &Node{Val: sum % 10, Next: next}
}
