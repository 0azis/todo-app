package models

// Note
// TODO: add the tags support

type Note struct {
	ID     string `json:"noteID" db:"noteid"`
	UserID int    `json:"userID" db:"userid"`
	Text   string `json:"text" db:"text"`
	Date   string `json:"date" db:"date"`
	// Tags []*Tag `json:"tags"`
}
