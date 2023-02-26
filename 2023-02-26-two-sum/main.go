package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 3))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(arr []int, tar int) []int {
	for i, v := range arr {
		for ii := i; ii < len(arr); ii++ {
			if i != ii && v+arr[ii] == tar {
				return []int{i, ii}
			}
		}
	}

	return nil
}
