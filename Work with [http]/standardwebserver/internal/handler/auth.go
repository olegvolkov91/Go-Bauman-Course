package handler

import (
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/middleware"
	"github.com/olegvolkov91/Go-Bauman-Course/tree/main/standardwebserver/internal/models"
	"net/http"
	"time"
)

// RegisterUser ... Регистрируем нового юзера
func (h *Handler) RegisterUser(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	h.api.Logger.Info("Register user POST /api/v1/register")

	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		h.api.Logger.Info("Invalid json received from client")
		msg := Message{
			Message:    "Provided json is invalid",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		respond(writer, http.StatusBadRequest, msg)
		return
	}

	// Находим пользователя в базе
	_, ok, err := h.api.Storage.User().FindByLogin(user.Login)
	if err != nil {
		h.api.Logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		respond(writer, http.StatusInternalServerError, msg)
		return
	}

	// Если такой пользователь есть, то регистрацию не делаем
	if ok {
		h.api.Logger.Info("User with that login already exists")
		msg := Message{
			Message:    "Authentication failed",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		respond(writer, http.StatusBadRequest, msg)
		return
	}

	// Хешируем пароль
	if err := user.HashPassword(); err != nil {
		h.api.Logger.Info("Password hashed with some errors")
		msg := Message{
			Message:    "Authentication failed",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		respond(writer, http.StatusBadRequest, msg)
		return
	}

	newUser, err := h.api.Storage.User().Create(&user)
	newUser.Sanitize()

	if err != nil {
		h.api.Logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		respond(writer, http.StatusInternalServerError, msg)
		return
	}

	msg := Message{
		Message:    fmt.Sprintf("User {login: %s}successfully registred!", newUser.Login),
		StatusCode: http.StatusCreated,
		IsError:    false,
	}
	respond(writer, http.StatusCreated, msg)
}

// AuthenticateUser ...
func (h *Handler) AuthenticateUser(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	h.api.Logger.Info("Authenticate user POST /api/v1/user/auth")

	var userFromJson *models.User
	if err := json.NewDecoder(req.Body).Decode(&userFromJson); err != nil {
		h.api.Logger.Info("Invalid json received from client")
		msg := Message{
			Message:    "Provided json is invalid",
			StatusCode: http.StatusBadRequest,
			IsError:    true,
		}
		respond(writer, http.StatusBadRequest, msg)
		return
	}

	// Теперь необходимо найти пользователя в БД с таким login в БД
	user, ok, err := h.api.Storage.User().FindByLogin(userFromJson.Login)

	if err != nil {
		h.api.Logger.Info("Troubles while accessing database table (users) with login. err:", err)
		msg := Message{
			Message:    "We have troubles to accessing database. Try again later",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		respond(writer, http.StatusInternalServerError, msg)
		return
	}

	// Если такого пользователя нет, то обрабатываем данную ситуацию
	if !ok {
		h.api.Logger.Info("User with that login does not exist")
		msg := Message{
			Message:    "Authentication failed",
			StatusCode: http.StatusUnauthorized,
			IsError:    true,
		}
		respond(writer, http.StatusUnauthorized, msg)
		return
	}

	// Если пользователь существует, проверяем пароли
	ok = user.ComparePass(user.Password)

	// Если пароли не одинаковы возвращаем пользователю соответствующий ответ
	if !ok {
		h.api.Logger.Info("Invalid credentials provided")
		msg := Message{
			Message:    "Authentication failed",
			StatusCode: http.StatusUnauthorized,
			IsError:    true,
		}
		respond(writer, http.StatusUnauthorized, msg)
		return
	}

	// Теперь можно вернуть токен после успешных проверок
	token := jwt.New(jwt.SigningMethodHS256)

	// дополнительно шифруем токен
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // задаём время жизни токена

	tokenString, err := token.SignedString(middleware.SecretKey)

	if err != nil {
		h.api.Logger.Info("Something went wrong : authenticate")
		msg := Message{
			Message:    "We have some troubles.Try again",
			StatusCode: http.StatusInternalServerError,
			IsError:    true,
		}
		respond(writer, http.StatusInternalServerError, msg)
		return
	}

	msg := Message{
		Message:    tokenString,
		StatusCode: http.StatusCreated,
		IsError:    false,
	}
	respond(writer, http.StatusCreated, msg)
}
