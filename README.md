# ApiEscuela - Sistema de Gestión de Visitas Educativas UTEQ

Un backend completo desarrollado en Go para la gestión integral de visitas educativas a la UTEQ, incluyendo estudiantes, instituciones, programas de visita, actividades, sistema de dudas y **autenticación JWT completa**.

## 🔐 **NUEVO: Sistema de Autenticación JWT**

**¡IMPORTANTE!** Todos los endpoints de la API ahora requieren autenticación JWT excepto los endpoints públicos de autenticación.

### 🚀 Características de Seguridad Implementadas

- **JWT Tokens**: Autenticación basada en tokens con expiración de 24 horas
- **Contraseñas Encriptadas**: Usando bcrypt con salt automático
- **Middleware de Protección**: Validación automática en todas las rutas protegidas
- **Gestión de Sesiones**: Login, logout, renovación de tokens
- **Información de Usuario**: Disponible en el contexto de cada petición

### 🌐 Nueva Estructura de URLs

#### **Rutas Públicas (Sin autenticación)**
- `POST /auth/login` - Iniciar sesión
- `POST /auth/register` - Registrar nuevo usuario
- `POST /auth/validate-token` - Validar token
- `GET /` - Página de bienvenida
- `GET /health` - Estado de salud de la API

#### **Rutas Protegidas (Requieren JWT)**
Todas las rutas de la API ahora están bajo el prefijo `/api` y requieren autenticación:
- `/api/auth/*` - Rutas de autenticación protegidas
- `/api/estudiantes/*` - Gestión de estudiantes
- `/api/personas/*` - Gestión de personas
- `/api/provincias/*` - Gestión de provincias
- `/api/ciudades/*` - Gestión de ciudades
- Y todas las demás rutas existentes...

### 🔑 Endpoints de Autenticación

#### 1. Login
```bash
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "nombre_usuario",
    "contraseña": "contraseña_usuario"
  }'
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

#### 2. Registro
```bash
curl -X POST http://localhost:3000/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "nuevo_usuario",
    "contraseña": "contraseña_segura",
    "persona_id": 1,
    "tipo_usuario_id": 1
  }'
```

#### 3. Usar Token en Peticiones Protegidas
```bash
curl -X GET http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer tu_token_jwt_aqui"
```

#### 4. Obtener Perfil del Usuario Autenticado
```bash
curl -X GET http://localhost:3000/api/auth/profile \
  -H "Authorization: Bearer tu_token_jwt_aqui"
```

#### 5. Cambiar Contraseña
```bash
curl -X POST http://localhost:3000/api/auth/change-password \
  -H "Authorization: Bearer tu_token_jwt_aqui" \
  -H "Content-Type: application/json" \
  -d '{
    "old_password": "contraseña_actual",
    "new_password": "nueva_contraseña"
  }'
```

#### 6. Renovar Token
```bash
curl -X POST http://localhost:3000/api/auth/refresh-token \
  -H "Authorization: Bearer tu_token_jwt_aqui"
```

### 🛡️ Cómo Migrar a la Nueva Autenticación

Si ya tienes código que usa la API, necesitas hacer estos cambios:

1. **Obtener un token primero:**
```javascript
const loginResponse = await fetch('/auth/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    usuario: 'mi_usuario',
    contraseña: 'mi_contraseña'
  })
});
const { token } = await loginResponse.json();
```

2. **Actualizar todas las URLs:** Agregar `/api` antes de la ruta
```javascript
// Antes
fetch('/estudiantes')

// Ahora
fetch('/api/estudiantes', {
  headers: { 'Authorization': `Bearer ${token}` }
})
```

3. **Incluir el token en todas las peticiones:**
```javascript
const response = await fetch('/api/estudiantes', {
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  }
});
```

## 🏗️ Estructura del Proyecto

```
ApiEscuela/
├── models/          # 15 modelos de datos (entidades del sistema)
├── repositories/    # Repositorios para acceso a datos
├── handlers/        # Controladores HTTP para todas las entidades
├── services/        # 🆕 Servicios de negocio (AuthService)
├── middleware/      # 🆕 Middleware de autenticación JWT
├── routers/         # Configuración consolidada de rutas
├── main.go         # Punto de entrada de la aplicación
├── config.env      # Variables de entorno
├── README.md       # Este archivo
└── AUTH_README.md  # 🆕 Documentación detallada de autenticación
```

## 📊 Modelos del Sistema

### Entidades Principales (15 modelos)

1. **Persona** - Información básica de personas
2. **Estudiante** - Estudiantes de instituciones educativas
3. **EstudianteUniversitario** - Estudiantes universitarios de la UTEQ
4. **AutoridadUTEQ** - Autoridades de la UTEQ
5. **Institucion** - Instituciones educativas visitantes
6. **ProgramaVisita** - Programas de visitas programadas
7. **DetalleAutoridadDetallesVisita** - **🆕 Relación muchos-a-muchos entre programas y autoridades**
8. **Actividad** - Actividades disponibles en las visitas
9. **Tematica** - Temáticas de las actividades
10. **VisitaDetalle** - Detalles de participación en visitas
11. **Dudas** - Sistema de preguntas y respuestas
12. **Usuario** - Usuarios del sistema con autenticación
13. **TipoUsuario** - Tipos de usuarios del sistema
14. **Ciudad** - Ciudades del país
15. **Provincia** - Provincias del país

### Ejemplo de Modelo Principal

```go
type Estudiante struct {
    gorm.Model
    PersonaID     uint   `json:"persona_id" gorm:"not null"`
    InstitucionID uint   `json:"institucion_id" gorm:"not null"`
    CiudadID      uint   `json:"ciudad_id" gorm:"not null"`
    Especialidad  string `json:"especialidad"`
    
    // Relaciones
    Persona     Persona     `json:"persona,omitempty"`
    Institucion Institucion `json:"institucion,omitempty"`
    Ciudad      Ciudad      `json:"ciudad,omitempty"`
    Dudas       []Dudas     `json:"dudas,omitempty"`
}
```

## 🚀 API Endpoints (80+ endpoints)

Base URL: `http://localhost:3000`

