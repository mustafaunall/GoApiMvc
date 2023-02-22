package router

import (
	"LearnGo/controller"
	"LearnGo/repository"
	"LearnGo/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ProductRoutes(route *gin.RouterGroup, db *gorm.DB) {
	validate := validator.New()
	// Repository
	productRepository := repository.NewProductRepositoryImpl(db)
	// Service
	productService := service.NewProductServiceImpl(productRepository, validate)
	// Controller
	productController := controller.NewProductController(productService)
	productRouter := route.Group("/product")
	{
		productRouter.GET("", productController.FindAll)
		productRouter.GET("/:productId", productController.FindById)
		productRouter.POST("", productController.Create)
		productRouter.PATCH("/:productId", productController.Update)
		productRouter.DELETE("/:productId", productController.Delete)
	}
}
