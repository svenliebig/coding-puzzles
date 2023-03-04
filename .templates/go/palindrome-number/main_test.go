package main

import (
  "testing"
)

func TestIsPalindrome(t *testing.T) {
  t.Run("case1", func(t *testing.T) {
    if !isPalindrome(121) {
      t.Error("expected true")
    }
  })

  t.Run("case2", func(t *testing.T) {
    if isPalindrome(-121) {
      t.Error("expected false")
    }
  })

  t.Run("case3", func(t *testing.T) {
    if isPalindrome(10) {
      t.Error("expected false")
    }
  })

  t.Run("case4", func(t *testing.T) {
    if !isPalindrome(1001) {
      t.Error("expected true")
    }
  })
}