### 🔐 **AUTENTICACIÓN (Rutas Públicas)**
- `POST /auth/login` - **🔑 Iniciar sesión**
- `POST /auth/register` - **👤 Registrar nuevo usuario**
- `POST /auth/validate-token` - **✅ Validar token**

### 🔒 **AUTENTICACIÓN PROTEGIDA (Requiere JWT)**
- `GET /api/auth/profile` - **👤 Obtener perfil del usuario**
- `POST /api/auth/change-password` - **🔒 Cambiar contraseña**
- `POST /api/auth/refresh-token` - **🔄 Renovar token**

### 📚 **ESTUDIANTES (Requiere JWT)**
- `POST /api/estudiantes` - Crear estudiante
- `GET /api/estudiantes` - Obtener todos los estudiantes activos
- `GET /api/estudiantes/all-including-deleted` - **📋 Obtener todos los estudiantes (activos + eliminados)** (NUEVO)
- `GET /api/estudiantes/deleted` - **🗑️ Obtener solo estudiantes eliminados** (NUEVO)
- `GET /api/estudiantes/:id` - Obtener estudiante por ID
- `PUT /api/estudiantes/:id` - Actualizar estudiante
- `DELETE /api/estudiantes/:id` - **🗑️ Eliminar estudiante, usuario y persona en cascada**
- `PUT /api/estudiantes/:id/restore` - **♻️ Restaurar estudiante, usuario y persona en cascada** (NUEVO)
- `GET /api/estudiantes/ciudad/:ciudad_id` - Filtrar por ciudad
- `GET /api/estudiantes/institucion/:institucion_id` - Filtrar por institución
- `GET /api/estudiantes/especialidad/:especialidad` - Filtrar por especialidad

### 👤 **PERSONAS (Requiere JWT)**
- `POST /api/personas` - Crear persona
- `GET /api/personas` - Obtener todas las personas
- `GET /api/personas/:id` - Obtener persona por ID
- `PUT /api/personas/:id` - Actualizar persona
- `DELETE /api/personas/:id` - Eliminar persona
- `GET /api/personas/cedula/:cedula` - Buscar por cédula
- `GET /api/personas/correo/:correo` - Buscar por correo

### 🔐 **SISTEMA DE USUARIOS (Requiere JWT)**
- `POST /api/usuarios` - Crear usuario
- `GET /api/usuarios` - Obtener todos los usuarios activos
- `GET /api/usuarios/all-including-deleted` - **📋 Obtener todos los usuarios (activos + eliminados)**
- `GET /api/usuarios/deleted` - **🗑️ Obtener solo usuarios eliminados**
- `GET /api/usuarios/:id` - Obtener usuario por ID
- `PUT /api/usuarios/:id` - Actualizar usuario
- `DELETE /api/usuarios/:id` - Eliminar usuario (soft delete)
- `PUT /api/usuarios/:id/restore` - **♻️ Restaurar usuario eliminado**
- `GET /api/usuarios/username/:username` - Buscar por nombre de usuario
- `GET /api/usuarios/tipo/:tipo_usuario_id` - Filtrar por tipo
- `GET /api/usuarios/persona/:persona_id` - Filtrar por persona

### 📅 **PROGRAMAS DE VISITA (Requiere JWT)**
- `POST /api/programas-visita` - Crear programa
- `GET /api/programas-visita` - Obtener todos los programas
- `GET /api/programas-visita/:id` - Obtener programa por ID
- `PUT /api/programas-visita/:id` - Actualizar programa
- `DELETE /api/programas-visita/:id` - Eliminar programa
- `GET /api/programas-visita/fecha/:fecha` - **Filtrar por fecha (YYYY-MM-DD)**
- `GET /api/programas-visita/autoridad/:autoridad_id` - Filtrar por autoridad
- `GET /api/programas-visita/institucion/:institucion_id` - Filtrar por institución
- `GET /api/programas-visita/rango-fecha?inicio=2024-01-01&fin=2024-12-31` - **Rango de fechas**

### 🔗 **DETALLE AUTORIDAD DETALLES VISITA (Requiere JWT)** (🆕 Relación Muchos-a-Muchos)
- `POST /api/detalle-autoridad-detalles-visita` - **Asignar autoridad a programa**
- `GET /api/detalle-autoridad-detalles-visita` - Obtener todas las asignaciones
- `GET /api/detalle-autoridad-detalles-visita/:id` - Obtener asignación por ID
- `PUT /api/detalle-autoridad-detalles-visita/:id` - Actualizar asignación
- `DELETE /api/detalle-autoridad-detalles-visita/:id` - Eliminar asignación
- `GET /api/detalle-autoridad-detalles-visita/programa-visita/:programa_visita_id` - **Autoridades por programa**
- `GET /api/detalle-autoridad-detalles-visita/autoridad/:autoridad_id` - **Programas por autoridad**

### 🎯 **ACTIVIDADES Y TEMÁTICAS (Requiere JWT)**
- `POST /api/actividades` - Crear actividad
- `GET /api/actividades` - Obtener todas las actividades
- `GET /api/actividades/:id` - Obtener actividad por ID
- `PUT /api/actividades/:id` - Actualizar actividad
- `DELETE /api/actividades/:id` - Eliminar actividad
- `GET /api/actividades/tematica/:tematica_id` - Filtrar por temática
- `GET /api/actividades/nombre/:nombre` - Buscar por nombre
- `GET /api/actividades/duracion?min=30&max=120` - **Filtrar por duración**

