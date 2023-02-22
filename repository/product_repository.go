package repository

import "LearnGo/model"

type ProductRepository interface {
	Save(tags model.Product)
	Update(tags model.Product)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Product, err error)
	FindAll() []model.Product
}
