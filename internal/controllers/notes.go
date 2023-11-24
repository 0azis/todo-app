package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"todo/internal/models"
	"todo/internal/pkg"
	"todo/internal/services"
)

type NotesControllers struct {
	NotesServices services.NotesServices
}

func (nc NotesControllers) CreateNote(c *fiber.Ctx) error {
	userID, err := pkg.GetIdentity(c.Cookies("token"))
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(500, "Ошибка при создании заметки")
	}

	var newNote models.Note
	c.BodyParser(&newNote)
	if err != nil {
		return fiber.NewError(500, "Ошибка при создании заметки")
	}

	err = nc.NotesServices.InsertOne(newNote, userID)
	if err != nil {
		return fiber.NewError(500, "Ошибка при создании заметки")
	}
	return fiber.NewError(200, "Заметка успешно создана")
}

func (nc NotesControllers) DeleteNote(c *fiber.Ctx) error {
	noteID := c.Query("note_id")
	err := nc.NotesServices.DeleteOne(noteID)
	if err != nil {
		return fiber.NewError(500, "Ошибка при удалении заметки")
	}
	return fiber.NewError(200, "Заметка успешно удалена")
}

func (nc NotesControllers) GetMyNotes(c *fiber.Ctx) error {
	userID, err := pkg.GetIdentity(c.Cookies("token"))
	if err != nil {
		return fiber.NewError(500, "Ошибка при создании заметки")
	}

	allNotes, err := nc.NotesServices.GetAllNotes(userID)
	if err != nil {
		return fiber.NewError(500, "Ошибка при получении заметок")
	}
	return c.JSON(allNotes)
}

func (nc NotesControllers) GetOneNote(c *fiber.Ctx) error {
	noteID := c.Params("note_id")
	oneNote, err := nc.NotesServices.GetNoteByID(noteID)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(500, "Ошибка при получении заметок")
	}
	return c.JSON(oneNote)
}
