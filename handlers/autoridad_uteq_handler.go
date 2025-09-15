package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AutoridadUTEQHandler struct {
	autoridadRepo *repositories.AutoridadUTEQRepository
}

func NewAutoridadUTEQHandler(autoridadRepo *repositories.AutoridadUTEQRepository) *AutoridadUTEQHandler {
	return &AutoridadUTEQHandler{autoridadRepo: autoridadRepo}
}

// CreateAutoridadUTEQ crea una nueva autoridad UTEQ
func (h *AutoridadUTEQHandler) CreateAutoridadUTEQ(c *fiber.Ctx) error {
	var autoridad models.AutoridadUTEQ
	
	if err := c.BodyParser(&autoridad); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.autoridadRepo.CreateAutoridadUTEQ(&autoridad); err != nil {
		switch err {
		case repositories.ErrAutoridadDuplicada:
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "autoridad ya existe"})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se puede crear la autoridad UTEQ"})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(autoridad)
}

// GetAutoridadUTEQ obtiene una autoridad UTEQ por ID
func (h *AutoridadUTEQHandler) GetAutoridadUTEQ(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de autoridad UTEQ inválido",
		})
	}

	autoridad, err := h.autoridadRepo.GetAutoridadUTEQByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Autoridad UTEQ no encontrada",
		})
	}

	return c.JSON(autoridad)
}

// GetAllAutoridadesUTEQ obtiene todas las autoridades UTEQ activas
func (h *AutoridadUTEQHandler) GetAllAutoridadesUTEQ(c *fiber.Ctx) error {
	autoridades, err := h.autoridadRepo.GetAllAutoridadesUTEQ()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las autoridades UTEQ",
		})
	}

	return c.JSON(autoridades)
}

// GetAllAutoridadesUTEQIncludingDeleted obtiene todas las autoridades UTEQ incluyendo las eliminadas
func (h *AutoridadUTEQHandler) GetAllAutoridadesUTEQIncludingDeleted(c *fiber.Ctx) error {
	autoridades, err := h.autoridadRepo.GetAllAutoridadesUTEQIncludingDeleted()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las autoridades UTEQ",
		})
	}

	return c.JSON(autoridades)
}

// GetDeletedAutoridadesUTEQ obtiene solo las autoridades UTEQ eliminadas
func (h *AutoridadUTEQHandler) GetDeletedAutoridadesUTEQ(c *fiber.Ctx) error {
	autoridades, err := h.autoridadRepo.GetDeletedAutoridadesUTEQ()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las autoridades UTEQ eliminadas",
		})
	}

	return c.JSON(autoridades)
}

// UpdateAutoridadUTEQ actualiza una autoridad UTEQ
func (h *AutoridadUTEQHandler) UpdateAutoridadUTEQ(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de autoridad UTEQ inválido",
		})
	}

	autoridad, err := h.autoridadRepo.GetAutoridadUTEQByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Autoridad UTEQ no encontrada",
		})
	}

	if err := c.BodyParser(autoridad); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.autoridadRepo.UpdateAutoridadUTEQ(autoridad); err != nil {
		switch err {
		case repositories.ErrAutoridadDuplicada:
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "autoridad ya existe"})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se puede actualizar la autoridad UTEQ"})
		}
	}

	return c.JSON(autoridad)
}

// DeleteAutoridadUTEQ elimina una autoridad UTEQ y en cascada su usuario y persona
func (h *AutoridadUTEQHandler) DeleteAutoridadUTEQ(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de autoridad UTEQ inválido",
		})
	}

	if err := h.autoridadRepo.DeleteAutoridadUTEQ(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede eliminar la autoridad UTEQ y sus datos relacionados",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Autoridad UTEQ, usuario y persona eliminados exitosamente",
	})
}

// RestoreAutoridadUTEQ restaura una autoridad UTEQ eliminada y en cascada su usuario y persona
func (h *AutoridadUTEQHandler) RestoreAutoridadUTEQ(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de autoridad UTEQ inválido",
		})
	}

	if err := h.autoridadRepo.RestoreAutoridadUTEQ(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede restaurar la autoridad UTEQ y sus datos relacionados",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Autoridad UTEQ, usuario y persona restaurados exitosamente",
	})
}

// GetAutoridadesUTEQByCargo obtiene autoridades por cargo
func (h *AutoridadUTEQHandler) GetAutoridadesUTEQByCargo(c *fiber.Ctx) error {
	cargo := c.Params("cargo")
	
	autoridades, err := h.autoridadRepo.GetAutoridadesUTEQByCargo(cargo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las autoridades UTEQ",
		})
	}

	return c.JSON(autoridades)
}

// GetAutoridadUTEQByPersona obtiene autoridad UTEQ por persona
func (h *AutoridadUTEQHandler) GetAutoridadUTEQByPersona(c *fiber.Ctx) error {
	personaID, err := strconv.Atoi(c.Params("persona_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de persona inválido",
		})
	}
	
	autoridad, err := h.autoridadRepo.GetAutoridadUTEQByPersona(uint(personaID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Autoridad UTEQ no encontrada",
		})
	}

	return c.JSON(autoridad)
}