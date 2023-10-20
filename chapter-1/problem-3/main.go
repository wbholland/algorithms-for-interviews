package main

import "fmt"

func nextGreatest(list []int, key int) int {
	upper := len(list) - 1
	lower := 0

	if upper < 0 || list[upper] <= key {
		return -1
	}

	for lower+1 < upper {
		mid := lower + (upper-lower)/2
		if list[mid] <= key {
			lower = mid
		} else {
			upper = mid
		}
	}

	return upper
}

func main() {
	type Case struct {
		list     []int
		key      int
		expected int
	}

	testCases := []Case{
		{[]int{}, 69, -1},
		{[]int{0, 1, 2}, 2, -1},
		{[]int{0, 0, 0, 0}, 4, -1},
		{[]int{0}, -1, 0},
		{[]int{0, 1, 2, 3, 4, 5}, 1, 2},
		{[]int{0, 1, 2, 3, 4}, 1, 2},
		{[]int{0, 1, 2, 2, 2, 2, 2, 2, 2, 3}, 2, 9},
	}

	for i, v := range testCases {
		output := fmt.Sprintf("Case %d: %d (%d)", i, nextGreatest(v.list, v.key), v.expected)
		fmt.Println(output)
	}
}
