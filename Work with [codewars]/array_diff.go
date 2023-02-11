package main

import "fmt"

func ArrayDiff(a, b []int) (result []int) {
	for _, v := range a {
		if !contains(b, v) {
			result = append(result, v)
		}
	}
	return result
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(ArrayDiff([]int{1, 2, 3, 4, 5, 2, 4, 5, 6, 1, 7, 8, 3, 5, 9, 2, 5, 0}, []int{1, 2, 3, 4, 5}))
}
