package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	fmt.Println(len(a)) // 3
	fmt.Println(cap(a)) // 3
	fmt.Println(a[0:])  // [1, 2, 3]
	fmt.Println(a[1:])  // [2, 3]
	fmt.Println(a[2:])  // [3]
	fmt.Println(a[3:])  // []
	fmt.Println(a[4:])  // slice bounds out of range [4:3]
}

func symmetricDifference(s ...[]int) []int {
	res := make([]int, 0)

	for _, arr := range s {
		arr = cleanup(arr)

		for _, item := range arr {
			contains := -1
			for i, resItem := range res {
				if resItem == item {
					contains = i
				}
			}

			if contains != -1 {
				res = append(res[:contains], res[contains+1:]...)
			} else {
				res = append(res, item)
			}
		}
	}

	return sortRange(res)
}

// removes duplicated numbers from a slice an returns it
func cleanup(a []int) []int {
	result := []int{}
	seen := map[int]bool{}

	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = true
		}
	}

	return result
}

func sortRange(a []int) []int {
	for i := range a {
		for r := range a[:i] {
			if i != r && a[r] > a[i] {
				t := a[i]
				a[i] = a[r]
				a[r] = t
			}
		}
	}
	return a
}

func sortFor(a []int) []int {
	for i := 0; i < len(a); i++ {
		for r := 0; r < i; r++ {
			if i != r && a[r] > a[i] {
				t := a[i]
				a[i] = a[r]
				a[r] = t
			}
		}
	}
	return a
}
