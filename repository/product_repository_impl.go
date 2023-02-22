package repository

import (
	"LearnGo/data/request"
	"LearnGo/helper"
	"LearnGo/model"
	"errors"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

// Delete implements ProductRepository
func (t *ProductRepositoryImpl) Delete(productId int) {
	var product model.Product
	result := t.Db.Where("id = ?", productId).Delete(&product)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ProductRepository
func (t *ProductRepositoryImpl) FindAll() []model.Product {
	var product []model.Product
	result := t.Db.Find(&product)
	helper.ErrorPanic(result.Error)
	return product
}

// FindById implements ProductRepository
func (t *ProductRepositoryImpl) FindById(productId int) (productResponse model.Product, err error) {
	var product model.Product
	result := t.Db.Find(&product, productId)
	if result != nil {
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

// Save implements ProductRepository
func (t *ProductRepositoryImpl) Save(product model.Product) {
	result := t.Db.Create(&product)
	helper.ErrorPanic(result.Error)
}

// Update implements ProductRepository
func (t *ProductRepositoryImpl) Update(product model.Product) {
	var updateProduct = request.UpdateProductRequest{
		Id:   product.Id,
		Name: product.Name,
	}
	result := t.Db.Model(&product).Updates(updateProduct)
	helper.ErrorPanic(result.Error)
}
