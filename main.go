package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/shahirov/hotel-reservation/api"
)

func main() {
	listenAdder := flag.String("listenAdder", ":3000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAdder)
}
