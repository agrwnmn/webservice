package models

type Players struct {
	School    string
	lastname  string
	firstname string
}

var (
	player []*Players
	nextID = 1
)
