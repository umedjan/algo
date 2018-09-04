package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{8, 10, 3, 7, 5, 6, 7, 1, 9, 9, 9}
	//a := []int{1, 3, 3}
	fmt.Println(QuickSort(a))
}

func QuickSort(A []int) []int {
	fmt.Print(A)
	if len(A) < 2 {
		fmt.Println()
		return A
	}

	l, r := 0, len(A)-1

	p := A[rand.Int()%len(A)]
	fmt.Println(p)

	for l <= r {
		for A[l] < p {
			l++
		}
		for A[r] > p {
			r--
		}
		if l < r {
			A[l], A[r] = A[r], A[l]
			l++
			r--
		} else if r == l {
			break
		}
	}

	A = append(QuickSort(A[:l]), QuickSort(A[l:])...)

	return A
}
