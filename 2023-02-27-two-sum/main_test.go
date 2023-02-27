package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		res := twoSum([]int{1, 3, 5, 6}, 7)
		exp := []int{0, 3}
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("expected %v to be %v", res, exp)
		}
	})

	t.Run("case2", func(t *testing.T) {
		res := twoSum([]int{3, 3}, 6)
		exp := []int{0, 1}
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("expected %v to be %v", res, exp)
		}
	})

	t.Run("case3", func(t *testing.T) {
		res := twoSum([]int{3, 2, 4}, 6)
		exp := []int{1, 2}
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("expected %v to be %v", res, exp)
		}
	})
}

// BenchmarkTwoSum-8   	15657832	        76.35 ns/op	      16 B/op	       1 allocs/op
func BenchmarkTwoSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		twoSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	}
}
