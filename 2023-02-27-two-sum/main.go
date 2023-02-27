package main

func twoSum(arr []int, tar int) []int {
	m := make(map[int]int)
	m[tar] = 0

	for i, v := range arr {
		for ii, vv := range arr {
			if _, ok := m[v+vv]; ok && i != ii {
				return []int{i, ii}
			}
		}
	}

	return []int{}
}
