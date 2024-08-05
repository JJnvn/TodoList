package main

import (
	"fmt"
	"todo-list/handler"
	"todo-list/repository"
	"todo-list/service"
	"todo-list/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func jwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		c.Locals("username", claims.Username)
		fmt.Println("hello")
		return c.Next()
	}
}

func main() {

	var err error
	db, err = sqlx.Open("mysql", "root:root@tcp(localhost:3306)/todo-list")
	if err != nil {
		panic(err)
	}
	usersRepo := repository.NewUserRepoImpl(db)
	usersService := service.NewUsersServiceImpl(usersRepo)
	usersHandler := handler.NewUsersHandler(usersService)
	app := fiber.New()

	// consider exported or non-exported scope in service etc.
	app.Use(cors.New())

	app.Post("/signup", usersHandler.HandleSignUp)
	app.Post("/todolist", jwtMiddleware(), usersHandler.HandleLogin)
	app.Post("/login", usersHandler.HandleLogin)
	app.Get("/hello", Hello)

	app.Listen(":8000")
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

type User struct {
	Id       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct {
	User    User   `json:"user"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
