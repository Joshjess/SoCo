package routes

type User struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Token    string
}
