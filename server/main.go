package main

import (
	"todo-list/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func main() {

	var err error
	db, err = sqlx.Open("mysql", "root:root@tcp(localhost:3306)/todo-list")
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	// add my own middleware to handle somthing
	app.Use(cors.New())

	app.Post("/signup", Signup)
	app.Post("/login", Login)
	app.Get("/hello", Hello)

	app.Listen(":8000")
}

func Signup(c *fiber.Ctx) error {

	request := SignupRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = utils.HandleEmptyUserOrPass(request.Username, request.Password)
	if err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	query := "insert users (username, password) values (?, ?)"
	result, err := db.Exec(query, request.Username, string(password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	response := SignupResponse{
		User: User{
			Id:       int(id),
			Username: request.Username,
			Password: string(password),
		},
		Message: "Sign up success!",
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func Login(c *fiber.Ctx) error {

	request := LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = utils.HandleEmptyUserOrPass(request.Username, request.Password)
	if err != nil {
		return err
	}

	user := User{}
	query := "select id, username, password FROM users WHERE username = ?"

	err = db.Get(&user, query, request.Username)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
	}

	return c.JSON(fiber.Map{"message": "Login successful"})
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
