package handlers

import (
	"strconv"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	LoginUserHandler(c *fiber.Ctx) error
	GetAllUsersHandler(c *fiber.Ctx) error
	UpdateUserHandler(c *fiber.Ctx) error
	DeleteUserHandler(c *fiber.Ctx) error
	CreateUserHandler(c *fiber.Ctx) error
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) LoginUserHandler(c *fiber.Ctx) error {
	var input dto.UserRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	token, err := h.userService.LoginUser(c.Context(), &input)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": token,
	})
}

func (h *userHandler) GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUser(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *userHandler) UpdateUserHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	var input dto.UserRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	user, err := h.userService.UpdateUser(c.Context(), uint(idUint), &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *userHandler) DeleteUserHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.userService.DeleteUser(c.Context(), uint(idUint))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusNoContent).JSON(nil)
}

func (h *userHandler) CreateUserHandler(c *fiber.Ctx) error {
	var input dto.UserRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	user, err := h.userService.CreateUser(c.Context(), &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
