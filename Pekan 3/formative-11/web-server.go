package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func kelilingAlasTabung(jariJari float64) float64 {
	return 2 * math.Pi * jariJari
}

func luasAlasTabung(jariJari float64) float64 {
	return math.Pi * jariJari * jariJari
}

func volumeTabung(jariJari float64, tinggi float64) float64 {
	return math.Pi * jariJari * jariJari * tinggi
}

func getPerhitunganTabung(w http.ResponseWriter, r *http.Request) {
	jariJari := 7.0
	tinggi := 10.0

	fmt.Fprintf(w, " jariJari : 7, tinggi: 10, volume : %f, luas alas: %f, keliling alas: %f", kelilingAlasTabung(jariJari), luasAlasTabung(jariJari), volumeTabung(jariJari, tinggi))
}

func main() {
	http.HandleFunc("/", getPerhitunganTabung)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
