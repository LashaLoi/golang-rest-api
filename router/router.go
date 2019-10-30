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
		middleware.ComposeMiddleware(controllers.Index)).Methods("GET")
	r.HandleFunc("/api/users",
		middleware.ComposeMiddleware(controllers.UsersGet(&users))).Methods("GET")
	r.HandleFunc("/api/users",
		middleware.ComposeMiddleware(controllers.UserCreate(&users))).Methods("POST")
	r.HandleFunc("/api/users/{id:[0-9]+}",
		middleware.ComposeMiddleware(controllers.UserGet(&users))).Methods("GET")
	r.HandleFunc("/api/users/{id:[0-9]+}",
		middleware.ComposeMiddleware(controllers.UserDelete(&users))).Methods("DELETE")

	return r
}
