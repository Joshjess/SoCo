package db_schema

import gorm "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"unique"`
	Password     string `gorm:"not null"`
	Status       string
	Token        string
	Comments     []Comment
	Posts        []Post
	VotePosts    []VotePost
	VoteComments []VoteComment
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Text    string `gorm:"not null"`
	Votes   []VotePost
	UserID  uint `gorm:"not null"`
	Commnts []Comment
}

type Comment struct {
	gorm.Model
	Text   string `gorm:"not null"`
	Votes  []VoteComment
	UserID uint `gorm:"not null"`
	PostID uint `gorm:"not null"`
}

type VotePost struct {
	gorm.Model
	PostID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
	Vote   bool
}

type VoteComment struct {
	gorm.Model
	CommentID uint `gorm:"not null"`
	UserID    uint `gorm:"not null"`
	Vote      bool
}
