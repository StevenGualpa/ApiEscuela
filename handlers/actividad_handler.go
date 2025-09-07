package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ActividadHandler struct {
	actividadRepo *repositories.ActividadRepository
}

func NewActividadHandler(actividadRepo *repositories.ActividadRepository) *ActividadHandler {
	return &ActividadHandler{actividadRepo: actividadRepo}
}

// CreateActividad crea una nueva actividad
func (h *ActividadHandler) CreateActividad(c *fiber.Ctx) error {
	var actividad models.Actividad
	
	if err := c.BodyParser(&actividad); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.actividadRepo.CreateActividad(&actividad); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede crear la actividad",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(actividad)
}

// GetActividad obtiene una actividad por ID
func (h *ActividadHandler) GetActividad(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de actividad inválido",
		})
	}

	actividad, err := h.actividadRepo.GetActividadByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Actividad no encontrada",
		})
	}

	return c.JSON(actividad)
}

// GetAllActividades obtiene todas las actividades
func (h *ActividadHandler) GetAllActividades(c *fiber.Ctx) error {
	actividades, err := h.actividadRepo.GetAllActividades()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las actividades",
		})
	}

	return c.JSON(actividades)
}

// UpdateActividad actualiza una actividad
func (h *ActividadHandler) UpdateActividad(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de actividad inválido",
		})
	}

	actividad, err := h.actividadRepo.GetActividadByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Actividad no encontrada",
		})
	}

	if err := c.BodyParser(actividad); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.actividadRepo.UpdateActividad(actividad); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede actualizar la actividad",
		})
	}

	return c.JSON(actividad)
}

// DeleteActividad elimina una actividad
func (h *ActividadHandler) DeleteActividad(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de actividad inválido",
		})
	}

	if err := h.actividadRepo.DeleteActividad(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede eliminar la actividad",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Actividad eliminada exitosamente",
	})
}

// GetActividadesByTematica obtiene actividades por temática
func (h *ActividadHandler) GetActividadesByTematica(c *fiber.Ctx) error {
	tematicaID, err := strconv.Atoi(c.Params("tematica_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de temática inválido",
		})
	}
	
	actividades, err := h.actividadRepo.GetActividadesByTematica(uint(tematicaID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las actividades",
		})
	}

	return c.JSON(actividades)
}

// GetActividadesByNombre busca actividades por nombre
func (h *ActividadHandler) GetActividadesByNombre(c *fiber.Ctx) error {
	nombre := c.Params("nombre")
	
	actividades, err := h.actividadRepo.GetActividadesByNombre(nombre)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las actividades",
		})
	}

	return c.JSON(actividades)
}

// GetActividadesByDuracion obtiene actividades por rango de duración
func (h *ActividadHandler) GetActividadesByDuracion(c *fiber.Ctx) error {
	duracionMin, err := strconv.Atoi(c.Query("min", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Duración mínima inválida",
		})
	}

	duracionMax, err := strconv.Atoi(c.Query("max", "999"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Duración máxima inválida",
		})
	}
	
	actividades, err := h.actividadRepo.GetActividadesByDuracion(duracionMin, duracionMax)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las actividades",
		})
	}

	return c.JSON(actividades)
}