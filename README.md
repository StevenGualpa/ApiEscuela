# ApiEscuela - Sistema de Gestión de Estudiantes

Un backend simple desarrollado en Go para la gestión de estudiantes, conectado a la base de datos de la UTEQ.

## Estructura del Proyecto

```
ApiEscuela/
├── models/          # Modelo Student
├── repositories/    # Repositorio para acceso a datos
├── handlers/        # Controlador HTTP para estudiantes
├── routers/         # Configuración de rutas
├── main.go         # Punto de entrada de la aplicación
├── config.env      # Variables de entorno
└── README.md       # Este archivo
```

## Modelo Student

```go
type Student struct {
    ID        uint   `json:"id"`
    Cedula    string `json:"cedula"`    // Único
    Nombres   string `json:"nombres"`
    Apellidos string `json:"apellidos"`
    Ciudad    string `json:"ciudad"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## Endpoints API

### Estudiantes
- `POST /api/v1/students` - Crear estudiante
- `GET /api/v1/students` - Obtener todos los estudiantes
- `GET /api/v1/students/:id` - Obtener estudiante por ID
- `PUT /api/v1/students/:id` - Actualizar estudiante
- `DELETE /api/v1/students/:id` - Eliminar estudiante
- `GET /api/v1/students/cedula/:cedula` - Obtener estudiante por cédula
- `GET /api/v1/students/ciudad/:ciudad` - Obtener estudiantes por ciudad

## Configuración de Base de Datos

El proyecto está configurado para conectarse a la base de datos de la UTEQ:
- **Host**: aplicaciones.uteq.edu.ec
- **Puerto**: 9010
- **Usuario**: aplicaciones
- **Base de datos**: bdrealidaduteq
- **SSL**: requerido

## Instalación y Ejecución

### Prerrequisitos
- Go 1.24+

### Pasos

1. **Instalar dependencias**
   ```bash
   go mod tidy
   ```

2. **Ejecutar la aplicación**
   ```bash
   go run main.go
   ```

La aplicación estará disponible en `http://localhost:3000`

## Ejemplos de Uso

### Crear un estudiante
```bash
curl -X POST http://localhost:3000/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{
    "cedula": "1234567890",
    "nombres": "Juan Carlos",
    "apellidos": "Pérez González",
    "ciudad": "Quevedo"
  }'
```

### Obtener todos los estudiantes
```bash
curl http://localhost:3000/api/v1/students
```

### Buscar por cédula
```bash
curl http://localhost:3000/api/v1/students/cedula/1234567890
```

### Buscar por ciudad
```bash
curl http://localhost:3000/api/v1/students/ciudad/Quevedo
```

## Tecnologías Utilizadas

- **Go 1.24**: Lenguaje de programación
- **Fiber v2**: Framework web HTTP
- **GORM**: ORM para Go
- **PostgreSQL**: Base de datos
- **Viper**: Gestión de configuración