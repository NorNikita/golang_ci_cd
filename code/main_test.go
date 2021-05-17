package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type testErrorResponse struct {
	Message string `json:"message"`
}

func TestGet(t *testing.T) {
	repo := &NoteRepository{
		notes: make(map[int64]*Note),
	}
	noteHandler := &NoteHandler{
		repo: repo,
	}

	r := httptest.NewRequest("GET", "/note/1", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})
	w := httptest.NewRecorder()

	noteHandler.GetNote(w, r)

	if w.Result().StatusCode != http.StatusNotFound {
		t.Errorf("status not match! actual: %d, expected: %d \n",
			w.Result().StatusCode, http.StatusNotFound)
		return
	}
	if getErrorResponseMessage(w) != "note does not exist!" {
		t.Errorf("error messages not match! actual: %s, expected: %s \n",
			getErrorResponseMessage(w), "note does not exist!")
		return
	}

	note := &Note{
		1,
		"some text",
		time.Now().Format("yyyy-MM-dd HH:mm:ss"),
		time.Now().Format("yyyy-MM-dd HH:mm:ss"),
	}
	repo.notes[1] = note
	r = httptest.NewRequest("GET", "/note/1", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})
	w = httptest.NewRecorder()

	noteHandler.GetNote(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("status not match! actual: %d, expected: %d \n",
			w.Result().StatusCode, http.StatusNotFound)
		return
	}
}

func TestCreateNote(t *testing.T) {
	repo := &NoteRepository{
		notes: make(map[int64]*Note),
	}
	noteHandler := &NoteHandler{
		repo: repo,
	}
	rawBody, err := json.Marshal(
		Note{
			-1,
			"text",
			time.Now().Format("yyyy-MM-dd HH:mm:ss"),
			time.Now().Format("yyyy-MM-dd HH:mm:ss"),
		})
	if err != nil {
		fmt.Println(err)
		return
	}

	r := httptest.NewRequest("POST", "/note", bytes.NewReader(rawBody))
	w := httptest.NewRecorder()

	noteHandler.CreateNote(w, r)

	if w.Result().StatusCode != http.StatusCreated {
		t.Errorf("status not match! actual: %d, expected: %d \n",
			w.Result().StatusCode, http.StatusCreated)
		return
	}
	if getSuccessResponseBody(w).ID != 1 {
		t.Errorf("incorrect id! actual: %d, expected: %d \n",
			getSuccessResponseBody(w).ID, 1)
		return
	}
}

func getSuccessResponseBody(w *httptest.ResponseRecorder) Note {
	note := &Note{}
	all, _ := ioutil.ReadAll(w.Body)
	err := json.Unmarshal(all, note)
	if err != nil {
		fmt.Println(err)
	}
	return *note
}

func getErrorResponseMessage(w *httptest.ResponseRecorder) string {
	res := new(testErrorResponse)
	all, _ := ioutil.ReadAll(w.Body)
	err := json.Unmarshal(all, res)
	if err != nil {
		fmt.Println(err)
	}
	return res.Message
}
