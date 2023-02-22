package service

import (
	"LearnGo/data/request"
	"LearnGo/data/response"
	"LearnGo/helper"
	"LearnGo/model"
	"LearnGo/repository"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductServiceImpl(tagRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: tagRepository,
		Validate:          validate,
	}
}

// Create implements ProductService
func (t *ProductServiceImpl) Create(product request.CreateProductRequest) {
	err := t.Validate.Struct(product)
	helper.ErrorPanic(err)
	productModel := model.Product{
		Name: product.Name,
	}
	t.ProductRepository.Save(productModel)
}

// Delete implements ProductService
func (t *ProductServiceImpl) Delete(productId int) {
	t.ProductRepository.Delete(productId)
}

// FindAll implements ProductService
func (t *ProductServiceImpl) FindAll() []response.ProductResponse {
	result := t.ProductRepository.FindAll()

	var product []response.ProductResponse
	for _, value := range result {
		tag := response.ProductResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		product = append(product, tag)
	}

	return product
}

// FindById implements ProductService
func (t *ProductServiceImpl) FindById(productId int) response.ProductResponse {
	productData, err := t.ProductRepository.FindById(productId)
	helper.ErrorPanic(err)

	productResponse := response.ProductResponse{
		Id:   productData.Id,
		Name: productData.Name,
	}
	return productResponse
}

// Update implements ProductService
func (t *ProductServiceImpl) Update(product request.UpdateProductRequest) {
	productData, err := t.ProductRepository.FindById(product.Id)
	helper.ErrorPanic(err)
	productData.Name = product.Name
	t.ProductRepository.Update(productData)
}
