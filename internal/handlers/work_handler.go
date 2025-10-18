package handlers

import (
	"strconv"

	"github.com/Auxesia23/velarsyapi/internal/services"
	"github.com/gofiber/fiber/v2"
)

type WorkHandler interface {
	CreateWorkHandler(c *fiber.Ctx) error
	GetAllWorkHandler(c *fiber.Ctx) error
	GetSingleWorkHandler(c *fiber.Ctx) error
	UpdateWorkHandler(c *fiber.Ctx) error
	DeleteWorkHandler(c *fiber.Ctx) error
}

type workHandler struct {
	workService services.WorkService
}

func NewWorkHandler(workService services.WorkService) WorkHandler {
	return &workHandler{
		workService: workService,
	}
}

func (h *workHandler) CreateWorkHandler(c *fiber.Ctx) error {
	title := c.FormValue("title")
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return fiber.ErrBadRequest
	}

	file, err := fileHeader.Open()
	if err != nil {
		return fiber.ErrBadRequest
	}
	defer file.Close()

	work, err := h.workService.CreateWork(c.Context(), &title, &file)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(work)
}

func (h *workHandler) GetAllWorkHandler(c *fiber.Ctx) error {
	works, err := h.workService.GetAllWork(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(works)
}

func (h *workHandler) GetSingleWorkHandler(c *fiber.Ctx) error {
	slug := c.Params("work_slug")
	

	work, err := h.workService.GetOneWork(c.Context(), &slug)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(work)
}

func (h *workHandler) UpdateWorkHandler(c *fiber.Ctx) error {
	id := c.Params("work_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}
	idUint := uint(idInt)

	title := c.FormValue("title")
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return fiber.ErrBadRequest
	}

	file, err := fileHeader.Open()
	if err != nil {
		return fiber.ErrBadRequest
	}
	defer file.Close()

	updatedWork, err := h.workService.UpdateWork(c.Context(), &title, &file, &idUint)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(updatedWork)
}

func (h *workHandler) DeleteWorkHandler(c *fiber.Ctx) error {
	id := c.Params("work_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fiber.ErrBadRequest
	}
	idUint := uint(idInt)

	err = h.workService.DeleteWork(c.Context(), &idUint)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(204).JSON(nil)
}