### 📋 **VISITA DETALLES Y ESTADÍSTICAS (Requiere JWT)** (🆕 Estructura Actualizada)
- `POST /api/visita-detalles` - Crear detalle
- `GET /api/visita-detalles` - Obtener todos los detalles
- `GET /api/visita-detalles/:id` - Obtener detalle por ID
- `PUT /api/visita-detalles/:id` - Actualizar detalle
- `DELETE /api/visita-detalles/:id` - Eliminar detalle
- `GET /api/visita-detalles/actividad/:actividad_id` - Filtrar por actividad
- `GET /api/visita-detalles/programa/:programa_id` - Filtrar por programa
- `GET /api/visita-detalles/participantes?min=10&max=50` - **Filtrar por participantes**
- `GET /api/visita-detalles/estadisticas` - **📊 Estadísticas de participación**

### 🆕 **ESTUDIANTES UNIVERSITARIOS EN PROGRAMAS DE VISITA (Requiere JWT)**
- `POST /api/visita-detalle-estudiantes-universitarios` - **Asignar estudiante a programa**
- `GET /api/visita-detalle-estudiantes-universitarios` - Obtener todas las asignaciones
- `GET /api/visita-detalle-estudiantes-universitarios/:id` - Obtener asignación por ID
- `PUT /api/visita-detalle-estudiantes-universitarios/:id` - Actualizar asignación
- `DELETE /api/visita-detalle-estudiantes-universitarios/:id` - Eliminar asignación
- `GET /api/visita-detalle-estudiantes-universitarios/programa-visita/:programa_visita_id` - **Estudiantes por programa**
- `GET /api/visita-detalle-estudiantes-universitarios/estudiante/:estudiante_id` - **Programas por estudiante**
- `DELETE /api/visita-detalle-estudiantes-universitarios/programa-visita/:programa_visita_id` - **Eliminar todos los estudiantes de un programa**
- `DELETE /api/visita-detalle-estudiantes-universitarios/estudiante/:estudiante_id` - **Eliminar todos los programas de un estudiante**
- `GET /api/visita-detalle-estudiantes-universitarios/estadisticas` - **📊 Estadísticas de participación estudiantil**

### ❓ **SISTEMA DE DUDAS CON PRIVACIDAD (Requiere JWT)**
- `POST /api/dudas` - Crear duda
- `GET /api/dudas` - Obtener todas las dudas
- `GET /api/dudas/:id` - Obtener duda por ID
- `PUT /api/dudas/:id` - Actualizar duda
- `DELETE /api/dudas/:id` - Eliminar duda
- `GET /api/dudas/estudiante/:estudiante_id` - Filtrar por estudiante
- `GET /api/dudas/autoridad/:autoridad_id` - Filtrar por autoridad
- `GET /api/dudas/sin-responder` - **📋 Dudas pendientes**
- `GET /api/dudas/respondidas` - **✅ Dudas respondidas**
- `GET /api/dudas/sin-asignar` - **⚠️ Dudas sin asignar**
- `GET /api/dudas/privacidad/:privacidad` - **🎯 Filtrar por privacidad (publico/privado)** (NUEVO)
- `GET /api/dudas/buscar/:termino` - **🔍 Búsqueda en preguntas**
- `PUT /api/dudas/:duda_id/asignar` - **👤 Asignar autoridad**
- `PUT /api/dudas/:duda_id/responder` - **💬 Responder duda**

### 🌍 **UBICACIONES GEOGRÁFICAS (Requiere JWT)**
- `GET /api/provincias` - Obtener todas las provincias
- `GET /api/ciudades` - Obtener todas las ciudades
- `GET /api/ciudades/provincia/:provincia_id` - Ciudades por provincia

### 🏫 **INSTITUCIONES Y AUTORIDADES (Requiere JWT)**
- `GET /api/instituciones` - Obtener todas las instituciones
- `GET /api/instituciones/nombre/:nombre` - Buscar por nombre

### 🎓 **AUTORIDADES UTEQ (Requiere JWT)**
- `POST /api/autoridades-uteq` - Crear autoridad UTEQ
- `GET /api/autoridades-uteq` - Obtener todas las autoridades activas
- `GET /api/autoridades-uteq/all-including-deleted` - **📋 Obtener todas las autoridades (activas + eliminadas)** (NUEVO)
- `GET /api/autoridades-uteq/deleted` - **🗑️ Obtener solo autoridades eliminadas** (NUEVO)
- `GET /api/autoridades-uteq/:id` - Obtener autoridad por ID
- `PUT /api/autoridades-uteq/:id` - Actualizar autoridad
- `DELETE /api/autoridades-uteq/:id` - **🗑️ Eliminar autoridad, usuario y persona en cascada** (NUEVO)
- `PUT /api/autoridades-uteq/:id/restore` - **♻️ Restaurar autoridad, usuario y persona en cascada** (NUEVO)
- `GET /api/autoridades-uteq/cargo/:cargo` - Filtrar por cargo
- `GET /api/autoridades-uteq/persona/:persona_id` - Filtrar por persona

## 📋 Estructuras JSON de los Modelos

### 🔐 **Autenticación**

#### Login Request
```json
{
  "usuario": "nombre_usuario",
  "contraseña": "contraseña_usuario"
}
```

#### Register Request
```json
{
  "usuario": "nuevo_usuario",
  "contraseña": "contraseña_segura",
  "persona_id": 1,
  "tipo_usuario_id": 1
}
```

#### Change Password Request
```json
{
  "old_password": "contraseña_actual",
  "new_password": "nueva_contraseña"
}
```

### Persona
```json
{
  "nombre": "Juan Carlos Pérez",
  "cedula": "1234567890",
  "correo": "juan.perez@email.com",
  "telefono": "0987654321",
  "fecha_nacimiento": "1990-05-15T00:00:00Z"
}
```

