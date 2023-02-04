package users

import (
	"encoding/json"
	"errors"
)

func CreateUser() ([]byte, error) {
	u := User{
		Id:       123,
		Name:     "Viktor",
		Username: "Volodar Kopalni",
		Address: Address{
			Street:  "unknown",
			Suite:   "unknown",
			City:    "Zhmerinka",
			Zipcode: "61070",
			Geo: Geo{
				Lat: "61,33,44",
				Lng: "39, 16, -120",
			},
		},
		Phone:   "0931234567",
		Website: "www.google.com",
		Company: Company{
			Name:        "SelfEmployed",
			CatchPhrase: "Do your best!",
			Bs:          "unknown",
		},
	}

	data, err := json.Marshal(u)

	// у json.Marshall есть ещё другая разновидность сериализации с добавлением отступов json.MarshalIndent
	//data, err := json.MarshalIndent(u, "", "  ")

	if err != nil {
		return nil, errors.New("can not convert struct to byte sequence")
	}

	return data, nil
}
