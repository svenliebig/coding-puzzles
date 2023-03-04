package main

func symmetricDifference(s ...[]int) []int {
  res := s[0]
  for _, v := range s[1:] {
    res = sumDifference(res, v)
  }
  return sort(res)
}

func sumDifference(a []int, b []int) []int {
  // giving this a bit of capactiy to avoid reallocation in some cases
  res := make([]int, 0, len(a))

  m := make(map[int]int)
  for i, v := range a {
    m[v] = i
  }

  mb := make(map[int]int)
  for i, v := range b {
    mb[v] = i
  }

  for i, v := range a {
    if _, ok := mb[v]; !ok && m[v] == i {
      res = append(res, v)
    }
  }

  for i, v := range b {
    if _, ok := m[v]; !ok && mb[v] == i {
      res = append(res, v)
    }
  }

  return res
}

func sort(a []int) []int {
  for i := range a {
    for ii := range a[i + 1:] {
      if a[ii + 1 + i] < a[i] {
        a[ii + 1 + i], a[i] = a[i], a[ii + 1 + i]
      }
    }
  }
  return a
}