### Estudiante
```json
{
  "persona_id": 1,
  "institucion_id": 1,
  "ciudad_id": 1,
  "especialidad": "Ingeniería en Sistemas"
}
```

### Usuario
```json
{
  "usuario": "jperez",
  "contraseña": "password123",
  "persona_id": 1,
  "tipo_usuario_id": 1
}
```

### Provincia
```json
{
  "provincia": "Los Ríos"
}
```

### Ciudad
```json
{
  "ciudad": "Quevedo",
  "provincia_id": 1
}
```

### Institución
```json
{
  "nombre": "Unidad Educativa San José",
  "autoridad": "Dr. María González",
  "contacto": "0987654321",
  "direccion": "Av. Principal 123, Quevedo"
}
```

### Programa de Visita
```json
{
  "fecha": "2024-03-15T09:00:00Z",
  "institucion_id": 1
}
```

### 🆕 Detalle Autoridad Detalles Visita (Relación Muchos-a-Muchos)
```json
{
  "programa_visita_id": 1,
  "autoridad_uteq_id": 2
}
```

**Ejemplo de Respuesta Completa**:
```json
{
  "id": 1,
  "programa_visita_id": 1,
  "autoridad_uteq_id": 2,
  "created_at": "2024-03-15T10:00:00Z",
  "updated_at": "2024-03-15T10:00:00Z",
  "programa_visita": {
    "id": 1,
    "fecha": "2024-03-15T09:00:00Z",
    "institucion_id": 1
  },
  "autoridad_uteq": {
    "id": 2,
    "persona_id": 5,
    "cargo": "Decano de Facultad"
  }
}
```

### Actividad
```json
{
  "actividad": "Visita a Laboratorio de Suelos",
  "tematica_id": 1,
  "duracion": 90
}
```

### Temática
```json
{
  "nombre": "Ingeniería Agrícola",
  "descripcion": "Temática sobre técnicas modernas de agricultura"
}
```

### 🆕 Duda con Privacidad
```json
{
  "pregunta": "¿Cuáles son los requisitos de ingreso?",
  "estudiante_id": 1,
  "privacidad": "publico"
}
```

**Campos Opcionales**:
- `fecha_pregunta`: Se establece automáticamente al crear la duda si no se proporciona
- `respuesta`: Campo opcional, se llena cuando se responde la duda
- `fecha_respuesta`: Se establece automáticamente al responder la duda
- `autoridad_uteq_id`: Campo opcional, se asigna cuando una autoridad toma la duda
- `privacidad`: Campo opcional, por defecto es "publico" (valores: "publico" o "privado")

**Ejemplo de Duda Completa** (después de ser asignada y respondida):
```json
{
  "id": 1,
  "pregunta": "¿Cuáles son los requisitos de ingreso?",
  "fecha_pregunta": "2024-03-15T10:30:00Z",
  "respuesta": "Los requisitos incluyen bachillerato completo...",
  "fecha_respuesta": "2024-03-15T14:45:00Z",
  "estudiante_id": 1,
  "autoridad_uteq_id": 2,
  "privacidad": "publico"
}
```

### 🆕 VisitaDetalle (Estructura Actualizada)
```json
{
  "actividad_id": 1,
  "programa_visita_id": 1,
  "participantes": 25
}
```

**Ejemplo de Respuesta Completa**:
```json
{
  "id": 1,
  "actividad_id": 1,
  "programa_visita_id": 1,
  "participantes": 25,
  "created_at": "2024-03-15T10:00:00Z",
  "updated_at": "2024-03-15T10:00:00Z",
  "actividad": {
    "id": 1,
    "actividad": "Visita a Laboratorio de Suelos",
    "tematica_id": 1,
    "duracion": 90
  },
  "programa_visita": {
    "id": 1,
    "fecha": "2024-03-15T09:00:00Z",
    "institucion_id": 1
  }
}
```

### 🆕 VisitaDetalleEstudiantesUniversitarios (Nueva Tabla de Relación)
```json
{
  "estudiante_universitario_id": 1,
  "programa_visita_id": 1
}
```

**Ejemplo de Respuesta Completa**:
```json
{
  "id": 1,
  "estudiante_universitario_id": 1,
  "programa_visita_id": 1,
  "created_at": "2024-03-15T10:00:00Z",
  "updated_at": "2024-03-15T10:00:00Z",
  "estudiante_universitario": {
    "id": 1,
    "persona_id": 2,
    "semestre": 5,
    "persona": {
      "id": 2,
      "nombre": "Ana María López",
      "cedula": "0987654321",
      "correo": "ana.lopez@uteq.edu.ec"
    }
  },
  "programa_visita": {
    "id": 1,
    "fecha": "2024-03-15T09:00:00Z",
    "institucion_id": 1,
    "institucion": {
      "id": 1,
      "nombre": "Unidad Educativa San José"
    }
  }
}
```

## 🗄️ Configuración de Base de Datos

El proyecto está configurado para conectarse a la base de datos de la UTEQ:

### Automigración
El sistema crea automáticamente todas las 15 tablas con sus relaciones al iniciar.

## ⚙️ Instalación y Ejecución

### Prerrequisitos
- Go 1.24+
- Acceso a la base de datos UTEQ

### Pasos

1. **Clonar el repositorio**
   ```bash
   git clone <repository-url>
   cd ApiEscuela
   ```

2. **Instalar dependencias**
   ```bash
   go mod tidy
   ```

3. **Configurar variables de entorno (opcional)**
   ```bash
   # Crear config.env si necesitas configuraciones personalizadas
   APP_PORT=3000
   APP_ENV=development
   JWT_SECRET=tu_clave_secreta_super_segura_aqui
   ```

