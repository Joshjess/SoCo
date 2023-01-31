package routes

import (
	Auth "backend/middleware/auth"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Post struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Title    string `json:"title" binding:"required"`
	Text     string `json:"text" binding:"required"`
	Upvote   int64  `json:"upvote" gorm:"-"`
	DownVote int64  `json:"downvote" gorm:"-"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"username" gorm:"-"` // Virtual field
}

type VotePost struct {
	PostID uint `json:"post_id"`
	UserID uint `json:"-"`
	Vote   bool `json:"vote"`
}

type User struct {
	gorm.Model
	Email string
}

type PostEnv struct {
	db *gorm.DB
}

func countLikesAndLikes(db *gorm.DB, postID uint) (int64, int64) {
	var likes int64
	var dislikes int64
	result := db.Model(&VotePost{}).Where("post_id = ? and vote = ?", postID, true).Count(&likes)
	result_1 := db.Model(&VotePost{}).Where("post_id = ? and vote = ?", postID, false).Count(&dislikes)
	if result.Error != nil && result_1.Error != nil {
		return 0, 0
	}
	return likes, dislikes
}

func addUsernameToPost(db *gorm.DB, post *Post) {
	var user User
	result := db.Where("id = ?", post.UserID).First(&user)
	if result.Error != nil {
		return
	}
	post.UserName = user.Email
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

func AddPostRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	posts := rg.Group("/posts")
	env := &PostEnv{db: db}
	//GET
	posts.GET("/", env.getAllPostsFunc)
	posts.GET("/:id", env.getPostFunc)
	posts.GET("/user/:id", env.getUserPostsFunc)
	//POST
	posts.Use(Auth.AuthRequired())
	{
		posts.POST("/create", env.createPostFunc)
		posts.POST("/vote", env.votePostFunc)
	}
}

func (e *PostEnv) getAllPostsFunc(c *gin.Context) {
	var posts []Post
	result := e.db.Find(&posts)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	if len(posts) == 0 {
		c.JSON(200, gin.H{"message": "No posts found"})
		return
	}
	// Add like and dislike count to post
	for i := 0; i < len(posts); i++ {
		likes, dislikes := countLikesAndLikes(e.db, posts[i].ID)
		posts[i].Upvote = likes
		posts[i].DownVote = dislikes
	}

	// add username to post
	for i := 0; i < len(posts); i++ {
		addUsernameToPost(e.db, &posts[i])
	}

	c.JSON(200, posts)
}

func (e *PostEnv) getUserPostsFunc(c *gin.Context) {
	id := c.Param("id")
	var posts []Post
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := e.db.Where("user_id = ?", idInt).Find(&posts)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	if len(posts) == 0 {
		c.JSON(200, gin.H{"message": "No posts found"})
		return
	}

	// Add like and dislike count to post
	for i := 0; i < len(posts); i++ {
		likes, dislikes := countLikesAndLikes(e.db, posts[i].ID)
		posts[i].Upvote = likes
		posts[i].DownVote = dislikes
	}

	// add username to post
	for i := 0; i < len(posts); i++ {
		addUsernameToPost(e.db, &posts[i])
	}

	c.JSON(200, posts)
}

func (e *PostEnv) getPostFunc(c *gin.Context) {
	id := c.Param("id")
	var post Post
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := e.db.Where("id = ?", idInt).First(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	// Add like and dislike count to post
	likes, dislikes := countLikesAndLikes(e.db, post.ID)
	post.Upvote = likes
	post.DownVote = dislikes

	// add username to post
	addUsernameToPost(e.db, &post)

	c.JSON(200, post)
}

func (e *PostEnv) createPostFunc(c *gin.Context) {
	var post Post
	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user_id, exists := getUintFromContext(c, "user_id")
	if !exists {
		return
	}

	post.UserID = user_id

	result := e.db.Create(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
}

func (e *PostEnv) votePostFunc(c *gin.Context) {
	var vote VotePost
	err := c.BindJSON(&vote)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user_id, exists := getUintFromContext(c, "user_id")
	if !exists {
		return
	}

	// Check if user has already voted
	result := e.db.Where("post_id = ? AND user_id = ?", vote.PostID, user_id).First(&vote)
	// Equality check because you want an error in this case cause that indicates no vote
	if result.Error == nil {
		c.JSON(400, gin.H{"error": "User has already voted"})
		return
	}
	// Update post value
	vote.UserID = user_id
	result = e.db.Create(&vote)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Already voted"})
		return
	}
}
