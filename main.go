package main

import (
	"ApiEscuela/handlers"
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"ApiEscuela/routers"
	"ApiEscuela/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Advertencia: No se pudo cargar el archivo .env: %v", err)
	}

	// Inicializar Fiber
	app := fiber.New(fiber.Config{
		AppName: "ApiEscuela v1.0",
		// Configurar para aceptar JSON automáticamente
		BodyLimit: 4 * 1024 * 1024, // 4MB
	})

	// Middleware para detectar JSON automáticamente
	app.Use(func(c *fiber.Ctx) error {
		// Si el body parece JSON pero no tiene Content-Type, lo establecemos
		if len(c.Body()) > 0 {
			body := c.Body()
			// Verificar si el body comienza con { o [ (JSON)
			if (body[0] == '{' || body[0] == '[') && c.Get("Content-Type") == "" {
				c.Request().Header.Set("Content-Type", "application/json")
			}
		}
		return c.Next()
	})

	// Configurar CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Configurar Viper
	config := viper.New()
	config.AutomaticEnv()
	config.SetDefault("APP_PORT", "3000")
	config.SetDefault("APP_ENV", "development")

	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath(".")
	config.AddConfigPath("/etc/secrets/")

	if err := config.ReadInConfig(); err != nil {
		log.Printf("Advertencia: No se pudo leer el archivo de configuración. %v", err)
	}

	// Configurar conexión a la base de datos
	var dsn string

	// Intentar usar DATABASE_URL primero (para producción/Neon)
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		dsn = databaseURL
		log.Printf("Usando DATABASE_URL (Neon) para conexión a la base de datos")
	} else {
		// Fallback a la configuración Neon por defecto
		// Configuración anterior (UTEQ) comentada:
		// dsn = "host=aplicaciones.uteq.edu.ec user=aplicaciones password=z8E9bYdQpHmOvtfH6Up5dE1HKCh35pgwlEDuZqMklOtg3Zm2UA dbname=bdrealidaduteq port=9010 sslmode=require"
		dsn = "postgresql://neondb_owner:npg_UkVLzt5h0Zyx@ep-noisy-heart-aehgd4hl-pooler.c-2.us-east-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
		log.Printf("Usando configuración Neon por defecto para conexión a la base de datos")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	// Automigración de todos los modelos
	if err := db.AutoMigrate(
		&models.Provincia{},
		&models.Ciudad{},
		&models.Persona{},
		&models.TipoUsuario{},
		&models.Usuario{},
		&models.Institucion{},
		&models.Estudiante{},
		&models.EstudianteUniversitario{},
		&models.AutoridadUTEQ{},
		&models.Tematica{},
		&models.Actividad{},
		&models.ProgramaVisita{},
		&models.DetalleAutoridadDetallesVisita{},
		&models.VisitaDetalle{},
		&models.Dudas{},
		&models.VisitaDetalleEstudiantesUniversitarios{},
		&models.CodigoUsuario{},
	); err != nil {
		log.Fatalf("Error en la automigración: %v", err)
	}

	// Inicializar repositorios
	estudianteRepo := repositories.NewEstudianteRepository(db)
	personaRepo := repositories.NewPersonaRepository(db)
	provinciaRepo := repositories.NewProvinciaRepository(db)
	ciudadRepo := repositories.NewCiudadRepository(db)
	institucionRepo := repositories.NewInstitucionRepository(db)
	tipoUsuarioRepo := repositories.NewTipoUsuarioRepository(db)
	usuarioRepo := repositories.NewUsuarioRepository(db)
	estudianteUnivRepo := repositories.NewEstudianteUniversitarioRepository(db)
	autoridadRepo := repositories.NewAutoridadUTEQRepository(db)
	tematicaRepo := repositories.NewTematicaRepository(db)
	actividadRepo := repositories.NewActividadRepository(db)
	programaVisitaRepo := repositories.NewProgramaVisitaRepository(db)
	detalleAutoridadDetallesVisitaRepo := repositories.NewDetalleAutoridadDetallesVisitaRepository(db)
	visitaDetalleRepo := repositories.NewVisitaDetalleRepository(db)
	dudasRepo := repositories.NewDudasRepository(db)
	visitaDetalleEstudiantesUniversitariosRepo := repositories.NewVisitaDetalleEstudiantesUniversitariosRepository(db)
	codigoUsuarioRepo := repositories.NewCodigoUsuarioRepository(db)

	// Inicializar handlers
	estudianteHandler := handlers.NewEstudianteHandler(estudianteRepo)
	personaHandler := handlers.NewPersonaHandler(personaRepo)
	provinciaHandler := handlers.NewProvinciaHandler(provinciaRepo)
	ciudadHandler := handlers.NewCiudadHandler(ciudadRepo)
	institucionHandler := handlers.NewInstitucionHandler(institucionRepo)
	tipoUsuarioHandler := handlers.NewTipoUsuarioHandler(tipoUsuarioRepo)
	usuarioHandler := handlers.NewUsuarioHandler(usuarioRepo)
	estudianteUnivHandler := handlers.NewEstudianteUniversitarioHandler(estudianteUnivRepo)
	autoridadHandler := handlers.NewAutoridadUTEQHandler(autoridadRepo)
	tematicaHandler := handlers.NewTematicaHandler(tematicaRepo)
	actividadHandler := handlers.NewActividadHandler(actividadRepo)
	programaVisitaHandler := handlers.NewProgramaVisitaHandler(programaVisitaRepo)
	detalleAutoridadDetallesVisitaHandler := handlers.NewDetalleAutoridadDetallesVisitaHandler(detalleAutoridadDetallesVisitaRepo)
	visitaDetalleHandler := handlers.NewVisitaDetalleHandler(visitaDetalleRepo)
	dudasHandler := handlers.NewDudasHandler(dudasRepo)
	visitaDetalleEstudiantesUniversitariosHandler := handlers.NewVisitaDetalleEstudiantesUniversitariosHandler(visitaDetalleEstudiantesUniversitariosRepo)

	// Inicializar servicios
	authService := services.NewAuthService(usuarioRepo, personaRepo, codigoUsuarioRepo)

	// Inicializar handler de autenticación
	authHandler := handlers.NewAuthHandler(authService)

	// Crear contenedor de todos los handlers
	allHandlers := routers.NewAllHandlers(
		estudianteHandler,
		personaHandler,
		provinciaHandler,
		ciudadHandler,
		institucionHandler,
		tipoUsuarioHandler,
		usuarioHandler,
		estudianteUnivHandler,
		autoridadHandler,
		tematicaHandler,
		actividadHandler,
		programaVisitaHandler,
		detalleAutoridadDetallesVisitaHandler,
		visitaDetalleHandler,
		dudasHandler,
		visitaDetalleEstudiantesUniversitariosHandler,
		authHandler,
	)

	// Configurar todas las rutas
	routers.SetupAllRoutes(app, allHandlers)

	// Ruta de bienvenida
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "¡Bienvenido a ApiEscuela!",
			"version": "1.0",
			"status":  "running",
		})
	})

	// Ruta de salud
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":   "healthy",
			"database": "connected",
		})
	})

	// Iniciar servidor
	port := config.GetString("APP_PORT")
	log.Printf("Servidor ApiEscuela iniciado en el puerto %s", port)
	log.Printf("Ambiente: %s", config.GetString("APP_ENV"))
	log.Printf("Conectado a la base de datos exitosamente")

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
