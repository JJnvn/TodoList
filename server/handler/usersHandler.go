package handler

import (
	"todo-list/service"
	"todo-list/utils"

	"github.com/gofiber/fiber/v2"
)

type usersHandler struct {
	service service.UsersService
}

func NewUsersHandler(service service.UsersService) *usersHandler {
	return &usersHandler{service: service}
}

func (u usersHandler) HandleSignUp(c *fiber.Ctx) error {
	s := u.service
	request := service.UsersRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "handler error")
	}
	response, err := s.UserSignUp(request)
	if err != nil {
		return err
	}
	return c.JSON(response)
}

func (u usersHandler) HandleLogin(c *fiber.Ctx) error {
	s := u.service
	request := service.UsersRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "handler error")
	}
	response, err := s.UserLogin(request)
	if err != nil {
		return err
	}
	token, err := utils.GenerateJWT(response.Username)
	utils.SetJWTInCookie(c, token)
	return c.JSON(response)
}
