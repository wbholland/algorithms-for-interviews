package main

import "fmt"

func squareRoot(input uint32) uint32 {
	var lower uint32 = 0
	var upper uint32 = 65536

	for lower+1 < upper {
		mid := lower + (upper-lower)/2
		midSquared := mid * mid
		if midSquared > input {
			upper = mid
		} else if midSquared == input {
			return mid
		} else {
			lower = mid
		}
	}
	return lower
}

func main() {
	testCases := []uint32{0, 1, 2, 3, 4, 25, 4294967295}
	for i, n := range testCases {
		output := fmt.Sprintf("Case %d: floor(sqrt(%d)) is %d", i, n, squareRoot(n))
		fmt.Println(output)
	}
}
