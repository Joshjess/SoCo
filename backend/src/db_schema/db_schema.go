package db_schema

import gorm "gorm.io/gorm"

type User struct {
	gorm.Model
	EmailAdress string
	Password    string
	Status      string
	Token       string
	Posts       []Post
}

type Post struct {
	gorm.Model
	Text     string
	Upvote   uint
	Downvote uint
	UserID   uint
	Commnts  []Comment
}

type Comment struct {
	gorm.Model
	Text     string
	Upvote   uint
	Downvote uint
	UserID   uint
	PostID   uint
}
