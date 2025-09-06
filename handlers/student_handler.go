package handlers

import (
	"ApiEscuela/models"
	"ApiEscuela/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	studentRepo *repositories.StudentRepository
}

func NewStudentHandler(studentRepo *repositories.StudentRepository) *StudentHandler {
	return &StudentHandler{studentRepo: studentRepo}
}

// CreateStudent crea un nuevo estudiante
func (h *StudentHandler) CreateStudent(c *fiber.Ctx) error {
	var student models.Student
	
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.studentRepo.CreateStudent(&student); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede crear el estudiante",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(student)
}

// GetStudent obtiene un estudiante por ID
func (h *StudentHandler) GetStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante inválido",
		})
	}

	student, err := h.studentRepo.GetStudentByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Estudiante no encontrado",
		})
	}

	return c.JSON(student)
}

// GetAllStudents obtiene todos los estudiantes
func (h *StudentHandler) GetAllStudents(c *fiber.Ctx) error {
	students, err := h.studentRepo.GetAllStudents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los estudiantes",
		})
	}

	return c.JSON(students)
}

// UpdateStudent actualiza un estudiante
func (h *StudentHandler) UpdateStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante inválido",
		})
	}

	student, err := h.studentRepo.GetStudentByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Estudiante no encontrado",
		})
	}

	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se puede procesar el JSON",
		})
	}

	if err := h.studentRepo.UpdateStudent(student); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede actualizar el estudiante",
		})
	}

	return c.JSON(student)
}

// DeleteStudent elimina un estudiante
func (h *StudentHandler) DeleteStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de estudiante inválido",
		})
	}

	if err := h.studentRepo.DeleteStudent(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se puede eliminar el estudiante",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Estudiante eliminado exitosamente",
	})
}

// GetStudentByCedula obtiene un estudiante por cédula
func (h *StudentHandler) GetStudentByCedula(c *fiber.Ctx) error {
	cedula := c.Params("cedula")
	
	student, err := h.studentRepo.GetStudentByCedula(cedula)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Estudiante no encontrado",
		})
	}

	return c.JSON(student)
}

// GetStudentsByCity obtiene estudiantes por ciudad
func (h *StudentHandler) GetStudentsByCity(c *fiber.Ctx) error {
	ciudad := c.Params("ciudad")
	
	students, err := h.studentRepo.GetStudentsByCity(ciudad)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pueden obtener los estudiantes",
		})
	}

	return c.JSON(students)
}