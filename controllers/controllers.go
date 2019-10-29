package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"../customtypes"
	"github.com/gorilla/mux"
)

// Info ...
func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go to /api/users to use API")
}

// GetUsers ...
func GetUsers(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(*users)
	}
}

// CreateUser ...
func CreateUser(users *[]customtypes.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user customtypes.User

		_ = json.NewDecoder(r.Body).Decode(&user)

		user.ID = rand.Intn(1000000)
		*users = append(*users, user)

		json.NewEncoder(w).Encode(user)
	}
}

// GetUser ...
func GetUser(users *[]customtypes.User) http.HandlerFunc {
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

// DeleteUser ...
func DeleteUser(users *[]customtypes.User) http.HandlerFunc {
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
					Code: 200,
					Ok:   true,
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
