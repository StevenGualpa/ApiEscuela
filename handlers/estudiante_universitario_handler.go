package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EstudianteUniversitarioHandler struct {
	estudianteUnivRepo *repositories.EstudianteUniversitarioRepository
}

func NewEstudianteUniversitarioHandler(estudianteUnivRepo *repositories.EstudianteUniversitarioRepository) *EstudianteUniversitarioHandler {
	return &EstudianteUniversitarioHandler{estudianteUnivRepo: estudianteUnivRepo}
}

// CreateEstudianteUniversitario crea un nuevo estudiante universitario
func (h *EstudianteUniversitarioHandler) CreateEstudianteUniversitario(c *fiber.Ctx) error {
	var estudiante models.EstudianteUniversitario
	
	if err := c.BodyParser(&estudiante); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.estudianteUnivRepo.CreateEstudianteUniversitario(&estudiante); err != nil {
		switch err {
		case repositories.ErrEstudianteUnivDuplicado:
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "estudiante universitario ya existe"})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se puede crear el estudiante universitario"})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(estudiante)
}

// GetEstudianteUniversitario obtiene un estudiante universitario por ID
func (h *EstudianteUniversitarioHandler) GetEstudianteUniversitario(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante universitario inválido",
		})
	}

	estudiante, err := h.estudianteUnivRepo.GetEstudianteUniversitarioByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Estudiante universitario no encontrado",
		})
	}

	return c.JSON(estudiante)
}

// GetAllEstudiantesUniversitarios obtiene todos los estudiantes universitarios
func (h *EstudianteUniversitarioHandler) GetAllEstudiantesUniversitarios(c *fiber.Ctx) error {
	estudiantes, err := h.estudianteUnivRepo.GetAllEstudiantesUniversitarios()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los estudiantes universitarios",
		})
	}

	return c.JSON(estudiantes)
}

// UpdateEstudianteUniversitario actualiza un estudiante universitario
func (h *EstudianteUniversitarioHandler) UpdateEstudianteUniversitario(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante universitario inválido",
		})
	}

	estudiante, err := h.estudianteUnivRepo.GetEstudianteUniversitarioByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Estudiante universitario no encontrado",
		})
	}

	if err := c.BodyParser(estudiante); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.estudianteUnivRepo.UpdateEstudianteUniversitario(estudiante); err != nil {
		switch err {
		case repositories.ErrEstudianteUnivDuplicado:
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "estudiante universitario ya existe"})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se puede actualizar el estudiante universitario"})
		}
	}

	return c.JSON(estudiante)
}

// DeleteEstudianteUniversitario elimina un estudiante universitario
func (h *EstudianteUniversitarioHandler) DeleteEstudianteUniversitario(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante universitario inválido",
		})
	}

	if err := h.estudianteUnivRepo.DeleteEstudianteUniversitario(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede eliminar el estudiante universitario",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Estudiante universitario eliminado exitosamente",
	})
}

// GetEstudiantesUniversitariosBySemestre obtiene estudiantes por semestre
func (h *EstudianteUniversitarioHandler) GetEstudiantesUniversitariosBySemestre(c *fiber.Ctx) error {
	semestre, err := strconv.Atoi(c.Params("semestre"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Semestre inválido",
		})
	}
	
	estudiantes, err := h.estudianteUnivRepo.GetEstudiantesUniversitariosBySemestre(semestre)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los estudiantes universitarios",
		})
	}

	return c.JSON(estudiantes)
}

// GetEstudianteUniversitarioByPersona obtiene estudiante universitario por persona
func (h *EstudianteUniversitarioHandler) GetEstudianteUniversitarioByPersona(c *fiber.Ctx) error {
	personaID, err := strconv.Atoi(c.Params("persona_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de persona inválido",
		})
	}
	
	estudiante, err := h.estudianteUnivRepo.GetEstudianteUniversitarioByPersona(uint(personaID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Estudiante universitario no encontrado",
		})
	}

	return c.JSON(estudiante)
}