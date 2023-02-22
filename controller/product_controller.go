package controller

import (
	"LearnGo/data/request"
	"LearnGo/data/response"
	"LearnGo/helper"
	"LearnGo/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}

// Create			godoc
// @Summary			Create product
// @Description		Save product data in Db.
// @Param			product body request.CreateProductRequest true "Create product"
// @Produce			application/json
// @Tags			product
// @Success			200 {object} response.Response{}
// @Router			/product [post]
func (controller *ProductController) Create(ctx *gin.Context) {
	log.Info().Msg("create product")
	createProductRequest := request.CreateProductRequest{}
	err := ctx.ShouldBindJSON(&createProductRequest)
	helper.ErrorPanic(err)

	controller.productService.Create(createProductRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update			godoc
// @Summary			Update product
// @Description		Update product data.
// @Param			productId path string true "update product by id"
// @Param			product body request.UpdateProductRequest true  "Update product"
// @Tags			product
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/product/{productId} [patch]
func (controller *ProductController) Update(ctx *gin.Context) {
	log.Info().Msg("update product")
	updateProductRequest := request.UpdateProductRequest{}
	err := ctx.ShouldBindJSON(&updateProductRequest)
	helper.ErrorPanic(err)

	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)
	updateProductRequest.Id = id

	controller.productService.Update(updateProductRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete			godoc
// @Summary			Delete product
// @Description		Remove product data by id.
// @Produce			application/json
// @Tags			product
// @Param			productId path string true "Product ID for Delete"
// @Success			200 {object} response.Response{}
// @Router			/product/{productId} [delete]
func (controller *ProductController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete product")
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)
	controller.productService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById 			godoc
// @Summary				Get Single product by id.
// @Param				productId path string true "update product by id"
// @Description			Return the that who productId value matches id.
// @Produce				application/json
// @Tags				product
// @Success				200 {object} response.Response{}
// @Router				/product/{productId} [get]
func (controller *ProductController) FindById(ctx *gin.Context) {
	log.Info().Msg("findById product")
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	helper.ErrorPanic(err)

	tagResponse := controller.productService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll 			godoc
// @Summary			Get All product.
// @Description		Return list of product.
// @Tags			product
// @Success			200 {object} response.Response{}
// @Router			/product [get]
func (controller *ProductController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll products")
	tagResponse := controller.productService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
