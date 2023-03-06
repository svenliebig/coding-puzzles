package main

import (
  "testing"
)

func TestIsSameTree(t *testing.T) {
  t.Run("case1", func(t *testing.T) {
    q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
    p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
    if !isSameTree(q, p) {
      t.Error("should have returned true")
    }
  })
  
  t.Run("case2", func(t *testing.T) {
    q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
    p := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
    if isSameTree(q, p) {
      t.Error("should have returned false")
    }
  })

  t.Run("case3", func(t *testing.T) {
    q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
    p := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
    if isSameTree(q, p) {
      t.Error("should have returned false")
    }
  })
}
