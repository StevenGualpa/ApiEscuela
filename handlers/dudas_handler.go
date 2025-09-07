package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DudasHandler struct {
	dudasRepo *repositories.DudasRepository
}

func NewDudasHandler(dudasRepo *repositories.DudasRepository) *DudasHandler {
	return &DudasHandler{dudasRepo: dudasRepo}
}

// CreateDudas crea una nueva duda
func (h *DudasHandler) CreateDudas(c *fiber.Ctx) error {
	var duda models.Dudas
	
	if err := c.BodyParser(&duda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.dudasRepo.CreateDudas(&duda); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede crear la duda",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(duda)
}

// GetDudas obtiene una duda por ID
func (h *DudasHandler) GetDudas(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de duda inválido",
		})
	}

	duda, err := h.dudasRepo.GetDudasByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Duda no encontrada",
		})
	}

	return c.JSON(duda)
}

// GetAllDudas obtiene todas las dudas
func (h *DudasHandler) GetAllDudas(c *fiber.Ctx) error {
	dudas, err := h.dudasRepo.GetAllDudas()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las dudas",
		})
	}

	return c.JSON(dudas)
}

// UpdateDudas actualiza una duda
func (h *DudasHandler) UpdateDudas(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de duda inválido",
		})
	}

	duda, err := h.dudasRepo.GetDudasByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Duda no encontrada",
		})
	}

	if err := c.BodyParser(duda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.dudasRepo.UpdateDudas(duda); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede actualizar la duda",
		})
	}

	return c.JSON(duda)
}

// DeleteDudas elimina una duda
func (h *DudasHandler) DeleteDudas(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de duda inválido",
		})
	}

	if err := h.dudasRepo.DeleteDudas(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede eliminar la duda",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Duda eliminada exitosamente",
	})
}

// GetDudasByEstudiante obtiene dudas por estudiante
func (h *DudasHandler) GetDudasByEstudiante(c *fiber.Ctx) error {
	estudianteID, err := strconv.Atoi(c.Params("estudiante_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante inválido",
		})
	}
	
	dudas, err := h.dudasRepo.GetDudasByEstudiante(uint(estudianteID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las dudas",
		})
	}

	return c.JSON(dudas)
}

// GetDudasByAutoridad obtiene dudas asignadas a una autoridad
func (h *DudasHandler) GetDudasByAutoridad(c *fiber.Ctx) error {
	autoridadID, err := strconv.Atoi(c.Params("autoridad_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de autoridad inválido",
		})
	}
	
	dudas, err := h.dudasRepo.GetDudasByAutoridad(uint(autoridadID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las dudas",
		})
	}

	return c.JSON(dudas)
}

// GetDudasSinResponder obtiene dudas sin respuesta
func (h *DudasHandler) GetDudasSinResponder(c *fiber.Ctx) error {
	dudas, err := h.dudasRepo.GetDudasSinResponder()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las dudas",
		})
	}

	return c.JSON(dudas)
}

// GetDudasRespondidas obtiene dudas con respuesta
func (h *DudasHandler) GetDudasRespondidas(c *fiber.Ctx) error {
	dudas, err := h.dudasRepo.GetDudasRespondidas()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las dudas",
		})
	}

	return c.JSON(dudas)
}

// GetDudasSinAsignar obtiene dudas sin autoridad asignada
func (h *DudasHandler) GetDudasSinAsignar(c *fiber.Ctx) error {
	dudas, err := h.dudasRepo.GetDudasSinAsignar()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las dudas",
		})
	}

	return c.JSON(dudas)
}

// BuscarDudasPorPregunta busca dudas por contenido de la pregunta
func (h *DudasHandler) BuscarDudasPorPregunta(c *fiber.Ctx) error {
	termino := c.Params("termino")
	
	dudas, err := h.dudasRepo.BuscarDudasPorPregunta(termino)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener las dudas",
		})
	}

	return c.JSON(dudas)
}

// AsignarAutoridadADuda asigna una autoridad a una duda
func (h *DudasHandler) AsignarAutoridadADuda(c *fiber.Ctx) error {
	dudaID, err := strconv.Atoi(c.Params("duda_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de duda inválido",
		})
	}

	var requestData struct {
		AutoridadID uint `json:"autoridad_id"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.dudasRepo.AsignarAutoridadADuda(uint(dudaID), requestData.AutoridadID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede asignar la autoridad a la duda",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Autoridad asignada exitosamente",
	})
}

// ResponderDuda actualiza la respuesta de una duda
func (h *DudasHandler) ResponderDuda(c *fiber.Ctx) error {
	dudaID, err := strconv.Atoi(c.Params("duda_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de duda inválido",
		})
	}

	var requestData struct {
		Respuesta string `json:"respuesta"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.dudasRepo.ResponderDuda(uint(dudaID), requestData.Respuesta); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede responder la duda",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Duda respondida exitosamente",
	})
}