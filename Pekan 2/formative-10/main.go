package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

func printData(kalimat string, tahun int) {
	fmt.Printf("[SOAL 1] Ini akan dijalankan di akhir: %s %d\n", kalimat, tahun)
}

func recoverError() {
	message := recover()
	if message != nil {
		fmt.Println("\nERROR : ", message)
	}
}

func kelilingSegitigaSamaSisi(sisi int, printKalimat bool) (string, error) {
	defer recoverError()
	if sisi == 0 && !printKalimat {
		defer panic("No Input")
		return "0.0", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}

	if printKalimat {
		if sisi == 0 {
			return "0.0", errors.New("Maaf pembagi tidak boleh NOL")
		}
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, 3*sisi), nil
	} else {
		return fmt.Sprintf("%d", 3*sisi), nil
	}

}

//Soal 5
func getMovies(ch chan string, movies ...string) {
	for _, movie := range movies {
		ch <- movie
	}
	close(ch)
}

//Soal 3
func cetakAngka(angka *int) {
	fmt.Printf("[SOAL 3] Jumlah Angka : %d\n", *angka)
}

func tambahAngka(penjumlah int, angka *int) {
	*angka += penjumlah
}

//Soal 4
func tambahDataPhone(data string, dataPhone *[]string) {
	*dataPhone = append(*dataPhone, data)
}

//Soal 5

func printDataPhone(i int, phone string, wg *sync.WaitGroup) {
	fmt.Printf("%d. %s\n", i+1, phone)
	time.Sleep(time.Second * 1)
	wg.Done()
}
func main() {
	//Deklarasi Angka dan Defer untuk Soal 3
	angka := 1
	defer cetakAngka(&angka)
	//Soal 1
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 1 : ")
	fmt.Println("---------------------")

	defer printData("Golang Backend Development", 2021)
	fmt.Println("Halo Program!")

	//Soal 2
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 2 : ")
	fmt.Println("---------------------")

	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	//Soal 3
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 3 : ")
	fmt.Println("---------------------")

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	//Soal 4
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 4 : ")
	fmt.Println("---------------------")

	var phones = []string{}
	tambahDataPhone("Xiaomi", &phones)
	tambahDataPhone("Asus", &phones)
	tambahDataPhone("Iphone", &phones)
	tambahDataPhone("Samsung", &phones)
	tambahDataPhone("Oppo", &phones)
	tambahDataPhone("Realme", &phones)
	tambahDataPhone("Vivo", &phones)

	sort.Strings(phones)
	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second * 1)
	}

	//Soal 5
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 5 : ")
	fmt.Println("---------------------")

	var wg sync.WaitGroup

	for i, phone := range phones {
		wg.Add(1)
		printDataPhone(i, phone, &wg)
	}

	//Soal 6
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Soal 6 : ")
	fmt.Println("---------------------")

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Output Defer : ")
	fmt.Println("---------------------")

}
