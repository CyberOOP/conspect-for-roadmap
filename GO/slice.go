package homework

import "github.com/k0kubun/pp"

func Slice() {
	nums := []int{12, 7, 3, 5, 9, 6}

	even := make([]int, 0)

	for _, value := range nums {
		if value%2 == 0 {
			even = append(even, value)
		}
	}

	pp.Println(nums)
	pp.Println(even)
}
