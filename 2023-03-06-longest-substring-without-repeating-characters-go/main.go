package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
  if len(s) <= 1 {
    return len(s)
  }

  lp := 1
  lngst := 0
  for i := 1; i < len(s); i++ {
    if strings.ContainsRune(s[i - lp:i], rune(s[i])) {
      if lp > lngst {
        lngst = lp
      }
      i = i - lp + 1
      lp = 1
    } else {
      lp++

      if i == len(s) - 1 && lp > lngst {
        lngst = lp
      }
    }
  }

  return lngst
}
