package note

type Note struct {
	Judul string
	Isi  string
}


note_repository.go

package note_repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"(link unavailable)"
)

type NoteRepository struct {
	filePath string
}

func NewNoteRepository(filePath string) *NoteRepository {
	return &NoteRepository{filePath: filePath}
}

func (r *NoteRepository) TambahCatatan(n note.Note) {
	catatan := r.bacaCatatan()
	catatan = append(catatan, n)
	r.simpanCatatan(catatan)
}

func (r *NoteRepository) LihatCatatan() []note.Note {
	return r.bacaCatatan()
}

func (r *NoteRepository) HapusCatatan(judul string) {
	catatan := r.bacaCatatan()
	var baru []note.Note

	for _, c := range catatan {
		if c.Judul != judul {
			baru = append(baru, c)
		}
	}

	r.simpanCatatan(baru)
}

func (r *NoteRepository) bacaCatatan() []note.Note {
	data, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return []note.Note{}
	}

	var catatan []note.Note
	json.Unmarshal(data, &catatan)
	return catatan
}

func (r *NoteRepository) simpanCatatan(catatan []note.Note) {
	data, err := json.MarshalIndent(catatan, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(r.filePath, data, os.ModeAppend)
	if err != nil {
		panic(err)
	}
}