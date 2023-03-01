package main

func symmetricDifference(s ...[]int) []int {
  if len(s) == 0 {
    return []int{}
  }

  res := s[0]
  for _, v := range s[1:] {
    res = getDifference(res, v)
  }

  return sort(res)
}

func getDifference(a, b []int) []int {
  ma := make(map[int]bool)
  mb := make(map[int]bool)
  mr := make(map[int]bool)

  for _, v := range a {
    ma[v] = false
  }

  for _, v := range b {
    mb[v] = false
  }

  for _, v := range a {
    if _, ok := mb[v]; !ok {
      mr[v] = false
    }
  }

  for _, v := range b {
    if _, ok := ma[v]; !ok {
      mr[v] = false
    }
  }

  res := make([]int, 0, len(mr))
  for k := range mr {
    res = append(res, k)
  }

  return res
}

func sort(a []int) []int {
  if len(a) < 2 {
    return a
  }

  for i := 0; i < len(a); i++ {
    for ii := len(a) - 1; ii > i; ii-- {
      if a[ii] < a[i] {
        t := a[ii]
        a[ii] = a[i]
        a[i] = t
      }
    }
  }

  return a
}
