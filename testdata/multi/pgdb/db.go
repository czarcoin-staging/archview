package pgdb

import (
	"example/site/comment"
	"example/site/user"
)

// DB implements postgres implementation of the main database.
//
// architecture: Database
type DB struct {
}

func New() *DB {
	return &DB{}
}

func (db *DB) Users() user.Repo       { return db }
func (db *DB) Comments() comment.Repo { return db }
