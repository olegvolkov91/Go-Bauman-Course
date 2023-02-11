package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// func CreatePhoneNumber(numbers [10]uint) string {
// 	// this is always the start of the numbers
// 	phoneNumber := "("
// 	// convert uint to string
// 	// sperate area code, middle 3 and the last 4 digits
// 	for i := 0; i < len(numbers); i++ {
// 		// convert to string. (utf-8 dec 48) is 0
// 		numberAsString := string(rune(numbers[i] + 48))
// 		// switch case lets us enter non int characters with precision
// 		switch i {
// 		// add ") " and then add the current number
// 		case 3:
// 			phoneNumber += ") " + numberAsString
// 			// add the dash and the current number
// 		case 6:
// 			phoneNumber += "-" + numberAsString
// 		// ship it
// 		default:
// 			phoneNumber += numberAsString
// 		}
// 	}
// 	// return the formated phone number as a string
// 	return phoneNumber
// }

func CreatePhoneNumber(numbers [10]uint) string {
	var format string = "(xxx) xxx-xxxx"

	for _, v := range numbers {
		format = strings.Replace(format, "x", strconv.Itoa(int(v)), 1)
	}
	return format
}

func main() {
	t := time.Now()
	CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	fmt.Println("It took: \n", time.Since(t))
}
