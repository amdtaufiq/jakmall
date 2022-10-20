package routes

import (
	"jakmall/app/controllers"
	"jakmall/app/repositories"
	"jakmall/app/services"

	"github.com/gin-gonic/gin"
)

func ReviewRoute(routerGroup *gin.RouterGroup) {
	reviewRepository := repositories.ReviewRepository()
	reviewService := services.ReviewService(reviewRepository)
	reviewController := controllers.ReviewController(reviewService)

	routerGroup.GET("/review", reviewController.GetAllReview)
	routerGroup.GET("/review/summary", reviewController.GetAllSummary)
	routerGroup.GET("/review/summary/:product_id", reviewController.GetSummaryByProductID)
}
