package apis

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Othayz/API/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (api *API) getStudents(c echo.Context) error {
	Students, err :=api.DB.GetStudentByID()
	if err != nil {
	  return c.String(http.StatusNotFound, "error getting students")
	}
	return c.JSON(http.StatusOK, Students)
  }
  
  func (api *API) createStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
	  return err
	}
	if err := api.DB.AddStudent(student); err != nil {
	  return c.String(http.StatusInternalServerError, "error adding student")
	}
  
	return c.String(http.StatusOK, "create a new student")
  }
  
  func (api *API) getStudentsByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student id")
	}
	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
	  return c.String(http.StatusNotFound, "student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to getting student")
	}
	return c.JSON(http.StatusOK, student)
  }
  func (api *API) updateStudentsByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student id")
	}

	receivedStudent := db.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
	  return err
	}

	updatingStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
	  return c.String(http.StatusNotFound, "student not found")
	}
	student := updateStudInfo(&receivedStudent, &updatingStudent)
	if err := api.DB.UpdateStudent(student); err != nil{
	  return c.String(http.StatusInternalServerError, "Failed to save student")
	}

	return c.JSON(http.StatusOK, "Student updated")
  }
  func (api *API) deleteStudentsByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student id")
	}
	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
	  return c.String(http.StatusNotFound, "student not found")
	}
	if err := api.DB.DeleteStudent(student); err != nil{
		return c.String(http.StatusInternalServerError, "Failed to delete student")
	  }
	  return c.JSON(http.StatusOK, "Student deleted")

  }
  func updateStudInfo(receivedStudent, updatingStudent *db.Student) db.Student {
	if receivedStudent.Name != "" {
	  updatingStudent.Name = receivedStudent.Name
	}
	if receivedStudent.CPF > 0 {
	  updatingStudent.CPF = receivedStudent.CPF

	}
	if receivedStudent.Email != "" {
	  updatingStudent.Email = receivedStudent.Email
	}
	if receivedStudent.Age > 0 {
	  updatingStudent.Age = receivedStudent.Age
	}
	if receivedStudent.Active != updatingStudent.Active {
	  updatingStudent.Active = receivedStudent.Active
	}
	return *updatingStudent
  }