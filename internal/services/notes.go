package services

import (
	"fmt"
	"todo/internal/models"
	"todo/internal/pkg"
	"todo/internal/store"
)

type NotesServices struct{}

func (ns NotesServices) InsertOne(note models.Note, userID int) error {
	db, err := store.NewConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	note.ID = pkg.GenerateUUID()
	_, err = db.Query(fmt.Sprintf("insert into notes values ('%s', %d, '%s')", note.ID, userID, note.Text))
	return err
}

func (ns NotesServices) DeleteOne(noteID string) error {
	db, err := store.NewConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query(fmt.Sprintf("delete from notes values where noteid = '%s'", noteID))
	return err
}

func (ns NotesServices) GetAllNotes(userID int) ([]models.Note, error) {
	var resultNotes []models.Note
	db, err := store.NewConnection()
	if err != nil {
		return resultNotes, err
	}
	defer db.Close()

	err = db.Select(&resultNotes, fmt.Sprintf("select * from notes where userid = %d", userID))
	return resultNotes, err
}

func (ns NotesServices) GetNoteByID(noteID string) (models.Note, error) {
	var resultNote models.Note
	db, err := store.NewConnection()
	if err != nil {
		return resultNote, err
	}
	defer db.Close()

	err = db.Get(&resultNote, fmt.Sprintf("select * from notes where noteid = '%s'", noteID))
	return resultNote, err
}
