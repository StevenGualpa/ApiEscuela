package handlers

import (
	"ApiEscuela/services"

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	// Validar campos requeridos
	if loginReq.Usuario == "" || loginReq.Contraseña == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Usuario y contraseña son requeridos",
		})
	}

	// Intentar login
	response, err := h.authService.Login(loginReq)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Register maneja el registro de nuevos usuarios
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