package main

import (
	"fmt"
	"reflect"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}
	var r = []int{}
	for _, s := range array1 {
		r = append(r, s*s)
	}
	sort.Ints(r)
	sort.Ints(array2)
	return reflect.DeepEqual(r, array2)
}

func main() {
	a := []int{2, 2, 3} // 11, 19, 19, 19, 121, 144, 144, 161
	b := []int{4, 9, 4} // 121, 361, 361, 361, 14641, 20736, 20736, 25921
	fmt.Println(Comp(a, b))
}
