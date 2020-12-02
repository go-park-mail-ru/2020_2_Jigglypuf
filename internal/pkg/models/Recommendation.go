package models

type RecommendationDataFrame struct {
	UserID           uint64
	MovieName        string
	MovieID          uint64
	UserRating       int
	MovieRating      float64
	MovieRatingCount int
}
