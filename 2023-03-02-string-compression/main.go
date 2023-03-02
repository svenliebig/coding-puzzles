package main

import (
	"fmt"
)

func compress(s []byte) int {
  ci := 0
  cc := s[0]
  ca := 1

  for _, v := range s[1:] {
    if cc == v {
      ca++
    }

    if cc != v {
      r := createCharCompress(cc, ca)
      s[ci] = r[0]
      for i, v := range r[1:] {
        s[ci + i + 1] = v
      }
      ci += len(r)
      cc = v
      ca = 1
    }
  }

  r := createCharCompress(cc, ca)
  s[ci] = r[0]
  for i, v := range r[1:] {
    s[ci + i + 1] = v
  }
  ci += len(r)

  s = s[:ci]
  return len(s)
}

func createCharCompress(c byte, a int) []byte {
  if a == 1 {
    return []byte{c}
  }

  return append([]byte{c}, []byte(fmt.Sprint(a))...)
}

