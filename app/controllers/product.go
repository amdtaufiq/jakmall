package controllers

import (
	"jakmall/app/helper"
	"jakmall/app/services"
	"jakmall/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService services.IProductService
}

func ProductController(productService services.IProductService) *productController {
	return &productController{productService}
}

func (c *productController) GetAllProduct(ctx *gin.Context) {

	products, err := c.productService.GetAllProduct()
	if err != nil {
		response := helper.APIResponse(utils.MsgBadRequest, http.StatusBadRequest, false, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(utils.MsgOk, http.StatusOK, true, products)
	ctx.JSON(http.StatusOK, response)
}
