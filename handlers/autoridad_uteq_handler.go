package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"regexp"
	"strconv"
	"strings"

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

	// Parsear JSON
	if err := c.BodyParser(&autoridad); err != nil {
		return SendError(c, 400, "invalid_json", "No se puede procesar el JSON. Verifique el formato de los datos", err.Error())
	}

	// Validar campos requeridos
	if validationErrors := h.validateAutoridadUTEQRequiredFields(&autoridad); len(validationErrors) > 0 {
		return SendValidationError(c, "Faltan campos requeridos", validationErrors)
	}

	// Validar datos completos
	if validationErrors := h.validateAutoridadUTEQ(&autoridad, false); len(validationErrors) > 0 {
		return SendValidationError(c, "Los datos proporcionados no son válidos", validationErrors)
	}

	// Limpiar datos
	autoridad.Cargo = strings.TrimSpace(autoridad.Cargo)

	// Crear autoridad
	if err := h.autoridadRepo.CreateAutoridadUTEQ(&autoridad); err != nil {
		switch err {
		case repositories.ErrAutoridadDuplicada:
			return SendError(c, 409, "duplicate_autoridad", "Ya existe una autoridad con esta persona", "La persona ya tiene un cargo asignado")
		default:
			return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo crear la autoridad UTEQ")
		}
	}

	return SendSuccess(c, 201, autoridad)
}

// GetAutoridadUTEQ obtiene una autoridad UTEQ por ID
func (h *AutoridadUTEQHandler) GetAutoridadUTEQ(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID de la autoridad UTEQ es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID de la autoridad UTEQ no es válido", "El ID debe ser un número entero positivo")
	}

	autoridad, err := h.autoridadRepo.GetAutoridadUTEQByID(uint(id))
	if err != nil {
		return SendError(c, 404, "autoridad_not_found", "No se encontró la autoridad UTEQ solicitada", "Verifique que el ID sea correcto")
	}

	return SendSuccess(c, 200, autoridad)
}

// GetAllAutoridadesUTEQ obtiene todas las autoridades UTEQ activas
func (h *AutoridadUTEQHandler) GetAllAutoridadesUTEQ(c *fiber.Ctx) error {
	autoridades, err := h.autoridadRepo.GetAllAutoridadesUTEQ()
	if err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudieron obtener las autoridades UTEQ")
	}

	return SendSuccess(c, 200, autoridades)
}

// GetAllAutoridadesUTEQIncludingDeleted obtiene todas las autoridades UTEQ incluyendo las eliminadas
func (h *AutoridadUTEQHandler) GetAllAutoridadesUTEQIncludingDeleted(c *fiber.Ctx) error {
	autoridades, err := h.autoridadRepo.GetAllAutoridadesUTEQIncludingDeleted()
	if err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudieron obtener las autoridades UTEQ")
	}

	return SendSuccess(c, 200, autoridades)
}

// GetDeletedAutoridadesUTEQ obtiene solo las autoridades UTEQ eliminadas
func (h *AutoridadUTEQHandler) GetDeletedAutoridadesUTEQ(c *fiber.Ctx) error {
	autoridades, err := h.autoridadRepo.GetDeletedAutoridadesUTEQ()
	if err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudieron obtener las autoridades UTEQ eliminadas")
	}

	return SendSuccess(c, 200, autoridades)
}

// UpdateAutoridadUTEQ actualiza una autoridad UTEQ
func (h *AutoridadUTEQHandler) UpdateAutoridadUTEQ(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID de la autoridad UTEQ es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID de la autoridad UTEQ no es válido", "El ID debe ser un número entero positivo")
	}

	// Verificar que la autoridad existe
	existingAutoridad, err := h.autoridadRepo.GetAutoridadUTEQByID(uint(id))
	if err != nil {
		return SendError(c, 404, "autoridad_not_found", "No se encontró la autoridad UTEQ solicitada", "Verifique que el ID sea correcto")
	}

	// Parsear datos de actualización
	var updateData models.AutoridadUTEQ
	if err := c.BodyParser(&updateData); err != nil {
		return SendError(c, 400, "invalid_json", "No se puede procesar el JSON. Verifique el formato de los datos", err.Error())
	}

	// Validar datos de actualización
	if validationErrors := h.validateAutoridadUTEQ(&updateData, true); len(validationErrors) > 0 {
		return SendValidationError(c, "Los datos proporcionados no son válidos", validationErrors)
	}

	// Actualizar campos
	existingAutoridad.PersonaID = updateData.PersonaID
	existingAutoridad.Cargo = strings.TrimSpace(updateData.Cargo)

	// Guardar cambios
	if err := h.autoridadRepo.UpdateAutoridadUTEQ(existingAutoridad); err != nil {
		switch err {
		case repositories.ErrAutoridadDuplicada:
			return SendError(c, 409, "duplicate_autoridad", "Ya existe una autoridad con esta persona", "La persona ya tiene un cargo asignado")
		default:
			return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo actualizar la autoridad UTEQ")
		}
	}

	return SendSuccess(c, 200, existingAutoridad)
}

