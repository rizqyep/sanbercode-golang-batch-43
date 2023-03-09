package main

import "fmt"

func main() {

	//Soal 1
	fmt.Println("---------------------")
	fmt.Println("Soal 1 : ")
	fmt.Println("---------------------")
	for i := 1; i <= 20; i++ {
		var kata string

		if i%2 != 0 && i%3 == 0 {
			kata = "I Love Coding."
		} else if i%2 != 0 {
			kata = "Santai"
		} else {
			kata = "Berkualitas"
		}

		fmt.Printf("%d - %s\n", i, kata)

	}

	// Soal 2
	fmt.Println("---------------------")
	fmt.Println("Soal 2 : ")
	fmt.Println("---------------------")
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Printf("\n")
	}

	// Soal 3

	fmt.Println("---------------------")
	fmt.Println("Soal 3 : ")
	fmt.Println("---------------------")

	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	// Soal 4
	fmt.Println("---------------------")
	fmt.Println("Soal 4 : ")
	fmt.Println("---------------------")

	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam")
	sayuran = append(sayuran, "Buncis")
	sayuran = append(sayuran, "Kangkung")
	sayuran = append(sayuran, "Kubis")
	sayuran = append(sayuran, "Seledri")
	sayuran = append(sayuran, "Tauge")
	sayuran = append(sayuran, "Timun")

	for i := 0; i < len(sayuran); i++ {
		fmt.Printf("%d. %s\n", i+1, sayuran[i])
	}

	// Soal 5
	fmt.Println("---------------------")
	fmt.Println("Soal 4 : ")
	fmt.Println("---------------------")

	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)

	}
	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Printf("volume balok = %d\n", volume)

}
