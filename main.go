package main

import (
	"log"

	"./config"
	"./router"
	"./server"
)

func main() {
	c := config.NewConfig()
	r := router.ConfigureRouter(c.Users)

	s := server.Server{
		Port:   c.Port,
		Router: r,
	}

	if err := s.NewServer(); err != nil {
		log.Fatal(err)
	}
}
