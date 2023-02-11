package fibo

import "fmt"

func fibonacci() func() int {
	res := -1
	fibo := []int{}

	return func() int {
		res++
		if res <= 2 {
			fibo = append(fibo, res)
			return res
		}
		fibo = append(fibo, fibo[res-2]+fibo[res-1])
		return fibo[res-1]
	}
}

func loop() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
