package routes

import (
	"log"
	"time"

	"backend/src/db_schema"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

type Env struct {
	db *gorm.DB
}

func _hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func _checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	users := rg.Group("/users")
	env := &Env{db: db}
	//GET
	// TODO: get user
	// TODO: get karma

	//POST
	users.POST("/login", env.loginUserFunc)
	users.POST("/signup", env.createUserFunc)
}

func (e *Env) loginUserFunc(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	password := user.Password
	result := e.db.Where("email = ?", user.Email).First(&user)
	if result.Error != nil {
		log.Println("Error finding user: ", result.Error.Error())
		c.JSON(401, gin.H{"error": "Error finding user in database"})
		return
	}
	if _checkPasswordHash(password, user.Password) == false {
		log.Println("Incorrect password")
		c.JSON(401, gin.H{"error": "Incorrect password"})
		return
	}

	// Create auth token for user
	expirationTime := time.Now().Add(5 * 24 * time.Hour)

	claims := &Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	// Write auth token to db
	user.Token = tokenString
	result = e.db.Model(&db_schema.User{}).Where("email = ?", user.Email).Update("token", tokenString)
	if result.Error != nil {
		log.Println("Error updating user token: ", result.Error.Error())
		c.JSON(500, gin.H{"error": "Error updating user token in database"})
		return
	}

	// If i am not a lacy fuck look into cookies
	c.JSON(200, gin.H{"token": tokenString})
}

func (e *Env) createUserFunc(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.Password, err = _hashPassword(user.Password)
	result := e.db.Create(&user)
	if result.Error != nil {
		log.Println("Error creating user: ", result.Error.Error())
		c.JSON(500, gin.H{"error": "Error creating user in database"})
		return
	}
	c.JSON(200, gin.H{"succes": "Created user"})
}
