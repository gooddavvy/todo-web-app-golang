package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	interfaceMod "github.com/gooddavvy/todo-web-app-golang/interface"
)

var (
	port = interfaceMod.PORT
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func TodoList(w http.ResponseWriter, r *http.Request) {
	list := []interfaceMod.TodoItem{
		{Title: "Todo Item 1", Desc: "My first todo item", DueDate: "2023-4-30", Completed: false},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/api/todoList", TodoList)

	fmt.Println("Server started on port " + port)
	http.ListenAndServe(":"+port, nil)
}
