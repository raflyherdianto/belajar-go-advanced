package main

import (
	"fmt"
	"go1/gorm/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic("Failed to connect to database!")
	}
	defer db.Close()

	fmt.Println("Successfully connected to database!")

	Migrate()
}

func Migrate() {
	db.AutoMigrate(&models.Student{})

	data := models.Student{}
	if db.Find(&data).RecordNotFound() {
		fmt.Println("Seeding database...")
		seederDB()
	}
}

func seederDB() {
	data := models.Student{
		Student_id:      1,
		Student_name:    "Nguyen Van A",
		Student_age:     20,
		Student_address: "Ha Noi",
		Student_phone:   "0123456789",
	}

	db.Create(&data)
}
