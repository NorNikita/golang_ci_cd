package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Note struct {
	ID       int64  `json:"id"`
	Text     string `json:"text"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

func main() {
	r := &NoteRepository{
		notes: make(map[int64]*Note),
	}

	handler := &NoteHandler{
		repo: r,
	}

	router := mux.NewRouter()
	router.HandleFunc("/note/{id}", handler.GetNote).Methods("GET")
	router.HandleFunc("/note", handler.CreateNote).Methods("POST")
	router.HandleFunc("/note/{id}", handler.UpdateNote).Methods("PUT")
	router.HandleFunc("/note/{id}", handler.DeleteNote).Methods("DELETE")
	router.HandleFunc("/note", handler.GetAll).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
