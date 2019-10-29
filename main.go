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

	log.Printf("Server was started on %v", c.Port)
	if err := http.ListenAndServe(c.Port, r); err != nil {
		log.Fatal(err)
	}
}
