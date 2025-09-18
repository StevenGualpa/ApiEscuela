package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"regexp"
	"strconv"
	"strings"
	"time"

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

	// Parsear JSON
	if err := c.BodyParser(&persona); err != nil {
		return SendError(c, 400, "invalid_json", "No se puede procesar el JSON. Verifique el formato de los datos", err.Error())
	}

	// Validar campos requeridos
	if validationErrors := h.validatePersonaRequiredFields(&persona); len(validationErrors) > 0 {
		return SendValidationError(c, "Faltan campos requeridos", validationErrors)
	}

	// Validar datos completos
	if validationErrors := h.validatePersona(&persona, false); len(validationErrors) > 0 {
		return SendValidationError(c, "Los datos proporcionados no son válidos", validationErrors)
	}

	// Limpiar datos
	persona.Nombre = strings.TrimSpace(persona.Nombre)
	persona.Cedula = strings.TrimSpace(persona.Cedula)
	persona.Correo = strings.TrimSpace(persona.Correo)
	persona.Telefono = strings.TrimSpace(persona.Telefono)

	// Crear persona
	if err := h.personaRepo.CreatePersona(&persona); err != nil {
		switch err {
		case repositories.ErrCedulaDuplicada:
			return SendError(c, 409, "duplicate_cedula", "Ya existe una persona con esta cédula", "La cédula debe ser única")
		case repositories.ErrCorreoDuplicado:
			return SendError(c, 409, "duplicate_email", "Ya existe una persona con este correo electrónico", "El correo debe ser único")
		case repositories.ErrPersonaYaExiste:
			return SendError(c, 409, "person_exists", "Ya existe una persona con estos datos", "Verifique los datos proporcionados")
		default:
			return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo crear la persona")
		}
	}

	return SendSuccess(c, 201, persona)
}

// GetPersona obtiene una persona por ID
func (h *PersonaHandler) GetPersona(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID de la persona es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID de la persona no es válido", "El ID debe ser un número entero positivo")
	}

	persona, err := h.personaRepo.GetPersonaByID(uint(id))
	if err != nil {
		return SendError(c, 404, "person_not_found", "No se encontró la persona solicitada", "Verifique que el ID sea correcto")
	}

	return SendSuccess(c, 200, persona)
}

// GetAllPersonas obtiene todas las personas
func (h *PersonaHandler) GetAllPersonas(c *fiber.Ctx) error {
	personas, err := h.personaRepo.GetAllPersonas()
	if err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudieron obtener las personas")
	}

	return SendSuccess(c, 200, personas)
}

// UpdatePersona actualiza una persona
func (h *PersonaHandler) UpdatePersona(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID de la persona es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID de la persona no es válido", "El ID debe ser un número entero positivo")
	}

	// Verificar que la persona existe
	persona, err := h.personaRepo.GetPersonaByID(uint(id))
	if err != nil {
		return SendError(c, 404, "person_not_found", "No se encontró la persona solicitada", "Verifique que el ID sea correcto")
	}

	// Parsear datos de actualización
	var updateData models.Persona
	if err := c.BodyParser(&updateData); err != nil {
		return SendError(c, 400, "invalid_json", "No se puede procesar el JSON. Verifique el formato de los datos", err.Error())
	}

	// Validar datos de actualización
	if validationErrors := h.validatePersona(&updateData, true); len(validationErrors) > 0 {
		return SendValidationError(c, "Los datos proporcionados no son válidos", validationErrors)
	}

	// Actualizar campos (mantener ID original)
	persona.Nombre = strings.TrimSpace(updateData.Nombre)
	persona.Cedula = strings.TrimSpace(updateData.Cedula)
	persona.Correo = strings.TrimSpace(updateData.Correo)
	persona.Telefono = strings.TrimSpace(updateData.Telefono)
	persona.FechaNacimiento = updateData.FechaNacimiento

	// Actualizar en base de datos
	if err := h.personaRepo.UpdatePersona(persona); err != nil {
		switch err {
		case repositories.ErrCedulaDuplicada:
			return SendError(c, 409, "duplicate_cedula", "Ya existe otra persona con esta cédula", "La cédula debe ser única")
		case repositories.ErrCorreoDuplicado:
			return SendError(c, 409, "duplicate_email", "Ya existe otra persona con este correo electrónico", "El correo debe ser único")
		case repositories.ErrPersonaYaExiste:
			return SendError(c, 409, "person_exists", "Ya existe una persona con estos datos", "Verifique los datos proporcionados")
		default:
			return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo actualizar la persona")
		}
	}

	return SendSuccess(c, 200, persona)
}

