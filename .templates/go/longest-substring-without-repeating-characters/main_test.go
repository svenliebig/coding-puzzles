package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
  t.Run("case1", func(t *testing.T) {
    res := lengthOfLongestSubstring("abcabcbb")
    if res != 3 {
      t.Errorf("expected 3, got %d", res)
    }
  })

  t.Run("case2", func(t *testing.T) {
    res := lengthOfLongestSubstring("bbbbb")
    if res != 1 {
      t.Errorf("expected 1, got %d", res)
    }
  })

  t.Run("case3", func(t *testing.T) {
    res := lengthOfLongestSubstring("pwwkew")
    if res != 3 {
      t.Errorf("expected 3, got %d", res)
    }
  })

  t.Run("case4", func(t *testing.T) {
    res := lengthOfLongestSubstring("aab")
    if res != 2 {
      t.Errorf("expected 2, got %d", res)
    }
  })

  t.Run("case5", func(t *testing.T) {
    res := lengthOfLongestSubstring("dvdf")
    if res != 3 {
      t.Errorf("expected 3, got %d", res)
    }
  })
}

// BenchmarkLengthOfLongestSubstring-8   	 9766644	       121.9 ns/op	      0 B/op	       0 allocs/op
func BenchmarkLengthOfLongestSubstring(b *testing.B) {
  for i := 0; i < b.N; i++ {
    lengthOfLongestSubstring("abcabcbb")
  }
}
