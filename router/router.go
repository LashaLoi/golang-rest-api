package router

import (
	"../controllers"
	"../customtypes"
	"../middleware"
	"github.com/gorilla/mux"
)

// ConfigureRouter ...
func ConfigureRouter(users []customtypes.User) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/",
		middleware.SetResponseHeader(controllers.Info)).Methods("GET")
	r.HandleFunc("/api/users",
		middleware.SetResponseHeader(controllers.GetUsers(&users))).Methods("GET")
	r.HandleFunc("/api/users",
		middleware.SetResponseHeader(controllers.CreateUser(&users))).Methods("POST")
	r.HandleFunc("/api/users/{id:[0-9]+}",
		middleware.SetResponseHeader(controllers.GetUser(&users))).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}",
		middleware.SetResponseHeader(controllers.DeleteUser(&users))).Methods("DELETE")

	return r
}
