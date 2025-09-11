package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type VisitaDetalleHandler struct {
	visitaDetalleRepo *repositories.VisitaDetalleRepository
}

func NewVisitaDetalleHandler(visitaDetalleRepo *repositories.VisitaDetalleRepository) *VisitaDetalleHandler {
	return &VisitaDetalleHandler{visitaDetalleRepo: visitaDetalleRepo}
}

// CreateVisitaDetalle crea un nuevo detalle de visita
func (h *VisitaDetalleHandler) CreateVisitaDetalle(c *fiber.Ctx) error {
	var detalle models.VisitaDetalle
	
	if err := c.BodyParser(&detalle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.visitaDetalleRepo.CreateVisitaDetalle(&detalle); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede crear el detalle de visita",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(detalle)
}

// GetVisitaDetalle obtiene un detalle de visita por ID
func (h *VisitaDetalleHandler) GetVisitaDetalle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de detalle de visita inválido",
		})
	}

	detalle, err := h.visitaDetalleRepo.GetVisitaDetalleByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Detalle de visita no encontrado",
		})
	}

	return c.JSON(detalle)
}

// GetAllVisitaDetalles obtiene todos los detalles de visita
func (h *VisitaDetalleHandler) GetAllVisitaDetalles(c *fiber.Ctx) error {
	detalles, err := h.visitaDetalleRepo.GetAllVisitaDetalles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los detalles de visita",
		})
	}

	return c.JSON(detalles)
}

// UpdateVisitaDetalle actualiza un detalle de visita
func (h *VisitaDetalleHandler) UpdateVisitaDetalle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de detalle de visita inválido",
		})
	}

	detalle, err := h.visitaDetalleRepo.GetVisitaDetalleByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Detalle de visita no encontrado",
		})
	}

	if err := c.BodyParser(detalle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.visitaDetalleRepo.UpdateVisitaDetalle(detalle); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede actualizar el detalle de visita",
		})
	}

	return c.JSON(detalle)
}

// DeleteVisitaDetalle elimina un detalle de visita
func (h *VisitaDetalleHandler) DeleteVisitaDetalle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de detalle de visita inválido",
		})
	}

	if err := h.visitaDetalleRepo.DeleteVisitaDetalle(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede eliminar el detalle de visita",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Detalle de visita eliminado exitosamente",
	})
}

// GetVisitaDetallesByActividad obtiene detalles por actividad
func (h *VisitaDetalleHandler) GetVisitaDetallesByActividad(c *fiber.Ctx) error {
	actividadID, err := strconv.Atoi(c.Params("actividad_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de actividad inválido",
		})
	}
	
	detalles, err := h.visitaDetalleRepo.GetVisitaDetallesByActividad(uint(actividadID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los detalles de visita",
		})
	}

	return c.JSON(detalles)
}

// GetVisitaDetallesByPrograma obtiene detalles por programa de visita
func (h *VisitaDetalleHandler) GetVisitaDetallesByPrograma(c *fiber.Ctx) error {
	programaID, err := strconv.Atoi(c.Params("programa_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de programa inválido",
		})
	}
	
	detalles, err := h.visitaDetalleRepo.GetVisitaDetallesByPrograma(uint(programaID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los detalles de visita",
		})
	}

	return c.JSON(detalles)
}

// GetVisitaDetallesByParticipantes obtiene detalles por rango de participantes
func (h *VisitaDetalleHandler) GetVisitaDetallesByParticipantes(c *fiber.Ctx) error {
	minParticipantes, err := strconv.Atoi(c.Query("min", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Número mínimo de participantes inválido",
		})
	}

	maxParticipantes, err := strconv.Atoi(c.Query("max", "999"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Número máximo de participantes inválido",
		})
	}
	
	detalles, err := h.visitaDetalleRepo.GetVisitaDetallesByParticipantes(minParticipantes, maxParticipantes)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los detalles de visita",
		})
	}

	return c.JSON(detalles)
}

// GetEstadisticasParticipacion obtiene estadísticas de participación
func (h *VisitaDetalleHandler) GetEstadisticasParticipacion(c *fiber.Ctx) error {
	estadisticas, err := h.visitaDetalleRepo.GetEstadisticasParticipacion()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las estadísticas",
		})
	}

	return c.JSON(estadisticas)
}