package handlers

import (
	"strconv"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ImageHandler interface {
	CreateImageHandler(c *fiber.Ctx) error
	DeleteImageHandler(c *fiber.Ctx) error
}

type imageHandler struct {
	imageService services.ImageService
}

func NewImageHandler(imageService services.ImageService) ImageHandler {
	return &imageHandler{
		imageService: imageService,
	}
}

func (h *imageHandler) CreateImageHandler(c *fiber.Ctx) error {
	projectId := c.Params("project_id")
	projectidUint, err := strconv.ParseUint(projectId, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	form, err := c.MultipartForm()
	if err != nil {
		return fiber.ErrBadRequest
	}
	filesHeader := form.File["images"]
	if len(filesHeader) == 0 {
		return fiber.ErrBadRequest
	}

	var response []dto.ImageResponse
	for _, fileHeader := range filesHeader {
		file, err := fileHeader.Open()
		if err != nil {
			return fiber.ErrBadRequest
		}
		defer file.Close()
		image, err := h.imageService.CreateImage(c.Context(), file, uint(projectidUint))
		if err != nil {
			return fiber.ErrInternalServerError
		}

		response = append(response, dto.ImageResponse{
			URL: image.URL,
		})
	}

	return c.JSON(response)
}

func (h *imageHandler) DeleteImageHandler(c *fiber.Ctx) error {
	projectId := c.Params("project_id")
	imageIdUint, err := strconv.ParseUint(projectId, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.imageService.Delete(c.Context(), uint(imageIdUint)); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(204).JSON(nil)
}
