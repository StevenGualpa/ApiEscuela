# ApiEscuela - Sistema de Gestión de Visitas Educativas UTEQ

Un backend completo desarrollado en Go para la gestión integral de visitas educativas a la UTEQ, incluyendo estudiantes, instituciones, programas de visita, actividades y sistema de dudas.

## 🏗️ Estructura del Proyecto

```
ApiEscuela/
├── models/          # 15 modelos de datos (entidades del sistema)
├── repositories/    # Repositorios para acceso a datos
├── handlers/        # Controladores HTTP para todas las entidades
├── routers/         # Configuración consolidada de rutas
├── main.go         # Punto de entrada de la aplicación
├── config.env      # Variables de entorno
└── README.md       # Este archivo
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

## 🚀 API Endpoints (70+ endpoints)

Base URL: `http://localhost:3000`

### 📚 Estudiantes
- `POST /estudiantes` - Crear estudiante
- `GET /estudiantes` - Obtener todos los estudiantes activos
- `GET /estudiantes/all-including-deleted` - **📋 Obtener todos los estudiantes (activos + eliminados)** (NUEVO)
- `GET /estudiantes/deleted` - **🗑️ Obtener solo estudiantes eliminados** (NUEVO)
- `GET /estudiantes/:id` - Obtener estudiante por ID
- `PUT /estudiantes/:id` - Actualizar estudiante
- `DELETE /estudiantes/:id` - **🗑️ Eliminar estudiante, usuario y persona en cascada**
- `PUT /estudiantes/:id/restore` - **♻️ Restaurar estudiante, usuario y persona en cascada** (NUEVO)
- `GET /estudiantes/ciudad/:ciudad_id` - Filtrar por ciudad
- `GET /estudiantes/institucion/:institucion_id` - Filtrar por institución
- `GET /estudiantes/especialidad/:especialidad` - Filtrar por especialidad

### 👤 Personas
- `POST /personas` - Crear persona
- `GET /personas` - Obtener todas las personas
- `GET /personas/:id` - Obtener persona por ID
- `PUT /personas/:id` - Actualizar persona
- `DELETE /personas/:id` - Eliminar persona
- `GET /personas/cedula/:cedula` - Buscar por cédula
- `GET /personas/correo/:correo` - Buscar por correo

### 🔐 Sistema de Usuarios y Autenticación
- `POST /usuarios` - Crear usuario
- `GET /usuarios` - Obtener todos los usuarios activos
- `GET /usuarios/all-including-deleted` - **📋 Obtener todos los usuarios (activos + eliminados)**
- `GET /usuarios/deleted` - **🗑️ Obtener solo usuarios eliminados**
- `GET /usuarios/:id` - Obtener usuario por ID
- `PUT /usuarios/:id` - Actualizar usuario
- `DELETE /usuarios/:id` - Eliminar usuario (soft delete)
- `PUT /usuarios/:id/restore` - **♻️ Restaurar usuario eliminado**
- `GET /usuarios/username/:username` - Buscar por nombre de usuario
- `GET /usuarios/tipo/:tipo_usuario_id` - Filtrar por tipo
- `GET /usuarios/persona/:persona_id` - Filtrar por persona
- `POST /usuarios/login` - **🔑 Autenticación de usuarios**

### 📅 Programas de Visita
- `POST /programas-visita` - Crear programa
- `GET /programas-visita` - Obtener todos los programas
- `GET /programas-visita/:id` - Obtener programa por ID
- `PUT /programas-visita/:id` - Actualizar programa
- `DELETE /programas-visita/:id` - Eliminar programa
- `GET /programas-visita/fecha/:fecha` - **Filtrar por fecha (YYYY-MM-DD)**
- `GET /programas-visita/autoridad/:autoridad_id` - Filtrar por autoridad
- `GET /programas-visita/institucion/:institucion_id` - Filtrar por institución
- `GET /programas-visita/rango-fecha?inicio=2024-01-01&fin=2024-12-31` - **Rango de fechas**

