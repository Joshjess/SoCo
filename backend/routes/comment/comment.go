package routes

import (
	Auth "backend/middleware/auth"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Comment struct {
	Text   string `json:"text" binding:"required"`
	UserID uint   `json:"-"`
	PostID uint   `json:"post_id"`
}

type VoteComment struct {
	CommentID uint `json:"comment_id"`
	UserID    uint `json:"-"`
	Vote      bool `json:"vote"`
}

type CommentEnv struct {
	db *gorm.DB
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
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
}
