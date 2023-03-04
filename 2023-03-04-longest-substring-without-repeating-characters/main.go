package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
  fmt.Println(s)
  if len(s) <= 1 {
    return len(s)
  }

  m := make(map[byte]int)
  for i := range s {
    if index, ok := m[s[i]]; ok {
      others := lengthOfLongestSubstring(s[index + 1:])
      if i > others {
        return i
      } else {
        return others
      }
    } else {
      m[s[i]] = i
    }
  }

  return len(s)
}