### 🔗 Detalle Autoridad Detalles Visita (🆕 Relación Muchos-a-Muchos)
- `POST /detalle-autoridad-detalles-visita` - **Asignar autoridad a programa**
- `GET /detalle-autoridad-detalles-visita` - Obtener todas las asignaciones
- `GET /detalle-autoridad-detalles-visita/:id` - Obtener asignación por ID
- `PUT /detalle-autoridad-detalles-visita/:id` - Actualizar asignación
- `DELETE /detalle-autoridad-detalles-visita/:id` - Eliminar asignación
- `GET /detalle-autoridad-detalles-visita/programa-visita/:programa_visita_id` - **Autoridades por programa**
- `GET /detalle-autoridad-detalles-visita/autoridad/:autoridad_id` - **Programas por autoridad**

### 🎯 Actividades y Temáticas
- `POST /actividades` - Crear actividad
- `GET /actividades` - Obtener todas las actividades
- `GET /actividades/:id` - Obtener actividad por ID
- `PUT /actividades/:id` - Actualizar actividad
- `DELETE /actividades/:id` - Eliminar actividad
- `GET /actividades/tematica/:tematica_id` - Filtrar por temática
- `GET /actividades/nombre/:nombre` - Buscar por nombre
- `GET /actividades/duracion?min=30&max=120` - **Filtrar por duración**

### 📋 Visita Detalles y Estadísticas
- `POST /visita-detalles` - Crear detalle
- `GET /visita-detalles` - Obtener todos los detalles
- `GET /visita-detalles/:id` - Obtener detalle por ID
- `PUT /visita-detalles/:id` - Actualizar detalle
- `DELETE /visita-detalles/:id` - Eliminar detalle
- `GET /visita-detalles/estudiante/:estudiante_id` - Filtrar por estudiante
- `GET /visita-detalles/actividad/:actividad_id` - Filtrar por actividad
- `GET /visita-detalles/programa/:programa_id` - Filtrar por programa
- `GET /visita-detalles/participantes?min=10&max=50` - **Filtrar por participantes**
- `GET /visita-detalles/estadisticas` - **📊 Estadísticas de participación**

### ❓ Sistema de Dudas
- `POST /dudas` - Crear duda
- `GET /dudas` - Obtener todas las dudas
- `GET /dudas/:id` - Obtener duda por ID
- `PUT /dudas/:id` - Actualizar duda
- `DELETE /dudas/:id` - Eliminar duda
- `GET /dudas/estudiante/:estudiante_id` - Filtrar por estudiante
- `GET /dudas/autoridad/:autoridad_id` - Filtrar por autoridad
- `GET /dudas/sin-responder` - **📋 Dudas pendientes**
- `GET /dudas/respondidas` - **✅ Dudas respondidas**
- `GET /dudas/sin-asignar` - **⚠️ Dudas sin asignar**
- `GET /dudas/buscar/:termino` - **🔍 Búsqueda en preguntas**
- `PUT /dudas/:duda_id/asignar` - **👤 Asignar autoridad**
- `PUT /dudas/:duda_id/responder` - **💬 Responder duda**

### 🌍 Ubicaciones Geográficas
- `GET /provincias` - Obtener todas las provincias
- `GET /ciudades` - Obtener todas las ciudades
- `GET /ciudades/provincia/:provincia_id` - Ciudades por provincia

### 🏫 Instituciones y Autoridades
- `GET /instituciones` - Obtener todas las instituciones
- `GET /instituciones/nombre/:nombre` - Buscar por nombre
- `GET /autoridades-uteq` - Obtener autoridades UTEQ
- `GET /autoridades-uteq/cargo/:cargo` - Filtrar por cargo

## 📋 Estructuras JSON de los Modelos

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

### Duda
```json
{
  "pregunta": "¿Cuáles son los requisitos de ingreso?",
  "estudiante_id": 1
}
```

**Campos Opcionales**:
- `fecha_pregunta`: Se establece automáticamente al crear la duda si no se proporciona
- `respuesta`: Campo opcional, se llena cuando se responde la duda
- `fecha_respuesta`: Se establece automáticamente al responder la duda
- `autoridad_uteq_id`: Campo opcional, se asigna cuando una autoridad toma la duda

**Ejemplo de Duda Completa** (después de ser asignada y respondida):
```json
{
  "id": 1,
  "pregunta": "¿Cuáles son los requisitos de ingreso?",
  "fecha_pregunta": "2024-03-15T10:30:00Z",
  "respuesta": "Los requisitos incluyen bachillerato completo...",
  "fecha_respuesta": "2024-03-15T14:45:00Z",
  "estudiante_id": 1,
  "autoridad_uteq_id": 2
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
   ```

