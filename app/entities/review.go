package entities

type Review struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	Rating    int8  `json:"rating"`
}