4. **Ejecutar la aplicación**
   ```bash
   go run main.go
   ```

La aplicación estará disponible en `http://localhost:3000`

## 📝 Ejemplos de Uso

### 🔐 **Flujo Completo de Autenticación**

#### 1. Registrar un nuevo usuario
```bash
curl -X POST http://localhost:3000/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "test_user",
    "contraseña": "password123",
    "persona_id": 1,
    "tipo_usuario_id": 1
  }'
```

#### 2. Hacer login
```bash
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "test_user",
    "contraseña": "password123"
  }'
```

**Respuesta:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "usuario": {...},
  "message": "Login exitoso"
}
```

#### 3. Usar el token para acceder a datos protegidos
```bash
# Reemplaza TOKEN con el token recibido del login
curl -X GET http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer TOKEN"
```

#### 4. Obtener perfil del usuario
```bash
curl -X GET http://localhost:3000/api/auth/profile \
  -H "Authorization: Bearer TOKEN"
```

#### 5. Cambiar contraseña
```bash
curl -X POST http://localhost:3000/api/auth/change-password \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "old_password": "password123",
    "new_password": "nueva_password456"
  }'
```

#### 6. Renovar token
```bash
curl -X POST http://localhost:3000/api/auth/refresh-token \
  -H "Authorization: Bearer TOKEN"
```

### 📚 **Gestión de Datos (Requiere Autenticación)**

**Nota:** Todos los siguientes ejemplos requieren el header `Authorization: Bearer TOKEN`

### Crear una Persona
```bash
curl -X POST http://localhost:3000/api/personas \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Juan Carlos Pérez",
    "cedula": "1234567890",
    "correo": "juan.perez@email.com",
    "telefono": "0987654321",
    "fecha_nacimiento": "1990-05-15T00:00:00Z"
  }'
```

### Crear un Estudiante
```bash
curl -X POST http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "persona_id": 1,
    "institucion_id": 1,
    "ciudad_id": 1,
    "especialidad": "Ingeniería en Sistemas"
  }'
```

### Crear un Programa de Visita
```bash
curl -X POST http://localhost:3000/api/programas-visita \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "fecha": "2024-03-15T09:00:00Z",
    "institucion_id": 1
  }'
```

### 🆕 Asignar Autoridad a Programa de Visita
```bash
curl -X POST http://localhost:3000/api/detalle-autoridad-detalles-visita \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "programa_visita_id": 1,
    "autoridad_uteq_id": 2
  }'
```

### Obtener Autoridades de un Programa
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/detalle-autoridad-detalles-visita/programa-visita/1
```

### Obtener Programas de una Autoridad
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/detalle-autoridad-detalles-visita/autoridad/2
```

### Crear un Usuario
```bash
curl -X POST http://localhost:3000/api/usuarios \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "jperez",
    "contraseña": "password123",
    "persona_id": 1,
    "tipo_usuario_id": 1
  }'
```

### 🗑️ Gestión de Usuarios Eliminados (Soft Delete)

### Obtener Todos los Usuarios (Incluyendo Eliminados)
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/usuarios/all-including-deleted
```

### Obtener Solo Usuarios Eliminados
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/usuarios/deleted
```

### Restaurar Usuario Eliminado
```bash
curl -X PUT -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/usuarios/5/restore
```

### Ejemplo de Flujo Completo de Soft Delete
```bash
# 1. Eliminar usuario (soft delete)
curl -X DELETE -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/usuarios/5

# 2. Verificar que aparece en usuarios eliminados
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/usuarios/deleted

# 3. Restaurar el usuario
curl -X PUT -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/usuarios/5/restore

# 4. Verificar que el usuario está activo nuevamente
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/usuarios/5
```

### 🆕 Gestión de Estudiantes con Eliminación en Cascada

### Eliminar Estudiante (Cascada: Estudiante + Usuario + Persona)
```bash
# Elimina el estudiante y automáticamente elimina su usuario y persona asociada
curl -X DELETE -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/3
```

**Respuesta**:
```json
{
  "message": "Estudiante, usuario y persona eliminados exitosamente"
}
```

### Restaurar Estudiante (Cascada: Estudiante + Usuario + Persona)
```bash
# Restaura el estudiante y automáticamente restaura su usuario y persona asociada
curl -X PUT -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/3/restore
```

**Respuesta**:
```json
{
  "message": "Estudiante, usuario y persona restaurados exitosamente"
}
```

### Ejemplo de Flujo Completo de Eliminación/Restauración en Cascada
```bash
# 1. Crear una persona
curl -X POST http://localhost:3000/api/personas \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "María González",
    "cedula": "0987654321",
    "correo": "maria.gonzalez@email.com",
    "telefono": "0987654321",
    "fecha_nacimiento": "1995-08-20T00:00:00Z"
  }'

# 2. Crear un usuario para esa persona
curl -X POST http://localhost:3000/api/usuarios \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "mgonzalez",
    "contraseña": "password123",
    "persona_id": 2,
    "tipo_usuario_id": 1
  }'

# 3. Crear un estudiante para esa persona
curl -X POST http://localhost:3000/api/estudiantes \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "persona_id": 2,
    "institucion_id": 1,
    "ciudad_id": 1,
    "especialidad": "Ingeniería Ambiental"
  }'

# 4. Eliminar el estudiante (elimina automáticamente usuario y persona)
curl -X DELETE -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/2

# 5. Verificar que el usuario también fue eliminado
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/usuarios/deleted

# 6. Restaurar el estudiante (restaura automáticamente usuario y persona)
curl -X PUT -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/2/restore

