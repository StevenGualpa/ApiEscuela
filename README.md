# ApiEscuela - Sistema de Gestión de Visitas Educativas UTEQ

Un backend completo desarrollado en Go para la gestión integral de visitas educativas a la UTEQ, incluyendo estudiantes, instituciones, programas de visita, actividades, sistema de dudas y **autenticación JWT completa**.

## 🚀 Características Principales

- **🔐 Autenticación JWT**: Sistema completo de seguridad con tokens
- **📊 15 Entidades**: Gestión completa de todos los aspectos del sistema
- **🌐 80+ Endpoints**: API REST completa y documentada
- **🛡️ Middleware de Seguridad**: Protección automática de rutas
- **🗑️ Soft Delete**: Eliminación segura con restauración
- **📈 Estadísticas**: Reportes y análisis integrados
- **🚨 Sistema de Errores Mejorado**: Respuestas estructuradas y específicas

## 📋 Tabla de Contenidos

1. [🔐 Sistema de Autenticación](#-sistema-de-autenticación)
2. [🏗️ Arquitectura](#️-arquitectura)
3. [🚀 API Endpoints](#-api-endpoints)
4. [🚨 Sistema de Errores](#-sistema-de-errores)
5. [📝 Ejemplos de Uso](#-ejemplos-de-uso)
6. [⚙️ Instalación](#️-instalación)
7. [🛠️ Tecnologías](#️-tecnologías)

## 🔐 Sistema de Autenticación

### 🌐 Estructura de URLs

#### **Rutas Públicas (Sin autenticación)**
- `POST /auth/login` - Iniciar sesión
- `POST /auth/register` - Registrar usuario
- `POST /auth/validate-token` - Validar token
- `GET /` - Página de bienvenida
- `GET /health` - Estado de salud

#### **Rutas Protegidas (Requieren JWT)**
Todas las rutas bajo `/api/*` requieren el header:
```
Authorization: Bearer tu_token_jwt_aqui
```

### 🔑 Flujo de Autenticación

1. **Login** para obtener token:
```bash
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"usuario": "tu_usuario", "contraseña": "tu_contraseña"}'
```

2. **Usar token** en peticiones protegidas:
```bash
curl -X GET http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer tu_token_jwt_aqui"
```

### 🛡️ Características de Seguridad

- **JWT Tokens**: Expiración de 24 horas
- **Contraseñas Encriptadas**: bcrypt con salt automático
- **Middleware Automático**: Validación en todas las rutas `/api/*`
- **Renovación de Tokens**: Endpoint `/api/auth/refresh-token`

## 🏗️ Arquitectura

### 📊 Entidades del Sistema (15 modelos)

| Entidad | Descripción |
|---------|-------------|
| **Persona** | Información básica de personas |
| **Estudiante** | Estudiantes de instituciones educativas |
| **EstudianteUniversitario** | Estudiantes universitarios de la UTEQ |
| **AutoridadUTEQ** | Autoridades de la UTEQ |
| **Usuario** | Sistema de autenticación |
| **TipoUsuario** | Tipos de usuarios |
| **Institucion** | Instituciones educativas visitantes |
| **ProgramaVisita** | Programas de visitas programadas |
| **Actividad** | Actividades disponibles |
| **Tematica** | Temáticas de actividades |
| **VisitaDetalle** | Detalles de participación |
| **Dudas** | Sistema de preguntas y respuestas |
| **Ciudad** | Ciudades del país |
| **Provincia** | Provincias del país |
| **DetalleAutoridadDetallesVisita** | Relación programas-autoridades |

### 🏛️ Patrón de Capas

```
ApiEscuela/
├── models/          # Entidades y relaciones
├── repositories/    # Acceso a datos con GORM
├── services/        # Lógica de negocio (AuthService)
├── handlers/        # Controladores HTTP
├── middleware/      # Autenticación JWT
├── routers/         # Configuración de rutas
└── main.go         # Punto de entrada
```

## 🚀 API Endpoints

Base URL: `http://localhost:3000`

### 🔐 Autenticación

| Método | Endpoint | Descripción | Auth |
|--------|----------|-------------|------|
| `POST` | `/auth/login` | Iniciar sesión | ❌ |
| `POST` | `/auth/register` | Registrar usuario | ❌ |
| `POST` | `/auth/validate-token` | Validar token | ❌ |
| `GET` | `/api/auth/profile` | Perfil del usuario | ✅ |
| `POST` | `/api/auth/change-password` | Cambiar contraseña | ✅ |
| `POST` | `/api/auth/refresh-token` | Renovar token | ✅ |

### 📚 Endpoints por Entidad

#### 👤 **Personas**
```
POST   /api/personas                    # Crear persona
GET    /api/personas                    # Obtener todas las personas
GET    /api/personas/:id                # Obtener persona por ID
PUT    /api/personas/:id                # Actualizar persona
DELETE /api/personas/:id                # Eliminar persona
GET    /api/personas/cedula/:cedula     # Buscar por cédula
GET    /api/personas/correo/:correo     # Buscar por correo
```

#### 🎓 **Estudiantes**
```
POST   /api/estudiantes                           # Crear estudiante
GET    /api/estudiantes                           # Obtener estudiantes activos
GET    /api/estudiantes/all-including-deleted     # Obtener todos (activos + eliminados)
GET    /api/estudiantes/deleted                   # Obtener solo eliminados
GET    /api/estudiantes/:id                       # Obtener estudiante por ID
PUT    /api/estudiantes/:id                       # Actualizar estudiante
DELETE /api/estudiantes/:id                       # Eliminar estudiante (cascada)
PUT    /api/estudiantes/:id/restore               # Restaurar estudiante (cascada)
GET    /api/estudiantes/ciudad/:ciudad_id         # Filtrar por ciudad
GET    /api/estudiantes/institucion/:institucion_id # Filtrar por institución
GET    /api/estudiantes/especialidad/:especialidad # Filtrar por especialidad
```

#### 🎓 **Estudiantes Universitarios**
```
POST   /api/estudiantes-universitarios            # Crear estudiante universitario
GET    /api/estudiantes-universitarios            # Obtener todos
GET    /api/estudiantes-universitarios/:id        # Obtener por ID
PUT    /api/estudiantes-universitarios/:id        # Actualizar
DELETE /api/estudiantes-universitarios/:id        # Eliminar
GET    /api/estudiantes-universitarios/semestre/:semestre # Filtrar por semestre
```

#### 👨‍🏫 **Autoridades UTEQ**
```
POST   /api/autoridades-uteq                      # Crear autoridad
GET    /api/autoridades-uteq                      # Obtener autoridades activas
GET    /api/autoridades-uteq/all-including-deleted # Obtener todas (activas + eliminadas)
GET    /api/autoridades-uteq/deleted              # Obtener solo eliminadas
GET    /api/autoridades-uteq/:id                  # Obtener autoridad por ID
PUT    /api/autoridades-uteq/:id                  # Actualizar autoridad
DELETE /api/autoridades-uteq/:id                  # Eliminar autoridad (cascada)
PUT    /api/autoridades-uteq/:id/restore          # Restaurar autoridad (cascada)
GET    /api/autoridades-uteq/cargo/:cargo         # Filtrar por cargo
GET    /api/autoridades-uteq/persona/:persona_id  # Filtrar por persona
```

#### 👥 **Usuarios**
```
POST   /api/usuarios                              # Crear usuario
GET    /api/usuarios                              # Obtener usuarios activos
GET    /api/usuarios/all-including-deleted        # Obtener todos (activos + eliminados)
GET    /api/usuarios/deleted                      # Obtener solo eliminados
GET    /api/usuarios/:id                          # Obtener usuario por ID
PUT    /api/usuarios/:id                          # Actualizar usuario
DELETE /api/usuarios/:id                          # Eliminar usuario (soft delete)
PUT    /api/usuarios/:id/restore                  # Restaurar usuario eliminado
GET    /api/usuarios/username/:username           # Buscar por nombre de usuario
GET    /api/usuarios/tipo/:tipo_usuario_id        # Filtrar por tipo
GET    /api/usuarios/persona/:persona_id          # Filtrar por persona
```

#### 🏢 **Tipos de Usuario**
```
POST   /api/tipos-usuario                         # Crear tipo de usuario
GET    /api/tipos-usuario                         # Obtener todos los tipos
GET    /api/tipos-usuario/:id                     # Obtener tipo por ID
PUT    /api/tipos-usuario/:id                     # Actualizar tipo
DELETE /api/tipos-usuario/:id                     # Eliminar tipo
```

#### 🏫 **Instituciones**
```
POST   /api/instituciones                         # Crear institución
GET    /api/instituciones                         # Obtener todas las instituciones
GET    /api/instituciones/:id                     # Obtener institución por ID
PUT    /api/instituciones/:id                     # Actualizar institución
DELETE /api/instituciones/:id                     # Eliminar institución
GET    /api/instituciones/nombre/:nombre          # Buscar por nombre
```

#### 📅 **Programas de Visita**
```
POST   /api/programas-visita                      # Crear programa
GET    /api/programas-visita                      # Obtener todos los programas
GET    /api/programas-visita/:id                  # Obtener programa por ID
PUT    /api/programas-visita/:id                  # Actualizar programa
DELETE /api/programas-visita/:id                  # Eliminar programa
GET    /api/programas-visita/fecha/:fecha         # Filtrar por fecha (YYYY-MM-DD)
GET    /api/programas-visita/autoridad/:autoridad_id # Filtrar por autoridad
GET    /api/programas-visita/institucion/:institucion_id # Filtrar por institución
GET    /api/programas-visita/rango-fecha?inicio=YYYY-MM-DD&fin=YYYY-MM-DD # Rango de fechas
```

#### 🔗 **Detalle Autoridad Detalles Visita** (Relación Muchos-a-Muchos)
```
POST   /api/detalle-autoridad-detalles-visita     # Asignar autoridad a programa
GET    /api/detalle-autoridad-detalles-visita     # Obtener todas las asignaciones
GET    /api/detalle-autoridad-detalles-visita/:id # Obtener asignación por ID
PUT    /api/detalle-autoridad-detalles-visita/:id # Actualizar asignación
DELETE /api/detalle-autoridad-detalles-visita/:id # Eliminar asignación
GET    /api/detalle-autoridad-detalles-visita/programa-visita/:programa_visita_id # Autoridades por programa
GET    /api/detalle-autoridad-detalles-visita/autoridad/:autoridad_id # Programas por autoridad
```

#### 🎯 **Actividades**
```
POST   /api/actividades                           # Crear actividad
GET    /api/actividades                           # Obtener todas las actividades
GET    /api/actividades/:id                       # Obtener actividad por ID
PUT    /api/actividades/:id                       # Actualizar actividad
DELETE /api/actividades/:id                       # Eliminar actividad
GET    /api/actividades/tematica/:tematica_id     # Filtrar por temática
GET    /api/actividades/nombre/:nombre            # Buscar por nombre
GET    /api/actividades/duracion?min=30&max=120   # Filtrar por duración
```

#### 📚 **Temáticas**
```
POST   /api/tematicas                             # Crear temática
GET    /api/tematicas                             # Obtener todas las temáticas
GET    /api/tematicas/:id                         # Obtener temática por ID
PUT    /api/tematicas/:id                         # Actualizar temática
DELETE /api/tematicas/:id                         # Eliminar temática
GET    /api/tematicas/nombre/:nombre              # Buscar por nombre
```

#### 📋 **Visita Detalles**
```
POST   /api/visita-detalles                       # Crear detalle
GET    /api/visita-detalles                       # Obtener todos los detalles
GET    /api/visita-detalles/:id                   # Obtener detalle por ID
PUT    /api/visita-detalles/:id                   # Actualizar detalle
DELETE /api/visita-detalles/:id                   # Eliminar detalle
GET    /api/visita-detalles/actividad/:actividad_id # Filtrar por actividad
GET    /api/visita-detalles/programa/:programa_id # Filtrar por programa
GET    /api/visita-detalles/participantes?min=10&max=50 # Filtrar por participantes
GET    /api/visita-detalles/estadisticas          # Estadísticas de participación
```

#### 🎓 **Visita Detalle Estudiantes Universitarios**
```
POST   /api/visita-detalle-estudiantes-universitarios # Asignar estudiante a programa
GET    /api/visita-detalle-estudiantes-universitarios # Obtener todas las asignaciones
GET    /api/visita-detalle-estudiantes-universitarios/:id # Obtener asignación por ID
PUT    /api/visita-detalle-estudiantes-universitarios/:id # Actualizar asignación
DELETE /api/visita-detalle-estudiantes-universitarios/:id # Eliminar asignación
GET    /api/visita-detalle-estudiantes-universitarios/programa-visita/:programa_visita_id # Estudiantes por programa
GET    /api/visita-detalle-estudiantes-universitarios/estudiante/:estudiante_id # Programas por estudiante
DELETE /api/visita-detalle-estudiantes-universitarios/programa-visita/:programa_visita_id # Eliminar todos los estudiantes de un programa
DELETE /api/visita-detalle-estudiantes-universitarios/estudiante/:estudiante_id # Eliminar todos los programas de un estudiante
GET    /api/visita-detalle-estudiantes-universitarios/estadisticas # Estadísticas de participación estudiantil
```

#### ❓ **Dudas**
```
POST   /api/dudas                                 # Crear duda
GET    /api/dudas                                 # Obtener todas las dudas
GET    /api/dudas/:id                             # Obtener duda por ID
PUT    /api/dudas/:id                             # Actualizar duda
DELETE /api/dudas/:id                             # Eliminar duda
GET    /api/dudas/estudiante/:estudiante_id       # Filtrar por estudiante
GET    /api/dudas/autoridad/:autoridad_id         # Filtrar por autoridad
GET    /api/dudas/sin-responder                   # Dudas pendientes
GET    /api/dudas/respondidas                     # Dudas respondidas
GET    /api/dudas/sin-asignar                     # Dudas sin asignar
GET    /api/dudas/privacidad/:privacidad          # Filtrar por privacidad (publico/privado)
GET    /api/dudas/buscar/:termino                 # Búsqueda en preguntas
PUT    /api/dudas/:duda_id/asignar                # Asignar autoridad
PUT    /api/dudas/:duda_id/responder              # Responder duda
```

#### 🌍 **Provincias**
```
POST   /api/provincias                            # Crear provincia
GET    /api/provincias                            # Obtener todas las provincias
GET    /api/provincias/:id                        # Obtener provincia por ID
PUT    /api/provincias/:id                        # Actualizar provincia
DELETE /api/provincias/:id                        # Eliminar provincia
```

#### 🏙️ **Ciudades**
```
POST   /api/ciudades                              # Crear ciudad
GET    /api/ciudades                              # Obtener todas las ciudades
GET    /api/ciudades/:id                          # Obtener ciudad por ID
PUT    /api/ciudades/:id                          # Actualizar ciudad
DELETE /api/ciudades/:id                          # Eliminar ciudad
GET    /api/ciudades/provincia/:provincia_id      # Ciudades por provincia
```

### 🔗 Resumen de Operaciones

**Total de Endpoints**: 80+

**Operaciones CRUD Estándar** (todas las entidades):
- `POST /api/{entidad}` - Crear
- `GET /api/{entidad}` - Listar todos
- `GET /api/{entidad}/:id` - Obtener por ID
- `PUT /api/{entidad}/:id` - Actualizar
- `DELETE /api/{entidad}/:id` - Eliminar

### 🆕 Funcionalidades Especiales

#### **Soft Delete y Restauración**
```bash
# Obtener eliminados
GET /api/{entidad}/deleted

# Obtener todos (activos + eliminados)
GET /api/{entidad}/all-including-deleted

# Restaurar eliminado
PUT /api/{entidad}/:id/restore
```

#### **Eliminación en Cascada**
- **Estudiantes**: Elimina estudiante → usuario → persona
- **Autoridades UTEQ**: Elimina autoridad → usuario → persona

#### **Filtros Avanzados**
```bash
# Por rango de fechas
GET /api/programas-visita/rango-fecha?inicio=2024-01-01&fin=2024-12-31

# Por duración
GET /api/actividades/duracion?min=30&max=120

# Por privacidad
GET /api/dudas/privacidad/publico
```

#### **Estadísticas**
```bash
GET /api/visita-detalles/estadisticas
GET /api/visita-detalle-estudiantes-universitarios/estadisticas
```

## 🚨 Sistema de Errores

### 📋 Estructura Estándar

```json
{
  "error": "Descripción corta",
  "error_code": "CODIGO_ESPECIFICO",
  "message": "Mensaje detallado con solución",
  "status_code": 401,
  "timestamp": "2024-01-15T10:30:00Z",
  "path": "/api/estudiantes",
  "method": "GET"
}
```

### 🔐 Códigos de Error Principales

| Código | Descripción | Solución |
|--------|-------------|----------|
| `AUTH_TOKEN_MISSING` | Sin token | Incluir header Authorization |
| `AUTH_TOKEN_EXPIRED` | Token expirado | Hacer login o refresh |
| `AUTH_TOKEN_MALFORMED` | Token inválido | Verificar formato JWT |
| `LOGIN_USER_NOT_FOUND` | Usuario no existe | Verificar credenciales |
| `LOGIN_PASSWORD_INCORRECT_HASH` | Contraseña incorrecta | Verificar contraseña |

### 🧪 Probar Errores

```bash
# Sin token
curl -X GET http://localhost:3000/api/estudiantes

# Token malformado
curl -X GET http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer token_invalido"

# Campos faltantes en login
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"usuario": ""}'
```

## 📝 Ejemplos de Uso

### 🔐 Flujo Completo de Autenticación

```bash
# 1. Login
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"usuario": "test_user", "contraseña": "password123"}'

# Respuesta: {"token": "eyJhbGciOiJIUzI1NiIs...", ...}

# 2. Usar token (reemplazar TOKEN)
curl -X GET http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer TOKEN"

# 3. Renovar token
curl -X POST http://localhost:3000/api/auth/refresh-token \
  -H "Authorization: Bearer TOKEN"
```

### 📚 Gestión de Datos

```bash
# Crear persona
curl -X POST http://localhost:3000/api/personas \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Juan Pérez",
    "cedula": "1234567890",
    "correo": "juan@email.com",
    "telefono": "0987654321"
  }'

# Crear estudiante
curl -X POST http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "persona_id": 1,
    "institucion_id": 1,
    "ciudad_id": 1,
    "especialidad": "Ingeniería en Sistemas"
  }'

# Obtener estudiantes eliminados
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/deleted

# Restaurar estudiante
curl -X PUT -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/1/restore
```

### 🔍 Búsquedas y Filtros

```bash
# Buscar por cédula
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/personas/cedula/1234567890

# Filtrar por especialidad
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/especialidad/Ingeniería

# Dudas por privacidad
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/dudas/privacidad/publico

# Programas por rango de fechas
curl -H "Authorization: Bearer TOKEN" \
  "http://localhost:3000/api/programas-visita/rango-fecha?inicio=2024-01-01&fin=2024-12-31"
```

### 📊 Estadísticas

```bash
# Estadísticas de visitas
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/visita-detalles/estadisticas

# Estadísticas de participación estudiantil
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/visita-detalle-estudiantes-universitarios/estadisticas
```

## ⚙️ Instalación

### Prerrequisitos
- Go 1.24+
- PostgreSQL
- Acceso a la base de datos UTEQ

### Pasos

```bash
# 1. Clonar repositorio
git clone <repository-url>
cd ApiEscuela

# 2. Instalar dependencias
go mod tidy

# 3. Configurar variables de entorno (opcional)
echo "APP_PORT=3000
JWT_SECRET=tu_clave_secreta_super_segura
APP_ENV=development" > config.env

# 4. Ejecutar aplicación
go run main.go
```

La aplicación estará disponible en `http://localhost:3000`

## 🛠️ Tecnologías

- **Go 1.24**: Lenguaje de programación
- **Fiber v2**: Framework web HTTP de alto rendimiento
- **GORM**: ORM con soporte completo para PostgreSQL
- **PostgreSQL**: Base de datos relacional
- **JWT (golang-jwt/jwt/v5)**: Autenticación basada en tokens
- **bcrypt**: Encriptación segura de contraseñas
- **Viper**: Gestión de configuración

## 🔒 Seguridad

### JWT Configuration
- **Algoritmo**: HS256 (HMAC SHA-256)
- **Expiración**: 24 horas
- **Claims**: user_id, username, tipo_usuario_id

### Password Security
- **Encriptación**: bcrypt con salt automático
- **Longitud mínima**: 6 caracteres

### Route Protection
- **Middleware automático**: Validación en rutas `/api/*`
- **Context injection**: Info del usuario disponible en handlers

## ⚠️ Migración desde Versión Anterior

### Cambios Importantes

1. **URLs actualizadas**: Agregar `/api/` antes de todas las rutas protegidas
2. **Header requerido**: `Authorization: Bearer TOKEN` en todas las peticiones
3. **Flujo obligatorio**: Login → Token → Peticiones autenticadas

### Ejemplo de Migración

```javascript
// Antes
fetch('/estudiantes')

// Ahora
const token = localStorage.getItem('token');
fetch('/api/estudiantes', {
  headers: { 'Authorization': `Bearer ${token}` }
})
```

## 🚀 Estado del Proyecto

✅ **Sistema Completo y Funcional**
- 15 entidades implementadas
- 80+ endpoints API
- Sistema de autenticación JWT completo
- Sistema de errores estructurado
- Documentación completa

## 📚 Documentación Adicional

- **`AUTH_README.md`**: Documentación detallada de autenticación
- **Ejemplos completos**: JavaScript/Fetch, cURL, Postman
- **Guías de migración**: Actualización de código existente

## 📞 Soporte

Para soporte técnico o consultas, contactar al equipo de desarrollo de la UTEQ.

### Problemas Comunes

1. **Error 401**: Verificar header `Authorization: Bearer TOKEN`
2. **Token expirado**: Usar `/api/auth/refresh-token` o hacer login
3. **URLs incorrectas**: Asegurar prefijo `/api/` en rutas protegidas

---

**Desarrollado para la Universidad Técnica Estatal de Quevedo (UTEQ)**

**🔐 Versión 2.0 - Sistema Completo con Autenticación JWT**