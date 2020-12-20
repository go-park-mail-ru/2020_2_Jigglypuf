package models



type ReplyInput struct {
	MovieID uint64
	Text string
}

type ReplyModel struct {
	MovieID uint64
	UserName string
	UserSurname string
	UserRating int64
	Text string
}
