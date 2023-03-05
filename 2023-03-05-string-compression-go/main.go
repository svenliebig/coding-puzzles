package main

import (
	"strconv"
)

// Complexity
// Time: O(n+a) where n is the number of elements in the array and a is the
//       amount of times when these elements reaccure 10 times in a row
// Space: O(1) we initialize two variables, c for counting reoccuring bytes and
//        li for the last index. However they are not created in a loop and just
//        take constant space, so therefor the complexity stays O(1)
func compress(s []byte) int {
  if len(s) <= 1 {
    return len(s)
  }

  c := 1
  li := 1
  for i, v := range s[1:] {
    if s[i] != v {
      if c == 1 {
        li++
        continue
      }

      for _, by := range []byte(strconv.Itoa(c)) {
        s[li] = by
        li++
      }

      c = 1
      li++
    } else {
      c++

      if i + 2 == len(s) {
        for _, by := range []byte(strconv.Itoa(c)) {
          s[li] = by
          li++
        }
      }
    }
  }

  s = s[:li]
  return len(s)
}
