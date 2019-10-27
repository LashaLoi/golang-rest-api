package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type notFoundError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type success struct {
	Code int  `json:"code"`
	Ok   bool `json:"ok"`
}

func main() {
	port := ":8080"

	r := mux.NewRouter()

	users := []user{
		{
			ID:    1,
			Name:  "Aliaksei",
			Email: "lashalo11409@gmail.com",
			Phone: "+375 33 603 80 02",
		},
		{
			ID:    2,
			Name:  "David",
			Email: "david@gmail.com",
			Phone: "+375 11 222 33 44",
		},
	}

	r.HandleFunc("/", info).Methods("GET")
	r.HandleFunc("/api/users", getUsers(users)).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id:[0-9]+}", getUser(users)).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}", deleteUser(users)).Methods("DELETE")

	log.Println("Server was started on", port)
	http.ListenAndServe(port, r)
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go to /api/users to use API")
}

func getUsers(users []user) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(users)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(users)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func getUser(users []user) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Fatal(err)
		}

		for _, item := range users {
			if id == item.ID {
				w.Header().Set("Content-Type", "application/json")

				json.NewEncoder(w).Encode(item)

				return
			}
		}

		json.NewEncoder(w).Encode(notFoundError{
			Code:    404,
			Message: "Element with this id is not found",
		})
	}
}

func deleteUser(users []user) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Fatal(err)
		}

		for i, item := range users {
			if id == item.ID {
				w.Header().Set("Content-Type", "application/json")

				users = append(users[:i], users[i+1:]...)

				fmt.Println(users)

				json.NewEncoder(w).Encode(success{
					Code: 200,
					Ok:   true,
				})

				return
			}
		}

		json.NewEncoder(w).Encode(notFoundError{
			Code:    404,
			Message: "Element with this id is not found",
		})
	}
}
