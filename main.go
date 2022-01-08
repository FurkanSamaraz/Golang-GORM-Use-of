package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Password string
}

func main() {
	db, _ := gorm.Open(sqlite.Open("veri.db"), &gorm.Config{})

	//db.AutoMigrate(&User{}) // oluşturma

	//db.Create(&User{Username: "eren", Password: "321"}) //ekleme

	//Tablo Gösterimi
	var users []User
	db.First(&users, 2)
	//db.Find(&users)
	fmt.Println(users)

	//güncelleme
	//db.First(&users, 2)
	//db.Model(&users).Update("username", "eren")
	//db.Model(&users).Updates(User{Username: "Golang", Password: "456"})
	//fmt.Println(users)

	//Silme işlemi
	//db.Delete(&users, 2)
}

// SİLİNEN VERİYİ TEKRAR KULLANMAYA ÇALIŞALIM

// VERİYİ VERİTABANINDAN YOK ETMİYOR AMA TEKRAR KULLANILMASINADA İZİN VERMİYOR..
