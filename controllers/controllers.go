package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"../customtypes"
	"github.com/gorilla/mux"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/api/users", http.StatusSeeOther)
}

// UsersGet ...
func UsersGet(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(*users)
	}
}

// UserCreate ...
func UserCreate(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user customtypes.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()

		user.ID = rand.Intn(1000000)
		*users = append(*users, user)

		json.NewEncoder(w).Encode(user)
	}
}

// UserGet ...
func UserGet(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Fatal(err)
		}

		for _, item := range *users {
			if id == item.ID {

				json.NewEncoder(w).Encode(item)

				return
			}
		}

		json.NewEncoder(w).Encode(customtypes.NotFoundError{
			Code:    404,
			Message: "Element with this id is not found",
		})
	}
}

// UserDelete ...
func UserDelete(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Fatal(err)
		}

		for i, item := range *users {
			if id == item.ID {

				*users = append((*users)[:i], (*users)[i+1:]...)

				json.NewEncoder(w).Encode(customtypes.Success{
					Code:    200,
					Message: "User was successfully deleted",
				})

				return
			}
		}

		json.NewEncoder(w).Encode(customtypes.NotFoundError{
			Code:    404,
			Message: "Element with this id is not found",
		})
	}
}
