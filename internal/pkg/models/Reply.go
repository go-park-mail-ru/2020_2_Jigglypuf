package models

type ReplyInput struct {
	MovieID uint64
	Text    string
}

type ReplyUpdateInput struct {
	ReplyID int
	NewText string
}

type ReplyUser struct {
	UserID uint64
	Name string
	Surname string
	AvatarPath string
}

type ReplyModel struct {
	ID			uint64
	MovieID     uint64
	User		ReplyUser
	UserRating  interface{}
	Text        string
}
