package utils

import (
	"net/http"
	"github.com/Arthelon/n10n/models"
	"github.com/gorilla/context"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func GetCurrentUser(r *http.Request) (*models.User, error) {
	user, ok := context.Get(r, "user").(*models.User)
	if !ok {
		return nil, errors.New("User claims not found in request context")
	}
	return user, nil
}

func SetCurrentUser(r *http.Request, user *models.User) {
	context.Set(r, "user", user)
}

func ClearRequest(r *http.Request) {
	context.Clear(r)
}

func CreateToken(id int64) (string, error) {
	claims := models.UserClaims{
		UserId: id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(models.Conf.SigningKey))
}

