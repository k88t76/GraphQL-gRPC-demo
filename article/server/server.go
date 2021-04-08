package server

import (
	"database/sql"
)

var db *sql.DB

type server struct {
}

type articleInput struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
