package service

import (
	"LearnGo/data/request"
	"LearnGo/data/response"
)

type ProductService interface {
	Create(product request.CreateProductRequest)
	Update(product request.UpdateProductRequest)
	Delete(productId int)
	FindById(productId int) response.ProductResponse
	FindAll() []response.ProductResponse
}
