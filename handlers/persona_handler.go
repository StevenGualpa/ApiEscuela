package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PersonaHandler struct {
	personaRepo *repositories.PersonaRepository
}

func NewPersonaHandler(personaRepo *repositories.PersonaRepository) *PersonaHandler {
	return &PersonaHandler{personaRepo: personaRepo}
}

// CreatePersona crea una nueva persona
func (h *PersonaHandler) CreatePersona(c *fiber.Ctx) error {
	var persona models.Persona
	
	if err := c.BodyParser(&persona); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.personaRepo.CreatePersona(&persona); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede crear la persona",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(persona)
}

// GetPersona obtiene una persona por ID
func (h *PersonaHandler) GetPersona(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de persona inválido",
		})
	}

	persona, err := h.personaRepo.GetPersonaByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Persona no encontrada",
		})
	}

	return c.JSON(persona)
}

// GetAllPersonas obtiene todas las personas
func (h *PersonaHandler) GetAllPersonas(c *fiber.Ctx) error {
	personas, err := h.personaRepo.GetAllPersonas()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las personas",
		})
	}

	return c.JSON(personas)
}

// UpdatePersona actualiza una persona
func (h *PersonaHandler) UpdatePersona(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de persona inválido",
		})
	}

	persona, err := h.personaRepo.GetPersonaByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Persona no encontrada",
		})
	}

	if err := c.BodyParser(persona); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.personaRepo.UpdatePersona(persona); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede actualizar la persona",
		})
	}

	return c.JSON(persona)
}

// DeletePersona elimina una persona
func (h *PersonaHandler) DeletePersona(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de persona inválido",
		})
	}

	if err := h.personaRepo.DeletePersona(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede eliminar la persona",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Persona eliminada exitosamente",
	})
}

// GetPersonaByCedula obtiene una persona por cédula
func (h *PersonaHandler) GetPersonaByCedula(c *fiber.Ctx) error {
	cedula := c.Params("cedula")
	
	persona, err := h.personaRepo.GetPersonaByCedula(cedula)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Persona no encontrada",
		})
	}

	return c.JSON(persona)
}

// GetPersonasByCorreo obtiene personas por correo
func (h *PersonaHandler) GetPersonasByCorreo(c *fiber.Ctx) error {
	correo := c.Params("correo")
	
	personas, err := h.personaRepo.GetPersonasByCorreo(correo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las personas",
		})
	}

	return c.JSON(personas)
}