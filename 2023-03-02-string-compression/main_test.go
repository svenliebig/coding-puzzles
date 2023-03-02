package main

import (
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
  t.Run("case1", func(t *testing.T) {
    given := []byte("aabbccc")
    want := []byte("a2b2c3")
    got := compress(given)
    if !reflect.DeepEqual(given[:got], want) {
      t.Errorf("got %v want %v", string(given[:got]), string(want))
    }
  })

  t.Run("case2", func(t *testing.T) {
    given := []byte("a")
    want := []byte("a")
    got := compress(given)
    if !reflect.DeepEqual(given[:got], want) {
      t.Errorf("got %v want %v", string(given[:got]), string(want))
    }
  })

  t.Run("case3", func(t *testing.T) {
    given := []byte("abbbbbbbbbbbb")
    want := []byte("ab12")
    got := compress(given)
    if !reflect.DeepEqual(given[:got], want) {
      t.Errorf("got %v want %v", string(given[:got]), string(want))
    }
  })
}
