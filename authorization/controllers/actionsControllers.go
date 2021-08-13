package controllers

import (
	"../models"
	u "../utils"
	"encoding/json"
	"net/http"
	"time"
)

var CreateAction = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	action := &models.Action{}

	err := json.NewDecoder(r.Body).Decode(action)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	action.UserId = user
	action.DateTime = time.Now().Format("2006-01-02 15:04:05")
	resp := action.Create()
	u.Respond(w, resp)
}

var GetActionsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetActions(id)
	userName := models.GetUser(id).Email
	resp := u.Message(true, "success")
	resp["data"] = data
	resp["user"] = userName
	u.Respond(w, resp)
}

