package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
  
type Student struct {
	gorm.Model
	Name string
	CPF int
	Email string
	Age int
	Active bool
}

func Init() *gorm.DB {
	db, err:= gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&Student{})

	return db
}

func AddStudent() {
	db := Init()

	student := Student{
		Name: "John Doe",
		CPF: 123456789,
		Email: "jhon@gmail.com",
		Age: 20,
		Active: true,
	}

	if result := db.Create(&student); result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("Student added")
}