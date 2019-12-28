package controllers

import (
	"encoding/json"
	"go-jwt-api/models"
	u "go-jwt-api/utils"
	"net/http"
)

// CreateAccount creates a new account
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create()
	u.Respond(w, resp)
}

// Authenticate authenticates user
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

// GetAccount return user information
var GetAccount = func(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "Logged In")
	resp["account"] = models.GetUser(r.Context().Value("user").(uint))

	u.Respond(w, resp)
}
