package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {

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

	if err = user.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	dataBase, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer dataBase.Close()

	repository := repository.NewUserRepository(dataBase)
	insertId, err := repository.Save(user)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = insertId
	response.Json(w, http.StatusCreated, user)
}

func FindByFilter(w http.ResponseWriter, r *http.Request) {

	userParam := strings.ToLower(r.URL.Query().Get("user"))

	log.Printf("user-controller - FindByFilter - param: %v", userParam)

	dataBase, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer dataBase.Close()

	repository := repository.NewUserRepository(dataBase)
	users, err := repository.FindByFilter(userParam)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.Json(w, http.StatusOK, users)

}

func FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	dataBase, err := db.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer dataBase.Close()

	repository := repository.NewUserRepository(dataBase)
	user, err := repository.FindById(userId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.Json(w, http.StatusOK, user)

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
