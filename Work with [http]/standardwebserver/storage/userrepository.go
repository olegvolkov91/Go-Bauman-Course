package storage

import (
	"fmt"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/models"
	"log"
)

// UserRepository ... Instance of User Repository (model interface)
type UserRepository struct {
	store *Storage
}

var (
	tableUser string = "users"
)

func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash) VALUES ($1, $2) RETURNING id", tableUser)

	row := ur.store.db.QueryRow(query, u.Login, u.PasswordHash)
	if err := row.Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()

	var founded bool
	if err != nil {
		return nil, founded, err
	}

	var userFound *models.User
	for _, u := range users {
		if u.Login == login {
			userFound = u
			founded = true
			break
		}
	}

	return userFound, founded, nil
}

func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)

	rows, err := ur.store.db.Query(query) // rows надо закрывать
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)

	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.Id, &u.Login, &u.PasswordHash)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}

	return users, nil
}
