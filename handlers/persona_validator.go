package handlers

import (
	"ApiEscuela/models"
	"regexp"
	"strings"
	"time"
)

// PersonaValidator maneja la validación de datos de persona
type PersonaValidator struct{}

// NewPersonaValidator crea una nueva instancia del validador
func NewPersonaValidator() *PersonaValidator {
	return &PersonaValidator{}
}

// ValidatePersona valida los datos de una persona
func (v *PersonaValidator) ValidatePersona(persona *models.Persona, isUpdate bool) []ValidationError {
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

// ValidateRequiredFields valida que los campos requeridos estén presentes
func (v *PersonaValidator) ValidateRequiredFields(persona *models.Persona) []ValidationError {
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
