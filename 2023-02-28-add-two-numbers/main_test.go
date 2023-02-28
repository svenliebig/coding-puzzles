package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		l1 := &Node{
			Val: 2,
			Next: &Node{
				Val: 4,
				Next: &Node{
					Val: 3,
				},
			},
		}

		l2 := &Node{
			Val: 5,
			Next: &Node{
				Val: 6,
				Next: &Node{
					Val: 4,
				},
			},
		}

		res := addTwoNumbers(l1, l2)

		result := make([]int, 0)

		for n := res; n != nil; n = n.Next {
			result = append(result, n.Val)
		}

		expected := []int{7, 0, 8}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v to be %v", result, expected)
		}
	})

	t.Run("case2", func(t *testing.T) {
		l1 := &Node{
			Val: 0,
		}

		l2 := &Node{
			Val: 0,
		}

		res := addTwoNumbers(l1, l2)

		result := make([]int, 0)

		for n := res; n != nil; n = n.Next {
			result = append(result, n.Val)
		}

		expected := []int{0}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v to be %v", result, expected)
		}
	})

	t.Run("case3", func(t *testing.T) {
		l1 := &Node{
			Val: 9,
			Next: &Node{
				Val: 9,
				Next: &Node{
					Val: 9,
					Next: &Node{
						Val: 9,
						Next: &Node{
							Val: 9,
							Next: &Node{
								Val: 9,
								Next: &Node{
									Val: 9,
								},
							},
						},
					},
				},
			},
		}

		l2 := &Node{
			Val: 9,
			Next: &Node{
				Val: 9,
				Next: &Node{
					Val: 9,
					Next: &Node{
						Val: 9,
					},
				},
			},
		}

		res := addTwoNumbers(l1, l2)

		result := make([]int, 0)

		for n := res; n != nil; n = n.Next {
			result = append(result, n.Val)
		}

		expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v to be %v", result, expected)
		}
	})

	t.Run("case4", func(t *testing.T) {
		l1 := &Node{
			Val: 2,
			Next: &Node{
				Val: 4,
				Next: &Node{
					Val: 3,
				},
			},
		}

		l2 := &Node{
			Val: 5,
			Next: &Node{
				Val: 6,
				Next: &Node{
					Val: 4,
					Next: &Node{
						Val: 3,
					},
				},
			},
		}

		res := addTwoNumbers(l1, l2)

		result := make([]int, 0)

		for n := res; n != nil; n = n.Next {
			result = append(result, n.Val)
		}

		expected := []int{7, 0, 8, 3}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v to be %v", result, expected)
		}
	})
}
