package main

import (
	"fmt"
	"strings"
	"time"
)

const PREV_CHAR_BYTES int = 96

func High(s string) string {
	slc := strings.Split(s, " ")
	wordIdx := 0
	maxScore := 0

	for idx, word := range slc {
		sum := 0
		for _, val := range []byte(word) {
			sum += (int(val) - PREV_CHAR_BYTES)
			if sum > maxScore {
				maxScore = sum
				wordIdx = idx
			}
		}
	}
	return slc[wordIdx]
}

func main() {
	t := time.Now()
	High("hello world")
	fmt.Printf("It took: %s\n", time.Since(t))
}