// DeletePersona elimina una persona
func (h *PersonaHandler) DeletePersona(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID de la persona es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID de la persona no es válido", "El ID debe ser un número entero positivo")
	}

	// Verificar que la persona existe antes de eliminar
	_, err = h.personaRepo.GetPersonaByID(uint(id))
	if err != nil {
		return SendError(c, 404, "person_not_found", "No se encontró la persona solicitada", "Verifique que el ID sea correcto")
	}

	if err := h.personaRepo.DeletePersona(uint(id)); err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo eliminar la persona")
	}

	return SendSuccess(c, 200, fiber.Map{
		"message": "Persona eliminada exitosamente",
		"id":      id,
	})
}

// GetPersonaByCedula obtiene una persona por cédula
func (h *PersonaHandler) GetPersonaByCedula(c *fiber.Ctx) error {
	cedula := c.Params("cedula")
	if cedula == "" {
		return SendError(c, 400, "missing_cedula", "La cédula es requerida", "Proporcione una cédula válida")
	}

	// Validar formato de cédula
	cedulaRegex := regexp.MustCompile(`^\d{10}$`)
	if !cedulaRegex.MatchString(strings.TrimSpace(cedula)) {
		return SendError(c, 400, "invalid_cedula", "El formato de la cédula no es válido", "La cédula debe tener exactamente 10 dígitos numéricos")
	}

	persona, err := h.personaRepo.GetPersonaByCedula(strings.TrimSpace(cedula))
	if err != nil {
		return SendError(c, 404, "person_not_found", "No se encontró la persona solicitada", "Verifique que la cédula sea correcta")
	}

	return SendSuccess(c, 200, persona)
}

// GetPersonasByCorreo obtiene personas por correo
func (h *PersonaHandler) GetPersonasByCorreo(c *fiber.Ctx) error {
	correo := c.Params("correo")
	if correo == "" {
		return SendError(c, 400, "missing_email", "El correo electrónico es requerido", "Proporcione un correo válido")
	}

	// Validar formato de correo
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(strings.TrimSpace(correo)) {
		return SendError(c, 400, "invalid_email", "El formato del correo electrónico no es válido", "Proporcione un correo con formato válido")
	}

	personas, err := h.personaRepo.GetPersonasByCorreo(strings.TrimSpace(correo))
	if err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudieron obtener las personas")
	}

	return SendSuccess(c, 200, personas)
}

