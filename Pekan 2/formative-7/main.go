package main

import "fmt"

//Soal 1
type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

//Soal 2

type segitiga struct {
	alas, tinggi int
}

func (s segitiga) hitungLuas() int {
	return s.alas * s.tinggi / 2
}

type persegi struct {
	sisi int
}

func (p persegi) hitungLuas() int {
	return p.sisi * p.sisi
}

type persegiPanjang struct {
	panjang, lebar int
}

func (p persegiPanjang) hitungLuas() int {
	return p.panjang * p.lebar
}

//Soal 3

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)

}

//Soal 4

type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string, duration int, genre string, year int, movies *[]movie) {
	movie := movie{
		title,
		genre,
		duration,
		year,
	}
	*movies = append(*movies, movie)
}

//Main Function
func main() {
	//Soal 1
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 1 : ")
	fmt.Println("---------------------")

	nanas := buah{
		nama:       "Nanas",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      9000,
	}

	jeruk := buah{
		nama:       "Jeruk",
		warna:      "Oranye",
		adaBijinya: true,
		harga:      8000,
	}

	semangka := buah{
		nama:       "Semangka",
		warna:      "Hijau",
		adaBijinya: true,
		harga:      10000,
	}

	pisang := buah{
		nama:       "Pisang",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      5000,
	}

	fmt.Println(nanas)
	fmt.Println(jeruk)
	fmt.Println(semangka)
	fmt.Println(pisang)

	//Soal 2
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 2 : ")
	fmt.Println("---------------------")

	persegi := persegi{sisi: 4}
	fmt.Println("Luas Persegi :", persegi.hitungLuas())

	segitiga := segitiga{alas: 4, tinggi: 5}
	fmt.Println("Luas Segitiga :", segitiga.hitungLuas())

	persegiPanjang := persegiPanjang{panjang: 5, lebar: 7}
	fmt.Println("Luas Persegi :", persegiPanjang.hitungLuas())

	//Soal 3
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 3 : ")
	fmt.Println("---------------------")

	phone := phone{name: "New Phone", brand: "A Brand", year: 2023, colors: []string{}}

	phone.addColor("blue")
	phone.addColor("red")
	fmt.Println("Phone data:")
	fmt.Println("Name :", phone.name)
	fmt.Println("Brand :", phone.brand)
	fmt.Println("Year :", phone.year)
	fmt.Println("Available Colors:", phone.colors)

	//Soal 4
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 3 : ")
	fmt.Println("---------------------")

	var dataFilm = []movie{}
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("%d. %2s %s\n", i+1, "title :", item.title)
		fmt.Printf("%2s %s %d jam\n", "", "duration :", item.duration)
		fmt.Printf("%2s %s %s\n", "", "genre :", item.genre)
		fmt.Printf("%2s %s %d\n\n", "", "year :", item.year)
	}
}
