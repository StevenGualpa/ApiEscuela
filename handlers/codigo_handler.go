package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CodigoHandler struct {
	codigoRepo *repositories.CodigoUsuarioRepository
	validator  *CodigoValidator
}

func NewCodigoHandler(codigoRepo *repositories.CodigoUsuarioRepository) *CodigoHandler {
	return &CodigoHandler{
		codigoRepo: codigoRepo,
		validator:  NewCodigoValidator(),
	}
}

// CreateCodigo crea un nuevo código
func (h *CodigoHandler) CreateCodigo(c *fiber.Ctx) error {
	var codigo models.CodigoUsuario

	// Parsear JSON
	if err := c.BodyParser(&codigo); err != nil {
		return SendError(c, 400, "invalid_json", "No se puede procesar el JSON. Verifique el formato de los datos", err.Error())
	}

	// Validar campos requeridos
	if validationErrors := h.validator.ValidateRequiredFields(&codigo); len(validationErrors) > 0 {
		return SendValidationError(c, "Faltan campos requeridos", validationErrors)
	}

	// Validar datos completos
	if validationErrors := h.validator.ValidateCodigo(&codigo, false); len(validationErrors) > 0 {
		return SendValidationError(c, "Los datos proporcionados no son válidos", validationErrors)
	}

	// Limpiar datos
	codigo.Codigo = strings.TrimSpace(codigo.Codigo)
	codigo.Estado = strings.TrimSpace(codigo.Estado)

	// Crear código
	if err := h.codigoRepo.Crear(codigo.UsuarioID, codigo.Codigo); err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo crear el código")
	}

	return SendSuccess(c, 201, fiber.Map{
		"message":    "Código creado exitosamente",
		"usuario_id": codigo.UsuarioID,
		"codigo":     codigo.Codigo,
	})
}

// GetCodigo obtiene un código por ID
func (h *CodigoHandler) GetCodigo(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID del código es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID del código no es válido", "El ID debe ser un número entero positivo")
	}

	codigo, err := h.codigoRepo.GetByID(uint(id))
	if err != nil {
		return SendError(c, 404, "codigo_not_found", "No se encontró el código solicitado", "Verifique que el ID sea correcto")
	}

	return SendSuccess(c, 200, codigo)
}

// GetAllCodigos obtiene todos los códigos
func (h *CodigoHandler) GetAllCodigos(c *fiber.Ctx) error {
	// Esta función no debería existir en producción por seguridad
	// pero la incluimos para completitud
	return SendError(c, 403, "forbidden", "Acceso denegado", "No se permite listar todos los códigos por seguridad")
}

// GetCodigosByUsuario obtiene códigos por usuario
func (h *CodigoHandler) GetCodigosByUsuario(c *fiber.Ctx) error {
	usuarioIDStr := c.Params("usuario_id")
	if usuarioIDStr == "" {
		return SendError(c, 400, "missing_usuario_id", "El ID del usuario es requerido", "Proporcione un ID de usuario válido")
	}

	usuarioID, err := strconv.Atoi(usuarioIDStr)
	if err != nil || usuarioID <= 0 {
		return SendError(c, 400, "invalid_usuario_id", "El ID del usuario no es válido", "El ID debe ser un número entero positivo")
	}

	// Esta función requeriría implementación en el repositorio
	return SendError(c, 501, "not_implemented", "Función no implementada", "Esta funcionalidad no está disponible")
}

// UpdateCodigo actualiza un código
func (h *CodigoHandler) UpdateCodigo(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID del código es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID del código no es válido", "El ID debe ser un número entero positivo")
	}

	// Verificar que el código existe
	codigo, err := h.codigoRepo.GetByID(uint(id))
	if err != nil {
		return SendError(c, 404, "codigo_not_found", "No se encontró el código solicitado", "Verifique que el ID sea correcto")
	}

	// Parsear datos de actualización
	var updateData models.CodigoUsuario
	if err := c.BodyParser(&updateData); err != nil {
		return SendError(c, 400, "invalid_json", "No se puede procesar el JSON. Verifique el formato de los datos", err.Error())
	}

	// Validar datos de actualización
	if validationErrors := h.validator.ValidateCodigo(&updateData, true); len(validationErrors) > 0 {
		return SendValidationError(c, "Los datos proporcionados no son válidos", validationErrors)
	}

	// Actualizar campos (mantener ID original)
	codigo.UsuarioID = updateData.UsuarioID
	codigo.Codigo = strings.TrimSpace(updateData.Codigo)
	codigo.Estado = strings.TrimSpace(updateData.Estado)
	codigo.ExpiraEn = updateData.ExpiraEn

	// Actualizar en base de datos
	if err := h.codigoRepo.Update(codigo); err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo actualizar el código")
	}

	return SendSuccess(c, 200, codigo)
}

