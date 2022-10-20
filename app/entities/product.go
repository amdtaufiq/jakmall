package entities

type (
	Product struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Price string `json:"price"`
	}

	Summary struct {
		TotalReviews  int64   `json:"total_reviews"`
		AverageRating float32 `json:"average_rating"`
		FiveStar      int64   `json:"5_star"`
		FourStar      int64   `json:"4_star"`
		ThreeStar     int64   `json:"3_star"`
		TwoStar       int64   `json:"2_star"`
		OneStar       int64   `json:"1_star"`
	}
)
