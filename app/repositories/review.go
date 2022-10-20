package repositories

import (
	"encoding/json"
	"io/ioutil"
	"jakmall/app/entities"
	"os"
)

type IReviewRepository interface {
	GetAllReview() ([]entities.Review, error)
}

type reviewRepository struct {
}

func ReviewRepository() *reviewRepository {
	return &reviewRepository{}
}

func (r *reviewRepository) GetAllReview() ([]entities.Review, error) {
	jsonFile, err := os.Open("./resource/reviews.json")
	var reviews []entities.Review

	if err != nil {
		return reviews, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &reviews)

	return reviews, nil
}
