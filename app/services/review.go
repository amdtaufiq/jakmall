package services

import (
	"jakmall/app/entities"
	"jakmall/app/repositories"
	"math"
)

type IReviewService interface {
	GetAllReview() ([]entities.Review, error)
	GetAllSummary() (entities.Summary, error)
	GetSummaryByProductID(productID int64) (entities.Summary, error)
}

type reviewService struct {
	reviewRepository repositories.IReviewRepository
}

func ReviewService(reviewRepository repositories.IReviewRepository) *reviewService {
	return &reviewService{reviewRepository}
}

func (s *reviewService) GetAllReview() ([]entities.Review, error) {
	return s.reviewRepository.GetAllReview()
}

func (s *reviewService) GetAllSummary() (entities.Summary, error) {
	var summary entities.Summary

	reviews, err := s.reviewRepository.GetAllReview()
	if err != nil {
		return summary, err
	}

	totalReview := 0
	totalRating := 0
	fiveStar := 0
	fourStar := 0
	threeStar := 0
	twoStar := 0
	oneStar := 0

	for _, data := range reviews {
		totalReview++
		totalRating += int(data.Rating)
		switch data.Rating {
		case 5:
			fiveStar++
		case 4:
			fourStar++
		case 3:
			threeStar++
		case 2:
			twoStar++
		case 1:
			oneStar++
		}
	}

	summary.TotalReviews = int64(totalReview)
	summary.AverageRating = float32(math.Round(float64(totalRating)/float64(totalReview)*100) / 100)
	summary.FiveStar = int64(fiveStar)
	summary.FourStar = int64(fourStar)
	summary.ThreeStar = int64(threeStar)
	summary.TwoStar = int64(twoStar)
	summary.OneStar = int64(oneStar)

	return summary, nil
}

func (s *reviewService) GetSummaryByProductID(productID int64) (entities.Summary, error) {
	var summary entities.Summary

	reviews, err := s.reviewRepository.GetAllReview()
	if err != nil {
		return summary, err
	}

	totalReview := 0
	totalRating := 0
	fiveStar := 0
	fourStar := 0
	threeStar := 0
	twoStar := 0
	oneStar := 0

	for _, data := range reviews {
		if data.ProductID == productID {
			totalReview++
			totalRating += int(data.Rating)
			switch data.Rating {
			case 5:
				fiveStar++
			case 4:
				fourStar++
			case 3:
				threeStar++
			case 2:
				twoStar++
			case 1:
				oneStar++
			}
		}
	}

	summary.TotalReviews = int64(totalReview)
	if totalReview != 0 {
		summary.AverageRating = float32(math.Round(float64(totalRating)/float64(totalReview)*100) / 100)
	} else {
		summary.AverageRating = 0
	}
	summary.FiveStar = int64(fiveStar)
	summary.FourStar = int64(fourStar)
	summary.ThreeStar = int64(threeStar)
	summary.TwoStar = int64(twoStar)
	summary.OneStar = int64(oneStar)

	return summary, nil
}
