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
		middleware.ComposeMiddleware(controllers.Info)).Methods("GET")
	r.HandleFunc("/api/users",
		middleware.ComposeMiddleware(controllers.GetUsers(&users))).Methods("GET")
	r.HandleFunc("/api/users",
		middleware.ComposeMiddleware(controllers.CreateUser(&users))).Methods("POST")
	r.HandleFunc("/api/users/{id:[0-9]+}",
		middleware.ComposeMiddleware(controllers.GetUser(&users))).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}",
		middleware.ComposeMiddleware(controllers.DeleteUser(&users))).Methods("DELETE")

	return r
}
