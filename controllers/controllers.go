package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"../customtypes"
	"../validation"
	"github.com/gorilla/mux"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/api/users", http.StatusSeeOther)
}

// UsersGet ...
func UsersGet(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(*users); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}
}

// UserCreate ...
func UserCreate(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user customtypes.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusExpectationFailed)
			return
		}
		defer r.Body.Close()

		if err := validation.UserValidation(user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user.ID = rand.Intn(1000000)
		*users = append(*users, user)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusExpectationFailed)
		}
	}
}

// UserGet ...
func UserGet(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, item := range *users {
			if id == item.ID {
				if err := json.NewEncoder(w).Encode(item); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				return
			}
		}

		notFount := customtypes.NotFoundError{
			Code:    404,
			Message: "Element with this id is not found",
		}

		if err := json.NewEncoder(w).Encode(notFount); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

// UserDelete ...
func UserDelete(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, item := range *users {
			if id == item.ID {
				*users = append((*users)[:i], (*users)[i+1:]...)

				success := customtypes.Success{
					Code:    200,
					Message: "User was successfully deleted",
				}

				if err := json.NewEncoder(w).Encode(success); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				return
			}
		}

		notFount := customtypes.NotFoundError{
			Code:    404,
			Message: "Element with this id is not found",
		}

		if err := json.NewEncoder(w).Encode(notFount); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
