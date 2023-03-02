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

// BenchmarkCompress-8   	 7306290	       161.4 ns/op	      24 B/op	         3 allocs/op
//
// BenchmarkCompress-8   	 1764495	       678.5 ns/op	      96 B/op	      12 allocs/op
// BenchmarkCompress-8   	 1761996	       678.3 ns/op	      96 B/op	      12 allocs/op
// BenchmarkCompress-8   	 1765428	       681.3 ns/op	      96 B/op	      12 allocs/op

// switching from using fmt to convert int to string to strconf.FormatInt
// BenchmarkCompress-8   	 3887724	       307.7 ns/op	      96 B/op	      12 allocs/op
func BenchmarkCompress(b *testing.B) {
  for i := 0; i < b.N; i++ {
    compress([]byte("aabbccccccdddddddddeeeeeeaaaaaakkkkkkjjjasjjjjssjsjjjjsjsjsjss"))
  }
}
