package routes

import (
	Auth "backend/middleware/auth"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Comment struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Text     string `json:"text" binding:"required"`
	UserID   uint   `json:"-"`
	PostID   uint   `json:"post_id"`
	Likes    int64  `json:"likes" gorm:"-"`
	Dislikes int64  `json:"dislikes" gorm:"-"`
	UserName string `json:"username" gorm:"-"`
}

type VoteComment struct {
	CommentID uint `json:"comment_id"`
	UserID    uint `json:"-"`
	Vote      bool `json:"vote"`
}

type User struct {
	gorm.Model
	Email string
}

type CommentEnv struct {
	db *gorm.DB
}

func countLikesAndLikes(db *gorm.DB, commentID uint) (int64, int64) {
	var likes int64
	var dislikes int64
	result := db.Model(&VoteComment{}).Where("comment_id = ? and vote = ?", commentID, true).Count(&likes)
	result_1 := db.Model(&VoteComment{}).Where("comment_id = ? and vote = ?", commentID, false).Count(&dislikes)
	if result.Error != nil && result_1.Error != nil {
		return 0, 0
	}
	return likes, dislikes
}

func addUsernameToPost(db *gorm.DB, comment *Comment) {
	var user User
	result := db.Where("id = ?", comment.UserID).First(&user)
	if result.Error != nil {
		return
	}
	comment.UserName = user.Email
}

func getUintFromContext(c *gin.Context, key string) (uint, bool) {
	data, exists := c.Get(key)
	if !exists {
		c.JSON(500, gin.H{"error": "Unknow"})
		return 0, false
	}
	//Convert data to int
	//Assert data is string
	dataString := data.(string)
	value, err := strconv.Atoi(dataString)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return 0, false
	}
	return uint(value), true
}

func AddCommentRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	comments := rg.Group("/comments")
	env := &CommentEnv{db: db}
	//GET
	comments.GET("/post/:id", env.getPostCommentsFunc)
	//POST
	comments.Use(Auth.AuthRequired())
	{
		comments.POST("/create", env.createCommentFunc)
		comments.POST("/vote", env.voteCommentFunc)
	}
}

func (e *CommentEnv) getPostCommentsFunc(c *gin.Context) {
	id := c.Param("id")
	var comments []Comment
	result := e.db.Where("post_id = ?", id).Find(&comments)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	if len(comments) == 0 {
		c.JSON(200, gin.H{"message": "No comments found"})
		return
	}
	for i := 0; i < len(comments); i++ {
		addUsernameToPost(e.db, &comments[i])
		likes, dislikes := countLikesAndLikes(e.db, comments[i].ID)
		comments[i].Likes = likes
		comments[i].Dislikes = dislikes
	}

	c.JSON(200, comments)
}

func (e *CommentEnv) createCommentFunc(c *gin.Context) {
	var comment Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user_id, exists := getUintFromContext(c, "user_id")
	if !exists {
		return
	}

	comment.UserID = user_id

	result := e.db.Create(&comment)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
}

func (e *CommentEnv) voteCommentFunc(c *gin.Context) {
	var vote VoteComment
	err := c.BindJSON(&vote)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user_id, exists := getUintFromContext(c, "user_id")
	if !exists {
		return
	}

	result := e.db.Where("comment_id = ? AND user_id = ?", vote.CommentID, user_id).First(&vote)
	if result.Error == nil {
		c.JSON(400, gin.H{"error": "User has already voted"})
		return
	}
	// Update post value
	vote.UserID = user_id
	result = e.db.Create(&vote)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Probaly comment does not exist"})
		return
	}
}
