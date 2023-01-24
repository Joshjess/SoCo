package main

import (
	"backend/db_schema"
	CommentRoutes "backend/routes/comment"
	PostRoutes "backend/routes/post"
	UserRoutes "backend/routes/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// function that handles options request
func optionsHandler(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
	}
}

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
		AllowMethods:     []string{"GET, POST, OPTIONS, PUT, PATCH, DELETE"},
		AllowHeaders:     []string{"Content-Type", "Origin, X-Requested-With, Content-Type, Accept, Authorization", "Accept-Encoding", "Cache-Control", "X-Requested-With", "accept", "origin, Access-Control-Allow-Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		ExposeHeaders:    []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization", "Accept-Encoding", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		// AllowAllOrigins:  true,
		MaxAge:        12 * time.Hour,
		AllowWildcard: true,
	})

	r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	v1 := r.Group("/v1")
	v1.Use(myCors)
	{
		v1.Use(optionsHandler)
		{
			UserRoutes.AddUserRoutes(v1, db)
			PostRoutes.AddPostRoutes(v1, db)
			CommentRoutes.AddCommentRoutes(v1, db)
		}
	}
	r.Run()
}
