package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
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
	userDb, err := repository.FindByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.VerifyPassword(userDb.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}
	token, _ := security.GenerateToken(userDb.ID)
	w.Write([]byte(token))
}
