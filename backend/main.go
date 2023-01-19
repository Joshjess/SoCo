package main

import (
	"backend/db_schema"
	CommentRoutes "backend/routes/comment"
	PostRoutes "backend/routes/post"
	UserRoutes "backend/routes/user"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password123 dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Amsterdam"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&db_schema.User{}, &db_schema.Post{}, &db_schema.Comment{}, db_schema.VotePost{}, db_schema.VoteComment{})

	//seup CORS
	myCors := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	v1.Use(myCors)
	{
		UserRoutes.AddUserRoutes(v1, db)
		PostRoutes.AddPostRoutes(v1, db)
		CommentRoutes.AddCommentRoutes(v1, db)
	}
	r.Run()
}
