package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func Deserialize(path string) (*Users, error) {
	start := time.Now()
	fmt.Println("Start deserializing ...")
	defer fmt.Printf("End of deserializing in %v\n", time.Now().Sub(start))

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, errors.New("can not open the file")
	}
	defer jsonFile.Close()

	// вычитываем jsonFile в виде последовательности байтш
	fileInfo, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.New("can not read the file")
	}
	var data *Users

	// проводим десериализацию
	if err := json.Unmarshal(fileInfo, &data); err != nil {
		return nil, errors.New("can not parse the file")
	}
	return data, nil
}
