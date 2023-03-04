package main

import (
	"strconv"
)

func compress(s []byte) int {
  if len(s) == 1 {
    return 1
  }

  ri := 1
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

