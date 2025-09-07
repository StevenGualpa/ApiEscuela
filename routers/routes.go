package routers

import (
	"ApiEscuela/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupAllRoutes configura todas las rutas de la aplicación
func SetupAllRoutes(app *fiber.App, handlers *AllHandlers) {
	// ==================== ESTUDIANTES ====================
	estudiantes := app.Group("/estudiantes")
	estudiantes.Post("/", handlers.EstudianteHandler.CreateEstudiante)
	estudiantes.Get("/", handlers.EstudianteHandler.GetAllEstudiantes)
	estudiantes.Get("/:id", handlers.EstudianteHandler.GetEstudiante)
	estudiantes.Put("/:id", handlers.EstudianteHandler.UpdateEstudiante)
	estudiantes.Delete("/:id", handlers.EstudianteHandler.DeleteEstudiante)
	estudiantes.Get("/ciudad/:ciudad_id", handlers.EstudianteHandler.GetEstudiantesByCity)
	estudiantes.Get("/institucion/:institucion_id", handlers.EstudianteHandler.GetEstudiantesByInstitucion)
	estudiantes.Get("/especialidad/:especialidad", handlers.EstudianteHandler.GetEstudiantesByEspecialidad)

	// ==================== PERSONAS ====================
	personas := app.Group("/personas")
	personas.Post("/", handlers.PersonaHandler.CreatePersona)
	personas.Get("/", handlers.PersonaHandler.GetAllPersonas)
	personas.Get("/:id", handlers.PersonaHandler.GetPersona)
	personas.Put("/:id", handlers.PersonaHandler.UpdatePersona)
	personas.Delete("/:id", handlers.PersonaHandler.DeletePersona)
	personas.Get("/cedula/:cedula", handlers.PersonaHandler.GetPersonaByCedula)
	personas.Get("/correo/:correo", handlers.PersonaHandler.GetPersonasByCorreo)

	// ==================== PROVINCIAS ====================
	provincias := app.Group("/provincias")
	provincias.Post("/", handlers.ProvinciaHandler.CreateProvincia)
	provincias.Get("/", handlers.ProvinciaHandler.GetAllProvincias)
	provincias.Get("/:id", handlers.ProvinciaHandler.GetProvincia)
	provincias.Put("/:id", handlers.ProvinciaHandler.UpdateProvincia)
	provincias.Delete("/:id", handlers.ProvinciaHandler.DeleteProvincia)
	provincias.Get("/nombre/:nombre", handlers.ProvinciaHandler.GetProvinciaByNombre)

	// ==================== CIUDADES ====================
	ciudades := app.Group("/ciudades")
	ciudades.Post("/", handlers.CiudadHandler.CreateCiudad)
	ciudades.Get("/", handlers.CiudadHandler.GetAllCiudades)
	ciudades.Get("/:id", handlers.CiudadHandler.GetCiudad)
	ciudades.Put("/:id", handlers.CiudadHandler.UpdateCiudad)
	ciudades.Delete("/:id", handlers.CiudadHandler.DeleteCiudad)
	ciudades.Get("/provincia/:provincia_id", handlers.CiudadHandler.GetCiudadesByProvincia)
	ciudades.Get("/nombre/:nombre", handlers.CiudadHandler.GetCiudadByNombre)

	// ==================== INSTITUCIONES ====================
	instituciones := app.Group("/instituciones")
	instituciones.Post("/", handlers.InstitucionHandler.CreateInstitucion)
	instituciones.Get("/", handlers.InstitucionHandler.GetAllInstituciones)
	instituciones.Get("/:id", handlers.InstitucionHandler.GetInstitucion)
	instituciones.Put("/:id", handlers.InstitucionHandler.UpdateInstitucion)
	instituciones.Delete("/:id", handlers.InstitucionHandler.DeleteInstitucion)
	instituciones.Get("/nombre/:nombre", handlers.InstitucionHandler.GetInstitucionesByNombre)
	instituciones.Get("/autoridad/:autoridad", handlers.InstitucionHandler.GetInstitucionesByAutoridad)

	// ==================== TIPOS DE USUARIO ====================
	tiposUsuario := app.Group("/tipos-usuario")
	tiposUsuario.Post("/", handlers.TipoUsuarioHandler.CreateTipoUsuario)
	tiposUsuario.Get("/", handlers.TipoUsuarioHandler.GetAllTiposUsuario)
	tiposUsuario.Get("/:id", handlers.TipoUsuarioHandler.GetTipoUsuario)
	tiposUsuario.Put("/:id", handlers.TipoUsuarioHandler.UpdateTipoUsuario)
	tiposUsuario.Delete("/:id", handlers.TipoUsuarioHandler.DeleteTipoUsuario)
	tiposUsuario.Get("/nombre/:nombre", handlers.TipoUsuarioHandler.GetTipoUsuarioByNombre)

	// ==================== USUARIOS ====================
	usuarios := app.Group("/usuarios")
	usuarios.Post("/", handlers.UsuarioHandler.CreateUsuario)
	usuarios.Get("/", handlers.UsuarioHandler.GetAllUsuarios)
	usuarios.Get("/:id", handlers.UsuarioHandler.GetUsuario)
	usuarios.Put("/:id", handlers.UsuarioHandler.UpdateUsuario)
	usuarios.Delete("/:id", handlers.UsuarioHandler.DeleteUsuario)
	usuarios.Get("/username/:username", handlers.UsuarioHandler.GetUsuarioByUsername)
	usuarios.Get("/tipo/:tipo_usuario_id", handlers.UsuarioHandler.GetUsuariosByTipo)
	usuarios.Get("/persona/:persona_id", handlers.UsuarioHandler.GetUsuariosByPersona)
	usuarios.Post("/login", handlers.UsuarioHandler.Login)

	// ==================== ESTUDIANTES UNIVERSITARIOS ====================
	estudiantesUniv := app.Group("/estudiantes-universitarios")
	estudiantesUniv.Post("/", handlers.EstudianteUnivHandler.CreateEstudianteUniversitario)
	estudiantesUniv.Get("/", handlers.EstudianteUnivHandler.GetAllEstudiantesUniversitarios)
	estudiantesUniv.Get("/:id", handlers.EstudianteUnivHandler.GetEstudianteUniversitario)
	estudiantesUniv.Put("/:id", handlers.EstudianteUnivHandler.UpdateEstudianteUniversitario)
	estudiantesUniv.Delete("/:id", handlers.EstudianteUnivHandler.DeleteEstudianteUniversitario)
	estudiantesUniv.Get("/semestre/:semestre", handlers.EstudianteUnivHandler.GetEstudiantesUniversitariosBySemestre)
	estudiantesUniv.Get("/persona/:persona_id", handlers.EstudianteUnivHandler.GetEstudianteUniversitarioByPersona)

	// ==================== AUTORIDADES UTEQ ====================
	autoridades := app.Group("/autoridades-uteq")
	autoridades.Post("/", handlers.AutoridadHandler.CreateAutoridadUTEQ)
	autoridades.Get("/", handlers.AutoridadHandler.GetAllAutoridadesUTEQ)
	autoridades.Get("/:id", handlers.AutoridadHandler.GetAutoridadUTEQ)
	autoridades.Put("/:id", handlers.AutoridadHandler.UpdateAutoridadUTEQ)
	autoridades.Delete("/:id", handlers.AutoridadHandler.DeleteAutoridadUTEQ)
	autoridades.Get("/cargo/:cargo", handlers.AutoridadHandler.GetAutoridadesUTEQByCargo)
	autoridades.Get("/persona/:persona_id", handlers.AutoridadHandler.GetAutoridadUTEQByPersona)

	// ==================== TEMÁTICAS ====================
	tematicas := app.Group("/tematicas")
	tematicas.Post("/", handlers.TematicaHandler.CreateTematica)
	tematicas.Get("/", handlers.TematicaHandler.GetAllTematicas)
	tematicas.Get("/:id", handlers.TematicaHandler.GetTematica)
	tematicas.Put("/:id", handlers.TematicaHandler.UpdateTematica)
	tematicas.Delete("/:id", handlers.TematicaHandler.DeleteTematica)
	tematicas.Get("/nombre/:nombre", handlers.TematicaHandler.GetTematicasByNombre)
	tematicas.Get("/descripcion/:descripcion", handlers.TematicaHandler.GetTematicasByDescripcion)

	// ==================== ACTIVIDADES ====================
	actividades := app.Group("/actividades")
	actividades.Post("/", handlers.ActividadHandler.CreateActividad)
	actividades.Get("/", handlers.ActividadHandler.GetAllActividades)
	actividades.Get("/:id", handlers.ActividadHandler.GetActividad)
	actividades.Put("/:id", handlers.ActividadHandler.UpdateActividad)
	actividades.Delete("/:id", handlers.ActividadHandler.DeleteActividad)
	actividades.Get("/tematica/:tematica_id", handlers.ActividadHandler.GetActividadesByTematica)
	actividades.Get("/nombre/:nombre", handlers.ActividadHandler.GetActividadesByNombre)
	actividades.Get("/duracion", handlers.ActividadHandler.GetActividadesByDuracion) // ?min=30&max=120

	// ==================== PROGRAMAS DE VISITA ====================
	programas := app.Group("/programas-visita")
	programas.Post("/", handlers.ProgramaVisitaHandler.CreateProgramaVisita)
	programas.Get("/", handlers.ProgramaVisitaHandler.GetAllProgramasVisita)
	programas.Get("/:id", handlers.ProgramaVisitaHandler.GetProgramaVisita)
	programas.Put("/:id", handlers.ProgramaVisitaHandler.UpdateProgramaVisita)
	programas.Delete("/:id", handlers.ProgramaVisitaHandler.DeleteProgramaVisita)
	programas.Get("/fecha/:fecha", handlers.ProgramaVisitaHandler.GetProgramasVisitaByFecha) // YYYY-MM-DD
	programas.Get("/autoridad/:autoridad_id", handlers.ProgramaVisitaHandler.GetProgramasVisitaByAutoridad)
	programas.Get("/institucion/:institucion_id", handlers.ProgramaVisitaHandler.GetProgramasVisitaByInstitucion)
	programas.Get("/rango-fecha", handlers.ProgramaVisitaHandler.GetProgramasVisitaByRangoFecha) // ?inicio=2024-01-01&fin=2024-12-31

	// ==================== VISITA DETALLES ====================
	detalles := app.Group("/visita-detalles")
	detalles.Post("/", handlers.VisitaDetalleHandler.CreateVisitaDetalle)
	detalles.Get("/", handlers.VisitaDetalleHandler.GetAllVisitaDetalles)
	detalles.Get("/:id", handlers.VisitaDetalleHandler.GetVisitaDetalle)
	detalles.Put("/:id", handlers.VisitaDetalleHandler.UpdateVisitaDetalle)
	detalles.Delete("/:id", handlers.VisitaDetalleHandler.DeleteVisitaDetalle)
	detalles.Get("/estudiante/:estudiante_id", handlers.VisitaDetalleHandler.GetVisitaDetallesByEstudiante)
	detalles.Get("/actividad/:actividad_id", handlers.VisitaDetalleHandler.GetVisitaDetallesByActividad)
	detalles.Get("/programa/:programa_id", handlers.VisitaDetalleHandler.GetVisitaDetallesByPrograma)
	detalles.Get("/participantes", handlers.VisitaDetalleHandler.GetVisitaDetallesByParticipantes) // ?min=10&max=50
	detalles.Get("/estadisticas", handlers.VisitaDetalleHandler.GetEstadisticasParticipacion)

	// ==================== DUDAS ====================
	dudas := app.Group("/dudas")
	dudas.Post("/", handlers.DudasHandler.CreateDudas)
	dudas.Get("/", handlers.DudasHandler.GetAllDudas)
	dudas.Get("/:id", handlers.DudasHandler.GetDudas)
	dudas.Put("/:id", handlers.DudasHandler.UpdateDudas)
	dudas.Delete("/:id", handlers.DudasHandler.DeleteDudas)
	dudas.Get("/estudiante/:estudiante_id", handlers.DudasHandler.GetDudasByEstudiante)
	dudas.Get("/autoridad/:autoridad_id", handlers.DudasHandler.GetDudasByAutoridad)
	dudas.Get("/sin-responder", handlers.DudasHandler.GetDudasSinResponder)
	dudas.Get("/respondidas", handlers.DudasHandler.GetDudasRespondidas)
	dudas.Get("/sin-asignar", handlers.DudasHandler.GetDudasSinAsignar)
	dudas.Get("/buscar/:termino", handlers.DudasHandler.BuscarDudasPorPregunta)
	dudas.Put("/:duda_id/asignar", handlers.DudasHandler.AsignarAutoridadADuda)
	dudas.Put("/:duda_id/responder", handlers.DudasHandler.ResponderDuda)
}

// AllHandlers contiene todos los handlers de la aplicación
type AllHandlers struct {
	EstudianteHandler      *handlers.EstudianteHandler
	PersonaHandler         *handlers.PersonaHandler
	ProvinciaHandler       *handlers.ProvinciaHandler
	CiudadHandler          *handlers.CiudadHandler
	InstitucionHandler     *handlers.InstitucionHandler
	TipoUsuarioHandler     *handlers.TipoUsuarioHandler
	UsuarioHandler         *handlers.UsuarioHandler
	EstudianteUnivHandler  *handlers.EstudianteUniversitarioHandler
	AutoridadHandler       *handlers.AutoridadUTEQHandler
	TematicaHandler        *handlers.TematicaHandler
	ActividadHandler       *handlers.ActividadHandler
	ProgramaVisitaHandler  *handlers.ProgramaVisitaHandler
	VisitaDetalleHandler   *handlers.VisitaDetalleHandler
	DudasHandler           *handlers.DudasHandler
}

// NewAllHandlers crea una instancia con todos los handlers
func NewAllHandlers(
	estudianteHandler *handlers.EstudianteHandler,
	personaHandler *handlers.PersonaHandler,
	provinciaHandler *handlers.ProvinciaHandler,
	ciudadHandler *handlers.CiudadHandler,
	institucionHandler *handlers.InstitucionHandler,
	tipoUsuarioHandler *handlers.TipoUsuarioHandler,
	usuarioHandler *handlers.UsuarioHandler,
	estudianteUnivHandler *handlers.EstudianteUniversitarioHandler,
	autoridadHandler *handlers.AutoridadUTEQHandler,
	tematicaHandler *handlers.TematicaHandler,
	actividadHandler *handlers.ActividadHandler,
	programaVisitaHandler *handlers.ProgramaVisitaHandler,
	visitaDetalleHandler *handlers.VisitaDetalleHandler,
	dudasHandler *handlers.DudasHandler,
) *AllHandlers {
	return &AllHandlers{
		EstudianteHandler:     estudianteHandler,
		PersonaHandler:        personaHandler,
		ProvinciaHandler:      provinciaHandler,
		CiudadHandler:         ciudadHandler,
		InstitucionHandler:    institucionHandler,
		TipoUsuarioHandler:    tipoUsuarioHandler,
		UsuarioHandler:        usuarioHandler,
		EstudianteUnivHandler: estudianteUnivHandler,
		AutoridadHandler:      autoridadHandler,
		TematicaHandler:       tematicaHandler,
		ActividadHandler:      actividadHandler,
		ProgramaVisitaHandler: programaVisitaHandler,
		VisitaDetalleHandler:  visitaDetalleHandler,
		DudasHandler:          dudasHandler,
	}
}