package main

import (
	"backend/db_schema"
	CommentRoutes "backend/routes/comment"
	PostRoutes "backend/routes/post"
	UserRoutes "backend/routes/user"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=" + os.Getenv("DATABASE_HOST") + " user=postgres password=password123 dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Amsterdam"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&db_schema.User{}, &db_schema.Post{}, &db_schema.Comment{}, db_schema.VotePost{}, db_schema.VoteComment{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	UserRoutes.AddUserRoutes(v1, db)
	PostRoutes.AddPostRoutes(v1, db)
	CommentRoutes.AddCommentRoutes(v1, db)
	r.Run()
}
