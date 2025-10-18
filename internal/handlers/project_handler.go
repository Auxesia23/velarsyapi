package handlers

import (
	"strconv"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Projecthandler interface {
	CreateProjectHandler(c *fiber.Ctx) error
	GetSingleProjectHandler(c *fiber.Ctx) error
	UpdateProjectHandler(c *fiber.Ctx) error
	DeleteProjectHandler(c *fiber.Ctx) error
}

type projecthandler struct {
	projectService services.ProjectService
}

func NewProjectHandler(projectService services.ProjectService) Projecthandler {
	return &projecthandler{projectService}
}

func (h *projecthandler) CreateProjectHandler(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return fiber.ErrBadRequest
	}
	var input dto.ProjectRequest
	if err := c.BodyParser(&input); err != nil {
		return fiber.ErrBadRequest
	}
	workId := c.Params("work_id")
	workIdInt, err := strconv.Atoi(workId)
	if err != nil {
		return fiber.ErrBadRequest
	}
	workIdUint := uint(workIdInt)

	fileHeader := form.File["image"]
	if len(fileHeader) == 0 {
		return fiber.ErrBadRequest
	}
	image, err := fileHeader[0].Open()
	if err != nil {
		return fiber.ErrBadRequest
	}
	createdProject, err := h.projectService.CreateProject(c.Context(), &input, &image, &workIdUint)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(createdProject)
}

func (h *projecthandler) GetSingleProjectHandler(c *fiber.Ctx) error {
	slug := c.Params("project_slug")

	project, err := h.projectService.GetSingleProject(c.Context(), &slug)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(project)
}

func (h *projecthandler) UpdateProjectHandler(c *fiber.Ctx) error {
	projectId := c.Params("project_id")
	projectIdInt, err := strconv.Atoi(projectId)
	if err != nil {
		return fiber.ErrBadRequest
	}
	projectIdUint := uint(projectIdInt)
	form, err := c.MultipartForm()
	if err != nil {
		return fiber.ErrBadRequest
	}
	var input dto.ProjectRequest
	if err := c.BodyParser(&input); err != nil {
		return fiber.ErrBadRequest
	}
	fileHeader := form.File["image"]
	if len(fileHeader) == 0 {
		return fiber.ErrBadRequest
	}
	image, err := fileHeader[0].Open()
	if err != nil {
		return fiber.ErrBadRequest
	}

	updatedProject, err := h.projectService.UpdateProject(c.Context(), &input, &image, &projectIdUint)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(updatedProject)
}

func (h *projecthandler) DeleteProjectHandler(c *fiber.Ctx) error {
	projectId := c.Params("project_id")
	projectIdInt, err := strconv.Atoi(projectId)
	if err != nil {
		return fiber.ErrBadRequest
	}
	projectIdUint := uint(projectIdInt)
	err = h.projectService.DeleteProject(c.Context(), &projectIdUint)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(204).JSON(nil)
}
