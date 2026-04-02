package services

import "errors"

var (
	ErrAuthorNotFound   = errors.New("author not found")
	ErrBookNotFound     = errors.New("book not found")
	ErrReviewNotFound   = errors.New("review not found")
	ErrInvalidMinRating = errors.New("min_rating must be between 0 and 5")
)
