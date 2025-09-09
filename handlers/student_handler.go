package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EstudianteHandler struct {
	estudianteRepo *repositories.EstudianteRepository
}

func NewEstudianteHandler(estudianteRepo *repositories.EstudianteRepository) *EstudianteHandler {
	return &EstudianteHandler{estudianteRepo: estudianteRepo}
}

// CreateEstudiante crea un nuevo estudiante
func (h *EstudianteHandler) CreateEstudiante(c *fiber.Ctx) error {
	var estudiante models.Estudiante
	
	if err := c.BodyParser(&estudiante); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.estudianteRepo.CreateEstudiante(&estudiante); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede crear el estudiante",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(estudiante)
}

// GetEstudiante obtiene un estudiante por ID
func (h *EstudianteHandler) GetEstudiante(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante inválido",
		})
	}

	estudiante, err := h.estudianteRepo.GetEstudianteByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Estudiante no encontrado",
		})
	}

	return c.JSON(estudiante)
}

// GetAllEstudiantes obtiene todos los estudiantes
func (h *EstudianteHandler) GetAllEstudiantes(c *fiber.Ctx) error {
	estudiantes, err := h.estudianteRepo.GetAllEstudiantes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los estudiantes",
		})
	}

	return c.JSON(estudiantes)
}

// UpdateEstudiante actualiza un estudiante
func (h *EstudianteHandler) UpdateEstudiante(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante inválido",
		})
	}

	estudiante, err := h.estudianteRepo.GetEstudianteByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Estudiante no encontrado",
		})
	}

	if err := c.BodyParser(estudiante); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.estudianteRepo.UpdateEstudiante(estudiante); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede actualizar el estudiante",
		})
	}

	return c.JSON(estudiante)
}

// DeleteEstudiante elimina un estudiante y en cascada su usuario y persona
func (h *EstudianteHandler) DeleteEstudiante(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante inválido",
		})
	}

	if err := h.estudianteRepo.DeleteEstudiante(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede eliminar el estudiante y sus datos relacionados",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Estudiante, usuario y persona eliminados exitosamente",
	})
}

// RestoreEstudiante restaura un estudiante eliminado y en cascada su usuario y persona
func (h *EstudianteHandler) RestoreEstudiante(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante inválido",
		})
	}

	if err := h.estudianteRepo.RestoreEstudiante(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede restaurar el estudiante y sus datos relacionados",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Estudiante, usuario y persona restaurados exitosamente",
	})
}

// GetEstudiantesByCity obtiene estudiantes por ciudad
func (h *EstudianteHandler) GetEstudiantesByCity(c *fiber.Ctx) error {
	ciudadID, err := strconv.Atoi(c.Params("ciudad_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de ciudad inválido",
		})
	}
	
	estudiantes, err := h.estudianteRepo.GetEstudiantesByCity(uint(ciudadID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los estudiantes",
		})
	}

	return c.JSON(estudiantes)
}

// GetEstudiantesByInstitucion obtiene estudiantes por institución
func (h *EstudianteHandler) GetEstudiantesByInstitucion(c *fiber.Ctx) error {
	institucionID, err := strconv.Atoi(c.Params("institucion_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de institución inválido",
		})
	}
	
	estudiantes, err := h.estudianteRepo.GetEstudiantesByInstitucion(uint(institucionID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los estudiantes",
		})
	}

	return c.JSON(estudiantes)
}

// GetEstudiantesByEspecialidad obtiene estudiantes por especialidad
func (h *EstudianteHandler) GetEstudiantesByEspecialidad(c *fiber.Ctx) error {
	especialidad := c.Params("especialidad")
	
	estudiantes, err := h.estudianteRepo.GetEstudiantesByEspecialidad(especialidad)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los estudiantes",
		})
	}

	return c.JSON(estudiantes)
}