package handlers

import (
	"ApiEscuela/models"
	"regexp"
	"strings"
	"time"
)

// CodigoValidator maneja la validación de datos de códigos
type CodigoValidator struct{}

// NewCodigoValidator crea una nueva instancia del validador
func NewCodigoValidator() *CodigoValidator {
	return &CodigoValidator{}
}

// ValidateCodigo valida los datos de un código
func (v *CodigoValidator) ValidateCodigo(codigo *models.CodigoUsuario, isUpdate bool) []ValidationError {
	var errors []ValidationError

	// Validar UsuarioID
	if codigo.UsuarioID == 0 {
		errors = append(errors, ValidationError{
			Field:   "usuario_id",
			Message: "El ID del usuario es requerido",
			Value:   "0",
		})
	}

	// Validar código
	if strings.TrimSpace(codigo.Codigo) == "" {
		errors = append(errors, ValidationError{
			Field:   "codigo",
			Message: "El código es requerido",
			Value:   codigo.Codigo,
		})
	} else {
		// Validar formato del código (6 dígitos numéricos)
		codigoRegex := regexp.MustCompile(`^\d{6}$`)
		if !codigoRegex.MatchString(strings.TrimSpace(codigo.Codigo)) {
			errors = append(errors, ValidationError{
				Field:   "codigo",
				Message: "El código debe tener exactamente 6 dígitos numéricos",
				Value:   codigo.Codigo,
			})
		}
	}

	// Validar estado
	if strings.TrimSpace(codigo.Estado) == "" {
		errors = append(errors, ValidationError{
			Field:   "estado",
			Message: "El estado es requerido",
			Value:   codigo.Estado,
		})
	} else {
		// Validar que el estado sea válido
		validStates := []string{"valido", "verificado", "expirado"}
		isValidState := false
		for _, state := range validStates {
			if strings.TrimSpace(codigo.Estado) == state {
				isValidState = true
				break
			}
		}
		if !isValidState {
			errors = append(errors, ValidationError{
				Field:   "estado",
				Message: "El estado debe ser uno de: valido, verificado, expirado",
				Value:   codigo.Estado,
			})
		}
	}

	// Validar ExpiraEn (opcional pero si se proporciona debe ser válida)
	if codigo.ExpiraEn != nil {
		// Verificar que la fecha no sea muy antigua (más de 1 año)
		oneYearAgo := time.Now().AddDate(-1, 0, 0)
		if codigo.ExpiraEn.Before(oneYearAgo) {
			errors = append(errors, ValidationError{
				Field:   "expira_en",
				Message: "La fecha de expiración no puede ser de hace más de 1 año",
				Value:   codigo.ExpiraEn.Format("2006-01-02 15:04:05"),
			})
		}

		// Verificar que la fecha no sea muy futura (más de 1 día)
		oneDayFromNow := time.Now().AddDate(0, 0, 1)
		if codigo.ExpiraEn.After(oneDayFromNow) {
			errors = append(errors, ValidationError{
				Field:   "expira_en",
				Message: "La fecha de expiración no puede ser de más de 1 día en el futuro",
				Value:   codigo.ExpiraEn.Format("2006-01-02 15:04:05"),
			})
		}
	}

	return errors
}

// ValidateRequiredFields valida que los campos requeridos estén presentes
func (v *CodigoValidator) ValidateRequiredFields(codigo *models.CodigoUsuario) []ValidationError {
	var errors []ValidationError

	if codigo.UsuarioID == 0 {
		errors = append(errors, ValidationError{
			Field:   "usuario_id",
			Message: "El campo usuario_id es requerido",
		})
	}

	if strings.TrimSpace(codigo.Codigo) == "" {
		errors = append(errors, ValidationError{
			Field:   "codigo",
			Message: "El campo código es requerido",
		})
	}

	return errors
}

// ValidateCodigoString valida un código como string
func (v *CodigoValidator) ValidateCodigoString(codigo string) []ValidationError {
	var errors []ValidationError

	if strings.TrimSpace(codigo) == "" {
		errors = append(errors, ValidationError{
			Field:   "codigo",
			Message: "El código es requerido",
			Value:   codigo,
		})
	} else {
		// Validar formato del código (6 dígitos numéricos)
		codigoRegex := regexp.MustCompile(`^\d{6}$`)
		if !codigoRegex.MatchString(strings.TrimSpace(codigo)) {
			errors = append(errors, ValidationError{
				Field:   "codigo",
				Message: "El código debe tener exactamente 6 dígitos numéricos",
				Value:   codigo,
			})
		}
	}

	return errors
}