// validatePersona valida los datos de una persona
func (h *PersonaHandler) validatePersona(persona *models.Persona, isUpdate bool) []ValidationError {
	var errors []ValidationError

	// Validar nombre
	if strings.TrimSpace(persona.Nombre) == "" {
		errors = append(errors, ValidationError{
			Field:   "nombre",
			Message: "El nombre es requerido",
			Value:   persona.Nombre,
		})
	} else if len(strings.TrimSpace(persona.Nombre)) < 2 {
		errors = append(errors, ValidationError{
			Field:   "nombre",
			Message: "El nombre debe tener al menos 2 caracteres",
			Value:   persona.Nombre,
		})
	} else if len(strings.TrimSpace(persona.Nombre)) > 100 {
		errors = append(errors, ValidationError{
			Field:   "nombre",
			Message: "El nombre no puede exceder 100 caracteres",
			Value:   persona.Nombre,
		})
	}

	// Validar cédula
	if strings.TrimSpace(persona.Cedula) == "" {
		errors = append(errors, ValidationError{
			Field:   "cedula",
			Message: "La cédula es requerida",
			Value:   persona.Cedula,
		})
	} else {
		// Validar formato de cédula ecuatoriana (10 dígitos)
		cedulaRegex := regexp.MustCompile(`^\d{10}$`)
		if !cedulaRegex.MatchString(strings.TrimSpace(persona.Cedula)) {
			errors = append(errors, ValidationError{
				Field:   "cedula",
				Message: "La cédula debe tener exactamente 10 dígitos numéricos",
				Value:   persona.Cedula,
			})
		}
	}

	// Validar correo (opcional pero si se proporciona debe ser válido)
	if strings.TrimSpace(persona.Correo) != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(strings.TrimSpace(persona.Correo)) {
			errors = append(errors, ValidationError{
				Field:   "correo",
				Message: "El formato del correo electrónico no es válido",
				Value:   persona.Correo,
			})
		} else if len(strings.TrimSpace(persona.Correo)) > 255 {
			errors = append(errors, ValidationError{
				Field:   "correo",
				Message: "El correo no puede exceder 255 caracteres",
				Value:   persona.Correo,
			})
		}
	}

	// Validar teléfono (opcional pero si se proporciona debe ser válido)
	if strings.TrimSpace(persona.Telefono) != "" {
		phoneRegex := regexp.MustCompile(`^[\d\s\-\+\(\)]{7,15}$`)
		if !phoneRegex.MatchString(strings.TrimSpace(persona.Telefono)) {
			errors = append(errors, ValidationError{
				Field:   "telefono",
				Message: "El formato del teléfono no es válido",
				Value:   persona.Telefono,
			})
		}
	}

	// Validar fecha de nacimiento (opcional pero si se proporciona debe ser válida)
	if !persona.FechaNacimiento.IsZero() {
		// Verificar que la fecha no sea futura
		if persona.FechaNacimiento.After(time.Now()) {
			errors = append(errors, ValidationError{
				Field:   "fecha_nacimiento",
				Message: "La fecha de nacimiento no puede ser futura",
				Value:   persona.FechaNacimiento.Format("2006-01-02"),
			})
		}

		// Verificar que la persona no sea muy joven (menos de 1 año)
		oneYearAgo := time.Now().AddDate(-1, 0, 0)
		if persona.FechaNacimiento.After(oneYearAgo) {
			errors = append(errors, ValidationError{
				Field:   "fecha_nacimiento",
				Message: "La fecha de nacimiento debe ser de al menos 1 año atrás",
				Value:   persona.FechaNacimiento.Format("2006-01-02"),
			})
		}

		// Verificar que la persona no sea muy vieja (más de 150 años)
		oneHundredFiftyYearsAgo := time.Now().AddDate(-150, 0, 0)
		if persona.FechaNacimiento.Before(oneHundredFiftyYearsAgo) {
			errors = append(errors, ValidationError{
				Field:   "fecha_nacimiento",
				Message: "La fecha de nacimiento no puede ser de hace más de 150 años",
				Value:   persona.FechaNacimiento.Format("2006-01-02"),
			})
		}
	}

	return errors
}

// validatePersonaRequiredFields valida que los campos requeridos estén presentes
func (h *PersonaHandler) validatePersonaRequiredFields(persona *models.Persona) []ValidationError {
	var errors []ValidationError

	if strings.TrimSpace(persona.Nombre) == "" {
		errors = append(errors, ValidationError{
			Field:   "nombre",
			Message: "El campo nombre es requerido",
		})
	}

	if strings.TrimSpace(persona.Cedula) == "" {
		errors = append(errors, ValidationError{
			Field:   "cedula",
			Message: "El campo cédula es requerido",
		})
	}

	return errors
}
