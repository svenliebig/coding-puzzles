package main

import (
	"strconv"
)

func compress(s []byte) int {
  if len(s) == 1 {
    return 1
  }

  // the running index where we are at
  ri := 1
  // the curret byte that we count
  // var rv byte
  // how often did we count
  rc := 1
  for i, v := range s[1:] {
    if s[ri - 1] != v {
      if rc > 1 {
        for _, b := range []byte(strconv.FormatInt(int64(rc), 10)) {
          s[ri] = b
          ri++
        }
        rc = 1
      }

      s[ri] = v
      ri++
    } else {
      rc++

      if i == len(s) - 2 {
        for _, b := range []byte(strconv.FormatInt(int64(rc), 10)) {
          s[ri] = b
          ri++
        }
      }
    }
  }

  s = s[:ri]
  return len(s)
}

