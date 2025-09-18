package handlers

import (
	"ApiEscuela/middleware"
	"ApiEscuela/services"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login maneja el inicio de sesión
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var loginReq services.LoginRequest

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(middleware.ErrorResponse{
			Error:      "Datos inválidos",
			ErrorCode:  "LOGIN_INVALID_JSON",
			Message:    "No se puede procesar el JSON. Verifique que el formato sea correcto",
			StatusCode: 400,
			Timestamp:  time.Now().Format(time.RFC3339),
			Path:       c.Path(),
			Method:     c.Method(),
		})
	}

	// Validar campos requeridos
	if loginReq.Usuario == "" || loginReq.Contraseña == "" {
		return c.Status(fiber.StatusBadRequest).JSON(middleware.ErrorResponse{
			Error:      "Campos requeridos faltantes",
			ErrorCode:  "LOGIN_MISSING_FIELDS",
			Message:    "Usuario y contraseña son requeridos",
			StatusCode: 400,
			Timestamp:  time.Now().Format(time.RFC3339),
			Path:       c.Path(),
			Method:     c.Method(),
		})
	}

	// Intentar login
	response, err := h.authService.Login(loginReq)
	if err != nil {
		// Determinar el código de error específico basado en el mensaje
		var errorCode string
		switch err.Error() {
		case "usuario no encontrado":
			errorCode = "LOGIN_USER_NOT_FOUND"
		case "usuario eliminado - contacte al administrador":
			errorCode = "LOGIN_USER_DELETED"
		case "contraseña incorrecta (texto plano)":
			errorCode = "LOGIN_PASSWORD_INCORRECT_PLAIN"
		case "contraseña incorrecta (hash bcrypt)":
			errorCode = "LOGIN_PASSWORD_INCORRECT_HASH"
		default:
			errorCode = "LOGIN_FAILED"
		}

		return c.Status(fiber.StatusUnauthorized).JSON(middleware.ErrorResponse{
			Error:      "Credenciales inválidas",
			ErrorCode:  errorCode,
			Message:    err.Error(),
			StatusCode: 401,
			Timestamp:  time.Now().Format(time.RFC3339),
			Path:       c.Path(),
			Method:     c.Method(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Register maneja el registro de nuevos usuarios
// RecoverPassword maneja la recuperación de contraseña por cédula (público)
func (h *AuthHandler) RecoverPassword(c *fiber.Ctx) error {
	var req struct {
		Cedula string `json:"cedula"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se puede procesar el JSON"})
	}
	if strings.TrimSpace(req.Cedula) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "La cédula es requerida"})
	}
	if err := h.authService.RecoverPassword(req.Cedula); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Si la cédula existe, se envió un correo con la contraseña temporal"})
}

// VerifyCode maneja la verificación del código temporal (público)
func (h *AuthHandler) VerifyCode(c *fiber.Ctx) error {
	var req struct {
		Codigo string `json:"codigo"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se puede procesar el JSON"})
	}

	result, err := h.authService.VerifyCodigoWithDetails(req.Codigo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	switch result.Estado {
	case "no existe":
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"estado": "no existe"})
	case "caducado":
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"estado": "caducado"})
	case "verificado":
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"estado":     "verificado",
			"usuario_id": result.UsuarioID,
			"cedula":     result.Cedula,
			"codigo_id":  result.CodigoID,
		})
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"estado": result.Estado})
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var registerReq services.RegisterRequest

	if err := c.BodyParser(&registerReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	// Validar campos requeridos
	if registerReq.Usuario == "" || registerReq.Contraseña == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Usuario y contraseña son requeridos",
		})
	}

	if len(registerReq.Contraseña) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "La contraseña debe tener al menos 6 caracteres",
		})
	}

	// Intentar registro
	usuario, err := h.authService.Register(registerReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuario registrado exitosamente",
		"usuario": usuario,
	})
}

// ChangePassword maneja el cambio de contraseña
// ResetPassword maneja el reseteo de contraseña por ID de código (público)
func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	var req struct {
		CodigoID  uint   `json:"codigo_id"`
		UsuarioID uint   `json:"usuario_id"`
		Clave     string `json:"clave"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se puede procesar el JSON"})
	}
	if req.CodigoID == 0 || req.UsuarioID == 0 || strings.TrimSpace(req.Clave) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "codigo_id, usuario_id y clave son requeridos"})
	}
	if len(req.Clave) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "La clave debe tener al menos 6 caracteres"})
	}
	if err := h.authService.ResetPasswordByCodigoID(req.CodigoID, req.UsuarioID, req.Clave); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "clave actualizada"})
}

func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	// Obtener ID del usuario del contexto (del JWT)
	userID := c.Locals("user_id").(uint)

	var changePasswordReq struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.BodyParser(&changePasswordReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	// Validar campos requeridos
	if changePasswordReq.OldPassword == "" || changePasswordReq.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Contraseña actual y nueva contraseña son requeridas",
		})
	}

	if len(changePasswordReq.NewPassword) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "La nueva contraseña debe tener al menos 6 caracteres",
		})
	}

	// Cambiar contraseña
	err := h.authService.ChangePassword(userID, changePasswordReq.OldPassword, changePasswordReq.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Contraseña cambiada exitosamente",
	})
}

// GetProfile obtiene el perfil del usuario autenticado
func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	username := c.Locals("username").(string)
	tipoUsuarioID := c.Locals("tipo_usuario_id").(uint)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id":         userID,
		"username":        username,
		"tipo_usuario_id": tipoUsuarioID,
	})
}

// RefreshToken genera un nuevo token para el usuario autenticado
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	username := c.Locals("username").(string)
	tipoUsuarioID := c.Locals("tipo_usuario_id").(uint)

	// Generar nuevo token
	token, err := h.authService.GenerateNewToken(userID, username, tipoUsuarioID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al generar nuevo token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":   token,
		"message": "Token renovado exitosamente",
	})
}

// ValidateToken valida si un token es válido (endpoint público para verificación)
func (h *AuthHandler) ValidateToken(c *fiber.Ctx) error {
	var tokenReq struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&tokenReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if tokenReq.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token es requerido",
		})
	}

	valid, claims := h.authService.ValidateToken(tokenReq.Token)
	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"valid": false,
			"error": "Token inválido o expirado",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"valid":           true,
		"user_id":         claims.UserID,
		"username":        claims.Username,
		"tipo_usuario_id": claims.TipoUsuarioID,
	})
}
