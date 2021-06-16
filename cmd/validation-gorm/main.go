package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Name   string
	Active bool
}

type Category struct {
	gorm.Model
	Name   string
	Active bool
}

type Product struct {
	gorm.Model
	Description string
	Active      bool
	Brand       Brand    `gorm:"foreignKey:ID"`
	Category    Category `gorm:"foreignKey:ID"`
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"), os.Getenv("SSLMODE"))
	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected. db.Error", db.Error)

	db.AutoMigrate(&Brand{}, &Category{}, &Product{})

	// db.Create(&Product{
	// 	Active:      true,
	// 	Description: "product 1 desc",
	// 	Brand:       Brand{Name: "brand1", Active: true},
	// 	Category:    Category{Name: "category", Active: true},
	// })
	brand := Brand{}
	db.First(&brand, 1)
	fmt.Println("category:", brand)

	category := Category{}
	db.First(&category, 1)
	fmt.Println("category:", category)

	db.Create(&Product{
		Active:      true,
		Description: "product 2 desc",
		Brand:       brand,
		Category:    category,
	})
}
