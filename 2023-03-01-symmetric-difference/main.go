package main

import "fmt"

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

  return quicksort(a)
}

func quicksort(a []int) []int {
  qs(a, 0, len(a)-1)
  return a
}

func qs(a []int, start, end int) {
  if start < end {
    p := part(a, start, end)
    qs(a, start, p-1)
    qs(a, p+1, end)
  }
}

func part(a []int, start, end int) int {
  i := start - 1
  pivot := a[end]

  for l := start; l < end; l++ {
    fmt.Println(a, a[l], pivot)
    if a[l] <= pivot {
      i++
      a[i], a[l] = a[l], a[i]
      fmt.Println("switch", i, l, a)
    }
  }

  a[i + 1], a[end] = a[end], a[i + 1] 
  fmt.Println(a)

  return i + 1
}
