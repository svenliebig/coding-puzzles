package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }

  m := make(map[int]rune)
  s := strconv.Itoa(x)

  for i, v := range s[:len(s)/2] {
    m[i] = v
  }

  end := len(s) - len(m)
  for i, v := range s[end:] {
    if v != m[len(m) - 1 - i] {
      return false
    }
  }

  return true
}
