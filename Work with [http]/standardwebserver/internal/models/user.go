package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"-"`
}

func (u *User) HashPassword() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)

		if err != nil {
			return err
		}
		u.PasswordHash = enc
	}
	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePass(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) == nil
}

func encryptString(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
