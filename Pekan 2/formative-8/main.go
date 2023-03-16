package main

import (
	"fmt"
	"math"
)

// Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

type persegiPanjang struct {
	panjang, lebar int
}

func (p persegiPanjang) luas() int {
	return p.lebar * p.panjang
}

func (p persegiPanjang) keliling() int {
	return (2 * p.lebar) + (2 * p.panjang)
}

type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.tinggi
}

type balok struct {
	panjang, lebar, tinggi int
}

func (b balok) luasPermukaan() float64 {
	return float64((2 * b.panjang * b.lebar) + (2 * b.panjang * b.tinggi) + (2 * b.lebar * b.tinggi))
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// Soal 2

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p phone) displayPhoneData() string {
	colorsDisplay := ""
	for i, color := range p.colors {
		if i != len(p.colors)-1 {
			colorsDisplay += fmt.Sprintf("%s, ", color)
		} else {
			colorsDisplay += color
		}
	}
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear:%d\nColors:%s\n", p.name, p.brand, p.year, colorsDisplay)
}

type phoneData interface {
	displayPhoneData() string
}

//Soal 3
func luasPersegi(sisi int, displayWords bool) interface{} {

	var returnValue interface{}

	if !displayWords && sisi == 0 {
		returnValue = nil
		return returnValue
	}

	if displayWords {
		if sisi != 0 {
			returnValue = fmt.Sprintf("Luas persegi dengan sisi %d adalah %d", sisi, sisi*sisi)
		} else {
			returnValue = "Maaf anda belum menginput sisi dari persegi"
		}
	} else {
		returnValue = sisi * sisi
	}
	return returnValue
}

func main() {
	// Soal 1
	var hitungPersegiPanjang hitungBangunDatar = persegiPanjang{5, 2}
	var hitungSegitiga hitungBangunDatar = segitigaSamaSisi{2, 4}
	var hitungBalok hitungBangunRuang = balok{5, 3, 2}
	var hitungTabung hitungBangunRuang = tabung{14, 8}

	fmt.Println("Luas Persegi Panjang :", hitungPersegiPanjang.luas())
	fmt.Println("Keliling Persegi Panjang :", hitungPersegiPanjang.keliling())
	fmt.Println("Luas Segitiga :", hitungSegitiga.luas())
	fmt.Println("Keliling Segitiga :", hitungSegitiga.keliling())

	fmt.Println("Luas Permukaan Balok :", hitungBalok.luasPermukaan())
	fmt.Println("Volume Balok :", hitungBalok.volume())
	fmt.Println("Luas Permukaan Tabung :", hitungTabung.luasPermukaan())
	fmt.Println("Volume Tabung :", hitungTabung.volume())

	// Soal 2
	var phone phoneData = phone{
		name:   "iPhone 13",
		brand:  "Apple",
		year:   2022,
		colors: []string{"Black", "White", "Blue"},
	}

	fmt.Println(phone.displayPhoneData())

	//Soal 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	//Soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var kumpulanAngka []int
	kumpulanAngka = append(kumpulanAngka, kumpulanAngkaPertama.([]int)...)
	kumpulanAngka = append(kumpulanAngka, kumpulanAngkaKedua.([]int)...)

	findArraySum := func(arr []int) int {
		res := 0
		for i := 0; i < len(arr); i++ {
			res += arr[i]
		}
		return res
	}

	jumlah := findArraySum(kumpulanAngka)

	for idx, angka := range kumpulanAngka {
		if idx != len(kumpulanAngka)-1 {
			prefix = prefix.(string) + fmt.Sprintf("%d + ", angka)
		} else {
			prefix = prefix.(string) + fmt.Sprintf("%d ", angka)
		}
	}
	// tulis jawaban anda disini

	prefix = prefix.(string) + fmt.Sprintf("= %d", jumlah)

	fmt.Println(prefix)
}
