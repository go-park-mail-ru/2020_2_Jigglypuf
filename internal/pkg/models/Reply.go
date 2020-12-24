package models

type ReplyInput struct {
	MovieID uint64
	Text    string
}

type ReplyUpdateInput struct {
	ReplyID int
	NewText string
}

type ReplyModel struct {
	MovieID     uint64
	UserName    string
	UserSurname string
	UserRating  interface{}
	Text        string
}
