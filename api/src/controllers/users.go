package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gustavomello-21/devbook/api/src/database"
	"github.com/gustavomello-21/devbook/api/src/models"
	"github.com/gustavomello-21/devbook/api/src/repositories"
	"github.com/gustavomello-21/devbook/api/src/response"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	err = user.Validate("cadastro")
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	user.ID, err = userRepository.Create(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	NameOrNick := r.URL.Query().Get("usuarios")
	NameOrNick = strings.ToLower(NameOrNick)

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	userReposity := repositories.NewUserRepository(db)
	user, err := userReposity.Find(NameOrNick)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	targetId, err := strconv.Atoi(params["userId"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	user, err := userRepository.FindById(targetId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	targetId, err := strconv.Atoi(params["userId"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}
	var user models.User
	err = json.Unmarshal(body, &user)

	err = user.Validate("edit")
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	userRepository := repositories.NewUserRepository(db)
	err = userRepository.Update(targetId, user)

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	targetId, err := strconv.Atoi(params["userId"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	userRepository := repositories.NewUserRepository(db)
	err = userRepository.Delete(targetId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusNoContent, nil)
}