// DeleteAutoridadUTEQ elimina una autoridad UTEQ y en cascada su usuario y persona
func (h *AutoridadUTEQHandler) DeleteAutoridadUTEQ(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID de la autoridad UTEQ es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID de la autoridad UTEQ no es válido", "El ID debe ser un número entero positivo")
	}

	// Verificar que la autoridad existe
	_, err = h.autoridadRepo.GetAutoridadUTEQByID(uint(id))
	if err != nil {
		return SendError(c, 404, "autoridad_not_found", "No se encontró la autoridad UTEQ solicitada", "Verifique que el ID sea correcto")
	}

	// Eliminar autoridad
	if err := h.autoridadRepo.DeleteAutoridadUTEQ(uint(id)); err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo eliminar la autoridad UTEQ y sus datos relacionados")
	}

	return SendSuccess(c, 200, fiber.Map{
		"message": "Autoridad UTEQ, usuario y persona eliminados exitosamente",
		"id":      id,
	})
}

// RestoreAutoridadUTEQ restaura una autoridad UTEQ eliminada y en cascada su usuario y persona
func (h *AutoridadUTEQHandler) RestoreAutoridadUTEQ(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID de la autoridad UTEQ es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID de la autoridad UTEQ no es válido", "El ID debe ser un número entero positivo")
	}

	// Restaurar autoridad
	if err := h.autoridadRepo.RestoreAutoridadUTEQ(uint(id)); err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo restaurar la autoridad UTEQ y sus datos relacionados")
	}

	return SendSuccess(c, 200, fiber.Map{
		"message": "Autoridad UTEQ, usuario y persona restaurados exitosamente",
		"id":      id,
	})
}

// GetAutoridadesUTEQByCargo obtiene autoridades por cargo
func (h *AutoridadUTEQHandler) GetAutoridadesUTEQByCargo(c *fiber.Ctx) error {
	cargo := c.Params("cargo")

	// Validar parámetro de búsqueda
	if validationErrors := h.validateAutoridadUTEQSearchParams(cargo); len(validationErrors) > 0 {
		return SendValidationError(c, "Los parámetros de búsqueda no son válidos", validationErrors)
	}

	autoridades, err := h.autoridadRepo.GetAutoridadesUTEQByCargo(cargo)
	if err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudieron obtener las autoridades UTEQ")
	}

	return SendSuccess(c, 200, autoridades)
}

// GetAutoridadUTEQByPersona obtiene autoridad UTEQ por persona
func (h *AutoridadUTEQHandler) GetAutoridadUTEQByPersona(c *fiber.Ctx) error {
	personaIDStr := c.Params("persona_id")
	if personaIDStr == "" {
		return SendError(c, 400, "missing_persona_id", "El ID de la persona es requerido", "Proporcione un ID válido")
	}

	personaID, err := strconv.Atoi(personaIDStr)
	if err != nil || personaID <= 0 {
		return SendError(c, 400, "invalid_persona_id", "El ID de la persona no es válido", "El ID debe ser un número entero positivo")
	}

	autoridad, err := h.autoridadRepo.GetAutoridadUTEQByPersona(uint(personaID))
	if err != nil {
		return SendError(c, 404, "autoridad_not_found", "No se encontró la autoridad UTEQ para esta persona", "Verifique que el ID de persona sea correcto")
	}

	return SendSuccess(c, 200, autoridad)
}

