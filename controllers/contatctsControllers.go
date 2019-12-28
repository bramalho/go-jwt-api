package controllers

import (
	"encoding/json"
	"go-jwt-api/models"
	u "go-jwt-api/utils"
	"net/http"
)

// CreateContact for user
var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error request body"))
		return
	}

	contact.UserID = user
	resp := contact.Create()
	u.Respond(w, resp)
}

// GetContacts for user
var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