# 7. Verificar que todo fue restaurado correctamente
curl -H "Authorization: Bearer TOKEN" http://localhost:3000/api/estudiantes/2
curl -H "Authorization: Bearer TOKEN" http://localhost:3000/api/usuarios/2
curl -H "Authorization: Bearer TOKEN" http://localhost:3000/api/personas/2
```

### 🆕 Gestión Completa de Estudiantes Eliminados (Similar a Usuarios)

### Obtener Todos los Estudiantes (Incluyendo Eliminados)
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/all-including-deleted
```

### Obtener Solo Estudiantes Eliminados
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/deleted
```

### Ejemplo de Flujo Completo de Gestión de Estudiantes Eliminados
```bash
# 1. Obtener todos los estudiantes activos
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes

# 2. Eliminar un estudiante (cascada: estudiante + usuario + persona)
curl -X DELETE -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/3

# 3. Verificar que ya no aparece en estudiantes activos
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes

# 4. Verificar que aparece en estudiantes eliminados
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/deleted

# 5. Obtener todos los estudiantes incluyendo eliminados
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/all-including-deleted

# 6. Restaurar el estudiante eliminado
curl -X PUT -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/3/restore

# 7. Verificar que vuelve a aparecer en estudiantes activos
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes

# 8. Verificar que ya no aparece en estudiantes eliminados
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/estudiantes/deleted
```

### ���� Gestión de Autoridades UTEQ con Eliminación en Cascada

### Crear una Autoridad UTEQ
```bash
curl -X POST http://localhost:3000/api/autoridades-uteq \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "persona_id": 3,
    "cargo": "Decano de Facultad de Ingeniería"
  }'
```

### Eliminar Autoridad UTEQ (Cascada: Autoridad + Usuario + Persona)
```bash
# Elimina la autoridad y automáticamente elimina su usuario y persona asociada
curl -X DELETE -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/2
```

**Respuesta**:
```json
{
  "message": "Autoridad UTEQ, usuario y persona eliminados exitosamente"
}
```

### Restaurar Autoridad UTEQ (Cascada: Autoridad + Usuario + Persona)
```bash
# Restaura la autoridad y automáticamente restaura su usuario y persona asociada
curl -X PUT -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/2/restore
```

**Respuesta**:
```json
{
  "message": "Autoridad UTEQ, usuario y persona restaurados exitosamente"
}
```

### Obtener Todas las Autoridades (Incluyendo Eliminadas)
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/all-including-deleted
```

### Obtener Solo Autoridades Eliminadas
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/deleted
```

### Ejemplo de Flujo Completo de Gestión de Autoridades UTEQ Eliminadas
```bash
# 1. Crear una persona para la autoridad
curl -X POST http://localhost:3000/api/personas \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Dr. Carlos Mendoza",
    "cedula": "1234567890",
    "correo": "carlos.mendoza@uteq.edu.ec",
    "telefono": "0987654321",
    "fecha_nacimiento": "1975-03-10T00:00:00Z"
  }'

# 2. Crear un usuario para esa persona
curl -X POST http://localhost:3000/api/usuarios \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "usuario": "cmendoza",
    "contraseña": "password123",
    "persona_id": 3,
    "tipo_usuario_id": 2
  }'

# 3. Crear una autoridad UTEQ para esa persona
curl -X POST http://localhost:3000/api/autoridades-uteq \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "persona_id": 3,
    "cargo": "Decano de Facultad de Ingeniería"
  }'

# 4. Obtener todas las autoridades activas
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq

# 5. Eliminar la autoridad (cascada: autoridad + usuario + persona)
curl -X DELETE -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/2

# 6. Verificar que ya no aparece en autoridades activas
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq

# 7. Verificar que aparece en autoridades eliminadas
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/deleted

# 8. Obtener todas las autoridades incluyendo eliminadas
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/all-including-deleted

# 9. Restaurar la autoridad eliminada
curl -X PUT -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/2/restore

# 10. Verificar que vuelve a aparecer en autoridades activas
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq

# 11. Verificar que ya no aparece en autoridades eliminadas
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/deleted

# 12. Verificar que todo fue restaurado correctamente
curl -H "Authorization: Bearer TOKEN" http://localhost:3000/api/autoridades-uteq/2
curl -H "Authorization: Bearer TOKEN" http://localhost:3000/api/usuarios/3
curl -H "Authorization: Bearer TOKEN" http://localhost:3000/api/personas/3
```

### Filtrar Autoridades por Cargo
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/cargo/Decano
```

### Obtener Autoridad por Persona
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/autoridades-uteq/persona/3
```

### Crear una Provincia
```bash
curl -X POST http://localhost:3000/api/provincias \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "provincia": "Los Ríos"
  }'
```

### Crear una Ciudad
```bash
curl -X POST http://localhost:3000/api/ciudades \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "ciudad": "Quevedo",
    "provincia_id": 1
  }'
```

### Crear una Institución
```bash
curl -X POST http://localhost:3000/api/instituciones \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Unidad Educativa San José",
    "autoridad": "Dr. María González",
    "contacto": "0987654321",
    "direccion": "Av. Principal 123, Quevedo"
  }'
```

### Crear una Temática
```bash
curl -X POST http://localhost:3000/api/tematicas \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Ingeniería Agrícola",
    "descripcion": "Temática sobre técnicas modernas de agricultura"
  }'
```

### Crear una Actividad
```bash
curl -X POST http://localhost:3000/api/actividades \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "actividad": "Visita a Laboratorio de Suelos",
    "tematica_id": 1,
    "duracion": 90
  }'
```

### 🆕 Crear una Duda con Privacidad
```bash
# Crear duda pública (por defecto)
curl -X POST http://localhost:3000/api/dudas \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "pregunta": "¿Cuáles son los requisitos de ingreso?",
    "estudiante_id": 1
  }'

# Crear duda privada
curl -X POST http://localhost:3000/api/dudas \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "pregunta": "¿Hay becas disponibles para estudiantes de bajos recursos?",
    "estudiante_id": 1,
    "privacidad": "privado"
  }'

