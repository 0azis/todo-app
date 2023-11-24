package server

import (
	"github.com/gofiber/fiber/v2"
	"todo/internal/controllers"
)

var authControllers controllers.AuthControllers

var notesControllers controllers.NotesControllers

func StartServer() {
	app := fiber.New()
	api := app.Group("api")

	auth := api.Group("auth")
	auth.Post("/signin", authControllers.SignIn)
	auth.Post("/signup", authControllers.SignUp)

	notes := api.Group("notes")
	notes.Post("/", notesControllers.CreateNote)
	notes.Delete("/", notesControllers.DeleteNote)
	notes.Get("/", notesControllers.GetMyNotes)
	notes.Get("/:note_id", notesControllers.GetOneNote)

	app.Listen(":8000")
}
