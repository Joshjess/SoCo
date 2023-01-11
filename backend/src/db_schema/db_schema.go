package db_schema

import gorm "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Status   string
	Token    string
	Posts    []Post
}

type Post struct {
	gorm.Model
	Text     string `gorm:"not null"`
	Upvote   uint
	Downvote uint
	UserID   uint `gorm:"not null"`
	Commnts  []Comment
}

type Comment struct {
	gorm.Model
	Text     string `gorm:"not null"`
	Upvote   uint
	Downvote uint
	UserID   uint `gorm:"not null"`
	PostID   uint `gorm:"not null"`
}
