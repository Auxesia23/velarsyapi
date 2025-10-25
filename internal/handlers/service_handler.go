package handlers

import (
	"strconv"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ServiceHandler interface {
	CreateServiceHandler(c *fiber.Ctx) error
	GetAllServicesHandler(c *fiber.Ctx) error
	UpdateServiceHandler(c *fiber.Ctx) error
	DeleteServiceHandler(c *fiber.Ctx) error
}

type serviceHandler struct {
	serviceService services.ServiceService
}

func NewServiceHandler(serviceService services.ServiceService) ServiceHandler {
	return &serviceHandler{
		serviceService: serviceService,
	}
}

func (h *serviceHandler) CreateServiceHandler(c *fiber.Ctx) error {
	var input dto.ServiceRequest
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	createdService, err := h.serviceService.CreateService(c.Context(), &input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(createdService)
}

func (h *serviceHandler) GetAllServicesHandler(c *fiber.Ctx) error {
	serviceList, err := h.serviceService.GetAllServices(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(200).JSON(serviceList)
}

func (h *serviceHandler) UpdateServiceHandler(c *fiber.Ctx) error {
	var input dto.ServiceRequest
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	updatedService, err := h.serviceService.UpdateService(c.Context(), &input, uint(idUint))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(updatedService)
}

func (h *serviceHandler) DeleteServiceHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.serviceService.DeleteService(c.Context(), uint(idUint))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(nil)
}
