package main

import (
	"fmt"
	. "formative-9/library"
)

func main() {
	// Soal 1
	var hitungPersegiPanjang = PersegiPanjang{5, 2}
	var hitungSegitiga HitungBangunDatar = SegitigaSamaSisi{2, 4}
	var hitungBalok HitungBangunRuang = Balok{5, 3, 2}
	var hitungTabung HitungBangunRuang = Tabung{14, 8}

	fmt.Println("Luas Persegi Panjang :", hitungPersegiPanjang.Luas())
	fmt.Println("Keliling Persegi Panjang :", hitungPersegiPanjang.Keliling())
	fmt.Println("Luas Segitiga :", hitungSegitiga.Luas())
	fmt.Println("Keliling Segitiga :", hitungSegitiga.Keliling())

	fmt.Println("Luas Permukaan Balok :", hitungBalok.LuasPermukaan())
	fmt.Println("Volume Balok :", hitungBalok.Volume())
	fmt.Println("Luas Permukaan Tabung :", hitungTabung.LuasPermukaan())
	fmt.Println("Volume Tabung :", hitungTabung.Volume())

	// Soal 2
	var phone PhoneData = Phone{
		Name:   "iPhone 13",
		Brand:  "Apple",
		Year:   2022,
		Colors: []string{"Black", "White", "Blue"},
	}

	fmt.Println(phone.DisplayPhoneData())

	//Soal 3
	fmt.Println(LuasPersegi(4, true))

	fmt.Println(LuasPersegi(8, false))

	fmt.Println(LuasPersegi(0, true))

	fmt.Println(LuasPersegi(0, false))

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

	prefix = prefix.(string) + fmt.Sprintf("= %d", jumlah)

	fmt.Println(prefix)
}
