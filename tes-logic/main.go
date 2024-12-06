package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"(link unavailable)"
	"(link unavailable)"
)

func main() {
	repo := note_repository.NewNoteRepository("notes.json")

	fmt.Println("Aplikasi Catatan")
	fmt.Println("-----------------")

	for {
		fmt.Println("1. Tambah Catatan")
		fmt.Println("2. Lihat Catatan")
		fmt.Println("3. Hapus Catatan")
		fmt.Println("4. Keluar")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahCatatan(repo)
		case 2:
			lihatCatatan(repo)
		case 3:
			hapusCatatan(repo)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahCatatan(repo *note_repository.NoteRepository) {
	var judul, isi string
	fmt.Print("Judul: ")
	fmt.Scanln(&judul)
	fmt.Print("Isi: ")
	fmt.Scanln(&isi)

	repo.TambahCatatan(note.Note{Judul: judul, Isi: isi})
	fmt.Println("Catatan berhasil ditambahkan")
}

func lihatCatatan(repo *note_repository.NoteRepository) {
	catatan := repo.LihatCatatan()
	for _, c := range catatan {
		fmt.Printf("Judul: %s\nIsi: %s\n\n", c.Judul, c.Isi)
	}
}

func hapusCatatan(repo *note_repository.NoteRepository) {
	var judul string
	fmt.Print("Judul catatan untuk dihapus: ")
	fmt.Scanln(&judul)

	repo.HapusCatatan(judul)
	fmt.Println("Catatan berhasil dihapus")
}