package main

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	fibo := make([]float64, n)

	for i := 0; i < n; i++ {
		if i < len(signature) {
			fibo[i] = signature[i]
		} else {
			fibo[i] = fibo[i-1] + fibo[i-2] + fibo[i-3]
		}
	}

	return fibo
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))
}