// DeleteCodigo elimina un código
func (h *CodigoHandler) DeleteCodigo(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID del código es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID del código no es válido", "El ID debe ser un número entero positivo")
	}

	// Verificar que el código existe antes de eliminar
	_, err = h.codigoRepo.GetByID(uint(id))
	if err != nil {
		return SendError(c, 404, "codigo_not_found", "No se encontró el código solicitado", "Verifique que el ID sea correcto")
	}

	// Los códigos no se eliminan físicamente, solo se marcan como expirados
	if err := h.codigoRepo.MarcarComoExpirado(uint(id)); err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo marcar el código como expirado")
	}

	return SendSuccess(c, 200, fiber.Map{
		"message": "Código marcado como expirado exitosamente",
		"id":      id,
	})
}

// VerifyCodigo verifica un código
func (h *CodigoHandler) VerifyCodigo(c *fiber.Ctx) error {
	var req struct {
		Codigo string `json:"codigo"`
	}

	// Parsear JSON
	if err := c.BodyParser(&req); err != nil {
		return SendError(c, 400, "invalid_json", "No se puede procesar el JSON. Verifique el formato de los datos", err.Error())
	}

	// Validar código
	if validationErrors := h.validator.ValidateCodigoString(req.Codigo); len(validationErrors) > 0 {
		return SendValidationError(c, "El código proporcionado no es válido", validationErrors)
	}

	// Buscar código
	codigo, err := h.codigoRepo.FindLatestByCodigo(strings.TrimSpace(req.Codigo))
	if err != nil {
		return SendError(c, 404, "codigo_not_found", "No se encontró el código", "Verifique que el código sea correcto")
	}

	// Verificar estado
	if codigo.Estado != "valido" {
		return SendError(c, 400, "codigo_invalid_state", "El código no está en estado válido", "El código debe estar en estado 'valido' para ser verificado")
	}

	// Verificar expiración
	if codigo.ExpiraEn != nil && codigo.ExpiraEn.Before(time.Now()) {
		// Marcar como expirado
		h.codigoRepo.MarcarComoExpirado(codigo.ID)
		return SendError(c, 400, "codigo_expired", "El código ha expirado", "Solicite un nuevo código")
	}

	return SendSuccess(c, 200, fiber.Map{
		"message":    "Código verificado exitosamente",
		"codigo_id":  codigo.ID,
		"usuario_id": codigo.UsuarioID,
		"estado":     codigo.Estado,
		"expira_en":  codigo.ExpiraEn,
	})
}

// MarcarComoVerificado marca un código como verificado
func (h *CodigoHandler) MarcarComoVerificado(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID del código es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID del código no es válido", "El ID debe ser un número entero positivo")
	}

	// Verificar que el código existe
	codigo, err := h.codigoRepo.GetByID(uint(id))
	if err != nil {
		return SendError(c, 404, "codigo_not_found", "No se encontró el código solicitado", "Verifique que el ID sea correcto")
	}

	// Verificar que el código esté en estado válido
	if codigo.Estado != "valido" {
		return SendError(c, 400, "codigo_invalid_state", "El código no está en estado válido", "Solo se pueden verificar códigos en estado 'valido'")
	}

	// Marcar como verificado
	if err := h.codigoRepo.MarcarComoVerificado(uint(id)); err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo marcar el código como verificado")
	}

	return SendSuccess(c, 200, fiber.Map{
		"message": "Código marcado como verificado exitosamente",
		"id":      id,
	})
}

// MarcarComoExpirado marca un código como expirado
func (h *CodigoHandler) MarcarComoExpirado(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return SendError(c, 400, "missing_id", "El ID del código es requerido", "Proporcione un ID válido")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return SendError(c, 400, "invalid_id", "El ID del código no es válido", "El ID debe ser un número entero positivo")
	}

	// Verificar que el código existe
	_, err = h.codigoRepo.GetByID(uint(id))
	if err != nil {
		return SendError(c, 404, "codigo_not_found", "No se encontró el código solicitado", "Verifique que el ID sea correcto")
	}

	// Marcar como expirado
	if err := h.codigoRepo.MarcarComoExpirado(uint(id)); err != nil {
		return SendError(c, 500, "database_error", "Error interno del servidor", "No se pudo marcar el código como expirado")
	}

	return SendSuccess(c, 200, fiber.Map{
		"message": "Código marcado como expirado exitosamente",
		"id":      id,
	})
}
