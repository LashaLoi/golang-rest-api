package main

import (
	"log"
	"net/http"

	"./config"
	"./router"
)

func main() {
	c := config.NewConfig()

	r := router.ConfigureRouter(c.Users)

	log.Println("Server was started on", c.Port)
	http.ListenAndServe(c.Port, r)
}
