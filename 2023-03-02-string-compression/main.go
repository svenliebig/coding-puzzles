package main

import (
	"strconv"
)

func compress(s []byte) int {
  ci := 0
  cc := s[0]
  ca := 1

  for i, v := range s[1:] {
    if cc == v {
      ca++
    }

    if cc != v || i == len(s) - 2 {
      if ca == 1 {
        s[ci] = cc
        ci++
      } else {
        r:= createCharCompress(cc, ca)
        for i, v := range r {
          s[ci + i] = v
        }
        ci += len(r)
        ca = 1
      }
      cc = v
    }
  }

  if ca == 1 {
    s[ci] = cc
    ci++
  } else {
    r:= createCharCompress(cc, ca)
    for i, v := range r {
      s[ci + i] = v
    }
    ci += len(r)
  }

  s = s[:ci]
  return len(s)
}

func createCharCompress(c byte, a int) []byte {
  return append([]byte{c}, []byte(strconv.FormatInt(int64(a), 10))...)
}

