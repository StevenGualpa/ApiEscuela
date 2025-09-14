package services

import (
	"ApiEscuela/middleware"
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	usuarioRepo *repositories.UsuarioRepository
}

func NewAuthService(usuarioRepo *repositories.UsuarioRepository) *AuthService {
	return &AuthService{
		usuarioRepo: usuarioRepo,
	}
}

// LoginRequest representa la estructura de datos para el login
type LoginRequest struct {
	Usuario    string `json:"usuario" validate:"required"`
	Contraseña string `json:"contraseña" validate:"required"`
}

// LoginResponse representa la respuesta del login
type LoginResponse struct {
	Token   string        `json:"token"`
	Usuario *models.Usuario `json:"usuario"`
	Message string        `json:"message"`
	RequiereCambioPassword bool `json:"requiere_cambio_password"`
}

// RegisterRequest representa la estructura de datos para el registro
type RegisterRequest struct {
	Usuario       string `json:"usuario" validate:"required"`
	Contraseña    string `json:"contraseña" validate:"required,min=6"`
	PersonaID     uint   `json:"persona_id" validate:"required"`
	TipoUsuarioID uint   `json:"tipo_usuario_id" validate:"required"`
}

// HashPassword encripta una contraseña
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword verifica si una contraseña coincide con su hash
func (s *AuthService) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Login autentica un usuario y devuelve un token JWT
func (s *AuthService) Login(loginReq LoginRequest) (*LoginResponse, error) {
	// Buscar usuario por nombre de usuario
	usuario, err := s.usuarioRepo.GetUsuarioByUsername(loginReq.Usuario)
	if err != nil {
		// Intentar buscar incluyendo eliminados para debugging
		usuarioDeleted, errDeleted := s.usuarioRepo.GetUsuarioByUsernameIncludingDeleted(loginReq.Usuario)
		if errDeleted == nil && usuarioDeleted != nil {
			// El usuario existe pero está eliminado
			return nil, errors.New("usuario eliminado - contacte al administrador")
		}
		// Usuario no existe
		return nil, errors.New("usuario no encontrado")
	}

	// Verificar si la contraseña está encriptada (hash bcrypt tiene al menos 60 caracteres)
	if len(usuario.Contraseña) < 60 {
		// Contraseña no está encriptada, comparar directamente
		if loginReq.Contraseña != usuario.Contraseña {
			return nil, errors.New("contraseña incorrecta (texto plano)")
		}
	} else {
		// Verificar contraseña encriptada
		if !s.CheckPassword(loginReq.Contraseña, usuario.Contraseña) {
			return nil, errors.New("contraseña incorrecta (hash bcrypt)")
		}
	}

	// Generar token JWT
	token, err := middleware.GenerateJWT(usuario.ID, usuario.Usuario, usuario.TipoUsuarioID)
	if err != nil {
		return nil, errors.New("error al generar token")
	}

	// Limpiar la contraseña antes de devolver el usuario
	usuario.Contraseña = ""

	// Verificar si el usuario necesita cambiar contraseña
	requiereCambioPassword := !usuario.Verificado
	message := "Login exitoso"
	if requiereCambioPassword {
		message = "Login exitoso - Debe cambiar su contraseña para continuar"
	}

	return &LoginResponse{
		Token:                  token,
		Usuario:                usuario,
		Message:                message,
		RequiereCambioPassword: requiereCambioPassword,
	}, nil
}

// Register registra un nuevo usuario
func (s *AuthService) Register(registerReq RegisterRequest) (*models.Usuario, error) {
	// Verificar si el usuario ya existe
	existingUser, _ := s.usuarioRepo.GetUsuarioByUsername(registerReq.Usuario)
	if existingUser != nil {
		return nil, errors.New("el usuario ya existe")
	}

	// Encriptar contraseña
	hashedPassword, err := s.HashPassword(registerReq.Contraseña)
	if err != nil {
		return nil, errors.New("error al encriptar contraseña")
	}

	// Crear nuevo usuario
	usuario := &models.Usuario{
		Usuario:       registerReq.Usuario,
		Contraseña:    hashedPassword,
		PersonaID:     registerReq.PersonaID,
		TipoUsuarioID: registerReq.TipoUsuarioID,
	}

	// Guardar usuario en la base de datos
	if err := s.usuarioRepo.CreateUsuario(usuario); err != nil {
		return nil, errors.New("error al crear usuario")
	}

	// Limpiar la contraseña antes de devolver
	usuario.Contraseña = ""

	return usuario, nil
}

// ChangePassword cambia la contraseña de un usuario
func (s *AuthService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	// Obtener usuario
	usuario, err := s.usuarioRepo.GetUsuarioByID(userID)
	if err != nil {
		return errors.New("usuario no encontrado")
	}

	// Verificar contraseña actual
	if !s.CheckPassword(oldPassword, usuario.Contraseña) {
		return errors.New("contraseña actual incorrecta")
	}

	// Encriptar nueva contraseña
	hashedPassword, err := s.HashPassword(newPassword)
	if err != nil {
		return errors.New("error al encriptar nueva contraseña")
	}

	// Actualizar contraseña y marcar como verificado
	usuario.Contraseña = hashedPassword
	usuario.Verificado = true
	if err := s.usuarioRepo.UpdateUsuario(usuario); err != nil {
		return errors.New("error al actualizar contraseña")
	}

	return nil
}

// GenerateNewToken genera un nuevo token JWT para un usuario
func (s *AuthService) GenerateNewToken(userID uint, username string, tipoUsuarioID uint) (string, error) {
	return middleware.GenerateJWT(userID, username, tipoUsuarioID)
}


// ValidateToken valida un token JWT y devuelve las claims
func (s *AuthService) ValidateToken(tokenString string) (bool, *middleware.JWTClaims) {
	claims, err := middleware.ValidateJWT(tokenString)
	if err != nil {
		return false, nil
	}
	return true, claims
}