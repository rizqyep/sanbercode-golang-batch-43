package main

import "fmt"

//Fungsi Soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return (2 * panjang) + (2 * lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

// Soal 2

func introduce(nama string, gender string, pekerjaan string, umur string) string {
	sapaan := "Pak"
	if gender == "perempuan" {
		sapaan = "Bu"
	}

	return fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", sapaan, nama, pekerjaan, umur)
}

// Soal 3

func buahFavorit(nama string, buah ...string) string {
	kalimat := fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah", nama)
	for _, element := range buah {
		kalimat += fmt.Sprintf(" \"%s\",", element)
	}

	return kalimat
}
func main() {

	// Soal 1
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 1 : ")
	fmt.Println("---------------------")

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// Soal 2
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 2 : ")
	fmt.Println("---------------------")

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 3 : ")
	fmt.Println("---------------------")

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// Soal 4
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 4 : ")
	fmt.Println("---------------------")
	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	tambahDataFilm := func(title string, jam string, genre string, tahun string) {
		film := map[string]string{
			"genre": genre,
			"jam":   jam,
			"tahun": tahun,
			"title": title,
		}

		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
