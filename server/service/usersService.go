package service

import (
	"todo-list/repository"
	"todo-list/utils"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

type usersServiceImpl struct {
	repo repository.UsersRepo
}

func NewUsersServiceImpl(repo repository.UsersRepo) *usersServiceImpl {
	return &usersServiceImpl{
		repo: repo,
	}
}

func (u *usersServiceImpl) userLogin(request UsersRequest) (*UsersResponse, error) {
	err := utils.HandleEmptyUserOrPass(request.Username, request.Password)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotAcceptable, "username of password cannot be empty")
	}

	repo := u.repo
	user, err := repo.GetUserByUsername(request.Username)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
	}

	response := UsersResponse{
		Username: user.Username,
		Message:  "Login successful",
	}

	return &response, nil
}

func (u *usersServiceImpl) userSignUp(request UsersRequest) (*UsersResponse, error) {
	err := utils.HandleEmptyUserOrPass(request.Username, request.Password)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotAcceptable, "username of password cannot be empty")
	}

	repo := u.repo
	response := UsersResponse{}

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user, err := repo.CreateUser(request.Username, string(hashed_password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "cannot create user or user already exist")
	}

	response = UsersResponse{
		Username: user.Username,
		Message:  "user created",
	}

	return &response, nil
}
