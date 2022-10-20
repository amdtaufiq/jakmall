package controllers

import (
	"jakmall/app/forms"
	"jakmall/app/helper"
	"jakmall/app/services"
	"jakmall/app/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type reviewController struct {
	reviewService services.IReviewService
}

func ReviewController(reviewService services.IReviewService) *reviewController {
	return &reviewController{reviewService}
}

func (c *reviewController) GetAllReview(ctx *gin.Context) {

	reviews, err := c.reviewService.GetAllReview()
	if err != nil {
		response := helper.APIResponse(utils.MsgBadRequest, http.StatusBadRequest, false, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(utils.MsgOk, http.StatusOK, true, reviews)
	ctx.JSON(http.StatusOK, response)
}

func (c *reviewController) GetAllSummary(ctx *gin.Context) {

	reviews, err := c.reviewService.GetAllSummary()
	if err != nil {
		response := helper.APIResponse(utils.MsgBadRequest, http.StatusBadRequest, false, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(utils.MsgOk, http.StatusOK, true, reviews)
	ctx.JSON(http.StatusOK, response)
}

func (c *reviewController) GetSummaryByProductID(ctx *gin.Context) {

	var input forms.GetProductIDInput

	err := ctx.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("please add product ID", http.StatusBadRequest, false, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, _ := strconv.Atoi(input.ID)
	reviews, err := c.reviewService.GetSummaryByProductID(int64(productID))
	if err != nil {
		response := helper.APIResponse(utils.MsgBadRequest, http.StatusBadRequest, false, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(utils.MsgOk, http.StatusOK, true, reviews)
	ctx.JSON(http.StatusOK, response)
}
