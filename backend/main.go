package main

import (
	"backend/db_schema"
	CommentRoutes "backend/routes/comment"
	PostRoutes "backend/routes/post"
	UserRoutes "backend/routes/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
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
	// myCors := cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: false,
	// 	MaxAge:           12 * time.Hour,
	// })

	r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	v1 := r.Group("/v1")
	v1.Use(CORSMiddleware())
	{
		UserRoutes.AddUserRoutes(v1, db)
		PostRoutes.AddPostRoutes(v1, db)
		CommentRoutes.AddCommentRoutes(v1, db)
	}
	r.Run()
}
