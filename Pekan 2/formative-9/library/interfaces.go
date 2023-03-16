package library

import (
	"fmt"
	"math"
)

// Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (p PersegiPanjang) Luas() int {
	return p.Lebar * p.Panjang
}

func (p PersegiPanjang) Keliling() int {
	return (2 * p.Lebar) + (2 * p.Panjang)
}

type Tabung struct {
	JariJari, Tinggi float64
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (t Tabung) Volume() float64 {
	return math.Pi * t.JariJari * t.Tinggi
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (b Balok) LuasPermukaan() float64 {
	return float64((2 * b.Panjang * b.Lebar) + (2 * b.Panjang * b.Tinggi) + (2 * b.Lebar * b.Tinggi))
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

// Soal 2

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

func (p Phone) DisplayPhoneData() string {
	colorsDisplay := ""
	for i, color := range p.Colors {
		if i != len(p.Colors)-1 {
			colorsDisplay += fmt.Sprintf("%s, ", color)
		} else {
			colorsDisplay += color
		}
	}
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear:%d\nColors:%s\n", p.Name, p.Brand, p.Year, colorsDisplay)
}

type PhoneData interface {
	DisplayPhoneData() string
}

//Soal 3
func LuasPersegi(sisi int, displayWords bool) interface{} {

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
