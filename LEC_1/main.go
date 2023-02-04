package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/LEC_1/internal/users"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world")
	jsonFile, err := os.Open("./users.json")
	if err != nil {
		log.Fatal(errors.New("can not open the file"))
	}

	fileInfo, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(errors.New("can not read the file"))
	}
	var data users.Users

	if err := json.Unmarshal(fileInfo, &data); err != nil {
		log.Fatal(errors.New("can not parse the file"))
	}

	fmt.Println(data)
}
