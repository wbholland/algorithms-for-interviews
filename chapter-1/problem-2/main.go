package main

import "fmt"

func find(list []int, key int) int {
	if len(list) == 0 {
		return -1
	}
	upper := len(list) - 1
	lower := 0
	for lower+1 < upper {
		mid := lower + (upper-lower)/2
		if key <= list[mid] {
			upper = mid
		} else {
			lower = mid
		}
	}
	if list[lower] == key {
		return lower
	} else if list[upper] == key {
		return upper
	} else {
		return -1
	}
}

func main() {
	type Case struct {
		list     []int
		key      int
		expected int
	}

	testCases := []Case{
		{[]int{0, 1, 2, 3, 4, 5}, 1, 1},
		{[]int{0, 1, 2, 3, 4, 5}, 5, 5},
		{[]int{0, 1, 2, 3, 4, 5}, 4, 4},
		{[]int{0, 0, 0, 0, 1}, 0, 0},
		{[]int{0, 0, 0, 0, 1}, 1, 4},
		{[]int{0, 0}, 1, -1},
		{[]int{0}, 0, 0},
		{[]int{0}, 1, -1},
		{[]int{}, 0, -1},
		{[]int{0, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5}, 3, 6},
		{[]int{0, 1, 1, 1}, 1, 1},
	}

	for i, v := range testCases {
		output := fmt.Sprintf("Case %d: %d (%d)", i, find(v.list, v.key), v.expected)
		fmt.Println(output)
	}
}
