package main

func isPalindrome(x int) bool {
  if x < 0 {
    return false
  }

  if x < 10 {
    return true
  }

  rev := 0
  xcp := x
  for x > 0 {
    rev = (rev * 10) + x % 10
    x = x / 10
  }

  return rev == xcp
}