// validateAutoridadUTEQ valida los datos de una autoridad UTEQ
func (h *AutoridadUTEQHandler) validateAutoridadUTEQ(autoridad *models.AutoridadUTEQ, isUpdate bool) []ValidationError {
	var errors []ValidationError

	// Validar PersonaID
	if autoridad.PersonaID == 0 {
		errors = append(errors, ValidationError{
			Field:   "persona_id",
			Message: "El ID de la persona es requerido",
		})
	}

	// Validar Cargo (opcional pero si se proporciona debe ser válido)
	if strings.TrimSpace(autoridad.Cargo) != "" {
		trimmedCargo := strings.TrimSpace(autoridad.Cargo)

		// Validar longitud mínima
		if len(trimmedCargo) < 2 {
			errors = append(errors, ValidationError{
				Field:   "cargo",
				Message: "El cargo debe tener al menos 2 caracteres",
				Value:   autoridad.Cargo,
			})
		}

		// Validar longitud máxima
		if len(trimmedCargo) > 100 {
			errors = append(errors, ValidationError{
				Field:   "cargo",
				Message: "El cargo no puede exceder 100 caracteres",
				Value:   autoridad.Cargo,
			})
		}

		// Validar que no contenga solo espacios o caracteres especiales
		if len(trimmedCargo) == 0 {
			errors = append(errors, ValidationError{
				Field:   "cargo",
				Message: "El cargo no puede contener solo espacios",
				Value:   autoridad.Cargo,
			})
		}

		// Validar formato del cargo (solo letras, espacios, guiones y puntos)
		cargoRegex := regexp.MustCompile(`^[a-zA-ZáéíóúÁÉÍÓÚñÑüÜ\s\-\.]+$`)
		if !cargoRegex.MatchString(trimmedCargo) {
			errors = append(errors, ValidationError{
				Field:   "cargo",
				Message: "El cargo solo puede contener letras, espacios, guiones y puntos",
				Value:   autoridad.Cargo,
			})
		}

		// Validar que no contenga números
		numberRegex := regexp.MustCompile(`\d`)
		if numberRegex.MatchString(trimmedCargo) {
			errors = append(errors, ValidationError{
				Field:   "cargo",
				Message: "El cargo no puede contener números",
				Value:   autoridad.Cargo,
			})
		}

		// Validar que no contenga caracteres especiales problemáticos
		specialCharsRegex := regexp.MustCompile(`[<>{}[\]\\|` + "`" + `~!@#$%^&*()+=;:'"<>?/]`)
		if specialCharsRegex.MatchString(trimmedCargo) {
			errors = append(errors, ValidationError{
				Field:   "cargo",
				Message: "El cargo no puede contener caracteres especiales",
				Value:   autoridad.Cargo,
			})
		}
	}

	return errors
}

// validateAutoridadUTEQRequiredFields valida que los campos requeridos estén presentes
func (h *AutoridadUTEQHandler) validateAutoridadUTEQRequiredFields(autoridad *models.AutoridadUTEQ) []ValidationError {
	var errors []ValidationError

	if autoridad.PersonaID == 0 {
		errors = append(errors, ValidationError{
			Field:   "persona_id",
			Message: "El campo persona_id es requerido",
		})
	}

	return errors
}

// validateAutoridadUTEQSearchParams valida los parámetros de búsqueda
func (h *AutoridadUTEQHandler) validateAutoridadUTEQSearchParams(cargo string) []ValidationError {
	var errors []ValidationError

	// Validar cargo de búsqueda
	if strings.TrimSpace(cargo) != "" {
		if len(strings.TrimSpace(cargo)) < 2 {
			errors = append(errors, ValidationError{
				Field:   "cargo",
				Message: "El término de búsqueda de cargo debe tener al menos 2 caracteres",
				Value:   cargo,
			})
		}

		// Validar que no contenga caracteres especiales problemáticos
		specialCharsRegex := regexp.MustCompile(`[<>{}[\]\\|` + "`" + `~!@#$%^&*()+=;:'"<>?/]`)
		if specialCharsRegex.MatchString(cargo) {
			errors = append(errors, ValidationError{
				Field:   "cargo",
				Message: "El término de búsqueda no puede contener caracteres especiales",
				Value:   cargo,
			})
		}
	}

	return errors
}
