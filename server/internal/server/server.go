package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"todo/internal/controllers"
	"todo/internal/middleware"
)

var authControllers controllers.AuthControllers

var notesControllers controllers.NotesControllers

func StartServer() {
	app := fiber.New()

	
	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "https://todo-app-jet-seven.vercel.app/",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	api := app.Group("api")

	auth := api.Group("auth")
	auth.Post("/signin", authControllers.SignIn)
	auth.Post("/signup", authControllers.SignUp)

	notes := api.Group("notes")
	notes.Post("/", middleware.Middleware, notesControllers.CreateNote)
	notes.Delete("/", middleware.Middleware, notesControllers.DeleteNote)
	notes.Get("/", middleware.Middleware, notesControllers.GetMyNotes)
	//notes.Get("/:note_id", middleware.Middleware, notesControllers.GetOneNote)

	app.Listen(":8000")
}
