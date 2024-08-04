package main

import (
	"todo-list/handler"
	"todo-list/repository"
	"todo-list/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

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

	// add my own middleware to handle somthing
	// consider exported or non-exported scope in service etc.
	app.Use(cors.New())

	app.Post("/signup", usersHandler.HandleSignUp)
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
