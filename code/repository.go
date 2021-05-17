package main

import (
	"errors"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type INoteRepository interface {
	GetNoteById(int64) (*Note, error)
	CreateNote(*Note) *Note
	UpdateNoteById(int64, *Note)
	DeleteNoteById(int64) error
	GetListNote(string) []Note
}

type NoteRepository struct {
	sync.RWMutex
	notes map[int64]*Note
}

var (
	noteId          int64 = 0
	autoincrementId       = func() int64 {
		atomic.AddInt64(&noteId, 1)
		return noteId
	}
)

func (r *NoteRepository) GetNoteById(id int64) (*Note, error) {
	r.RLock()
	note, ok := r.notes[id]
	if !ok {
		return nil, errors.New("note does not exist!")
	}
	r.RUnlock()
	return note, nil
}

func (r *NoteRepository) CreateNote(newNote *Note) *Note {
	r.Lock()
	defer r.Unlock()

	id := autoincrementId()
	newNote.ID = id
	newNote.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	r.notes[id] = newNote

	return newNote
}

func (r *NoteRepository) UpdateNoteById(id int64, updatedNote *Note) {
	r.Lock()
	defer r.Unlock()

	old := r.notes[id]
	old.Text = updatedNote.Text
	old.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	r.notes[id] = old
}

func (r *NoteRepository) DeleteNoteById(id int64) error {
	r.Lock()
	_, ok := r.notes[id]
	if !ok {
		return errors.New("notes doest not exist")
	}
	delete(r.notes, id)
	r.Unlock()

	return nil
}

func (r *NoteRepository) GetListNote(order string) []Note {
	var result []Note

	for _, v := range r.notes {
		result = append(result, *v)
	}

	return sortNoteBy(result, order)
}

func sortNoteBy(tasks []Note, order string) []Note {
	sort.Slice(tasks, func(i, j int) bool {
		switch order {
		case "text":
			{
				return strings.Compare(tasks[i].Text, tasks[j].Text) != 1
			}
		case "create_at":
			{
				return strings.Compare(tasks[i].CreateAt, tasks[j].CreateAt) != 1
			}
		case "update_at":
			{
				return strings.Compare(tasks[i].UpdateAt, tasks[j].UpdateAt) != 1
			}
		default:
			{
				return tasks[i].ID < tasks[j].ID
			}
		}
	})
	return tasks
}
