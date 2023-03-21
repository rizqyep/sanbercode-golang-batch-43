package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type NilaiMahasiswa struct {
	ID          uint   `json:"id,omitempty"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai,omitempty"`
	Nilai       uint   `json:"nilai"`
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if username == "rep" && password == "rep123" {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Username atau Password tidak sesuai"))
	})
}

func GetIndeksNilai(nilai uint) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	} else {
		return "E"
	}
}

func AddNilaiEntry(w http.ResponseWriter, r *http.Request) (NilaiMahasiswa, error) {
	var newEntry NilaiMahasiswa
	err := json.NewDecoder(r.Body).Decode(&newEntry)
	if err != nil {
		return NilaiMahasiswa{}, err
	}

	newEntry.ID = uint(rand.Uint32())
	newEntry.IndeksNilai = GetIndeksNilai(newEntry.Nilai)
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newEntry)
	return newEntry, nil
}

func MahasiswaHandlers(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := Response{
			Data:    nilaiNilaiMahasiswa,
			Message: "Data nilai mahasiswa berhasil di fetch !",
		}
		json.NewEncoder(w).Encode(response)
	case "POST":
		newEntry, err := AddNilaiEntry(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := Response{
			Data:    newEntry,
			Message: "Berhasil menambahkan data nilai baru!",
		}
		json.NewEncoder(w).Encode(response)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	server := &http.Server{
		Addr: ":8080",
	}
	http.Handle("/mahasiswa", Auth(http.HandlerFunc(MahasiswaHandlers)))

	fmt.Println("[OK] Server Running at http://localhost:8080")
	server.ListenAndServe()
}
