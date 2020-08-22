package main

import "fmt"

func arrangeCoins(n int) int {
	it := 1
	if n == it || n < it {
		return n
	}
	for ; it < n; it++ {
		n -= it
		if n <= it {
			break
		}
	}
	return it
}
func prisonAfterNDays(cells []int, N int) []int {
	var cur [8]int
	copy(cur[:], cells)
	iterateFunc := func() {
		tmp := [8]int{}
		isEqual := func(i, j int) bool {
			return i^j == 0
		}
		for i := 1; i < len(cur)-1; i++ {
			if isEqual(cur[i-1], cur[i+1]) {
				tmp[i] = 1
			}
		}
		copy(cur[:], tmp[:])
	}
	days := [8]int{}
	for d := 0; d < N; d++ {
		iterateFunc()
		if d == 0 {
			days = cur
		} else if cur == days {
			d *= ((N - 1) / d)
		}
	}

	copy(cells, cur[:])

	return cells
}
func plusOne(digits []int) []int {
	size := len(digits)
	for i := size - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}

	if digits[0] == 0 {
		tmp := digits[0]
		digits[0] = 1
		digits = append(digits, tmp)
	}

	return digits
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	sequence := make([]int, n, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &sequence[i])
	}

	iterate := func(i int) bool {
		return (i >= 10 && i < 100) && (i%8 == 0)
	}
	for i := 0; i < n; i++ {
		if iterate(sequence[i]) {
			sum += sequence[i]

		}
	}

	fmt.Printf("Sum : %d\n", sum)
}
