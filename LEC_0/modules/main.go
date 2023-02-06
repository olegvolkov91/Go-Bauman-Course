package main

import (
	"fmt"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/LEC_0/modules/calculator/calculator"
)

func main() {
	// Добавим первый пакет калькулятор и вызовем метод Add
	res := calculator.Add(5, 10)
	fmt.Println(res) // <- 15 [Успех]
}
