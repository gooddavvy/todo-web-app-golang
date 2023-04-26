package main

var (
	PORT = "1038"
)

type TodoItem struct {
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	DueDate   string `json:"due_date"`
	Completed bool   `json:"completed"`
}
