package controllers

import (
	"net/http"
	"github.com/Arthelon/n10n/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/Arthelon/n10n/models"
	"github.com/Arthelon/n10n/controllers/api"
	"fmt"
)

func RequireUserToken(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenStr, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(models.Conf.SigningKey), nil
		})
		if err != nil  {
			fmt.Printf("Error occured while parsing user token: %v\n", err)
			failResp := models.Response{Success: false, Message: "Invalid user token"}
			api.JSONResponse(w, failResp, 400)
			return
		}
		claims, ok := token.Claims.(*models.UserClaims)
		if !ok  {
			failResp := models.Response{Success: false, Message: "Invalid user token"}
			api.JSONResponse(w, failResp, 400)
			return
		}
		user, err := models.GetUserById(claims.UserId)
		if err != nil {
			failResp := models.Response{Success: false, Message: "Invalid user token"}
			api.JSONResponse(w, failResp, 400)
			return
		}
		utils.SetCurrentUser(r, user)
		handler.ServeHTTP(w, r)
		utils.ClearRequest(r)
	}
}