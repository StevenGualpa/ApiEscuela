package routers

import (
	"ApiEscuela/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupStudentRoutes(app *fiber.App, studentHandler *handlers.StudentHandler) {
	api := app.Group("/api/v1")
	students := api.Group("/students")

	// CRUD básico
	students.Post("/", studentHandler.CreateStudent)
	students.Get("/", studentHandler.GetAllStudents)
	students.Get("/:id", studentHandler.GetStudent)
	students.Put("/:id", studentHandler.UpdateStudent)
	students.Delete("/:id", studentHandler.DeleteStudent)

	// Rutas adicionales
	students.Get("/cedula/:cedula", studentHandler.GetStudentByCedula)
	students.Get("/ciudad/:ciudad", studentHandler.GetStudentsByCity)
}