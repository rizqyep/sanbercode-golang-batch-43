package main

import (
	"fmt"
	"strings"
)

func main() {

	// Soal 1
	bootcamp := "Bootcamp"
	digital := "Digital"
	skill := "Skill"
	sanbercode := "Sanbercode"
	golang := "Golang"

	fmt.Printf("%s %s %s %s %s\n", bootcamp, digital, skill, sanbercode, golang)

	//Soal 2

	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)

	fmt.Println(halo)

	//Soal 3

	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.Replace(kataKedua, "s", "S", 1)
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Printf("%s %s %s %s\n", kataPertama, kataKedua, kataKetiga, kataKeempat)
}
