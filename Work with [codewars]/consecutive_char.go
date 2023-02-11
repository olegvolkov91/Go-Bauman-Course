package main

import "fmt"

type Result struct {
	C rune
	L int
}

func LongestRepetition(text string) Result {
	runeSample := []rune(text)
	res := Result{0, 0}
	temp := Result{0, 0}
	for _, v := range runeSample {
		if temp.C == v {
			temp.C = v
			temp.L = temp.L + 1
		} else {
			temp.C = v
			temp.L = 1
		}
		if res.L < temp.L {
			res.C = temp.C
			res.L = temp.L
		}
	}
	return res
}

func main() {
	res := LongestRepetition("bbbaaabaaaa")
	fmt.Println(string(res.C), res.L)
}

/*


func LongestRepetition(text string) Result {
	slc := strings.Split(text, "")
	res := Result{"", 0}
	temp := Result{"", 0}
	var prevChar string
	for i, v := range slc {
		if i > 0 {
			prevChar = slc[i-1]
		} else {
			prevChar = slc[0]
		}

		if prevChar == v {
			temp.C = prevChar
			temp.L = temp.L + 1
		} else {
			if res.L < temp.L {
				res.C = temp.C
				res.L = temp.L
			}
			temp.C = v
			temp.L = 1
		}
	}

	return res
}

*/
