package main

import "fmt"

func main() {
	for _, i := range []int{
		1,
		1 << 1,
		1 << 2,
		1 << 3,
		1 << 4,
		1 << 5,
		5695939,
		5000000,
		5111111,
		5222222,
		5999999,
		3000000000,
		25 << 7,
		798 << 10} {
		fmt.Printf("for %d %d bits\n", i, countbits(i))
	}
}

func countbits(x int) int {
	numbits := 0
	for i := 0; i <= x; i++ {
		numbits += x & 1
		x >>= 1
	}

	return numbits
}