4. **Ejecutar la aplicación**
   ```bash
   go run main.go
   ```

La aplicación estará disponible en `http://localhost:3000`

## 📝 Ejemplos de Uso

### Autenticación
```bash
curl -X POST http://localhost:3000/usuarios/login \
  -d '{
    "usuario": "admin",
    "contraseña": "password123"
  }'
```

### Crear una Persona
```bash
curl -X POST http://localhost:3000/personas \
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
curl -X POST http://localhost:3000/estudiantes \
  -d '{
    "persona_id": 1,
    "institucion_id": 1,
    "ciudad_id": 1,
    "especialidad": "Ingeniería en Sistemas"
  }'
```

### Crear un Programa de Visita
```bash
curl -X POST http://localhost:3000/programas-visita \
  -d '{
    "fecha": "2024-03-15T09:00:00Z",
    "institucion_id": 1
  }'
```

### 🆕 Asignar Autoridad a Programa de Visita
```bash
curl -X POST http://localhost:3000/detalle-autoridad-detalles-visita \
  -d '{
    "programa_visita_id": 1,
    "autoridad_uteq_id": 2
  }'
```

### Obtener Autoridades de un Programa
```bash
curl http://localhost:3000/detalle-autoridad-detalles-visita/programa-visita/1
```

### Obtener Programas de una Autoridad
```bash
curl http://localhost:3000/detalle-autoridad-detalles-visita/autoridad/2
```

### Crear un Usuario
```bash
curl -X POST http://localhost:3000/usuarios \
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
curl http://localhost:3000/usuarios/all-including-deleted
```

### Obtener Solo Usuarios Eliminados
```bash
curl http://localhost:3000/usuarios/deleted
```

### Restaurar Usuario Eliminado
```bash
curl -X PUT http://localhost:3000/usuarios/5/restore
```

### Ejemplo de Flujo Completo de Soft Delete
```bash
# 1. Eliminar usuario (soft delete)
curl -X DELETE http://localhost:3000/usuarios/5

# 2. Verificar que aparece en usuarios eliminados
curl http://localhost:3000/usuarios/deleted

# 3. Restaurar el usuario
curl -X PUT http://localhost:3000/usuarios/5/restore

# 4. Verificar que el usuario está activo nuevamente
curl http://localhost:3000/usuarios/5
```

### 🆕 Gestión de Estudiantes con Eliminación en Cascada

### Eliminar Estudiante (Cascada: Estudiante + Usuario + Persona)
```bash
# Elimina el estudiante y automáticamente elimina su usuario y persona asociada
curl -X DELETE http://localhost:3000/estudiantes/3
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
curl -X PUT http://localhost:3000/estudiantes/3/restore
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
curl -X POST http://localhost:3000/personas \
  -d '{
    "nombre": "María González",
    "cedula": "0987654321",
    "correo": "maria.gonzalez@email.com",
    "telefono": "0987654321",
    "fecha_nacimiento": "1995-08-20T00:00:00Z"
  }'

# 2. Crear un usuario para esa persona
curl -X POST http://localhost:3000/usuarios \
  -d '{
    "usuario": "mgonzalez",
    "contraseña": "password123",
    "persona_id": 2,
    "tipo_usuario_id": 1
  }'

# 3. Crear un estudiante para esa persona
curl -X POST http://localhost:3000/estudiantes \
  -d '{
    "persona_id": 2,
    "institucion_id": 1,
    "ciudad_id": 1,
    "especialidad": "Ingeniería Ambiental"
  }'

# 4. Eliminar el estudiante (elimina automáticamente usuario y persona)
curl -X DELETE http://localhost:3000/estudiantes/2

# 5. Verificar que el usuario también fue eliminado
curl http://localhost:3000/usuarios/deleted

# 6. Restaurar el estudiante (restaura automáticamente usuario y persona)
curl -X PUT http://localhost:3000/estudiantes/2/restore

# 7. Verificar que todo fue restaurado correctamente
curl http://localhost:3000/estudiantes/2
curl http://localhost:3000/usuarios/2
curl http://localhost:3000/personas/2
```

### 🆕 Gestión Completa de Estudiantes Eliminados (Similar a Usuarios)

### Obtener Todos los Estudiantes (Incluyendo Eliminados)
```bash
curl http://localhost:3000/estudiantes/all-including-deleted
```

### Obtener Solo Estudiantes Eliminados
```bash
curl http://localhost:3000/estudiantes/deleted
```

