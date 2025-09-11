package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims representa las claims del JWT
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	TipoUsuarioID uint `json:"tipo_usuario_id"`
	jwt.RegisteredClaims
}

// JWTSecret es la clave secreta para firmar los tokens (en producción debe estar en variables de entorno)
var JWTSecret = []byte("tu_clave_secreta_super_segura_aqui_cambiar_en_produccion")

// GenerateJWT genera un nuevo token JWT
func GenerateJWT(userID uint, username string, tipoUsuarioID uint) (string, error) {
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		TipoUsuarioID: tipoUsuarioID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token válido por 24 horas
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ApiEscuela",
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// ValidateJWT valida un token JWT
func ValidateJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// JWTMiddleware es el middleware que protege las rutas
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obtener el token del header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token de autorización requerido",
			})
		}

		// Verificar que el header tenga el formato "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Formato de token inválido. Use: Bearer <token>",
			})
		}

		tokenString := tokenParts[1]

		// Validar el token
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido o expirado",
			})
		}

		// Guardar la información del usuario en el contexto
		c.Locals("user_id", claims.UserID)
		c.Locals("username", claims.Username)
		c.Locals("tipo_usuario_id", claims.TipoUsuarioID)

		return c.Next()
	}
}

// OptionalJWTMiddleware es un middleware opcional que no bloquea si no hay token
func OptionalJWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader != "" {
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
				tokenString := tokenParts[1]
				claims, err := ValidateJWT(tokenString)
				if err == nil {
					c.Locals("user_id", claims.UserID)
					c.Locals("username", claims.Username)
					c.Locals("tipo_usuario_id", claims.TipoUsuarioID)
				}
			}
		}
		return c.Next()
	}
}