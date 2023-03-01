package main

import (
	"reflect"
	"testing"
)

func TestSymmetricDifference(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		res := symmetricDifference([]int{1, 2, 3}, []int{5, 2, 1, 4})
		if len(res) != 3 || res[0] != 3 || res[1] != 4 || res[2] != 5 {
			t.Errorf("expected result to be [3,4,5] but was %v", res)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		res := symmetricDifference([]int{1, 2, 3, 3}, []int{5, 2, 1, 4})
		if len(res) != 3 || res[0] != 3 || res[1] != 4 || res[2] != 5 {
			t.Errorf("expected result to be [3,4,5] but was %v", res)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		res := symmetricDifference([]int{1, 2, 3}, []int{5, 2, 1, 4, 5})
		if len(res) != 3 || res[0] != 3 || res[1] != 4 || res[2] != 5 {
			t.Errorf("expected result to be [3,4,5] but was %v", res)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		res := symmetricDifference([]int{1, 2, 5}, []int{2, 3, 5}, []int{3, 4, 5})
		if len(res) != 3 || res[0] != 1 || res[1] != 4 || res[2] != 5 {
			t.Errorf("expected result to be [1,4,5] but was %v", res)
		}
	})

	t.Run("case 5", func(t *testing.T) {
		res := symmetricDifference([]int{1, 1, 2, 5}, []int{2, 2, 3, 5}, []int{3, 4, 5, 5})
		if len(res) != 3 || res[0] != 1 || res[1] != 4 || res[2] != 5 {
			t.Errorf("expected result to be [1,4,5] but was %v", res)
		}
	})

	t.Run("case 6", func(t *testing.T) {
		res := symmetricDifference([]int{3, 3, 3, 2, 5}, []int{2, 1, 5, 7}, []int{3, 4, 6, 6}, []int{1, 2, 3})
		if len(res) != 5 || res[0] != 2 || res[1] != 3 || res[2] != 4 || res[3] != 6 || res[4] != 7 {
			t.Errorf("expected result to be [2,3,5,6,7] but was %v", res)
		}
	})

	t.Run("case 7", func(t *testing.T) {
		res := symmetricDifference([]int{3, 3, 3, 2, 5}, []int{2, 1, 5, 7}, []int{3, 4, 6, 6}, []int{1, 2, 3}, []int{5, 3, 9, 8}, []int{1})
		exp := []int{1, 2, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("expected result to be %v but was %v", exp, res)
		}
	})
}
