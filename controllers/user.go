package controllers

import (
	"encoding/json"
	"net/http"

	repositories "github.com/kaetaen/bookworm/repositories"
)

type User struct{}

type UserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	var ur repositories.User

	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := ur.Save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := UserResponse{
		ID:      int(userId),
		Message: "User saved successfully",
	}

	js, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
