package main

func main() {
	addTwoNumbers([]int{1, 2, 3}, []int{2, 4})
}

func addTwoNumbers(l1 []int, l2 []int) []int {
	var res []int

	plus := false
	if len(l1) >= len(l2) {
		res = make([]int, len(l1))

		for i, v := range l1 {
			sum := v

			if plus {
				sum++
				plus = false
			}

			if i < len(l2) {
				sum += l2[i]
			}

			if sum > 9 {
				plus = true
				sum -= 10
			}

			res[i] = sum
		}
	} else {
		res = make([]int, len(l2))

		for i, v := range l2 {
			sum := v

			if plus {
				sum++
				plus = false
			}

			if i < len(l1) {
				sum += l1[i]
			}

			if sum > 9 {
				plus = true
				sum -= 10
			}

			res[i] = sum
		}
	}

	if plus {
		res = append(res, 1)
	}

	return res
}
