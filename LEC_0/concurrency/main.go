package main

import (
	"fmt"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello world")
}

func main() {
	fmt.Println("main started")
	go HelloWorld() // <- данная горутина не успеет даже инициализироваться
	// так как функция main уже закончит своё выполнения
	// Для примера в данном случае, можно сделать следующее:
	time.Sleep(1 * time.Second)
	fmt.Println("main finished")
}
