package models

import "gorm.io/gorm"

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