# Crear duda pública explícitamente
curl -X POST http://localhost:3000/api/dudas \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "pregunta": "¿Cuándo son las inscripciones?",
    "estudiante_id": 1,
    "privacidad": "publico"
  }'
```

### 🆕 Filtrar Dudas por Privacidad
```bash
# Filtrar por privacidad específica
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/dudas/privacidad/publico

curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/dudas/privacidad/privado
```

### Responder una Duda
```bash
curl -X PUT http://localhost:3000/api/dudas/1/responder \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "respuesta": "Los requisitos incluyen bachillerato completo y aprobar el examen de admisión."
  }'
```

### 🆕 Crear VisitaDetalle (Estructura Actualizada)
```bash
# Crear detalle de visita (sin estudiantes universitarios)
curl -X POST http://localhost:3000/api/visita-detalles \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "actividad_id": 1,
    "programa_visita_id": 1,
    "participantes": 25
  }'
```

### 🆕 Asignar Estudiante Universitario a Programa de Visita
```bash
# Asignar estudiante universitario a programa de visita
curl -X POST http://localhost:3000/api/visita-detalle-estudiantes-universitarios \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "estudiante_universitario_id": 1,
    "programa_visita_id": 1
  }'
```

### 🆕 Obtener Estudiantes de un Programa de Visita
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/visita-detalle-estudiantes-universitarios/programa-visita/1
```

### 🆕 Obtener Programas de Visita de un Estudiante
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/visita-detalle-estudiantes-universitarios/estudiante/1
```

### 🆕 Eliminar Todos los Estudiantes de un Programa
```bash
curl -X DELETE -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/visita-detalle-estudiantes-universitarios/programa-visita/1
```

### 🆕 Obtener Estadísticas de Participación Estudiantil
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/visita-detalle-estudiantes-universitarios/estadisticas
```

**Ejemplo de Respuesta de Estadísticas**:
```json
{
  "total_participaciones": 45,
  "total_estudiantes_unicos": 15,
  "total_programas_con_estudiantes": 8,
  "promedio_estudiantes_por_programa": 5.625
}
```

### Obtener Estadísticas de Participación
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/visita-detalles/estadisticas
```

### Buscar Dudas Pendientes
```bash
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:3000/api/dudas/sin-responder
```

### Filtrar Actividades por Duración
```bash
curl -H "Authorization: Bearer TOKEN" \
  "http://localhost:3000/api/actividades/duracion?min=30&max=120"
```

### Obtener Programas por Rango de Fechas
```bash
curl -H "Authorization: Bearer TOKEN" \
  "http://localhost:3000/api/programas-visita/rango-fecha?inicio=2024-01-01&fin=2024-12-31"
