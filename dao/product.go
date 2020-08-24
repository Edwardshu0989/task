package dao

import (
	"awesomeProject/model"
)

func (d *Db) AddProduct(product *model.Product) error {
	return d.db.Create(product).Error
}

func (d *Db) GetProduct(productName string) (product *model.Product) {
	product = &model.Product{}
	d.db.Where("product_name = ?", productName).First(&product)
	return product
}
