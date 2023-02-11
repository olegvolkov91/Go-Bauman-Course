package main

import "fmt"

const BACKSPACE = "#"

func CleanString(s string) (res string) {
	for i := range s {
		str := fmt.Sprintf("%c", s[i])
		sz := len(res)
		if str == BACKSPACE {
			if sz != 0 {
				res = res[:sz-1]
			}
		} else {
			res += str
		}
	}
	return
}

func main() {
	fmt.Println(CleanString("2$#fs@d####s@#"))
}
