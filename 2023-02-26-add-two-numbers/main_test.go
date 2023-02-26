package main

import (
	"reflect"
	"testing"
)

func TestAddTwo(t *testing.T) {
	t.Run("case4", func(t *testing.T) {
		res := addTwoNumbers([]int{1, 2}, []int{2, 4})
		exp := []int{3, 6}
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("expected %v to equal %v", res, exp)
		}
	})

	t.Run("9 + 2 = 11", func(t *testing.T) {
		res := addTwoNumbers([]int{9}, []int{2})
		exp := []int{1, 1}
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("expected %v to equal %v", res, exp)
		}
	})

	t.Run("case1", func(t *testing.T) {
		res := addTwoNumbers([]int{2, 4, 3}, []int{5, 6, 4})
		exp := []int{7, 0, 8}
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("expected %v to equal %v", res, exp)
		}
	})

	t.Run("case2", func(t *testing.T) {
		res := addTwoNumbers([]int{0}, []int{0})
		exp := []int{0}
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("expected %v to equal %v", res, exp)
		}
	})

	t.Run("case3", func(t *testing.T) {
		res := addTwoNumbers([]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9})
		exp := []int{8, 9, 9, 9, 0, 0, 0, 1}
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("expected %v to equal %v", res, exp)
		}
	})
}
