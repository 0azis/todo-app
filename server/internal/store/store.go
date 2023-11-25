package store

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// NewConnection opens a new instance of database (new connection)
func NewConnection() (*sqlx.DB, error) {
	return sqlx.Open("pgx", "postgres://notes_db_sebp_user:mGA7RJ0scHFknI60vZ6RGkQmaEovnTNU@dpg-clgdq9ef27hc739jfplg-a.frankfurt-postgres.render.com/notes_db_sebp")
}
