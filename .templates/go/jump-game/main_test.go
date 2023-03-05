package main

import (
  "testing"
)

func TestMinJumps(t *testing.T) {
  t.Run("case1", func(t *testing.T) {
    arr := []int{1, 2, 3, 1}
    got := minJumps(arr)
    want := 1
    if got != want {
      t.Errorf("got %d want %d", got, want)
    }
  })

  t.Run("case2", func(t *testing.T) {
    arr := []int{3, 2, 3, 1}
    got := minJumps(arr)
    want := 2
    if got != want {
      t.Errorf("got %d want %d", got, want)
    }
  })

  t.Run("case3", func(t *testing.T) {
    arr := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
    got := minJumps(arr)
    want := 3
    if got != want {
      t.Errorf("got %d want %d", got, want)
    }
  })

  t.Run("case4", func(t *testing.T) {
    arr := []int{1, 2, 3, 1, 1, 3}
    got := minJumps(arr)
    want := 2
    if got != want {
      t.Errorf("got %d want %d", got, want)
    }
  })

  t.Run("case5", func(t *testing.T) {
    arr := []int{1, 2, 0, 0, 2, 3}
    got := minJumps(arr)
    want := 3
    if got != want {
      t.Errorf("got %d want %d", got, want)
    }
  })

  t.Run("case6", func(t *testing.T) {
    arr := []int{1, 9, 0, 0, 0, 7, 8, 1, 0, 0, 7, 3}
    got := minJumps(arr)
    want := 5
    if got != want {
      t.Errorf("got %d want %d", got, want)
    }
  })

  t.Run("case7", func(t *testing.T) {
    arr := []int{68,-94,-44,-18,-1,18,-87,29,-6,-87,-27,37,-57,7,18,68,-59,29,7,53,-27,-59,18,-1,18,-18,-59,-1,-18,-84,-20,7,7,-87,-18,-84,-20,-27}
    got := minJumps(arr)
    want := 5
    if got != want {
      t.Errorf("got %d want %d", got, want)
    }
  })
}
