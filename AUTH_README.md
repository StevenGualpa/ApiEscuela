# Sistema de Autenticación JWT - ApiEscuela

## Descripción

Se ha implementado un sistema de autenticación completo basado en JWT (JSON Web Tokens) para proteger todos los endpoints de la API. Ahora todos los endpoints requieren autenticación excepto los endpoints públicos de autenticación.

## Estructura de Rutas

### Rutas Públicas (Sin autenticación)
- `POST /auth/login` - Iniciar sesión
- `POST /auth/register` - Registrar nuevo usuario
- `POST /auth/validate-token` - Validar token
- `GET /` - Página de bienvenida
- `GET /health` - Estado de salud de la API

### Rutas Protegidas (Requieren JWT)
Todas las rutas de la API ahora están bajo el prefijo `/api` y requieren autenticación:
- `/api/auth/*` - Rutas de autenticación protegidas
- `/api/estudiantes/*` - Gestión de estudiantes
- `/api/personas/*` - Gestión de personas
- `/api/provincias/*` - Gestión de provincias
- `/api/ciudades/*` - Gestión de ciudades
- Y todas las demás rutas existentes...

## Endpoints de Autenticación

### 1. Login
```http
POST /auth/login
Content-Type: application/json

{
  "usuario": "nombre_usuario",
  "contraseña": "contraseña_usuario"
}
```

**Respuesta exitosa:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "usuario": {
    "ID": 1,
    "usuario": "nombre_usuario",
    "persona_id": 1,
    "tipo_usuario_id": 1,
    "persona": {...},
    "tipo_usuario": {...}
  },
  "message": "Login exitoso"
}
```

### 2. Registro
```http
POST /auth/register
Content-Type: application/json

{
  "usuario": "nuevo_usuario",
  "contraseña": "contraseña_segura",
  "persona_id": 1,
  "tipo_usuario_id": 1
}
```

**Respuesta exitosa:**
```json
{
  "message": "Usuario registrado exitosamente",
  "usuario": {
    "ID": 2,
    "usuario": "nuevo_usuario",
    "persona_id": 1,
    "tipo_usuario_id": 1
  }
}
```

### 3. Validar Token
```http
POST /auth/validate-token
Content-Type: application/json

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Respuesta exitosa:**
```json
{
  "valid": true,
  "user_id": 1,
  "username": "nombre_usuario",
  "tipo_usuario_id": 1
}
```

## Endpoints Protegidos

### 4. Obtener Perfil
```http
GET /api/auth/profile
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Respuesta:**
```json
{
  "user_id": 1,
  "username": "nombre_usuario",
  "tipo_usuario_id": 1
}
```

### 5. Cambiar Contraseña
```http
POST /api/auth/change-password
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Content-Type: application/json

{
  "old_password": "contraseña_actual",
  "new_password": "nueva_contraseña"
}
```

### 6. Renovar Token
```http
POST /api/auth/refresh-token
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Respuesta:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "message": "Token renovado exitosamente"
}
```

## Cómo usar la autenticación

### 1. Obtener un token
Primero, debes hacer login para obtener un token JWT:

```bash
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "tu_usuario",
    "contraseña": "tu_contraseña"
  }'
```

### 2. Usar el token en las peticiones
Para acceder a cualquier endpoint protegido, incluye el token en el header Authorization:

```bash
curl -X GET http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer tu_token_jwt_aqui"
```

### 3. Ejemplo con JavaScript/Fetch
```javascript
// Login
const loginResponse = await fetch('/auth/login', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    usuario: 'mi_usuario',
    contraseña: 'mi_contraseña'
  })
});

const { token } = await loginResponse.json();

// Usar el token para acceder a endpoints protegidos
const estudiantesResponse = await fetch('/api/estudiantes', {
  headers: {
    'Authorization': `Bearer ${token}`
  }
});

const estudiantes = await estudiantesResponse.json();
```

## Características de Seguridad

### JWT Token
- **Expiración**: Los tokens expiran en 24 horas
- **Algoritmo**: HS256 (HMAC SHA-256)
- **Claims incluidos**:
  - `user_id`: ID del usuario
  - `username`: Nombre de usuario
  - `tipo_usuario_id`: Tipo de usuario
  - `exp`: Tiempo de expiración
  - `iat`: Tiempo de emisión
  - `iss`: Emisor (ApiEscuela)

### Contraseñas
- **Encriptación**: bcrypt con salt automático
- **Longitud mínima**: 6 caracteres
- **Verificación**: Comparación segura con hash almacenado

### Middleware de Protección
- **Validación automática**: Todos los endpoints bajo `/api` requieren token válido
- **Información de usuario**: El middleware inyecta la información del usuario en el contexto
- **Manejo de errores**: Respuestas claras para tokens inválidos o expirados

## Códigos de Error

### 400 - Bad Request
- Datos de entrada inválidos
- JSON malformado
- Campos requeridos faltantes

### 401 - Unauthorized
- Token faltante o inválido
- Token expirado
- Credenciales incorrectas

### 404 - Not Found
- Usuario no encontrado

### 500 - Internal Server Error
- Error del servidor
- Error de base de datos

## Configuración de Seguridad

### Variables de Entorno Recomendadas
```env
JWT_SECRET=tu_clave_secreta_super_segura_aqui
JWT_EXPIRATION=24h
```

**Nota**: Actualmente la clave JWT está hardcodeada en el código. En producción, debes:
1. Mover la clave a una variable de entorno
2. Usar una clave más larga y compleja
3. Considerar rotación de claves

## Migración de Usuarios Existentes

Si ya tienes usuarios en la base de datos con contraseñas sin encriptar, necesitarás:

1. Crear un script de migración para encriptar las contraseñas existentes
2. O requerir que los usuarios cambien sus contraseñas

## Ejemplo de Uso Completo

```bash
# 1. Registrar un nuevo usuario
curl -X POST http://localhost:3000/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "test_user",
    "contraseña": "password123",
    "persona_id": 1,
    "tipo_usuario_id": 1
  }'

# 2. Hacer login
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "test_user",
    "contraseña": "password123"
  }'

# 3. Usar el token para acceder a datos (reemplaza TOKEN con el token recibido)
curl -X GET http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer TOKEN"

# 4. Obtener perfil del usuario
curl -X GET http://localhost:3000/api/auth/profile \
  -H "Authorization: Bearer TOKEN"
```

## Notas Importantes

1. **Todos los endpoints existentes ahora requieren autenticación** excepto los de autenticación pública
2. **Las URLs han cambiado**: Los endpoints protegidos ahora están bajo `/api/`
3. **El token debe incluirse en cada petición** a endpoints protegidos
4. **Los tokens expiran en 24 horas** - usa el endpoint de refresh para renovar
5. **Las contraseñas se encriptan automáticamente** al registrar o cambiar contraseña