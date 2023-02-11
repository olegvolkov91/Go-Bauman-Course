package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Order(sentence string) string {
	words := []string{}

	for i := range [10]int{} {
		for _, word := range strings.Split(sentence, " ") {
			if res := regexp.MustCompile(fmt.Sprintf("%d", i+1)); res.MatchString(word) {
				words = append(words, word)
			}
		}
	}

	return strings.Join(words, " ")
}

func main() {
	println(Order("4of Fo1r pe6ople g3ood th5e the2"))
}

/*
"is2 Thi1s T4est 3a"  -->  "Thi1s is2 3a T4est"
"4of Fo1r pe6ople g3ood th5e the2"  -->  "Fo1r the2 g3ood 4of th5e pe6ople"
""  -->  ""
*/