```

## 🛠️ Tecnologías Utilizadas

- **Go 1.24**: Lenguaje de programación
- **Fiber v2**: Framework web HTTP de alto rendimiento
- **GORM**: ORM para Go con soporte completo para PostgreSQL
- **PostgreSQL**: Base de datos relacional
- **Viper**: Gestión de configuración y variables de entorno
- **🆕 JWT (golang-jwt/jwt/v5)**: Autenticación basada en tokens
- **🆕 bcrypt**: Encriptación segura de contraseñas

## 🏛️ Arquitectura del Sistema

### Patrón de Capas
- **Models**: Definición de entidades y relaciones
- **Repositories**: Capa de acceso a datos con GORM
- **🆕 Services**: Lógica de negocio (AuthService)
- **Handlers**: Controladores HTTP con validación
- **🆕 Middleware**: Middleware de autenticación JWT
- **Routers**: Configuración consolidada de rutas

### Características Técnicas
- **CRUD Completo**: Para todas las 15 entidades
- **🔐 Autenticación JWT**: Sistema completo de autenticación
- **🔒 Rutas Protegidas**: Middleware automático de validación
- **Relaciones Complejas**: Preloading automático de relaciones
- **Filtros Avanzados**: Búsquedas por múltiples criterios
- **Validación**: Validación de datos de entrada
- **Manejo de Errores**: Respuestas de error estructuradas
- **CORS**: Configurado para desarrollo y producción
- **🎯 JSON Automático**: No requiere header `Content-Type: application/json`

### 🚀 Facilidad de Uso
- **Sin Headers Requeridos**: Envía JSON directamente sin especificar Content-Type
- **Detección Automática**: El servidor detecta automáticamente contenido JSON
- **Compatible con Postman**: Funciona perfectamente sin configuración adicional
- **🔐 Autenticación Transparente**: Solo incluye el token Bearer en el header

## 📊 Funcionalidades del Sistema

### 🔐 **Sistema de Autenticación JWT**
- **Login/Logout**: Autenticación segura con tokens
- **Registro de Usuarios**: Creación de nuevas cuentas
- **Gestión de Sesiones**: Tokens con expiración automática
- **Cambio de Contraseñas**: Actualización segura de credenciales
- **Renovación de Tokens**: Extensión de sesiones activas
- **Validación de Tokens**: Verificación de autenticidad
- **Protección de Rutas**: Middleware automático de seguridad

### 🎯 Gestión de Visitas
- Programación de visitas educativas
- Asignación de actividades y temáticas
- Registro de participantes
- Estadísticas de participación

### 👥 Gestión de Usuarios
- Sistema de autenticación
- Tipos de usuarios diferenciados
- Gestión de perfiles
- **🗑️ Soft Delete**: Eliminación segura con posibilidad de restauración
- **♻️ Restauración**: Recuperación de usuarios eliminados accidentalmente
- **📋 Auditoría**: Visualización de usuarios eliminados para control administrativo

### 🆕 Gestión de Estudiantes con Eliminación en Cascada
- **🔗 Eliminación en Cascada**: Al eliminar un estudiante, automáticamente se eliminan:
  - El registro del estudiante
  - Todos los usuarios asociados a la persona del estudiante
  - El registro de la persona del estudiante
- **♻️ Restauración en Cascada**: Al restaurar un estudiante, automáticamente se restauran:
  - El registro del estudiante
  - Todos los usuarios asociados a la persona
  - El registro de la persona
- **🔒 Transacciones**: Todas las operaciones usan transacciones para garantizar integridad
- **🛡️ Rollback Automático**: Si alguna operación falla, se deshacen todos los cambios
- **💾 Soft Delete**: Los datos no se eliminan físicamente, solo se marcan como eliminados
- **🔍 Recuperación Completa**: La restauración recupera todos los datos relacionados

### 🆕 Gestión de Autoridades UTEQ con Eliminación en Cascada
- **🔗 Eliminación en Cascada**: Al eliminar una autoridad UTEQ, automáticamente se eliminan:
  - El registro de la autoridad UTEQ
  - Todos los usuarios asociados a la persona de la autoridad
  - El registro de la persona de la autoridad
- **♻️ Restauración en Cascada**: Al restaurar una autoridad UTEQ, automáticamente se restauran:
  - El registro de la autoridad UTEQ
  - Todos los usuarios asociados a la persona
  - El registro de la persona
- **🔒 Transacciones**: Todas las operaciones usan transacciones para garantizar integridad
- **🛡️ Rollback Automático**: Si alguna operación falla, se deshacen todos los cambios
- **💾 Soft Delete**: Los datos no se eliminan físicamente, solo se marcan como eliminados
- **🔍 Recuperación Completa**: La restauración recupera todos los datos relacionados
- **📋 Auditoría**: Visualización de autoridades eliminadas para control administrativo
- **🎯 Filtros Avanzados**: Búsqueda por cargo y persona

### ❓ Sistema de Dudas
- Creación de dudas por estudiantes
- Asignación a autoridades UTEQ
- Seguimiento de respuestas
- Estados de dudas (pendientes, respondidas, sin asignar)

### 📈 Reportes y Estadísticas
- Estadísticas de participación en visitas
- Filtros por fechas, instituciones, actividades
- Búsquedas avanzadas

### 🌍 Gestión Geográfica
- Provincias y ciudades
- Filtros por ubicación

## 🔒 Seguridad Implementada

### JWT Token Security
- **Algoritmo**: HS256 (HMAC SHA-256)
- **Expiración**: 24 horas
- **Claims incluidos**:
  - `user_id`: ID del usuario
  - `username`: Nombre de usuario
  - `tipo_usuario_id`: Tipo de usuario
  - `exp`: Tiempo de expiración
  - `iat`: Tiempo de emisión
  - `iss`: Emisor (ApiEscuela)

### Password Security
- **Encriptación**: bcrypt con salt automático
- **Longitud mínima**: 6 caracteres
- **Verificación**: Comparación segura con hash almacenado

### Route Protection
- **Middleware automático**: Validación en todas las rutas `/api/*`
- **Context injection**: Información del usuario disponible en handlers
- **Error handling**: Respuestas claras para tokens inválidos

## ⚠️ Notas Importantes de Migración

### 🚨 **CAMBIOS IMPORTANTES**

1. **Todas las rutas existentes ahora requieren autenticación** excepto:
   - `GET /`
   - `GET /health`
   - `POST /auth/login`
   - `POST /auth/register`
   - `POST /auth/validate-token`

2. **Las URLs han cambiado**: Todos los endpoints protegidos ahora están bajo `/api/`
   ```
   Antes: GET /estudiantes
   Ahora:  GET /api/estudiantes (con Authorization header)
   ```

3. **Header de autorización requerido**: Todas las peticiones a `/api/*` deben incluir:
   ```
   Authorization: Bearer tu_token_jwt_aqui
   ```

4. **Flujo de autenticación obligatorio**:
   - Primero hacer login para obtener token
   - Incluir token en todas las peticiones subsecuentes
   - Renovar token antes de que expire (24 horas)

### 🔧 **Configuración Recomendada para Producción**

```env
JWT_SECRET=tu_clave_secreta_super_larga_y_compleja_aqui
JWT_EXPIRATION=24h
APP_ENV=production
```

## 🚀 Estado del Proyecto

✅ **Sistema Completo y Funcional**
- ✅ 15 entidades implementadas
- ✅ 80+ endpoints API
- ✅ **Sistema de autenticación JWT completo**
- ✅ **Middleware de seguridad automático**
- ✅ **Gestión de contraseñas encriptadas**
- ✅ Filtros y búsquedas avanzadas
- ✅ Estadísticas integradas
- ✅ Documentación completa
- ✅ **Documentación de autenticación detallada**

## 📚 Documentación Adicional

- **`AUTH_README.md`**: Documentación detallada del sistema de autenticación
- **Ejemplos de código**: JavaScript/Fetch, cURL, Postman
- **Códigos de error**: Documentación completa de respuestas de error
- **Guías de migración**: Cómo actualizar código existente

## 📞 Soporte

Para soporte técnico o consultas sobre el sistema, contactar al equipo de desarrollo de la UTEQ.

### 🔐 Soporte de Autenticación

Si tienes problemas con la autenticación:
1. Verifica que estés usando las nuevas URLs con `/api/`
2. Confirma que incluyes el header `Authorization: Bearer TOKEN`
3. Verifica que el token no haya expirado (24 horas)
4. Usa el endpoint `/auth/refresh-token` para renovar tokens

---

**Desarrollado para la Universidad Técnica Estatal de Quevedo (UTEQ)**

**🔐 Versión 2.0 - Con Sistema de Autenticación JWT Completo**