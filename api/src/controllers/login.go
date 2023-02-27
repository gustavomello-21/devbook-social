package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gustavomello-21/devbook/api/src/authenticate"
	"github.com/gustavomello-21/devbook/api/src/database"
	"github.com/gustavomello-21/devbook/api/src/models"
	"github.com/gustavomello-21/devbook/api/src/repositories"
	"github.com/gustavomello-21/devbook/api/src/response"
	"github.com/gustavomello-21/devbook/api/src/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
	}
	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewUserRepository(db)
	userInDatabase, err := userRepository.FindByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusNotFound, nil)
		return
	}

	err = security.CheckPasswordHash(user.Password, userInDatabase.Password)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authenticate.Generate(userInDatabase.ID)

	w.Write([]byte(token))
}
