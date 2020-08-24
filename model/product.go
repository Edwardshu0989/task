package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductName string
}

type AddProduct struct {
	ProductName string `json:"product_name"`
}
