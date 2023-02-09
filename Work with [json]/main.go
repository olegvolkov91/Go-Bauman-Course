package main

import (
	"fmt"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/LEC_1/internal/users"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world")
	data, err := users.Deserialize("./internal/users/users.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	bytesArr, err := users.CreateUser()

	if err := os.WriteFile("output.json", bytesArr, 0664); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved")
}