### Ejemplo de Flujo Completo de Gestión de Estudiantes Eliminados
```bash
# 1. Obtener todos los estudiantes activos
curl http://localhost:3000/estudiantes

# 2. Eliminar un estudiante (cascada: estudiante + usuario + persona)
curl -X DELETE http://localhost:3000/estudiantes/3

# 3. Verificar que ya no aparece en estudiantes activos
curl http://localhost:3000/estudiantes

# 4. Verificar que aparece en estudiantes eliminados
curl http://localhost:3000/estudiantes/deleted

# 5. Obtener todos los estudiantes incluyendo eliminados
curl http://localhost:3000/estudiantes/all-including-deleted

# 6. Restaurar el estudiante eliminado
curl -X PUT http://localhost:3000/estudiantes/3/restore

# 7. Verificar que vuelve a aparecer en estudiantes activos
curl http://localhost:3000/estudiantes

# 8. Verificar que ya no aparece en estudiantes eliminados
curl http://localhost:3000/estudiantes/deleted
```

### Crear una Provincia
```bash
curl -X POST http://localhost:3000/provincias \
  -d '{
    "provincia": "Los Ríos"
  }'
```

### Crear una Ciudad
```bash
curl -X POST http://localhost:3000/ciudades \
  -d '{
    "ciudad": "Quevedo",
    "provincia_id": 1
  }'
```

### Crear una Institución
```bash
curl -X POST http://localhost:3000/instituciones \
  -d '{
    "nombre": "Unidad Educativa San José",
    "autoridad": "Dr. María González",
    "contacto": "0987654321",
    "direccion": "Av. Principal 123, Quevedo"
  }'
```

### Crear una Temática
```bash
curl -X POST http://localhost:3000/tematicas \
  -d '{
    "nombre": "Ingeniería Agrícola",
    "descripcion": "Temática sobre técnicas modernas de agricultura"
  }'
```

### Crear una Actividad
```bash
curl -X POST http://localhost:3000/actividades \
  -d '{
    "actividad": "Visita a Laboratorio de Suelos",
    "tematica_id": 1,
    "duracion": 90
  }'
```

### Crear una Duda
```bash
curl -X POST http://localhost:3000/dudas \
  -d '{
    "pregunta": "¿Cuáles son los requisitos de ingreso?",
    "estudiante_id": 1
  }'
```

### Responder una Duda
```bash
curl -X PUT http://localhost:3000/dudas/1/responder \
  -d '{
    "respuesta": "Los requisitos incluyen bachillerato completo y aprobar el examen de admisión."
  }'
```

### Obtener Estadísticas de Participación
```bash
curl http://localhost:3000/visita-detalles/estadisticas
```

### Buscar Dudas Pendientes
```bash
curl http://localhost:3000/dudas/sin-responder
```

### Filtrar Actividades por Duración
```bash
curl "http://localhost:3000/actividades/duracion?min=30&max=120"
```

### Obtener Programas por Rango de Fechas
```bash
curl "http://localhost:3000/programas-visita/rango-fecha?inicio=2024-01-01&fin=2024-12-31"
```

## 🛠️ Tecnologías Utilizadas

- **Go 1.24**: Lenguaje de programación
- **Fiber v2**: Framework web HTTP de alto rendimiento
- **GORM**: ORM para Go con soporte completo para PostgreSQL
- **PostgreSQL**: Base de datos relacional
- **Viper**: Gestión de configuración y variables de entorno

## 🏛️ Arquitectura del Sistema

### Patrón de Capas
- **Models**: Definición de entidades y relaciones
- **Repositories**: Capa de acceso a datos con GORM
- **Handlers**: Controladores HTTP con validación
- **Routers**: Configuración consolidada de rutas

### Características Técnicas
- **CRUD Completo**: Para todas las 15 entidades
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

## 📊 Funcionalidades del Sistema

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

## 🚀 Estado del Proyecto

✅ **Sistema Completo y Funcional**
- entidades implementadas
- endpoints API
- Sistema de autenticación
- Filtros y búsquedas avanzadas
- Estadísticas integradas
- Documentación completa

## 📞 Soporte

Para soporte técnico o consultas sobre el sistema, contactar al equipo de desarrollo de la UTEQ.

---

**Desarrollado para la Universidad Técnica Estatal de Quevedo (UTEQ)**