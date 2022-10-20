package forms

type GetProductIDInput struct {
	ID string `uri:"product_id" binding:"required"`
}
