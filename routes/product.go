package routes

import (
	"jakmall/app/controllers"
	"jakmall/app/repositories"
	"jakmall/app/services"

	"github.com/gin-gonic/gin"
)

func ProductRoute(routerGroup *gin.RouterGroup) {
	productRepository := repositories.ProductRepository()
	productService := services.ProductService(productRepository)
	productController := controllers.ProductController(productService)

	routerGroup.GET("/product", productController.GetAllProduct)
}
