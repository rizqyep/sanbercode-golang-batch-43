package main

import "fmt"

//Fungsi Soal 1
const PI float64 = 3.14

func hitungKelilingLingkaran(jariJari *float64, kelilingLingkaran *float64) {
	*kelilingLingkaran = 2 * PI * *jariJari
}

func hitungLuasLingkaran(jariJari *float64, luasLingkaran *float64) {
	*luasLingkaran = PI * *jariJari * *jariJari
}

// Soal 2

func introduce(sentence *string, nama string, gender string, pekerjaan string, umur string) {
	sapaan := "Pak"
	if gender == "perempuan" {
		sapaan = "Bu"
	}

	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", sapaan, nama, pekerjaan, umur)
}

// Soal 3

func tambahDataBuah(buah *[]string) {
	*buah = append(*buah, "semangka", "jeruk", "melon", "pepaya")
}

//Soal 4

func tambahDataFilm(title string, jam string, genre string, tahun string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"genre": genre,
		"jam":   jam,
		"tahun": tahun,
		"title": title,
	}

	*dataFilm = append(*dataFilm, film)
}

// Main Function
func main() {

	// Soal 1
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 1 : ")
	fmt.Println("---------------------")

	var jariJari float64 = 14

	var luasLingkaran float64
	var kelilingLingkaran float64

	hitungLuasLingkaran(&jariJari, &luasLingkaran)
	hitungKelilingLingkaran(&jariJari, &kelilingLingkaran)
	fmt.Println("Nilai PI = ", PI)
	fmt.Println("Luas Lingkaran:", luasLingkaran)
	fmt.Println("Keliling Lingakran:", kelilingLingkaran)

	// Soal 2
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 2 : ")
	fmt.Println("---------------------")

	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 3 : ")
	fmt.Println("---------------------")

	var buah = []string{}

	tambahDataBuah(&buah)

	for i, element := range buah {
		fmt.Printf("%d. %s\n", i+1, element)
	}

	// Soal 4
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 4 : ")
	fmt.Println("---------------------")
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("%d. %2s %s\n", i+1, "title :", item["title"])
		fmt.Printf("%2s %s %s jam\n", "", "durasi :", item["jam"])
		fmt.Printf("%2s %s %s\n", "", "genre :", item["genre"])
		fmt.Printf("%2s %s %s\n\n", "", "year :", item["tahun"])
	}
}
