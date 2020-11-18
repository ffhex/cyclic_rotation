package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{3, 8, 9, 7, 6}, 3))
}

func Solution(A []int, K int) []int {
	if len(A) == K || len(A) == 1 {
		return A
	}

	for i := 0; i < K; i++ {
		m := make([]int, len(A))
		for j := 0; j < len(A); j++ {
			if j+1 > len(A)-1 {
				m[0] = A[j]
			} else {
				m[j+1] = A[j]
			}
		}
		A = m
	}

	return A
}

