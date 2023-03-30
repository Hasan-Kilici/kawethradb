package main

import (
	"fmt"
	kawethradb "main/KawethraDB"
)

type Ogrenci struct {
	ID    int
	Ad    string
	Soyad string
	Sinif int
}

func main() {
	ogrenciler := []Ogrenci{
		{ID: 1, Ad: "Ali", Soyad: "Veli", Sinif: 9},
		{ID: 2, Ad: "Ahmet", Soyad: "Mehmet", Sinif: 10},
		{ID: 3, Ad: "Ayşe", Soyad: "Fatma", Sinif: 11},
		{ID: 4, Ad: "Hasan", Soyad: "KILICI", Sinif: 12},
	}

	err := kawethradb.CreateDB("Ogrenciler", "./data/Ogrenciler.csv", ogrenciler)
	if err != nil {
		fmt.Println("DB oluşturma hatası:", err)
		return
	}

	fmt.Println("DB oluşturuldu!")
	yeniVeri := []string{"2", "Hasan", "Kılıcı", "12"}
	err = kawethradb.Update("./data/Ogrenciler.csv", "ID", 2, yeniVeri)
	if err != nil {
		fmt.Println("Kayıt güncellenirken bir hata oluştu:", err)
		return
	}

	fmt.Println("Kayıt başarıyla güncellendi.")
	kawethradb.Delete("./data/Ogrenciler.csv", "ID", 2)
}
