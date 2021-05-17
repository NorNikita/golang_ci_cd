package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type NoteHandler struct {
	repo INoteRepository
}

func (s *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	note, errNote := s.repo.GetNoteById(int64(id))
	if errNote != nil {
		jsonError(w, http.StatusNotFound, errNote.Error())
		return
	}

	body, _ := json.Marshal(note)
	successJson(w, http.StatusOK, body)
}

func (s *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	bodyNote, err := ioutil.ReadAll(r.Body)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	newNote := new(Note)
	errUnmarshal := json.Unmarshal(bodyNote, newNote)
	if errUnmarshal != nil {
		jsonError(w, http.StatusInternalServerError, errUnmarshal.Error())
		return
	}

	note := s.repo.CreateNote(newNote)

	body, _ := json.Marshal(note)
	successJson(w, http.StatusCreated, body)
}

func (s *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	byteNote, err := ioutil.ReadAll(r.Body)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	updated := new(Note)
	errUnmarshal := json.Unmarshal(byteNote, updated)
	if errUnmarshal != nil {
		jsonError(w, http.StatusInternalServerError, errUnmarshal.Error())
		return
	}

	s.repo.UpdateNoteById(int64(id), updated)
	w.WriteHeader(http.StatusAccepted)
}

func (s *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	errDelete := s.repo.DeleteNoteById(int64(id))
	if errDelete != nil {
		jsonError(w, http.StatusInternalServerError, errDelete.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *NoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	order := r.URL.Query()["order_by"][0]

	notes := s.repo.GetListNote(order)

	marshalNotes, err := json.Marshal(notes)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	successJson(w, http.StatusOK, marshalNotes)
}

func jsonError(w http.ResponseWriter, status int, msg string) {
	resp, _ := json.Marshal(map[string]interface{}{
		"message": msg,
	})
	w.WriteHeader(status)

	_, err := w.Write(resp)
	if err != nil {
		fmt.Println(err)
	}
}

func successJson(w http.ResponseWriter, status int, body []byte) {
	w.WriteHeader(status)
	_, _ = w.Write(body)
}
