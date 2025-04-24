package Service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidateReviewRating(review *models.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("a avaliação deve estar entre 1 e 5")
	}
	return nil
}
